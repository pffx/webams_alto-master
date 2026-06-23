package rpcOnusModels

import (
	"alto_server/constants"
	"strconv"
)

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
//****************************************************************************************

// onus/onu
type Onu struct {
	XmlnsXc     string `xml:"xmlns:xc,attr"`
	Operation   string `xml:"xc:operation,omitempty,attr"` // may be should be added automaticlly
	Name        string `xml:"name"`
	Description string `xml:"description"`
	// Usage         interface{}    `xml:",innerxml"` // can not insert namespace for sigle tag, so use innerxml instead of it
	Usage *struct {
		Name          string `xml:",chardata"`
		XmlnsTemplate string `xml:"xmlns:template-common,attr"`
	} `xml:"usage,omitempty"`
	OnuManagement *OnuManagement `xml:"onu-management,omitempty"`
	Root          *Root          `xml:"root,omitempty"`
}

func NewONU(action string) *Onu {
	a := new(Onu)
	a.XmlnsXc = constants.XMLNS_VERSION
	if action == constants.OPERATION_DELETE {
		a.Operation = constants.OPERATION_DELETE
	}
	if action != constants.OPERATION_DELETE {
		a.Usage = &struct {
			Name          string `xml:",chardata"`
			XmlnsTemplate string `xml:"xmlns:template-common,attr"`
		}{
			XmlnsTemplate: `http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-template-common`,
			Name:          "template-common:node-actual-usage",
		}

	}
	//a.OnuManagement = &OnuManagement{TargetSoftwareControl: "immediate-activation"}
	// a.Root = NewRoot()

	// fmt.Println(a)
	return a
}
func NewONU2(action string) *Onu {
	a := new(Onu)
	a.XmlnsXc = constants.XMLNS_VERSION
	if action == constants.OPERATION_DELETE {
		a.Operation = constants.OPERATION_DELETE
	}
	if action != constants.OPERATION_DELETE {
		a.Usage = &struct {
			Name          string `xml:",chardata"`
			XmlnsTemplate string `xml:"xmlns:template-common,attr"`
		}{
			XmlnsTemplate: `http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-template-common`,
			Name:          "template-common:node-actual-usage",
		}

	}
	return a
}
func NewGetONU() *Onu {
	a := new(Onu)
	a.XmlnsXc = constants.XMLNS_VERSION
	// fmt.Println(a)
	return a
}

// onus/onu/onu-management
type OnuManagement struct {
	TargetActiveSoftware  *TargetActiveSoftware `xml:"target-active-software,omitempty"`
	TargetSoftwareControl string                `xml:"target-software-control,omitempty"`
}

func NewOnuManagement() *OnuManagement {
	a := new(OnuManagement)
	return a
}

// onus/onu/onu-management/target-active-software
type TargetActiveSoftware struct {
	Release string `xml:"release,omitempty"`
}

func NewTargetActiveSoftware() *TargetActiveSoftware {
	a := new(TargetActiveSoftware)
	return a
}

// onus/onu/root
type Root struct {
	Classifiers      *Classifiers      `xml:"classifiers"`
	Policies         *Policies         `xml:"policies"`
	TmProfiles       *TmProfiles       `xml:"tm-profiles"`
	Hardware         *Hardware         `xml:"hardware"`
	HardwareStateOnu *HardwareStateOnu `xml:"hardware-state"`
	InterfacesOnu    *InterfacesOnu    `xml:"interfaces"`
	InterfacesState  *InterfacesState  `xml:"interfaces-state"`
}

func NewRoot() *Root {
	a := new(Root)
	return a
}

// onus/onu/root/classifiers
type Classifiers struct {
	Xmlns           string            `xml:"xmlns,attr"`
	ClassifierEntry []ClassifierEntry `xml:"classifier-entry"`
}

func NewClassifiers() *Classifiers {
	a := new(Classifiers)
	a.Xmlns = "urn:bbf:yang:bbf-qos-classifiers-mounted"
	return a
}

// onus/onu/root/policies
type Policies struct {
	Xmlns  string   `xml:"xmlns,attr"`
	Policy []Policy `xml:"policy"`
}

func NewPolicies() *Policies {
	a := new(Policies)
	a.Xmlns = "urn:bbf:yang:bbf-qos-policies-mounted"
	return a
}

