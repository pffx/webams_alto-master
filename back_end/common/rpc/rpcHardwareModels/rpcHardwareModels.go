package rpcHardwareModels

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

//hardware/component
type Component struct {
	Name         string                 `xml:"name,omitempty"`
	Class        *C_TagWithAttrAndValue `xml:"class,omitempty"`
	Parent       string                 `xml:"parent,omitempty"`
	ParentRelPos string                 `xml:"parent-rel-pos,omitempty"`
	AdminState   string                 `xml:"admin-state,omitempty"`
	ModelName    *C_TagWithAttrAndValue `xml:"model-name,omitempty"`
}

func NewComponent() *Component {
	a := new(Component)
	return a
}
