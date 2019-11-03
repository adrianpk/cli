package generator

import (
	"gitlab.com/mikrowezel/backend/cli/internal/inflector"
)

func (md *Metadata) genInflections() {
	md.genPlural()
	md.genSingularLowercase()
	md.genPluralLowercase()
	md.genSingularUppercase()
	md.genPluralUppercase()
	md.genSingularCamelCase()
	md.genPluralCamelCase()
	md.genSingularPascalCase()
	md.genPluralPascalCase()
	md.genSingularSnakeCase()
	md.genPluralSnakeCase()
	md.genSingularDashed()
	md.genPluralDashed()
}

func (md *Metadata) genPlural() {
	md.Infl.Plural = inflector.Pluralize(md.ResName)
}

// Lowercase
func (md *Metadata) genSingularLowercase() {
	md.Infl.SingularLowercase = inflector.ToSingularLowercase(md.ResName)
}

func (md *Metadata) genPluralLowercase() {
	md.Infl.PluralLowercase = inflector.ToPluralLowercase(md.ResName)
}

// Uppercase
func (md *Metadata) genSingularUppercase() {
	md.Infl.SingularUppercase = inflector.ToSingularUppercase(md.ResName)
}

func (md *Metadata) genPluralUppercase() {
	md.Infl.PluralLowercase = inflector.ToPluralUppercase(md.ResName)
}

// CamelCase
func (md *Metadata) genSingularCamelCase() {
	md.Infl.SingularCamelCase = inflector.ToSingularCamelCase(md.ResName)
}

func (md *Metadata) genPluralCamelCase() {
	md.Infl.PluralCamelCase = inflector.ToPluralCamelCase(md.ResName)
}

// PascalCase
func (md *Metadata) genSingularPascalCase() {
	md.Infl.SingularPascalCase = inflector.ToSingularPascalCase(md.ResName)
}

func (md *Metadata) genPluralPascalCase() {
	md.Infl.PluralPascalCase = inflector.ToPluralPascalCase(md.ResName)
}

// SnakeCase
func (md *Metadata) genSingularSnakeCase() {
	md.Infl.SingularSnakeCase = inflector.ToSingularSnakeCase(md.ResName)
}

func (md *Metadata) genPluralSnakeCase() {
	md.Infl.PluralSnakeCase = inflector.ToPluralSnakeCase(md.ResName)
}

// Dashed
func (md *Metadata) genSingularDashed() {
	md.Infl.SingularDashedCase = inflector.ToSingularDashedCase(md.ResName)
}

func (md *Metadata) genPluralDashed() {
	md.Infl.PluralDashedCase = inflector.ToPluralDashedCase(md.ResName)
}