// onus/onu/root/tm-profiles
type TmProfiles struct {
	Xmlns                      string                       `xml:"xmlns,attr,omitempty"`
	TcId2QueueIdMappingProfile []TcId2QueueIdMappingProfile `xml:"tc-id-2-queue-id-mapping-profile"`
}

func NewTmProfiles() *TmProfiles {
	a := new(TmProfiles)
	a.Xmlns = "urn:bbf:yang:bbf-qos-traffic-mngt-mounted"
	return a
}

// onus/onu/root/hardware
type Hardware struct {
	Xmlns     string      `xml:"xmlns,attr"`
	Component []Component `xml:"component"`
}

func NewHardware() *Hardware {
	a := new(Hardware)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-hardware-mounted"
	return a
}

// onus/onu/root/hardware-state
type HardwareStateOnu struct {
	Xmlns     string      `xml:"xmlns,attr"`
	Component []Component `xml:"component"`
}

func NewHardwareStateOnu() *HardwareStateOnu {
	a := new(HardwareStateOnu)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-hardware-mounted"
	return a
}

// onus/onu/root/interfaces-state
type InterfacesState struct {
	Xmlns     string      `xml:"xmlns,attr"`
	Interface []Interface `xml:"interface"`
}

func NewInterfacesState() *InterfacesState {
	a := new(InterfacesState)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-interfaces-mounted"
	return a
}

// onus/onu/root/interfaces-state/interface
type Interface struct {
	Name string `xml:"name"`
	Type *struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	} `xml:"type,omitempty"`
	AdminStatus string       `xml:"admin-status,omitempty"`
	OperStatus  string       `xml:"oper-status,omitempty"`
	PhysAddress string       `xml:"phys-address,omitempty"`
	Speed       string       `xml:"speed,omitempty"`
	Ethernet    *Ethernet    `xml:"Ethernet,omitempty"`
	Performance *Performance `xml:"performance,omitempty"`
}

func NewInterface() *Interface {
	a := new(Interface)
	return a
}

// onus/onu/root/interfaces-state/interfaces/Ethernet
type Ethernet struct {
	Duplex string `xml:"duplex,omitempty"`
	Speed  string `xml:"speed,omitempty"`
	Fec    *struct {
		Name  string `xml:",chardata"`
		Xmlns string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	} `xml:"fec,omitempty"`
}

// onus/onu/root/interfaces-state/interfaces/performance
type Performance struct {
	Intervals15min *Intervals15min `xml:"intervals-15min,omitempty"`
}

func NewPerformance() *Performance {
	a := new(Performance)
	return a
}

// onus/onu/root/interfaces-state/interfaces/performance/intervals-15min
type Intervals15min struct {
	Current *PerformanceStatics `xml:"current,omitempty"`
	History *PerformanceStatics `xml:"history,omitempty"`
}

func NewIntervals15min() *Intervals15min {
	a := new(Intervals15min)
	return a
}

// onus/onu/root/interfaces-state/interfaces/performance/intervals-15min/current
// onus/onu/root/interfaces-state/interfaces/performance/intervals-15min/history
type PerformanceStatics struct {
	InOctets     string `xml:"in-octets,omitempty"`
	OutOctets    string `xml:"out-octets,omitempty"`
	InTotalPkts  string `xml:"in-total-pkts,omitempty"`
	OutTotalPkts string `xml:"out-total-pkts,omitempty"`
	Xpon         *Xpon  `xml:"xpon,omitempty"`
}

func NewPerformanceStatics() *PerformanceStatics {
	a := new(PerformanceStatics)
	return a
}

// onus/onu/root/interfaces-state/interfaces/performance/intervals-15min/current/xpon
type Xpon struct {
	Phy *Phy `xml:"phy,omitempty"`
}

func NewXpon() *Xpon {
	a := new(Xpon)
	return a
}

// onus/onu/root/interfaces-state/interfaces/performance/intervals-15min/current/xpon/phy
type Phy struct {
	InBipErrors string `xml:"in-bip-errors,omitempty"`
}

func NewPhy() *Phy {
	a := new(Phy)
	return a
}

// onus/onu/root/interfaces
type InterfacesOnu struct {
	Xmlns            string            `xml:"xmlns,attr"`
	InterfaceOnu     []InterfaceOnu    `xml:"interface"`
	UniLoopDetection *UniLoopDetection `xml:"uni-loop-detection,omitempty"`
}

