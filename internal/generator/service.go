package generator

import (
	"fmt"
	"log"
	"path/filepath"
	"text/template"
)

type (
	serviceGenerator struct {
		Meta  *Metadata
		force bool
	}
)

func (g *Generator) GenService() error {
	md := g.Meta
	sg := serviceGenerator{
		Meta:  md,
		force: g.force,
	}

	err := sg.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
		return err
	}

	log.Println("Done!")
	return nil
}

func (sg *serviceGenerator) write() error {
	md := sg.Meta

	n := fmt.Sprintf("%s.go", md.Infl.SingularLowercase)
	f := filepath.Join("pkg", md.Pkg.ServicePath, "service", n)

	log.Printf("Service file: %s\n", f)

	w, err := fileWriter(f, sg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := sg.template()
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (sg *serviceGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/service.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
