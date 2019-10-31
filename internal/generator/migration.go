package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"time"
)

type (
	migrationGenerator struct {
		Meta  *Metadata
		force bool
	}
)

func (g *Generator) GenMigration() error {
	md := g.Meta
	mg := migrationGenerator{
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

func (mg *migrationGenerator) updateMetadata() {
	mg.makeCreateSt()
	mg.makeDropSt()
	mg.makeFKAlterSt()
}

func (mg *migrationGenerator) makeCreateSt() {
	md := mg.Meta
	props := md.PropDefs
	var createSQL bytes.Buffer
	createSQL.WriteString(fmt.Sprintf("CREATE TABLE %s\n\t(\n", md.Infl.PluralSnakeCase))
	last := len(props) - 1
	for i := range props {
		prop := props[i]
		var ending string
		if i < last {
			if prop.Col.Modifier != "" {
				ending = fmt.Sprintf("%s %s,\n", prop.Col.Type, prop.Col.Modifier)
			} else {
				ending = fmt.Sprintf("%s,\n", prop.Col.Type)
			}
		} else {
			if prop.Col.Modifier != "" {
				ending = fmt.Sprintf("%s %s\n", prop.Col.Type, prop.Col.Modifier)
			} else {
				ending = fmt.Sprintf("%s\n", prop.Col.Type)
			}
		}
		createSQL.WriteString(fmt.Sprintf("\t %s %s", prop.Col.Name, ending))
	}
	createSQL.WriteString("\t);")
	md.SQL.CreateSt = createSQL.String()
}

func (mg *migrationGenerator) makeDropSt() {
	md := mg.Meta
	var dropSQL bytes.Buffer
	dropSQL.WriteString(fmt.Sprintf("DROP TABLE %s CASCADE;", md.Infl.PluralSnakeCase))
	md.SQL.DropSt = dropSQL.String()
}

func (mg *migrationGenerator) makeFKAlterSt() {
	md := mg.Meta
	props := md.NonVirtualPropDefs
	for i := range props {
		prop := props[i]
		var alterSQL bytes.Buffer
		if prop.Ref.FKName != "" {
			var a bytes.Buffer
			a.WriteString(fmt.Sprintf("ALTER TABLE %s\n", md.Infl.PluralSnakeCase))
			a.WriteString(fmt.Sprintf("\tADD CONSTRAINT %s\n", prop.Ref.FKName))
			a.WriteString(fmt.Sprintf("\tFOREIGN KEY (%s)\n", prop.Col.Name))
			a.WriteString(fmt.Sprintf("\tREFERENCES %s\n", prop.Ref.TrgTable))
			a.WriteString("\tON DELETE CASCADE;")
			alterSQL.WriteString(a.String())
			md.SQL.AlterSt = append(md.SQL.AlterSt, alterSQL.String())
		}
	}
}

func (mg *migrationGenerator) write() error {
	md := mg.Meta

	n := fmt.Sprintf("%screatetable%s.go", filePrefix(), md.Infl.PluralPascalCase)
	f := filepath.Join("internal", "migration", n)

	log.Printf("Migration file: %s\n", f)

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

func (mg *migrationGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/migration.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}

func filePrefix() string {
	t := time.Now()
	return t.Format("20060102150405")
}
