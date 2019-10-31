package generator

func (md *Metadata) addIdentification() {
	p := makePropDef("ID", "primary_key", 36, false, true, true, false, true, true, "")
	md.PropDefs = append([]PropDef{p}, md.PropDefs...)
	p = makePropDef("TenantID", "string", 128, false, true, true, false, true, true, "")
	md.PropDefs = append([]PropDef{p}, md.PropDefs...)
	p = makePropDef("Slug", "string", 36, false, false, true, false, true, false, "")
	md.PropDefs = append(md.PropDefs, p)
}

func (md *Metadata) addAudit() {
	p := makePropDef("CreatedBy", "uuid", 36, false, false, false, false, true, true, "")
	p.Col.Name = "created_by_id"
	// PropDef.Ref = PropertyRef{Property: "id"}
	md.PropDefs = append(md.PropDefs, p)
	p = makePropDef("UpdatedBy", "uuid", 36, false, false, false, false, true, true, "")
	p.Col.Name = "updated_by_id"
	// PropDef.Ref = PropertyRef{Property: "id"}
	md.PropDefs = append(md.PropDefs, p)
	p = makePropDef("CreatedAt", "timestamptz", 0, false, false, false, false, true, true, "")
	md.PropDefs = append(md.PropDefs, p)
	p = makePropDef("UpdatedAt", "timestamptz", 0, false, false, false, false, true, true, "")
	md.PropDefs = append(md.PropDefs, p)
}
