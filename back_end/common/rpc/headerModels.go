package rpc

// 1. common rpc information stuct
// 2. the root nodes of  netconf yang rpc. the child nodes are defined in rpc**Models folder
//
//

import (
	"alto_server/common/rpc/rpcAlarmsModels"
	"alto_server/common/rpc/rpcConfigurePortModels"
	"alto_server/common/rpc/rpcControlledChannelPairStateModels"
	"alto_server/common/rpc/rpcForwardingModels"
	"alto_server/common/rpc/rpcForwardingStateModels"
	"alto_server/common/rpc/rpcHardwareModels"
	"alto_server/common/rpc/rpcHardwareProtectionElementStateModels"
	"alto_server/common/rpc/rpcHardwareStateModels"
	"alto_server/common/rpc/rpcInterfacesModels"
	"alto_server/common/rpc/rpcInterfacesStateModels"
	"alto_server/common/rpc/rpcOnusModels"
	"alto_server/common/rpc/rpcOperDataFormatCliBlock"
	"alto_server/common/rpc/rpcStatePortModels"
	"alto_server/common/rpc/rpcSyslogModels"
	"alto_server/common/rpc/rpcSystemModels"
	"alto_server/common/rpc/rpcSystemStateModels"
	rpcTmProfiles "alto_server/common/rpc/rpcTmProfilesModels"
	"alto_server/common/rpc/rpcXpongemtcontStateModels"

	"alto_server/constants"
	"encoding/xml"

	netconf "alto_server/common/netconf2"
)

type RpcError struct {
	ErrorType     string `xml:"error-type"`
	ErrorTag      string `xml:"error-tag"`
	ErrorSeverity string `xml:"error-severity"`
	ErrorAppTag   string `xml:"error-app-tag"`
	ErrorMessage  string `xml:"error-message"`
	ErrorInfo     string `xml:"error-info"`
}

// information about rpc connection
type RPCInfor struct {
	Account        string
	Password       string
	IP             string
	Port           string
	Version        string
	TemplatePath   string
	NetconfSession *netconf.Session
}

// information about ont part in rpc
type RPCOntInfor struct {
	OntModel  string //Device type: G-140W-G
	OntSn     string // serial number: ALCL****
	CageIndex string //cage number: C1 or C2
	LTIndex   string //LT number: LT1 or LT2
	OntIndex  string //ONU ID in lightspan
	OnuName   string
	// OntSecIndex        string //ONU ID in lightspan
	ChannelTermination string // Termination of this onu, it can be null in database
	ChannelPair        string // ChannelPair of this onu, it can be null in database
	ChannelPartition   string // ChannelPartition of this onu, it can be null in database
	ChannelGroup       string // ChannelGroup of this onu, it can be null in database
	OntPonType         string // GPON or XGSPON or 25GSPON
	OntRouterMode      string // TRUE or FALSE
	LanNum             int    // hardware capability
	XlanNum            int    // hardware capability
	TelNum             int    // hardware capability
	ServiceId          string // service profile name
	Poe                string
	SfpName            string
}

// used for create edit rpc, should use constructor to create
type EditConfig struct {
	XMLName xml.Name `xml:"edit-config"`
	Target  Target   `xml:"target"`
	Config  Config   `xml:"config"`
	Comment string   `xml:",comment"` // can output a comment in xml file
}

type EditCandidateConfig struct {
	XMLName xml.Name        `xml:"edit-config"`
	Target  TargetCandidate `xml:"target"`
	Config  Config          `xml:"config"`
	Comment string          `xml:",comment"` // can output a comment in xml file
}

type Target struct {
	Running string `xml:"running"`
}
type TargetCandidate struct {
	Candidate string `xml:"candidate"`
}

// used for create get rpc, should wait for the  rpc reply,
type GetConfig struct {
	XMLName xml.Name `xml:"get-config"`
	Source  *Source  `xml:"source,omitempty"`
	Filter  Filter   `xml:"filter"`
	Comment string   `xml:",comment"` // can output a comment in xml file
}
type Get struct {
	XMLName xml.Name  `xml:"get"`
	Xmlns   string    `xml:"xmlns,omitempty,attr"`
	Filter  GetFilter `xml:"filter"`
	Comment string    `xml:",comment"` // can output a comment in xml file
}
type Commit struct {
	XMLName xml.Name `xml:"commit"`
	Xmlns   string   `xml:"xmlns,omitempty,attr"`
}

type DiscardChanges struct {
	XMLName xml.Name `xml:"discard-changes"`
	Xmlns   string   `xml:"xmlns,omitempty,attr"`
}

