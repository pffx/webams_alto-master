package rpcTmProfiles

type TmProfiles struct {
	ShaperProfile []ShaperProfile `xml:"shaper-profile,omitempty"`
}

//tm-profiles/shaper-profile
type ShaperProfile struct {
	Name string `xml:"name,omitempty"`
}