func NewInterfacesOnu() *InterfacesOnu {
	a := new(InterfacesOnu)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-interfaces-mounted"
	return a
}

// onus/onu/root/classifiers/classifier-entry
type ClassifierEntry struct {
	Name                     string                   `xml:"name"`
	FilterOperation          string                   `xml:"filter-operation"`
	MatchCriteria            MatchCriteria            `xml:"match-criteria"`
	ClassifierActionEntryCfg ClassifierActionEntryCfg `xml:"classifier-action-entry-cfg"`
}

func NewClassifierEntry() *ClassifierEntry {
	a := new(ClassifierEntry)
	a.FilterOperation = "match-all-filter"
	return a
}
func NewTagClassifierEntry() *ClassifierEntry {
	a := new(ClassifierEntry)
	a.FilterOperation = "match-all-filter"
	a.MatchCriteria = *NewTagMatchCriteria()
	return a
}
func NewUntagClassifierEntry() *ClassifierEntry {
	a := new(ClassifierEntry)
	a.FilterOperation = "match-all-filter"
	a.MatchCriteria = *NewUntagMatchCriteria()
	return a
}

// onus/onu/root/classifiers/classifier-entry/match-criteria
type MatchCriteria struct {
	Tag *struct {
		Index      string `xml:"index"`
		InPbitList string `xml:"in-pbit-list"`
	} `xml:"tag,omitempty"`
	Untagged *struct {
		Name string `xml:",chardata"`
	} `xml:"untagged,omitempty"`
	DscpRange       string           `xml:"dscp-range,omitempty"`
	PbitMarkingList *PbitMarkingList `xml:"pbit-marking-list,omitempty"`
}

func NewMatchCriteria() *MatchCriteria {
	a := new(MatchCriteria)
	return a
}
func NewTagMatchCriteria() *MatchCriteria {
	a := new(MatchCriteria)
	a.Tag = &struct {
		Index      string `xml:"index"`
		InPbitList string `xml:"in-pbit-list"`
	}{
		Index:      "0",
		InPbitList: "0-7",
	}
	return a
}
func NewUntagMatchCriteria() *MatchCriteria {
	a := new(MatchCriteria)
	a.Untagged = &struct {
		Name string `xml:",chardata"`
	}{
		Name: "",
	}
	return a
}
func NewDhcpMatchCriteria() *MatchCriteria {
	a := new(MatchCriteria)
	a.DscpRange = "any"
	a.PbitMarkingList = NewPbitMarkingList()
	return a
}

// onus/onu/root/classifiers/classifier-entry/classifier-action-entry-cfg
type ClassifierActionEntryCfg struct {
	ActionType             string `xml:"action-type"`
	SchedulingTrafficClass string `xml:"scheduling-traffic-class,omitempty"`
	PbitMarkingCfg         *struct {
		PbitMarkingList PbitMarkingList `xml:"pbit-marking-list"`
	} `xml:"pbit-marking-cfg,omitempty"`
}

func NewClassifierActionEntryCfg() *ClassifierActionEntryCfg {
	a := new(ClassifierActionEntryCfg)
	return a
}
func NewPbitClassifierActionEntryCfg() *ClassifierActionEntryCfg {
	a := new(ClassifierActionEntryCfg)
	a.ActionType = "pbit-marking"
	a.PbitMarkingCfg = &struct {
		PbitMarkingList PbitMarkingList `xml:"pbit-marking-list"`
	}{
		PbitMarkingList: *NewPbitMarkingList(),
	}
	return a
}
func NewSchedulingClassifierActionEntryCfg() *ClassifierActionEntryCfg {
	a := new(ClassifierActionEntryCfg)
	a.ActionType = "scheduling-traffic-class"
	a.SchedulingTrafficClass = "0"
	return a
}

// onus/onu/root/classifiers/classifier-entry/classifier-action-entry-cfg/pbit-marking-list
// onus/onu/root/classifiers/classifier-entry/match-criteria/pbit-marking-list
type PbitMarkingList struct {
	Xmlns     string `xml:"xmlns,attr,omitempty"`
	Index     string `xml:"index"`
	PbitValue string `xml:"pbit-value"`
}

