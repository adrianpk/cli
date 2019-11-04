package generator

import (
	"fmt"
	"log"
	"path/filepath"
	"text/template"
)

type (
	serverGenerator struct {
		Meta  *Metadata
		force bool
	}
)

func (g *Generator) GenServer() error {
	md := g.Meta
	sg := serverGenerator{
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

func (sg *serverGenerator) write() error {
	md := sg.Meta

	n := fmt.Sprintf("%ssrv.go", md.Infl.SingularLowercase)
	f := filepath.Join("pkg", md.Pkg.ServicePath, n)

	log.Printf("Server file: %s\n", f)

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

func (sg *serverGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/server.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
