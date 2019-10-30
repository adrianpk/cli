package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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
	}

	Metadata struct {
		// ResName is the name of the resource to be created.
		ResName string  `yaml:"name"`
		Pkg     PkgData `yaml:"pkg"`

		// API data
		API APIData `yaml:"api"`

		// Inflections
		ResInflections InflectionData `yaml:"resInflections""`

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
		SingularDashed string `yaml:"singularDashed"`
		PluralDashed   string `yaml:"pluralDashed"`
	}

	PropDef struct {
		Name            string  `yaml:"name"`
		Type            string  `yaml:"type"`
		Length          int     `yaml:"length"`
		IsVirtual       bool    `yaml:"isVirtual"`
		IsKey           bool    `yaml:"isKey"`
		IsUnique        bool    `yaml:"isUnique"`
		AdmitNull       bool    `yaml:"admitNull"`
		Ref             PropRef `yaml:"ref"`
		IsEmbedded      bool
		IsBackendOnly   bool
		ModelType       string
		NullType        string
		NullTypeMaker   string
		NameInflections InflectionData
		SQL             SQLColData
		Value           interface{}
	}

	PropRef struct {
		Model    string `yaml:"model"`
		Property string `yaml:"property"`
		FKName   string
		TrgTable string
	}

	ModelData struct {
		MatchCondFx string
	}

	SQLData struct {
		CreateSt string
		AlterSt  []string
		DropSt   string
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
		ColName  string
		Type     string
		Modifier string
	}
)

func MakeGenerator(resource, pkg string) (*Generator, error) {
	r := inflector.ToSingularSnakeCase(resource)

	g := Generator{
		res: r,
		pkg: pkg,
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

	md := Metadata{}
	err := yaml.Unmarshal(g.data, &md)
	if err != nil {
		return err
	}

	//log.Println(spew.Sdump(md))

	g.Meta = &md
	return nil
}

func (g *Generator) procMetadata() error {
	panic("not implemented")
	// return nil
}

func projectRootDir() (dir string, err error) {
	return os.Getwd()
}

func errMsg(err error) string {
	return strings.Title(strings.ToLower(err.Error()))
}
