package generator

import (
	"bytes"
	"fmt"
	"time"
)

func (sg *serviceGenerator) genTestMaps() {
	md := sg.Meta
	props := md.ClientPropDefs
	end := ",\n"

	// Maps
	var crt bytes.Buffer
	var upd bytes.Buffer
	var smpl1 bytes.Buffer
	var smpl2 bytes.Buffer

	//pc := md.Infl.SingularPascalCase
	cc := md.Infl.SingularCamelCase

	// Create
	crt.WriteString(fmt.Sprintf("\tcreate%sValid = map[string]interface{}{\n", cc))

	for i, p := range props {
		if p.IsEmbedded || p.IsBackendOnly {
			continue
		}

		prop := props[i]
		var line string
		varName := p.Infl.SingularCamelCase
		// Create
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		crt.WriteString(fmt.Sprintf("%s", line))

		// Update
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		upd.WriteString(fmt.Sprintf("%s", line))

		// Sample 1
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		smpl1.WriteString(fmt.Sprintf("%s", line))

		// Sample 2
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		smpl2.WriteString(fmt.Sprintf("%s", line))
	}

	// Create
	crt.WriteString("\t}")

	md.ServiceTest.CreateMap = crt.String()
	md.ServiceTest.UpdateMap = upd.String()
	md.ServiceTest.SampleMap = []string{smpl1.String(), smpl2.String()}
}

func sampleFormattedVal(p PropDef) string {
	return formatVal(sampleVal(p), p)
}

func sampleVal(p PropDef) string {
	if p.isKeyType() {
		u := generateUUIDString()
		return fmt.Sprintf("\"%s\"", u)
	} else if p.isUUIDType() {
		u := generateUUIDString()
		return fmt.Sprintf("\"%s\"", u)
	} else if p.isTextType() {
		return fmt.Sprintf("\"%s\"", sampleString(p.Name, 4, 8))
	} else if p.isIntType() {
		return fmt.Sprintf("%d", sampleInt(2))
	} else if p.isDecimalType() {
		return fmt.Sprintf("%f", sampleDecimal(4))
	} else if p.isBooleanType() {
		return fmt.Sprintf("%t", bg.SampleBool())
	} else if p.isTimeType() {
		return fmt.Sprintf("\"%s\"", time.Now().Format(time.RFC3339))
	}
	return fmt.Sprintf("\"%s\"", "")
}

func formatVal(val interface{}, p PropDef) string {
	if p.isKeyType() {
		return fmt.Sprintf("\"%s\"", val)
	} else if p.isUUIDType() {
		return fmt.Sprintf("\"%s\"", val)
	} else if p.isTextType() {
		return fmt.Sprintf("\"%s\"", val)
	} else if p.isIntType() {
		return fmt.Sprintf("%d", val)
	} else if p.isDecimalType() {
		return fmt.Sprintf("%f", val)
	} else if p.isBooleanType() {
		return fmt.Sprintf("%t", val)
	} else if p.isTimeType() {
		return fmt.Sprintf("\"%s\"", val)
	}
	return fmt.Sprintf("\"%s\"", val)
}

func (sg *serviceGenerator) genTestStructs() {
	md := sg.Meta
	props := md.ClientPropDefs

	//l := len(props) - 1
	end := ",\n"

	// Maps
	var crt bytes.Buffer
	var upd bytes.Buffer
	var smpl1 bytes.Buffer
	var smpl2 bytes.Buffer

	pc := md.Infl.SingularPascalCase
	cc := md.Infl.SingularCamelCase

	// Create
	crt.WriteString(fmt.Sprintf("\treq := Create%sReq{\n", pc))
	crt.WriteString(fmt.Sprintf("\t\t%s{\n", pc))

	// Update
	upd.WriteString(fmt.Sprintf("\treq := Update%sReq{\n", pc))
	upd.WriteString(fmt.Sprintf("\t\t%s.Identifier{\n", "mod"))
	upd.WriteString(fmt.Sprintf("\t\t\tSlug: %s.Slug.String,\n", cc))
	upd.WriteString(fmt.Sprintf("\t\t},\n"))
	upd.WriteString(fmt.Sprintf("\t\t%s{\n", pc))

	// Sample1
	smpl1.WriteString(fmt.Sprintf("\t%s1 := &model.%s{\n", cc, pc))

	// Sample2
	smpl2.WriteString(fmt.Sprintf("\t%s2 := &model.%s{\n", cc, pc))

	for i, p := range props {
		if p.IsBackendOnly {
			fmt.Printf("BEO %+v\n\n", p)
			continue
		}

		p := props[i]
		var line string
		field := p.Infl.SingularPascalCase
		key := p.Name
		stm := p.NullTypeMaker

		// Create
		line = fmt.Sprintf("\t\t\t%s:\tcreate%sValid[\"%s\"]%s", field, pc, key, end)
		crt.WriteString(fmt.Sprintf("%s", line))

		// Update
		line = fmt.Sprintf("\t\t\t%s:\tupdate%sValid[\"%s\"]%s", field, pc, key, end)
		upd.WriteString(fmt.Sprintf("%s", line))

		// Sample 1
		line = fmt.Sprintf("\t\t\t%s:\t%s(sample%s1[\"%s\"]) %s", field, stm, pc, key, end)
		smpl1.WriteString(fmt.Sprintf("%s", line))

		// Sample 2
		line = fmt.Sprintf("\t\t\t%s:\t%s(sample%s2[\"%s\"]) %s", field, stm, pc, key, end)
		smpl2.WriteString(fmt.Sprintf("%s", line))
	}

	// Create
	crt.WriteString("\t\t},\n")
	crt.WriteString("\t}")

	// Update
	upd.WriteString("\t\t},\n")
	upd.WriteString("\t}\n")

	// Sample1
	smpl1.WriteString("\t}\n")

	// Sample2
	smpl2.WriteString("\t}\n")

	md.ServiceTest.CreateStruct = crt.String()
	md.ServiceTest.UpdateStruct = upd.String()
	md.ServiceTest.SampleStruct = []string{smpl1.String(), smpl2.String()}
}