func NewGet() *Get {
	a := new(Get)
	a.Xmlns = "urn:ietf:params:xml:ns:netconf:base:1.0"
	a.Filter.Type = "subtree"
	return a
}
func NewGetRunningConfig() *GetConfig {
	a := new(GetConfig)
	a.Source = NewSource()
	a.Filter.Type = "subtree"
	return a
}
func NewGetConfig() *GetConfig {
	a := new(GetConfig)
	a.Filter.Type = "subtree"
	return a
}
func NewCommit() *Commit {
	a := new(Commit)
	a.Xmlns = "urn:ietf:params:xml:ns:netconf:base:1.0"
	return a
}
func NewEditConfig() *EditConfig {
	a := new(EditConfig)
	return a
}
func NewCandidateConfig() *EditCandidateConfig {
	a := new(EditCandidateConfig)
	return a
}
func NewDiscardChanges() *DiscardChanges {
	a := new(DiscardChanges)
	a.Xmlns = "urn:ietf:params:xml:ns:netconf:base:1.0"
	return a
}

type Source struct {
	Running string `xml:"running"`
}

func NewSource() *Source {
	a := new(Source)
	return a
}

type GetFilter struct {
	Type            string           `xml:"type,attr,omitempty"`
	InterfacesState *InterfacesState `xml:"interfaces-state,omitempty"`
	Interfaces      *Interfaces      `xml:"interfaces,omitempty"`
	Hardware        *Hardware        `xml:"hardware,omitempty"`
	Onus            *Onus            `xml:"onus,omitempty"`
	Alarms          *Alarms          `xml:"alarms,omitempty"`
	HardwareState   *HardwareState   `xml:"hardware-state,omitempty"`
	SystemState     *SystemState     `xml:"system-state,omitempty"`
	Forwarding      *Forwarding      `xml:"forwarding,omitempty"`
	ForwardingState *ForwardingState `xml:"forwarding-state,omitempty"`
	// used for ihub
	Configure                      *Configure                      `xml:"configure,omitempty"`
	State                          *State                          `xml:"state,omitempty"`
	Xpongemtcont                   *Xpongemtcont                   `xml:"xpongemtcont,omitempty"`
	HardwareProtectionElementState *HardwareProtectionElementState `xml:"hardware-protection-element-state,omitempty"`
}
type Filter struct {
	Type         string        `xml:"type,attr,omitempty"`
	Onus         *Onus         `xml:"onus,omitempty"`
	Interfaces   *Interfaces   `xml:"interfaces,omitempty"`
	Hardware     *Hardware     `xml:"hardware,omitempty"`
	Forwarding   *Forwarding   `xml:"forwarding,omitempty"`
	Xpongemtcont *Xpongemtcont `xml:"xpongemtcont,omitempty"`
	Syslog       *Syslog       `xml:"syslog,omitempty"`
	System       *System       `xml:"system,omitempty"`
}

// hardware/
type Hardware struct {
	Xmlns     string                        `xml:"xmlns,attr"`
	Component []rpcHardwareModels.Component `xml:"component,omitempty"`
}

func NewHardware() *Hardware {
	a := new(Hardware)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-hardware"
	return a
}

// hardware-state
type HardwareState struct {
	Xmlns     string                             `xml:"xmlns,attr"`
	Component []rpcHardwareStateModels.Component `xml:"component,omitempty"`
}

func NewHardwareStatee() *HardwareState {
	a := new(HardwareState)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-hardware"
	return a
}

// hardware-protection-element-state
type HardwareProtectionElementState struct {
	Xmlns     string                                              `xml:"xmlns,attr"`
	Component []rpcHardwareProtectionElementStateModels.Component `xml:"component,omitempty"`
}

func NewHardwareProtectionElementState() *HardwareProtectionElementState {
	a := new(HardwareProtectionElementState)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-protection-management"
	return a
}

// system-state
type SystemState struct {
	Xmlns            string                                 `xml:"xmlns,attr"`
	Platform         *rpcSystemStateModels.Platform         `xml:"platform,omitempty"`
	Clock            *rpcSystemStateModels.Clock            `xml:"clock,omitempty"`
	RadiusStatistics *rpcSystemStateModels.RadiusStatistics `xml:"radius-statistics,omitempty"`
}

func NewSystemState() *SystemState {
	a := new(SystemState)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-system"
	return a
}

// interfaces-state/
type InterfacesState struct {
	Xmlns     string                               `xml:"xmlns,omitempty,attr"`
	Interface []rpcInterfacesStateModels.Interface `xml:"interface,omitempty"`
}

