package generator

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"text/template"
)

type (
	repoGenerator struct {
		Meta  *Metadata
		force bool
	}
)

func (g *Generator) GenRepo() error {
	md := g.Meta
	mg := repoGenerator{
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

func (mg *repoGenerator) updateMetadata() {
	mg.genInsertSt()
	mg.genSelectAllSt()
	mg.genSelectByIDSt()
	mg.genSelectBySlugSt()
	//mg.genUpdateSt()
	//mg.genDeleteByIDSt()
	//mg.genDeleteBySlugSt()
}

func (mg *repoGenerator) genInsertSt() {
	md := mg.Meta
	props := md.NonVirtualPropDefs

	var st bytes.Buffer
	var vals bytes.Buffer

	st.WriteString(fmt.Sprintf("INSERT INTO %s (", md.Infl.PluralSnakeCase))
	vals.WriteString("VALUES (")

	last := len(props) - 1

	for i := range props {
		prop := props[i]
		var col string
		var val string

		if i < last {
			col = fmt.Sprintf("%s, ", prop.Col.Name)
			val = fmt.Sprintf(":%s, ", prop.Col.Name)

		} else {
			col = fmt.Sprintf("%s) ", prop.Col.Name)
			val = fmt.Sprintf(":%s);", prop.Col.Name)
		}

		st.WriteString(col)
		vals.WriteString(val)
	}
	st.WriteString(vals.String())
	md.SQL.InsertSt = st.String()
}

func (mg *repoGenerator) genSelectAllSt() {
	md := mg.Meta
	md.SQL.SelectAllSt = fmt.Sprintf("SELECT * FROM %s;", md.Infl.PluralSnakeCase)
}

func (mg *repoGenerator) genSelectByIDSt() {
	md := mg.Meta
	md.SQL.SelectByIDSt = fmt.Sprintf("SELECT * FROM %s WHERE id = '%%s' LIMIT 1;", md.Infl.PluralSnakeCase)
}

func (mg *repoGenerator) genSelectBySlugSt() {
	md := mg.Meta
	md.SQL.SelectByIDSt = fmt.Sprintf("SELECT * FROM %s WHERE slug = '%%s' LIMIT 1;", md.Infl.PluralSnakeCase)
}

//func (mg *repoGenerator) genUpdateSt() {
//md := mg.Meta
//pos := 0

//changes := mg.updateChanges()

//qty := len(md.ClientPropDefs)
//last := qty < 2

//var st bytes.Buffer

//st.WriteString(fmt.Sprintf("UPDATE %s SET ", md.Infl.PluralSnakeCase))

//for p, sp := range changes {
//var partial string
//if last {
//partial = fmt.Sprintf("%v = %v ", p, sp)
//} else {
//partial = fmt.Sprintf("%v = %v, ", p, sp)
//}

//st.WriteString(partial)

//pos = pos + 1
//last = pos == qty-1
//}

//st.WriteString("WHERE id = '%%s';")

//return err
//}

func (mg *repoGenerator) write() error {
	md := mg.Meta

	n := fmt.Sprintf("%s.go", md.Infl.SingularLowercase)
	f := filepath.Join("internal", "repo", n)

	log.Printf("Repo file: %s\n", f)

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

func (mg *repoGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/repo.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
