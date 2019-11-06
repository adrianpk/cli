package generator

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"text/template"
)

type (
	jsonrestGenerator struct {
		Meta  *Metadata
		force bool
	}
)

func (g *Generator) GenModel() error {
	md := g.Meta
	jrg := jsonrestGenerator{
		Meta:  md,
		force: g.force,
	}

	jrg.updateMetadata()

	err := jrg.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
		return err
	}

	log.Println("Done!")
	return nil
}

func (jrg *jsonrestGenerator) updateMetadata() {
	jrg.genMatchCondition()
}

func (jrg *jsonrestGenerator) genMatchCondition() {
	md := jrg.Meta
	props := md.ClientPropDefs
	l := len(props) - 1
	var mcond bytes.Buffer
	mcond.WriteString(fmt.Sprintf("%s.Identification.Match(tc.Identification) &&\n", md.Infl.SingularCamelCase))
	for i, prop := range props {
		if prop.Name != "ID" {
			prop := props[i]
			var line string
			if l == i {
				line = fmt.Sprintf("\t\t%s.%s == tc.%s\n", md.Infl.SingularCamelCase, prop.Name, prop.Name)
			} else {
				line = fmt.Sprintf("\t\t%s.%s == tc.%s &&\n", md.Infl.SingularCamelCase, prop.Name, prop.Name)
			}
			mcond.WriteString(line)
		}
	}
	md.Model.MatchCond = mcond.String()
}

func (jrg *jsonrestGenerator) write() error {
	md := jrg.Meta

	n := fmt.Sprintf("%s.go", md.Infl.SingularLowercase)
	f := filepath.Join("pkg", md.Pkg.ServicePath, "jsonrest", n)

	log.Printf("Service file: %s\n", f)

	w, err := fileWriter(f, jrg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := jrg.template()
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (jrg *jsonrestGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/jsonrest.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