func NewInterfacesState() *InterfacesState {
	a := new(InterfacesState)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-interfaces"
	return a
}

type AlarmNotification struct {
	XMLName xml.Name `xml:"alarm-notification"`
	//Id      int      `xml:"id,attr"`
	Resource          string `xml:"resource"`
	AlarmTypeId       string `xml:"alarm-type-id"`
	Time              string `xml:"time"`
	PerceivedSeverity string `xml:"perceived-severity"`
	AlarmText         string `xml:"alarm-text"`
}

// used for read  alarm reply,
type NotificationReply struct {
	XMLName           xml.Name           `xml:"notification"`
	EventTime         string             `xml:"eventTime"`
	AlarmNotification *AlarmNotification `xml:"alarm-notification,omitempty"`
}

// used for read  rpc reply,
type FilterRPCReply struct {
	XMLName         xml.Name         `xml:"data"`
	Onus            *Onus            `xml:"onus,omitempty"`
	InterfacesState *InterfacesState `xml:"interfaces-state,omitempty"`
	Interfaces      *Interfaces      `xml:"interfaces,omitempty"`
	Hardware        *Hardware        `xml:"hardware,omitempty"`

	SystemState       *SystemState       `xml:"system-state,omitempty"`
	System            *System            `xml:"system,omitempty"`
	HardwareState     *HardwareState     `xml:"hardware-state,omitempty"`
	Forwarding        *Forwarding        `xml:"forwarding,omitempty"`
	ForwardingState   *ForwardingState   `xml:"forwarding-state,omitempty"`
	Alarms            *Alarms            `xml:"alarms,omitempty"`
	Xpongemtcont      *Xpongemtcont      `xml:"xpongemtcont,omitempty"`
	TmProfiles        *TmProfiles        `xml:"tm-profiles,omitempty"`
	XpongemtcontState *XpongemtcontState `xml:"xpongemtcont-state,omitempty"`
	Comment           string             `xml:",comment"` // can output a comment in xml file

	State     *State     `xml:"state,omitempty"`
	Configure *Configure `xml:"configure,omitempty"`
	Syslog    *Syslog    `xml:"syslog,omitempty"`

	OperDataFormatCliBlock         *OperDataFormatCliBlock         `xml:"oper-data-format-cli-block,omitempty"`
	ControlledChannelPairState     *ControlledChannelPairState     `xml:"controlled-channel-pair-state,omitempty"`
	ControlledChannelPair          *ControlledChannelPair          `xml:"controlled-channel-pair,omitempty"`
	HardwareProtectionElementState *HardwareProtectionElementState `xml:"hardware-protection-element-state,omitempty"`
}

type Config struct {
	Onus         *Onus         `xml:"onus,omitempty"`
	Interfaces   *Interfaces   `xml:"interfaces,omitempty"`
	Forwarding   *Forwarding   `xml:"forwarding,omitempty"`
	Hardware     *Hardware     `xml:"hardware,omitempty"`
	Syslog       *Syslog       `xml:"syslog,omitempty"`
	Xpongemtcont *Xpongemtcont `xml:"xpongemtcont,omitempty"`
	System       *System       `xml:"system,omitempty"`
	Configure    *Configure    `xml:"configure,omitempty"`
}

type ControlledChannelPairState struct {
	Xmlns                  string                                                      `xml:"xmlns,attr"`
	ControlledChannelPairs []rpcControlledChannelPairStateModels.ControlledChannelPair `xml:"controlled-channel-pair,omitempty"`
}

type ControlledChannelPair struct {
	Xmlns                  string                                                      `xml:"xmlns,attr"`
	ControlledChannelPairs []rpcControlledChannelPairStateModels.ControlledChannelPair `xml:"controlled-channel-pair,omitempty"`
}

// ********************************ONUS part start****************************************
// ****************************************************************************************
// onus/
type Onus struct {
	Xmlns string `xml:"xmlns,attr"`
	// ONU   struct {
	// 	XmlnsXc   string `xml:"xmlns:xc,attr"`
	// 	Operation string `xml:"xc:operation,omitempty,attr"` // may be should be added automaticlly
	// 	Name      string `xml:"name"`
	// 	// Usage         interface{}    `xml:",innerxml"` // can not insert namespace for sigle tag, so use innerxml instead of it
	// 	Usage struct {
	// 		Name          string `xml:",chardata"`
	// 		XmlnsTemplate string `xml:"xmlns:template-common,attr"`
	// 	} `xml:"usage,omitempty"`
	// 	OnuManagement *OnuManagement `xml:"onu-management,omitempty"`
	// 	Root          *Root          `xml:"root,omitempty"`
	// } `xml:"onu"`
	Onu []rpcOnusModels.Onu `xml:"onu,omitempty"`
}

