package generator

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"text/template"
)

type (
	modelGenerator struct {
		Meta  *Metadata
		force bool
	}
)

func (g *Generator) GenModel() error {
	md := g.Meta
	mg := modelGenerator{
		Meta:  md,
		force: g.force,
	}

	mg.updateMetadata()

	err := mg.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
		return err
	}

	log.Println("Done!")
	return nil
}

func (mg *modelGenerator) updateMetadata() {
	mg.genMatchCondition()
}

func (mg *modelGenerator) genMatchCondition() {
	md := mg.Meta
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

func (mg *modelGenerator) write() error {
	md := mg.Meta

	n := fmt.Sprintf("%s.go", md.Infl.SingularLowercase)
	f := filepath.Join("internal", "model", n)

	log.Printf("Model file: %s\n", f)

	w, err := fileWriter(f, mg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := mg.template()
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (mg *modelGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/model.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
