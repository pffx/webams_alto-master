package rpcForwardingModels

import "alto_server/constants"

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

//forwarding/forwarding-databases
type ForwardingDatabases struct {
	ForwardingDatabase []ForwardingDatabase `xml:"forwarding-database,omitempty"`
}

//forwarding/forwarders
type Forwarders struct {
	Forwarder []Forwarder `xml:"forwarder,omitempty"`
}

func NewForwarders() *Forwarders {
	a := new(Forwarders)
	return a
}

//forwarding/forwarders/forwarder
type Forwarder struct {
	Name string `xml:"name,omitempty"`
	//todo: change to pointer
	Ports Ports `xml:"ports,omitempty"`
}

func NewForwarder() *Forwarder {
	a := new(Forwarder)
	return a
}

//forwarding/forwarders/forwarder/ports
type Ports struct {
	Port []Port `xml:"port,omitempty"`
}

func NewPorts() *Ports {
	a := new(Ports)
	return a
}

//forwarding/forwarders/forwarder/ports/port
type Port struct {
	Operation    string `xml:"operation,omitempty,attr"`
	Name         string `xml:"name,omitempty"`
	SubInterface string `xml:"sub-interface,omitempty"`
}

func NewPort() *Port {
	a := new(Port)
	return a
}
func NewPortDel() *Port {
	a := new(Port)
	a.Operation = constants.OPERATION_DELETE
	return a
}

func NewForwardingDatabases() *ForwardingDatabases {
	a := new(ForwardingDatabases)
	return a
}

//forwarding/forwarding-databases/forwarding-database
type ForwardingDatabase struct {
	Name                  string              `xml:"name,omitempty"`
	AgingTimer            string              `xml:"aging-timer,omitempty"`
	MaxNumberMacAddresses string              `xml:"max-number-mac-addresses,omitempty"`
	MacLearningControl    *MacLearningControl `xml:"mac-learning-control,omitempty"`
	StaticMacAddress      []StaticMacAddress  `xml:"static-mac-address,omitempty"`
}

func NewForwardingDatabase() *ForwardingDatabase {
	a := new(ForwardingDatabase)
	return a
}

//forwarding/forwarding-databases/forwarding-database/static-mac-address
type StaticMacAddress struct {
	MacAddress             string                  `xml:"mac-address,omitempty"`
	StaticForwarderPortRef *StaticForwarderPortRef `xml:"static-forwarder-port-ref,omitempty"`
}

func NewStaticMacAddress() *StaticMacAddress {
	a := new(StaticMacAddress)
	return a
}

//forwarding/forwarding-databases/forwarding-database/static-mac-address/static-forwarder-port-ref
type StaticForwarderPortRef struct {
	Forwarder string `xml:"forwarder,omitempty"`
	Port      string `xml:"port,omitempty"`
}

func NewStaticForwarderPortRef() *StaticForwarderPortRef {
	a := new(StaticForwarderPortRef)
	return a
}

//forwarding/forwarding-databases/forwarding-database/mac-learning-control
type MacLearningControl struct {
	GenerateMacLearningAlarm  string `xml:"generate-mac-learning-alarm,omitempty"`
	MacLearningControlProfile string `xml:"mac-learning-control-profile,omitempty"`
}

func NewMacLearningControl(action string) *MacLearningControl {
	a := new(MacLearningControl)
	return a
}

//forwarding/mac-learning-control-profiles
type MacLearningControlProfiles struct {
	MacLearningControlProfile []MacLearningControlProfile `xml:"mac-learning-control-profile,omitempty"`
}

func NewMacLearningControlProfiles(action string) *MacLearningControlProfiles {
	a := new(MacLearningControlProfiles)
	return a
}

//forwarding/mac-learning-control-profiles/mac-learning-control-profile
type MacLearningControlProfile struct {
	Name            string            `xml:"name,omitempty"`
	MacLearningRule []MacLearningRule `xml:"mac-learning-rule,omitempty"`
}

func NewMacLearningControlProfile(action string) *MacLearningControlProfile {
	a := new(MacLearningControlProfile)
	return a
}

//forwarding/mac-learning-control-profiles/mac-learning-control-profile/mac-learning-rule
type MacLearningRule struct {
	ReceivingInterfaceUsage string   `xml:"receiving-interface-usage,omitempty"`
	MacCanNotMoveTo         []string `xml:"mac-can-not-move-to,omitempty"`
}

func NewMacLearningRule(action string) *MacLearningRule {
	a := new(MacLearningRule)
	return a
}

// ********************************forwarding part end****************************************
//****************************************************************************************