func NewOnus() *Onus {
	a := new(Onus)
	a.Xmlns = "urn:bbf:params:xml:ns:yang:bbf-fiber-onu-emulated-mount"
	return a
}

// ********************************ONUS part end****************************************
//****************************************************************************************

// ********************************Interface part start****************************************
//****************************************************************************************

// interfaces/
type Interfaces struct {
	Xmlns      string                           `xml:"xmlns,attr"`
	Interface1 []rpcInterfacesModels.Interface1 `xml:"interface,omitempty"`
}

func NewInterfaces() *Interfaces {
	a := new(Interfaces)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-interfaces"
	return a
}

// ********************************Interface part end****************************************
//****************************************************************************************

// ********************************system part start****************************************
//****************************************************************************************

// system/
type System struct {
	Xmlns  string                  `xml:"xmlns,attr"`
	Radius *rpcSystemModels.Radius `xml:"radius,omitempty"`
}

func NewSystem() *System {
	a := new(System)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-system"
	return a
}

// ********************************system part end****************************************
//****************************************************************************************

// ********************************syslog part end****************************************
// ****************************************************************************************
type Syslog struct {
	Xmlns   string                   `xml:"xmlns,attr,omitempty"`
	Actions *rpcSyslogModels.Actions `xml:"actions,omitempty"`
}

func NewSyslog() *Syslog {
	a := new(Syslog)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-syslog"
	return a
}

// ********************************syslog part end****************************************
//****************************************************************************************

// ********************************Forwarding part start****************************************
//****************************************************************************************

// forwarding/
type Forwarding struct {
	Xmlns                      string                                          `xml:"xmlns,attr,omitempty"`
	MacLearningControlProfiles *rpcForwardingModels.MacLearningControlProfiles `xml:"mac-learning-control-profiles,omitempty"`
	ForwardingBases            *rpcForwardingModels.ForwardingDatabases        `xml:"forwarding-databases,omitempty"`
	Forwarders                 *rpcForwardingModels.Forwarders                 `xml:"forwarders,omitempty"`
}

func NewForwarding() *Forwarding {
	a := new(Forwarding)
	a.Xmlns = "urn:bbf:yang:bbf-l2-forwarding"
	return a
}

// ********************************Forwarding part end****************************************
//****************************************************************************************

// ********************************Forwarding-state part start****************************************
//****************************************************************************************

// forwarding-state/
type ForwardingState struct {
	Xmlns           string                                        `xml:"xmlns,attr,omitempty"`
	ForwardingBases *rpcForwardingStateModels.ForwardingDatabases `xml:"forwarding-databases,omitempty"`
}

func NewForwardingState() *ForwardingState {
	a := new(ForwardingState)
	a.Xmlns = "urn:bbf:yang:bbf-l2-forwarding"
	return a
}

// ********************************Forwarding-state part end****************************************
//****************************************************************************************

// ********************************Alarms part start****************************************
//****************************************************************************************

// alarms/
type Alarms struct {
	Xmlns     string                     `xml:"xmlns,attr,omitempty"`
	AlarmList *rpcAlarmsModels.AlarmList `xml:"alarm-list,omitempty"`
	Summary   *rpcAlarmsModels.Summary   `xml:"summary,omitempty"`
	// used for ihub yang
	AlarmSummary []rpcAlarmsModels.AlarmSummary `xml:"alarm-summary,omitempty"`
}

func NewAlarms() *Alarms {
	a := new(Alarms)
	a.Xmlns = "urn:ietf:params:xml:ns:yang:ietf-alarms"
	return a
}
func NewAlarmsIhub() *Alarms {
	a := new(Alarms)
	a.Xmlns = "urn:bbf:yang:bbf-alarm-management"
	return a
}

// ********************************Alarms part end****************************************
//****************************************************************************************

// ********************************xpongemtcont-state and xpongemtcont part start****************************************
// ****************************************************************************************
// xpongemtcont-state
type XpongemtcontState struct {
	Tconts   *rpcXpongemtcontStateModels.Tconts   `xml:"tconts,omitempty"`
	Gemports *rpcXpongemtcontStateModels.Gemports `xml:"gemports,omitempty"`
}

