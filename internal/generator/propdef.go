package generator

import (
	"bytes"
	"fmt"
	"strings"

	"gitlab.com/mikrowezel/backend/cli/internal/inflector"
)

func makePropDef(name, propType string, length int, isVirtual, isKey, isUnique, admitNull, isEmbedded, isBackendOnly bool, value interface{}) PropDef {
	return PropDef{
		Name:          name,
		Type:          propType,
		Length:        length,
		IsVirtual:     isVirtual,
		IsKey:         isKey,
		IsUnique:      isUnique,
		AdmitNull:     admitNull,
		IsEmbedded:    isEmbedded,
		IsBackendOnly: isBackendOnly,
		Value:         value,
	}
}

func (p *PropDef) setTypes() {
	pt := p.Type

	switch pt {
	case "id":
		p.ValAccessor = "Int64"
		p.NullType = "sql.NullInt64"
		p.NullTypeMaker = "db.ToNullInt64"

	case "uuid":
		p.Type = "uuid.UUID"
		p.ValAccessor = ""
		p.NullType = "uuid.UUID"
		p.NullTypeMaker = ""

	case "binary":
		p.ValAccessor = "ByteSlice"
		p.NullType = "sql.NullByteSlice"
		p.NullTypeMaker = "db.ToNullByteSlice"

	case "boolean":
		p.ValAccessor = "Bool"
		p.NullType = "sql.NullBool"
		p.NullTypeMaker = "db.ToNullBool"

	case "date":
		p.ValAccessor = "Time"
		p.NullType = "sql.NullTime"
		p.NullTypeMaker = "db.ToNullTime"

	case "datetime":
		p.ValAccessor = "Time"
		p.NullType = "sql.NullTime"
		p.NullTypeMaker = "db.ToNullTime"

	case "decimal":
		p.ValAccessor = "Float"
		p.NullType = "sql.NullFloat64"
		p.NullTypeMaker = "db.ToNullFloat64"

	case "float":
		p.ValAccessor = "Float"
		p.NullType = "sql.NullFloat64"
		p.NullTypeMaker = "db.ToNullFloat64"

	case "geolocation":
		p.ValAccessor = "Point"
		p.NullType = "sql.NullPoint"
		p.NullTypeMaker = "db.ToNullPoint"

	case "integer":
		p.ValAccessor = "Int64"
		p.NullType = "sql.NullInt64"
		p.NullTypeMaker = "db.ToNullInt64"

	case "json":
		p.ValAccessor = "String"
		p.NullType = "sqlxtypes.JSONText"
		p.NullTypeMaker = "db.ToNullsByteSlice"

	case "primary_key":
		p.Type = "uuid.UUID"
		p.ValAccessor = ""
		p.NullType = "uuid.UUID"
		p.NullTypeMaker = ""

	case "string":
		p.ValAccessor = "String"
		p.NullType = "sql.NullString"
		p.NullTypeMaker = "db.ToNullString"

	case "text":
		p.ValAccessor = "String"
		p.NullType = "sql.NullString"
		p.NullTypeMaker = "db.ToNullString"

	case "password":
		p.ValAccessor = "String"
		p.NullType = "sql.NullString"
		p.NullTypeMaker = "db.ToNullString"

	case "password_confirmation":
		p.ValAccessor = "String"
		p.NullType = "sql.NullString"
		p.NullTypeMaker = "db.ToNullString"

	case "time":
		p.ValAccessor = "Time"
		p.NullType = "sql.NullTime"
		p.NullTypeMaker = "db.ToNullTime"

	case "timestamp":
		p.ValAccessor = "Time"
		p.NullType = "sql.NullTime"
		p.NullTypeMaker = "db.ToNullTime"

	case "timestamptz":
		p.ValAccessor = "Time"
		p.NullType = "sql.NullTime"
		p.NullTypeMaker = "db.ToNullTime"

	default:
		p.ValAccessor = "String"
		p.NullType = "sql.NullString"
		p.NullTypeMaker = "db.ToNullString"
	}
}

func (p *PropDef) updateColName() {
	p.Col.Name = inflector.ToSingularSnakeCase(p.Name)
}

func (p *PropDef) updateColType() {
	pt := p.Type
	//propSize := prop.Length
	switch pt {
	case "binary":
		p.Col.Type = "BYTEA"
	case "boolean":
		p.Col.Type = "BOOLEAN"
	case "date":
		p.Col.Type = "DATE"
	case "datetime":
		p.Col.Type = "TIMESTAMP"
	case "decimal":
		p.Col.Type = "FLOAT(24)"
	case "float":
		p.Col.Type = "FLOAT(24)"
	case "geolocation":
		p.Col.Type = "GEOGRAPHY(Point,4326)"
	case "integer":
		p.Col.Type = "BIGINT"
	case "json":
		p.Col.Type = "JSONB"
	case "primary_key":
		p.Col.Type = "UUID"
	case "string":
		p.Col.Type = "VARCHAR(64)"
	case "text":
		p.Col.Type = "TEXT"
	case "time":
		p.Col.Type = "TIME"
	case "timestamp":
		p.Col.Type = "TIMESTAMP"
	case "timestamptz":
		p.Col.Type = "TIMESTAMP WITH TIME ZONE"
	case "uuid.UUID":
		p.Col.Type = "UUID"
	default:
		p.Col.Type = "VARCHAR(64)"
	}
}

func (p *PropDef) updateColModifiers() {
	mq := 0
	var m bytes.Buffer
	if p.IsKey {
		mq = mq + 1
		m.WriteString("PRIMARY KEY")
	}
	if p.IsUnique {
		if mq > 0 {
			m.WriteString(" ")
		}
		m.WriteString("UNIQUE")
	}
	if p.AdmitNull {
		if mq > 0 {
			m.WriteString(" ")
		}
		m.WriteString("NULL")
	}
	p.Col.Modifier = m.String()
}

func (p *PropDef) updateFK() {
	ref := &p.Ref
	emptyRef := &PropRef{}

	if *emptyRef != *ref {
		refModel := inflector.Pluralize(ref.Model)
		refProp := ref.Property
		ref.FKName = strings.ToLower(fmt.Sprintf("%s_%s_fkey", refModel, refProp))
		ref.TrgTable = strings.ToLower(p.Ref.TrgTable)
	}
}

func (p *PropDef) updateShowInClient() {
	if !p.IsEmbedded {
		p.IsBackendOnly = false
	}
}

func (p *PropDef) hasNullTypeMaker() bool {
	return p.NullTypeMaker != ""
}
