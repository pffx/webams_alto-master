package rpcHardwareProtectionElementStateModels

//Demo   <onu-name xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-if-xponvani-aug">C12_ONT2</onu-name>

type C_TagWithAttrAndValue struct {
	Name            string `xml:",chardata"`
	Xmlns           string `xml:"xmlns,attr,omitempty"`
	XmlnsBbfXPonift string `xml:"xmlns:bbf-xponift,attr,omitempty"`
	XmlnsBbfHwt     string `xml:"xmlns:bbf-hwt,attr,omitempty"`
	XmlnsNokiaHwi   string `xml:"xmlns:nokia-hwi,attr,omitempty"`
}

func NewC_TagWithAttrAndValue() *C_TagWithAttrAndValue {
	a := new(C_TagWithAttrAndValue)
	return a
}

// hardware-protection-element-state/component
type Component struct {
	HardwareReference                            string `xml:"hardware-reference,omitempty"`
	ProtectionGroupReference                     string `xml:"protection-group-reference,omitempty"`
	ProtectionElementStandbyStatus               string `xml:"protection-element-standby-status,omitempty"`
	StandbyStatusChangeReason                    string `xml:"standby-status-change-reason,omitempty"`
	StandbyStateChangeTime                       string `xml:"standby-state-change-time,omitempty"`
	StandbyStateChangeTroubleShootingInformation string `xml:"standby-state-change-trouble-shooting-information,omitempty"`
}

func NewComponent() *Component {
	a := new(Component)
	return a
}