func NewPbitMarkingList() *PbitMarkingList {
	a := new(PbitMarkingList)
	a.Index = "0"
	a.PbitValue = "0"
	return a
}

// onus/onu/root/policies/policy
type Policy struct {
	Name        string `xml:"name"`
	Classifiers []struct {
		Name string `xml:"name"`
	} `xml:"classifiers"`
}

func NewPolicy() *Policy {
	a := new(Policy)
	return a
}
func NewPolicyWithClassifiers(classifilesName []string) *Policy {
	a := new(Policy)
	for _, name := range classifilesName {
		a.Classifiers = append(a.Classifiers, struct {
			Name string `xml:"name"`
		}{
			Name: name,
		})
	}
	return a
}

// onus/onu/root/tm-profiles/tc-id-2-queue-id-mapping-profile
type TcId2QueueIdMappingProfile struct {
	Name         string         `xml:"name"`
	MappingEntry []MappingEntry `xml:"mapping-entry"`
}

func NewTcId2QueueIdMappingProfile() *TcId2QueueIdMappingProfile {
	a := new(TcId2QueueIdMappingProfile)
	return a
}

// onus/onu/root/tm-profiles/tc-id-2-queue-id-mapping-profile/mapping-entry
type MappingEntry struct {
	TrafficClassId string `xml:"traffic-class-id"`
	LocalQueueId   string `xml:"local-queue-id"`
}

func NewMappingEntry() *MappingEntry {
	a := new(MappingEntry)
	return a
}

type ComponentClassStruct struct {
	Name        string `xml:",chardata"`
	XmlnsHwt    string `xml:"xmlns:bbf-hwt,attr,omitempty"`
	XmlnsHwi    string `xml:"xmlns:nokia-hwi,attr,omitempty"`
	XmlnsIanahw string `xml:"xmlns:ianahw,attr,omitempty"`
}

func NewComponentClassStruct() *ComponentClassStruct {
	a := new(ComponentClassStruct)
	return a
}

// onus/onu/root/hardware-state/component
// onus/onu/root/hardware/component
type Component struct {
	//???????
	Name string `xml:"name,omitempty"`
	// Class        interface{} `xml:",innerxml"` // can not insert namespace for sigle tag, so use innerxml instead of it
	Class           *ComponentClassStruct `xml:"class,omitempty"`
	Parent          string                `xml:"parent,omitempty"`
	ModelName       string                `xml:"model-name,omitempty"`
	EquipmentId     string                `xml:"equipment-id,omitempty"`
	SerialNum       string                `xml:"serial-num,omitempty"`
	MacAddress      string                `xml:"mac-address,omitempty"`
	ParentRelPos    string                `xml:"parent-rel-pos,omitempty"`
	AdminState      string                `xml:"admin-state,omitempty"`
	SoftwaresOnu    *SoftwaresOnu         `xml:"software,omitempty"`
	TransceiverLink *TransceiverLink      `xml:"transceiver-link,omitempty"`
	Poe             *Poe                  `xml:"poe,omitempty"`
}

func NewComponent() *Component {
	a := new(Component)
	a.Class = NewComponentClassStruct()
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	return a
}
func NewComponent2() *Component {
	a := new(Component)
	return a
}
func NewANIComponent() *Component {
	a := new(Component)
	a.Name = "ANIPORT"
	a.Class = NewComponentClassStruct()
	a.Class.Name = "bbf-hwt:transceiver-link"
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	a.Parent = "SFP"
	a.ParentRelPos = "1"
	return a
}
func NewCageComponent() *Component {
	a := new(Component)
	a.Name = "CAGE"
	a.Class = NewComponentClassStruct()
	a.Class.Name = "bbf-hwt:cage"
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	a.Parent = "CHASSIS"
	a.ParentRelPos = "0"
	return a
}
func NewCardComponent(index string) *Component {
	a := new(Component)
	a.Name = "CARD_" + index
	a.Class = NewComponentClassStruct()
	a.Class.Name = "bbf-hwt:board"
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	a.Parent = "CHASSIS"
	a.ParentRelPos = index
	a.AdminState = "unlocked"
	return a
}

