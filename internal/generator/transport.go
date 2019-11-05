package generator

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

type (
	transportGenerator struct {
		Meta  *Metadata
		force bool
	}
)

func (g *Generator) GenTransport() error {
	md := g.Meta
	tg := transportGenerator{
		Meta:  md,
		force: g.force,
	}

	tg.updateMetadata()

	err := tg.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
	}

	log.Println("Done!")
	return nil
}

func (tg *transportGenerator) updateMetadata() {
	tg.genToModel()
	tg.genFromModel()
}

func (tg *transportGenerator) genToModel() {
	md := tg.Meta
	var tm strings.Builder

	tm.WriteString(fmt.Sprintf("model.%s{\n", md.Infl.SingularPascalCase))

	for _, prop := range md.TransportPropDefs {
		spc := prop.Infl.SingularPascalCase
		ntm := prop.NullTypeMaker

		// NOTE: Check if there is a better way to do this
		// Get rid of the embedded Identification and
		// replicate its behavior in each generated model?
		if prop.Name == "Slug" {
			tm.WriteString("\t\tIdentification: m.Identification{\n")
			tm.WriteString("\t\t\tSlug: db.ToNullString(req.Identifier.Slug),\n")
			tm.WriteString("\t\t},\n")
			continue
		}

		if prop.hasNullTypeMaker() {
			tm.WriteString(fmt.Sprintf("\t\t%s:\t%s(req.%s),\n", spc, ntm, spc))

		} else {
			tm.WriteString(fmt.Sprintf("\t\t%s:\treq.%s,\n", spc, spc))
		}
	}
	tm.WriteString("\t}")

	md.Transport.ToModelStruct = tm.String()
}

func (tg *transportGenerator) genFromModel() {
	md := tg.Meta
	var fm strings.Builder

	fm.WriteString(fmt.Sprintf("%s{\n", md.Infl.SingularPascalCase))

	for _, prop := range md.TransportPropDefs {
		spc := prop.Infl.SingularPascalCase
		v := prop.Infl.SingularPascalCase
		va := prop.ValAccessor
		if prop.hasNullTypeMaker() {
			fm.WriteString(fmt.Sprintf("\t\t\t%s:\tm.%s.%s,\n", spc, spc, va))
		} else {
			fm.WriteString(fmt.Sprintf("\t\t\t%s:\tm.%s,\n", spc, v))
		}
	}
	fm.WriteString("\t\t}")

	md.Transport.FromModelStruct = fm.String()
}

func (tg *transportGenerator) write() error {
	err := tg.writeFile("", "transport")
	if err != nil {
		return err
	}

	return tg.writeFile("cnv", "transportcnv")
}

func (tg *transportGenerator) writeFile(sufix, template string) error {
	md := tg.Meta

	n := fmt.Sprintf("%s%s.go", md.Infl.SingularLowercase, sufix)
	f := filepath.Join("pkg", md.Pkg.ServicePath, "transport", n)

	log.Printf("Transport file: %s\n", f)

	w, err := fileWriter(f, tg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := tg.template(template)
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (tg *transportGenerator) template(name string) (*template.Template, error) {
	path := fmt.Sprintf("assets/templates/%s.tmpl", name)
	res, err := Asset(path)
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
