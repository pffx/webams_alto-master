package rpcForwardingStateModels

// ********************************Common struct part start****************************************
//****************************************************************************************

//Demo   <onu-name xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-if-xponvani-aug">C12_ONT2</onu-name>

type C_TagWithAttrAndValue struct {
	Name            string `xml:",chardata"`
	Xmlns           string `xml:"xmlns,attr,omitempty"`
	XmlnsBbfXPonift string `xml:"xmlns:bbf-xponift,attr,omitempty"`
}

func NewC_TagWithAttrAndValue() *C_TagWithAttrAndValue {
	a := new(C_TagWithAttrAndValue)
	return a
}

//onus/onu/root/interfaces/interface/performance
//interfaces/interface/performance
//interfaces/interface/statistics

//Demo    <statistics xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-interfaces-statistics">
//          <enable>best-effort</enable>
//        </statistics>
type C_TagWithAttrAndEnable struct {
	Xmlns  string `xml:"xmlns,attr"`
	Enable string `xml:"enable"`
}

func NewC_TagWithAttrAndEnable() *C_TagWithAttrAndEnable {
	a := new(C_TagWithAttrAndEnable)
	return a
}

// ********************************Common struct part end****************************************
//****************************************************************************************

//forwarding-state/forwarding-databases
type ForwardingDatabases struct {
	ForwardingDatabase []ForwardingDatabase `xml:"forwarding-database,omitempty"`
}

//forwarding-state/forwarders
type Forwarders struct {
	Forwarder []Forwarder `xml:"forwarder,omitempty"`
}

//forwarding-state/forwarders/forwarder
type Forwarder struct {
	Name  string `xml:"name,omitempty"`
	Ports *Ports `xml:"ports,omitempty"`
}

//forwarding-state/forwarders/forwarder/ports
type Ports struct {
	Port []Port `xml:"port,omitempty"`
}

//forwarding-state/forwarders/forwarder/ports/port
type Port struct {
	Name         string `xml:"name,omitempty"`
	SubInterface string `xml:"sub-interface,omitempty"`
}

func NewForwardingDatabases() *ForwardingDatabases {
	a := new(ForwardingDatabases)
	return a
}

//forwarding-state/forwarding-databases/forwarding-database
type ForwardingDatabase struct {
	Name         string        `xml:"name,omitempty"`
	MacAddresses *MacAddresses `xml:"mac-addresses,omitempty"`
}

func NewForwardingDatabase() *ForwardingDatabase {
	a := new(ForwardingDatabase)
	return a
}

//forwarding-state/forwarding-databases/forwarding-database/mac-addresses
type MacAddresses struct {
	MacAddress []MacAddress `xml:"mac-address,omitempty"`
}

func NewMacAddresses() *MacAddresses {
	a := new(MacAddresses)
	return a
}

//forwarding-state/forwarding-databases/forwarding-database/mac-addresses/mac-address
type MacAddress struct {
	MacAddress   string                 `xml:"mac-address,omitempty"`
	Forwarder    string                 `xml:"forwarder,omitempty"`
	Port         string                 `xml:"port,omitempty"`
	SubInterface *C_TagWithAttrAndValue `xml:"sub-interface,omitempty"`
}

func NewMacAddress() *MacAddress {
	a := new(MacAddress)
	return a
}

// ********************************forwarding-state part end****************************************
//****************************************************************************************