func NewRJ451GComponent(index string) *Component {
	a := new(Component)
	a.Name = constants.UNI_LAN + index
	a.Class = NewComponentClassStruct()
	a.Class.Name = "nokia-hwi:rj45-1G"
	a.Class.XmlnsHwi = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-identities"
	a.Parent = "CARD_1"
	a.ParentRelPos = index
	// a.AdminState = "unlocked"
	return a
}

func NewChassisComponent() *Component {
	a := new(Component)
	a.Name = "CHASSIS"
	a.Class = NewComponentClassStruct()
	a.Class.Name = "ianahw:chassis"
	a.Class.XmlnsIanahw = "urn:ietf:params:xml:ns:yang:iana-hardware"
	a.AdminState = "unlocked"
	return a
}
func NewSFPComponent() *Component {
	a := new(Component)
	a.Name = "SFP"
	a.Class = NewComponentClassStruct()
	a.Class.Name = "bbf-hwt:transceiver"
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	a.Parent = "CAGE"
	a.ParentRelPos = "0"
	return a
}
func NewUniLanComponent(index string) *Component {
	a := new(Component)
	a.Name = constants.UNI_LAN + index
	a.Class = NewComponentClassStruct()
	a.Class.Name = "bbf-hwt:transceiver-link"
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	// a.Parent = "CHASSIS" from 2309, olt can config it itself, but use CARD 1 now
	a.Parent = "CARD_1"
	a.ParentRelPos = index
	return a
}
func NewUniTelComponent(index string) *Component {
	a := new(Component)
	a.Name = "UNI_TEL" + index
	a.Class = NewComponentClassStruct()
	a.Class.Name = "bbf-hwt:transceiver-link"
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	a.Parent = "CARD_2"
	a.ParentRelPos = index
	return a
}
func NewVeipComponent() *Component {
	a := new(Component)
	a.Name = "VEIP"
	a.Class = NewComponentClassStruct()
	a.Class.Name = "nokia-hwi:virtual-port"
	a.Class.XmlnsHwi = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-identities"
	// a.Parent = "CARD_14" from 2309, olt can config it itself
	a.Parent = "CHASSIS"
	a.ParentRelPos = "1"
	return a
}
func NewUni10GComponent() *Component {
	a := new(Component)
	a.Name = "UNI_10G"
	a.Class = NewComponentClassStruct()
	a.Class.Name = "bbf-hwt:transceiver-link"
	a.Class.XmlnsHwt = "urn:bbf:yang:bbf-hardware-types"
	a.Parent = "CARD_10"
	a.ParentRelPos = "1"
	return a
}

func NewUni25GComponent() *Component {
	a := new(Component)
	a.Name = "UNI_25G"
	a.Class = NewComponentClassStruct()
	a.Class.Name = "nokia-hwi:rj45-1G"
	a.Class.XmlnsHwi = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-identities"
	a.Parent = "CHASSIS"
	a.ParentRelPos = "1"
	return a
}

// onus/onu/root/hardware-state/component/transceiver-link
type TransceiverLink struct {
	Diagnostics *Diagnostics `xml:"diagnostics"`
}

// onus/onu/root/hardware-state/component/transceiver-link/diagnostics
type Diagnostics struct {
	TxBias           string `xml:"tx-bias,omitempty"`
	TxPower          string `xml:"tx-power,omitempty"`
	RxPower          string `xml:"rx-power,omitempty"`
	LaserTemperature string `xml:"laser-temperature,omitempty"`
	TxPowerDbm       string `xml:"tx-power-dbm,omitempty"`
	RxPowerDbm       string `xml:"rx-power-dbm,omitempty"`
}

// onus/onu/root/hardware-state/component/poe
type Poe struct {
	Xmlns           string `xml:"xmlns,attr,omitempty"`
	Enable          string `xml:"enable,omitempty"`
	PseClassControl string `xml:"pse-class-control,omitempty"`
	PowerPriority   string `xml:"power-priority,omitempty"`
}

func NewPoe() *Poe {
	a := new(Poe)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-power-over-ethernet-control-mounted"
	return a
}
func NewHighPoe() *Poe {
	a := new(Poe)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-power-over-ethernet-control-mounted"
	a.Enable = "true"
	a.PowerPriority = "high"
	a.PseClassControl = "0"
	return a
}

// onus/onu/root/hardware-state/component/software
type SoftwaresOnu struct {
	SoftwareOnu []SoftwareOnu `xml:"software"`
}

