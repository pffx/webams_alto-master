package rpcXpongemtcontStateModels

import "alto_server/constants"

//tconts
type Tconts struct {
	Tcont []Tcont `xml:"tcont,omitempty"`
}

func NewTconts() *Tconts {
	a := new(Tconts)
	return a
}

//tconts/tcont
type Tcont struct {
	Operation                   string `xml:"operation,omitempty,attr"`
	Name                        string `xml:"name,omitempty"`
	ActualAllocId               string `xml:"actual-alloc-id,omitempty"`
	InterfaceReference          string `xml:"interface-reference,omitempty"`
	TrafficDescriptorProfileRef string `xml:"traffic-descriptor-profile-ref,omitempty"`
}

func NewTcontDel() *Tcont {
	a := new(Tcont)
	a.Operation = constants.OPERATION_DELETE
	return a
}

//gemports
type Gemports struct {
	Gemport []Gemport `xml:"gemport,omitempty"`
}

func NewGemports() *Gemports {
	a := new(Gemports)
	return a
}

//gemports/gemport
type Gemport struct {
	Operation       string `xml:"operation,omitempty,attr"`
	Name            string `xml:"name,omitempty"`
	ActualGemportId string `xml:"actual-gemport-id,omitempty"`
}

func NewGemportDel() *Gemport {
	a := new(Gemport)
	a.Operation = constants.OPERATION_DELETE
	return a
}

//traffic-descriptor-profiles
type TrafficDescriptorProfiles struct {
	TrafficDescriptorProfile []TrafficDescriptorProfile `xml:"traffic-descriptor-profile,omitempty"`
}

//traffic-descriptor-profiles/traffic-descriptor-profile
type TrafficDescriptorProfile struct {
	Name string `xml:"name,omitempty"`
}
