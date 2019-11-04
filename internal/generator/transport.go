package generator

import (
	"fmt"
	"log"
	"path/filepath"
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

	err := tg.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
		return err
	}

	log.Println("Done!")
	return nil
}

func (tg *transportGenerator) write() error {
	md := tg.Meta

	n := fmt.Sprintf("%s.go", md.Infl.SingularLowercase)
	f := filepath.Join("pkg", md.Pkg.ServicePath, "transport", n)

	log.Printf("Transport file: %s\n", f)

	w, err := fileWriter(f, tg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := tg.template()
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (tg *transportGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/transport.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