// onus/onu/root/hardware-state/component/software/software
type SoftwareOnu struct {
	Name      string    `xml:"name"`
	Revisions Revisions `xml:"revisions"`
}

// onus/onu/root/hardware-state/component/software/software/revisions
type Revisions struct {
	Revision []Revision `xml:"revision"`
}

// onus/onu/root/hardware-state/component/software/software/revisions/revision
type Revision struct {
	Name        string `xml:"name"`
	Version     string `xml:"version"`
	IsValid     string `xml:"is-valid"`
	IsCommitted string `xml:"is-committed"`
	IsActive    string `xml:"is-active"`
}

// onus/onu/root/interfaces/interface
type InterfaceOnu struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Type        *struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	} `xml:"type,omitempty"`
	Enabled      string                 `xml:"enabled,omitempty"`
	PortLayerIf  *C_TagWithAttrAndValue `xml:"port-layer-if,omitempty"`
	PhysVoiceItf *struct {
		Xmlns       string `xml:"xmlns,attr"`
		PortLayerIf string `xml:"port-layer-if"`
	} `xml:"phys-voice-itf,omitempty"`
	Performance *C_TagWithAttrAndEnable `xml:"performance"`
	OnuVEnet    *struct {
		Ani   string `xml:"ani"`
		Xmlns string `xml:"xmlns,attr,omitempty"`
	} `xml:"onu-v-enet,omitempty"`
	OnuVVrefpoint *struct {
		Xmlns      string `xml:"xmlns,attr"`
		RelatedOnu string `xml:"related-onu"`
	} `xml:"onu-v-vrefpoint"`
	Pae *Pae `xml:"pae,omitempty"`
}

