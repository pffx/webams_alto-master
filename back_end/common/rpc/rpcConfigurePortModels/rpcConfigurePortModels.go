package rpcConfigurePortModels

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

// Demo    <statistics xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-interfaces-statistics">
//
//	  <enable>best-effort</enable>
//	</statistics>
type C_TagWithAttrAndEnable struct {
	Xmlns  string `xml:"xmlns,attr"`
	Enable string `xml:"enable"`
}

func NewC_TagWithAttrAndEnable() *C_TagWithAttrAndEnable {
	a := new(C_TagWithAttrAndEnable)
	return a
}

// ********************************Common struct part end****************************************
// ****************************************************************************************
// configure/port
type Port struct {
	PortId      string    `xml:"port-id,omitempty" json:"PortId,omitempty"`
	Name        string    `xml:"name,omitempty" json:"Name,omitempty"`
	AdminState  string    `xml:"admin-state,omitempty"`
	Description string    `xml:"description,omitempty"`
	Ethernet    *Ethernet `xml:"ethernet,omitempty"`
}

func NewPort() *Port {
	a := new(Port)
	return a
}

// configure/lag
type Lag struct {
	LagIndex   string     `xml:"lag-index,omitempty"`
	AdminState string     `xml:"admin-state,omitempty"`
	LagPort    []LagPort  `xml:"port,omitempty"`
	SubGroup   []SubGroup `xml:"sub-group,omitempty"`
}

func NewLag() *Lag {
	a := new(Lag)
	return a
}

// configure/lag/port
type LagPort struct {
	PortId   string `xml:"port-id,omitempty"`
	SubGroup string `xml:"sub-group,omitempty"`
}

func NewLagPort() *LagPort {
	a := new(LagPort)
	return a
}

// configure/lag/sub-group
type SubGroup struct {
	SubGroupId string `xml:"sub-group-id,omitempty"`
	Preference string `xml:"preference,omitempty"`
}

func NewSubGroup() *SubGroup {
	a := new(SubGroup)
	return a
}

// configure/service
type Service struct {
	Vpls []Vpls `xml:"vpls,omitempty"`
}

func NewService() *Service {
	a := new(Service)
	return a
}

// configure/service/vpls
type Vpls struct {
	ServiceName string `xml:"service-name,omitempty"`
	AdminState  string `xml:"admin-state,omitempty"`
	Description string `xml:"description,omitempty"`
	ServiceId   string `xml:"service-id,omitempty"`
	Customer    string `xml:"customer,omitempty"`
	VVpls       string `xml:"v-vpls,omitempty"`
	Vlan        string `xml:"vlan,omitempty"`
	UserUserCom string `xml:"user-user-com,omitempty"`
	Fdb         *Fdb   `xml:"fdb,omitempty"`
	Sap         []Sap  `xml:"sap,omitempty"`
}

func NewVpls() *Vpls {
	a := new(Vpls)
	return a
}

// configure/service/vpls/fdb
// configure/service/vpls/sap/fdb
type Fdb struct {
	MacLearning *MacLearning `xml:"mac-learning,omitempty"`
}

func NewFdb() *Fdb {
	a := new(Fdb)
	return a
}

// configure/service/vpls/sap
type Sap struct {
	SapId      string `xml:"sap-id,omitempty"`
	AdminState string `xml:"admin-state,omitempty"`
	Fdb        *Fdb   `xml:"fdb,omitempty"`
}

func NewSap() *Sap {
	a := new(Sap)
	return a
}

// configure/service/vpls/fdb/
// configure/service/vpls/sap/fdb/mac-learning
type MacLearning struct {
	Learning bool `xml:"learning,omitempty"`
}

func NewMacLearning() *MacLearning {
	a := new(MacLearning)
	return a
}

// ********************************Common part end****************************************
//****************************************************************************************

type Ethernet struct {
	Autonegotiate     string `xml:"autonegotiate,omitempty"`
	Dot1qEtype        string `xml:"dot1q-etype,omitempty"`
	Mode              string `xml:"mode,omitempty"`
	EncapType         string `xml:"encap-type,omitempty"`
	Category          string `xml:"category,omitempty"`
	Remark            string `xml:"remark,omitempty"`
	UseVlanDot1qEtype string `xml:"use-vlan-dot1q-etype,omitempty"`
	Speed             string `xml:"speed,omitempty"`
}