// xpongemtcont
type Xpongemtcont struct {
	Xmlns                     string                                                `xml:"xmlns,attr,omitempty"`
	Tconts                    *rpcXpongemtcontStateModels.Tconts                    `xml:"tconts,omitempty"`
	Gemports                  *rpcXpongemtcontStateModels.Gemports                  `xml:"gemports,omitempty"`
	TrafficDescriptorProfiles *rpcXpongemtcontStateModels.TrafficDescriptorProfiles `xml:"traffic-descriptor-profiles,omitempty"`
}

func NewXpongemtcont() *Xpongemtcont {
	a := new(Xpongemtcont)
	a.Xmlns = "urn:bbf:yang:bbf-xpongemtcont"
	return a
}
func NewXpongemtcont2() *Xpongemtcont {
	a := new(Xpongemtcont)
	a.Xmlns = "urn:bbf:yang:bbf-xpongemtcont"
	a.Gemports = rpcXpongemtcontStateModels.NewGemports()
	a.Tconts = rpcXpongemtcontStateModels.NewTconts()
	return a
}

// ********************************xpongemtcont-state and xpongemtcont part end****************************************
//****************************************************************************************

// ********************************tm-profiles part start****************************************
// ****************************************************************************************
// tm-profiles
type TmProfiles struct {
	ShaperProfile []rpcTmProfiles.ShaperProfile `xml:"shaper-profile,omitempty"`
}
type OperDataFormatCliBlock struct {
	Item []rpcOperDataFormatCliBlock.Item `xml:"item"`
}

// ********************************tm-profiles part end****************************************
//****************************************************************************************

// ********************************Configure part start****************************************
//****************************************************************************************

// configure/
type Configure struct {
	Xmlns   string                          `xml:"xmlns,attr,omitempty"`
	Port    []rpcConfigurePortModels.Port   `xml:"port,omitempty"`
	Service *rpcConfigurePortModels.Service `xml:"service,omitempty"`
	Lag     []rpcConfigurePortModels.Lag    `xml:"lag,omitempty"`
}

func NewConfigure() *Configure {
	a := new(Configure)
	a.Xmlns = "urn:nokia.com:sros:ns:yang:sr:conf"
	return a
}

// ********************************Configure part end****************************************
//****************************************************************************************

// ********************************Configure part start****************************************
//****************************************************************************************

// state/
type State struct {
	Xmlns string                    `xml:"xmlns,attr,omitempty"`
	Port  []rpcStatePortModels.Port `xml:"port,omitempty"` // 改为数组类型
}

func NewState() *State {
	a := new(State)
	a.Xmlns = "urn:nokia.com:sros:ns:yang:sr:state"
	return a
}

// ********************************Configure part end****************************************
//****************************************************************************************

// Constructor start
// ****************************************************************************************
func NewRpcError() *RpcError {
	a := new(RpcError)
	return a
}
func NewRPCInfor() *RPCInfor {
	a := new(RPCInfor)
	return a
}
func NewRPCOntInfor() *RPCOntInfor {
	a := new(RPCOntInfor)
	return a
}

func NewEditInterfacesONUsConfigRPC() *EditConfig {
	a := new(EditConfig)
	a.Config.Onus = NewOnus()
	a.Config.Interfaces = NewInterfaces()
	// fmt.Println(a)
	return a
}
func NewEditONUsConfigRPC() *EditConfig {
	a := new(EditConfig)
	a.Config.Onus = NewOnus()
	// fmt.Println(a)
	return a
}

func NewEditHardwareConfigRPC() *EditConfig {
	a := new(EditConfig)
	a.Config.Hardware = NewHardware()
	// fmt.Println(a)
	return a
}
func NewEditInterfacesConfigRPC() *EditConfig {
	a := new(EditConfig)
	a.Config.Interfaces = NewInterfaces()
	// fmt.Println(a)
	return a
}

func NewEditConfigDeleteInterfacesONUsRPC() *EditConfig {
	a := NewEditInterfacesONUsConfigRPC()
	onu := rpcOnusModels.NewONU(constants.OPERATION_DELETE)
	onu.XmlnsXc = constants.XMLNS_VERSION
	onu.Operation = constants.OPERATION_DELETE
	// can not insert namespace for sigle tag, so use innerxml instead of it
	// onu.Usage = `<usage xmlns:template-common="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-template-common">template-common:node-actual-usage</usage>`
	// onu.Usage.XmlnsTemplate = `http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-template-common`
	// onu.Usage.Name = "template-common:node-actual-usage"
	// a.Config.Onus.Onu = onu
	a.Config.Onus.Onu = append(a.Config.Onus.Onu, *onu)
	// fmt.Println(a)
	return a
}

// Constructor end
// ****************************************************************************************
