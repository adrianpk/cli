package inflector

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"

	"github.com/fatih/camelcase"
)

func ToSingular(str string) string {
	return Singularize(str)
}

func ToPlural(str string) string {
	return Pluralize(str)
}

func ToSingularLowercase(str string) string {
	return strings.ToLower(ToSingular(str))
}

func ToPluralLowercase(str string) string {
	return strings.ToLower(ToPlural(str))
}

func ToSingularUppercase(str string) string {
	return strings.ToLower(ToSingular(str))
}

func ToPluralUppercase(str string) string {
	return strings.ToUpper(ToPlural(str))
}

func ToSingularCamelCase(str string) string {
	return ToCamelCase(ToSingular(str))
}

func ToPluralCamelCase(str string) string {
	return ToCamelCase(ToPlural(str))
}

func ToSingularPascalCase(str string) string {
	return ToPascalCase(ToSingular(str))
}

func ToPluralPascalCase(str string) string {
	return ToPascalCase(ToPlural(str))
}

func ToSingularSnakeCase(str string) string {
	return ToSnakeCase(ToSingular(str))
}

func ToPluralSnakeCase(str string) string {
	return ToSnakeCase(ToPlural(str))
}

func ToSingularDashedCase(str string) string {
	return ToDashedCase(ToSingular(str))
}

func ToPluralDashedCase(str string) string {
	return ToDashedCase(ToPlural(str))
}

func UpercaseFirst(str string) string {
	temp := []rune(str)
	temp[0] = unicode.ToUpper(temp[0])
	return string(temp)
}

func LowercaseFirst(str string) string {
	temp := []rune(str)
	temp[0] = unicode.ToLower(temp[0])
	return string(temp)
}

func ToCamelCase(str string) string {
	camelCased := toCamelCaseString(str)
	splitted := camelcase.Split(camelCased)
	splitted[0] = strings.ToLower(splitted[0])
	return strings.Join(splitted, "")
}

func ToPascalCase(str string) string {
	camelCased := toCamelCaseString(str)
	splitted := camelcase.Split(camelCased)
	splitted[0] = UpercaseFirst(splitted[0])
	return strings.Join(splitted, "")
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToDashedCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	dashed := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	dashed = matchAllCap.ReplaceAllString(dashed, "${1}-${2}")
	return strings.ToLower(dashed)
}

func toCamelCaseString(str string) string {
	var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")
	byteSrc := []byte(str)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	//fmt.Printf("Original: %s - Camel: %s", str, string(bytes.Join(chunks, nil)))
	return string(bytes.Join(chunks, nil))
}
