package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"gitlab.com/mikrowezel/backend/cli/internal/inflector"
	"gopkg.in/yaml.v2"
)

type (
	Generator struct {
		res        string
		pkg        string
		inFilePath string
		data       []byte
		Meta       *Metadata
		force      bool
	}

	Metadata struct {
		// ResName is the name of the resource to be created.
		ResName string  `yaml:"name"`
		Pkg     PkgData `yaml:"pkg"`

		// API data
		API APIData `yaml:"api"`

		// Inflections
		Infl InflectionData `yaml:"resInflections""`

		// Property definitions
		PropDefs           []PropDef `yaml:"propDefs"`
		NonVirtualPropDefs []PropDef
		ClientPropDefs     []PropDef

		// Model
		Model ModelData

		// SQL
		SQL SQLData

		// REST Client
		REST RESTClientData

		// Service
		Service ServiceData
	}

	APIData struct {
		Version string `yaml:"version"`
	}

	PkgData struct {
		Name        string
		Dir         string
		ServicePath string `yaml:"servicePath"`
	}

	InflectionData struct {
		Singular string
		Plural   string `yaml:"plural"`
		// Lowercase
		SingularLowercase string `yaml:"singularLowercase"`
		PluralLowercase   string `yaml:"pluralLowercase"`
		// Uppercase
		SingularUppercase string `yaml:"singularUppercase"`
		PluralUppercase   string `yaml:"pluralUppercase"`
		// Camel case
		SingularCamelCase string `yaml:"singularCamelCase"`
		PluralCamelCase   string `yaml:"pluralCamelCase"`
		// Pascal case
		SingularPascalCase string `yaml:"singulaPascalCase"`
		PluralPascalCase   string `yaml:"pluralPascalCase"`
		// Snake case
		SingularSnakeCase string `yaml:"singularSnakeCase"`
		PluralSnakeCase   string `yaml:"pluralSnakeCase"`
		// Dashed
		SingularDashedCase string `yaml:"singularDashed"`
		PluralDashedCase   string `yaml:"pluralDashed"`
	}

	PropDef struct {
		Name          string  `yaml:"name"`
		Type          string  `yaml:"type"`
		Length        int     `yaml:"length"`
		IsVirtual     bool    `yaml:"isVirtual"`
		IsKey         bool    `yaml:"isKey"`
		IsUnique      bool    `yaml:"isUnique"`
		AdmitNull     bool    `yaml:"admitNull"`
		Ref           PropRef `yaml:"ref"`
		IsEmbedded    bool
		IsBackendOnly bool
		ModelType     string
		NullType      string
		NullTypeMaker string
		Infl          InflectionData
		Col           SQLColData
		Value         interface{}
	}

	PropRef struct {
		Model    string `yaml:"model"`
		Property string `yaml:"property"`
		FKName   string
		TrgTable string
	}

	ModelData struct {
		MatchCond string
	}

	SQLData struct {
		ColModifier string
		CreateSt    string
		AlterSt     []string
		DropSt      string
		// Repo
		InsertSt       string
		SelectAllSt    string
		SelectByIDSt   string
		SelectBySlugSt string
		UpdateSt       string
		DeleteByIDSt   string
		DeleteBySlugSt string
	}

	RESTClientData struct {
		CreateJSON string
		UpdateJSON string
	}

	ServiceData struct {
		Test TestData
	}

	TestData struct {
		CreateMap string
		UpdateMap string
		SampleMap []string

		// Struct
		CreateStruct string
		UpdateStruct string
		SampleStruct []string
	}

	SQLColData struct {
		Name     string
		Type     string
		Modifier string
		Value    interface{}
	}
)

func MakeGenerator(resource, pkg string, force bool) (*Generator, error) {
	r := inflector.ToSingularSnakeCase(resource)

	g := Generator{
		res:   r,
		pkg:   pkg,
		force: force,
	}

	err := g.LoadMeta()
	if err != nil {
		return &g, err
	}

	return &g, nil
}

// LoadMeta parse model metadata YAML file
// to build generator metadata.
func (g *Generator) LoadMeta() error {
	g.inFilePath = fmt.Sprintf("assets/gen/%s.yaml", g.res)

	return g.genMeta()
}

func (g *Generator) genMeta() error {
	err := g.readFile()
	if err != nil {
		return err
	}

	err = g.parseData()
	if err != nil {
		return err
	}

	err = g.procMetadata()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) readFile() error {
	log.Printf("Reading input file: '%s'\n", g.inFilePath)

	data, err := ioutil.ReadFile(g.inFilePath)
	if err != nil {
		return fmt.Errorf("cannot read input file: %s", g.inFilePath)
	}

	g.data = data

	return nil
}

func (g *Generator) parseData() error {
	log.Println("Generating metadata")

	md := Metadata{
		Pkg:      PkgData{},
		API:      APIData{},
		Infl:     InflectionData{},
		PropDefs: []PropDef{},
		Model:    ModelData{},
		SQL:      SQLData{},
		REST:     RESTClientData{},
		Service:  ServiceData{},
	}

	err := yaml.Unmarshal(g.data, &md)
	if err != nil {
		return err
	}

	log.Println(spew.Sdump(md))

	g.Meta = &md
	return nil
}

func (g *Generator) procMetadata() error {
	if g.pkg != "" {
		// Use flag value instead value from YAML file.
		g.Meta.Pkg.Name = g.pkg
	}

	md := g.Meta
	md.ResName = inflector.UpercaseFirst(md.ResName)

	md.genInflections()
	md.addIdentification()
	md.addAudit()

	props := md.PropDefs
	for i := range props {
		p := &props[i]
		p.setTypes() //SafeType = safeType(prop)

		p.Infl.SingularCamelCase = inflector.ToSingularCamelCase(p.Name)
		p.Infl.SingularPascalCase = inflector.ToSingularPascalCase(p.Name)
		p.Infl.SingularDashedCase = inflector.ToSingularDashedCase(p.Name)
		p.updateColName()
		p.updateColType()

		p.updateColModifiers()
		p.updateFK()
		p.updateShowInClient()
	}

	md.selectNonVirtualPropDefs()
	md.selectClientPropDefs()

	return nil
}

func fileWriter(outputFile string, force bool) (*os.File, error) {
	return fileWriterWithPerm(outputFile, force, 0666)
}

func fileWriterWithPerm(outputFile string, force bool, perm int) (*os.File, error) {
	flag := os.O_CREATE | os.O_EXCL
	if force {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	p := os.FileMode(perm)
	return os.OpenFile(outputFile, flag, p)
}

func projectRootDir() (dir string, err error) {
	return os.Getwd()
}

func errMsg(err error) string {
	return strings.Title(strings.ToLower(err.Error()))
}
