package generator

func (md *Metadata) selectNonVirtualPropDefs() {
	props := md.PropDefs
	for i := range md.PropDefs {
		prop := props[i]
		if !prop.IsVirtual {
			md.NonVirtualPropDefs = append(md.NonVirtualPropDefs, prop)
		}
	}
}

func (md *Metadata) selectClientPropDefs() {
	props := md.PropDefs
	for i := range md.PropDefs {
		prop := props[i]
		if !(prop.IsKey || prop.IsBackendOnly) {
			md.ClientPropDefs = append(md.ClientPropDefs, prop)
		}
	}
}

func (md *Metadata) selectTransportPropDefs() {
	props := md.PropDefs
	for i := range md.PropDefs {
		prop := props[i]
		if !(prop.IsKey || prop.IsBackendOnly || prop.Name == "Slug") {
			md.TransportPropDefs = append(md.TransportPropDefs, prop)
		}
	}
}