func NewInterfaceOnu() *InterfaceOnu {
	a := new(InterfaceOnu)
	return a
}
func NewANIInterfaceOnu() *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = "ANI"
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:     "bbf-xponift-mounted:ani",
		XmlnsPon: "urn:bbf:yang:bbf-xpon-if-type-mounted",
	}
	a.Enabled = "true"
	a.PortLayerIf = &C_TagWithAttrAndValue{
		Name:  "ANIPORT",
		Xmlns: "urn:bbf:yang:bbf-interface-port-reference-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	return a
}
func NewUniLanInterfaceOnu(index string) *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = constants.UNI_LAN + index
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:  "ianaift-mounted:ethernetCsmacd",
		Xmlns: "urn:ietf:params:xml:ns:yang:iana-if-type-mounted",
	}
	a.Enabled = "false"
	a.PortLayerIf = &C_TagWithAttrAndValue{
		Name:  constants.UNI_LAN + index,
		Xmlns: "urn:bbf:yang:bbf-interface-port-reference-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	return a
}
func NewVoipInterfaceOnu() *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = "UNI_VOIP"
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:     "bbf-xponift-mounted:onu-v-enet",
		XmlnsPon: "urn:bbf:yang:bbf-xpon-if-type-mounted",
	}
	a.Enabled = "true"
	a.OnuVEnet = &struct {
		Ani   string `xml:"ani"`
		Xmlns string `xml:"xmlns,attr,omitempty"`
	}{
		Ani:   "ANI",
		Xmlns: "urn:bbf:yang:bbf-xponani-mounted",
	}
	return a
}
func NewUniTelInterfaceXGPonOnu(index string) *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = "UNI_TEL" + index
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:     "ianaift-mounted:voiceFXS",
		XmlnsPon: "urn:bbf:yang:bbf-xpon-if-type-mounted",
	}
	a.Enabled = "true"
	a.PhysVoiceItf = &struct {
		Xmlns       string `xml:"xmlns,attr"`
		PortLayerIf string `xml:"port-layer-if"`
	}{
		PortLayerIf: "UNI_TEL" + index,
		Xmlns:       "urn:bbf:yang:bbf-sip-voip-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	return a
}
func NewUniTelInterfacePonOnu(index string) *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = "UNI_TEL" + index
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:  "ianaift-mounted:voiceFXS",
		Xmlns: "urn:ietf:params:xml:ns:yang:iana-if-type-mounted",
	}
	a.Enabled = "true"
	a.PhysVoiceItf = &struct {
		Xmlns       string `xml:"xmlns,attr"`
		PortLayerIf string `xml:"port-layer-if"`
	}{
		PortLayerIf: "UNI_TEL" + index,
		Xmlns:       "urn:bbf:yang:bbf-sip-voip-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	return a
}
func NewUni10GIndexInterfaceOnu(index int) *InterfaceOnu {
	suffix := strconv.Itoa(index)
	a := new(InterfaceOnu)
	if index <= 1 {
		suffix = ""
	}
	a.Name = "UNI_10G" + suffix
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:  "ianaift-mounted:ethernetCsmacd",
		Xmlns: "urn:ietf:params:xml:ns:yang:iana-if-type-mounted",
	}
	a.Enabled = "false"
	a.PortLayerIf = &C_TagWithAttrAndValue{
		Name:  "UNI_10G" + suffix,
		Xmlns: "urn:bbf:yang:bbf-interface-port-reference-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	return a
}
func NewUni10GInterfaceOnu() *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = "UNI_10G"
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:  "ianaift-mounted:ethernetCsmacd",
		Xmlns: "urn:ietf:params:xml:ns:yang:iana-if-type-mounted",
	}
	a.Enabled = "false"
	a.PortLayerIf = &C_TagWithAttrAndValue{
		Name:  "UNI_10G",
		Xmlns: "urn:bbf:yang:bbf-interface-port-reference-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	return a
}
func NewUni25GInterfaceOnu() *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = "UNI_25G"
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:  "ianaift-mounted:ethernetCsmacd",
		Xmlns: "urn:ietf:params:xml:ns:yang:iana-if-type-mounted",
	}
	a.Enabled = "false"
	a.PortLayerIf = &C_TagWithAttrAndValue{
		Name:  "UNI_25G",
		Xmlns: "urn:bbf:yang:bbf-interface-port-reference-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	return a
}
func NewVeipInterfaceOnu() *InterfaceOnu {
	a := new(InterfaceOnu)
	a.Name = "NNI_VEIP"
	a.Type = &struct {
		Name     string `xml:",chardata"`
		XmlnsPon string `xml:"xmlns:bbf-xponift-mounted,attr,omitempty"`
		Xmlns    string `xml:"xmlns:ianaift-mounted,attr,omitempty"`
	}{
		Name:     "bbf-xponift-mounted:onu-v-vrefpoint",
		XmlnsPon: "urn:bbf:yang:bbf-xpon-if-type-mounted",
	}
	a.Enabled = "false"
	a.PortLayerIf = &C_TagWithAttrAndValue{
		Name:  "VEIP",
		Xmlns: "urn:bbf:yang:bbf-interface-port-reference-mounted",
	}
	a.Performance = &C_TagWithAttrAndEnable{
		Enable: "false",
		Xmlns:  "urn:bbf:yang:bbf-interfaces-performance-management-mounted",
	}
	a.OnuVVrefpoint = &struct {
		Xmlns      string `xml:"xmlns,attr"`
		RelatedOnu string `xml:"related-onu"`
	}{
		RelatedOnu: "ANI",
		Xmlns:      "urn:bbf:yang:bbf-xponani-mounted",
	}
	return a
}

// onus/onu/root/interfaces/uni-loop-detection
type UniLoopDetection struct {
	Xmlns       string `xml:"xmlns,attr,omitempty"`
	Enable      string `xml:"enable,omitempty"`
	AutoShutOff string `xml:"auto-shut-off,omitempty"`
}

func NewUniLoopDetection() *UniLoopDetection {
	a := new(UniLoopDetection)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-uni-loop-detection-mounted"
	return a
}

// onus/onu/root/interfaces/interface/pae
type Pae struct {
	Xmlns            string            `xml:"xmlns,attr,omitempty"`
	Operation        string            `xml:"xc:operation,omitempty,attr"`
	PortCapabilities *PortCapabilities `xml:"port-capabilities,omitempty"`
}

func NewPae() *Pae {
	a := new(Pae)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-802dot1x-mounted"
	return a
}

// onus/onu/root/interfaces/interface/pae/port-capabilities
type PortCapabilities struct {
	Auth string `xml:"auth,omitempty"`
}

func NewPortCapabilities() *PortCapabilities {
	a := new(PortCapabilities)
	return a
}

// ********************************ONUS part end****************************************
//****************************************************************************************
