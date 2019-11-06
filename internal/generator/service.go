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

	sg.updateMetadata()

	err := sg.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
		return err
	}

	log.Println("Done!")
	return nil
}

func (sg *serviceGenerator) updateMetadata() {
	sg.genTestMaps()
	sg.genTestStructs()
}

func (sg *serviceGenerator) write() error {
	err := sg.writeFile("", "service")
	if err != nil {
		return err
	}

	return sg.writeFile("_test", "servicetest")
}

func (sg *serviceGenerator) writeFile(sufix, template string) error {
	md := sg.Meta

	n := fmt.Sprintf("%s%s.go", md.Infl.SingularLowercase, sufix)
	f := filepath.Join("pkg", md.Pkg.ServicePath, "service", n)

	log.Printf("Service file: %s\n", f)

	w, err := fileWriter(f, sg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := sg.template(template)
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (sg *serviceGenerator) template(name string) (*template.Template, error) {
	path := fmt.Sprintf("assets/templates/%s.tmpl", name)
	res, err := Asset(path)
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
