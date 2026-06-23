package rpcInterfacesModels

import "alto_server/constants"

// ********************************Common struct part start****************************************
//****************************************************************************************

//Demo   <onu-name xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-if-xponvani-aug">C12_ONT2</onu-name>

type C_TagWithAttrAndValue struct {
	Name            string `xml:",chardata"`
	Xmlns           string `xml:"xmlns,attr,omitempty"`
	XmlnsBbfXPonift string `xml:"xmlns:bbf-xponift,attr,omitempty"`
	XmlnsIanaIft    string `xml:"xmlns:ianaift,attr,omitempty"`
	Operation       string `xml:"operation,omitempty,attr"`
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

//interfaces/interface
type Interface1 struct {
	XmlnsXc     string `xml:"xmlns:xc,omitempty,attr"`
	Operation   string `xml:"xc:operation,omitempty,attr"`
	Name        string `xml:"name,omitempty"`
	Description string `xml:"description,omitempty"`
	// Type        interface{} `xml:",innerxml,omitempty"` // <type xmlns:bbf-xponift="urn:bbf:yang:bbf-xpon-if-type">bbf-xponift:v-ani</type> it is for xgspon
	Type                  *C_TagWithAttrAndValue  `xml:"type,omitempty"`
	Enable                string                  `xml:"enabled,omitempty"` // can not set omitempty, the value  false will cause tge tag does not display in xml
	Statistics            *C_TagWithAttrAndEnable `xml:"statistics,omitempty"`
	PortLayerIf           *C_TagWithAttrAndValue  `xml:"port-layer-if,omitempty"`
	Performance           *C_TagWithAttrAndEnable `xml:"performance,omitempty"`
	VAni                  *VAni                   `xml:"v-ani,omitempty"`
	OltVEnet              *OltVEnet               `xml:"olt-v-enet,omitempty"`
	SubifLowerLayer       *SubifLowerLayer        `xml:"subif-lower-layer,omitempty"`
	InlineFrameProcessing *InlineFrameProcessing  `xml:"inline-frame-processing,omitempty"`
	TMRoot                *TMRoot                 `xml:"tm-root,omitempty"`
	ChannelTermination    *ChannelTermination     `xml:"channel-termination,omitempty"`
	ChannelPair           *ChannelPair            `xml:"channel-pair,omitempty"`
	ChannelPartition      *ChannelPartition       `xml:"channel-partition,omitempty"`
	Pae                   *Pae                    `xml:"pae,omitempty"`
	Ipv4Security          *Ipv4Security           `xml:"ipv4-security,omitempty"`

	AggregationPort *AggregationPort `xml:"aggregation-port,omitempty"`
	Aggregator      *Aggregator      `xml:"aggregator,omitempty"`

	InterfaceUsage *InterfaceUsage `xml:"interface-usage,omitempty"`
}

func NewInterface1(action string) *Interface1 {
	a := new(Interface1)
	a.XmlnsXc = constants.XMLNS_VERSION
	if action == "remove" {
		a.Operation = constants.OPERATION_DELETE
	}
	return a
}

// interfaces/interface/ipv4-security
type Ipv4Security struct {
	Xmlns                      string `xml:"xmlns,attr,omitempty"`
	PreventIpv4AddressSpoofing string `xml:"prevent-ipv4-address-spoofing,omitempty"`
	MaxAddress                 string `xml:"max-address,omitempty"`

	Address []Address `xml:"address,omitempty"`
}

func NewIpv4Security() *Ipv4Security {
	a := new(Ipv4Security)
	a.Xmlns = "uri:http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ipv4-address-spoofing-prevention"
	return a
}

// interfaces/interface/aggregator
type InterfaceUsage struct {
	Xmlns          string `xml:"xmlns,attr,omitempty"`
	InterfaceUsage string `xml:"interface-usage,omitempty"`
}

func NewInterfaceUsage() *InterfaceUsage {
	a := new(InterfaceUsage)
	a.Xmlns = "urn:bbf:yang:bbf-interface-usage"
	return a
}

// interfaces/interface/aggregator
type Aggregator struct {
	Xmlns             string `xml:"xmlns,attr,omitempty"`
	Name              string `xml:"name,omitempty"`
	AggSystemName     string `xml:"agg-system-name,omitempty"`
	NonRevertiveLag   string `xml:"non-revertive-lag,omitempty"`
	Mode              string `xml:"mode,omitempty"`
	MaxActiveNumber   string `xml:"max-active-number,omitempty"`
	PrimaryLagPortRef string `xml:"primary-lag-port-ref,omitempty"`

	AggregatorLacp *AggregatorLacp `xml:"aggregator-lacp,omitempty"`
}

func NewAggregator() *Aggregator {
	a := new(Aggregator)
	a.Xmlns = "urn:ieee:std:802.1AX:yang:ieee802-dot1ax"
	return a
}

// interfaces/interface/aggregation-port
type AggregationPort struct {
	Xmlns               string               `xml:"xmlns,attr,omitempty"`
	Operation           string               `xml:"xc:operation,omitempty,attr"`
	AggregationPortLacp *AggregationPortLacp `xml:"aggregation-port-lacp,omitempty"`
}

func NewAggregationPort() *AggregationPort {
	a := new(AggregationPort)
	a.Xmlns = "urn:ieee:std:802.1AX:yang:ieee802-dot1ax"
	return a
}
func NewAggregationPortRemove() *AggregationPort {
	a := new(AggregationPort)
	a.Operation = constants.OPERATION_DELETE
	a.Xmlns = "urn:ieee:std:802.1AX:yang:ieee802-dot1ax"
	return a
}

// interfaces/interface/ipv4-security/address
// interfaces/interface/ipv6-security/address
type Address struct {
	Operation             string `xml:"xc:operation,omitempty,attr"`
	IP                    string `xml:"ip,omitempty"`
	Netmask               string `xml:"netmask,omitempty"`
	PrefixLength          string `xml:"prefix-length,omitempty"`
	IpAddressOrigin       string `xml:"ip-address-origin,omitempty"`
	Ipv6AddressOrigin     string `xml:"ipv6-address-origin,omitempty"`
	LeaseTime             string `xml:"lease-time,omitempty"`
	IpAddressExpiryTime   string `xml:"ip-address-expiry-time,omitempty"`
	Ipv6AddressExpiryTime string `xml:"ipv6-address-expiry-time,omitempty"`
	Chaddr                string `xml:"chaddr,omitempty"`
}

func NewAddress() *Address {
	a := new(Address)
	return a
}

//interfaces/interface/tm-root
type TMRoot struct {
	Xmlns                          string `xml:"xmlns,attr"`
	TcId2QueueIdMappingProfileName string `xml:"tc-id-2-queue-id-mapping-profile-name,omitempty"`

	SchedulerNode       []SchedulerNode       `xml:"scheduler-node"`
	ChildSchedulerNodes []ChildSchedulerNodes `xml:"child-scheduler-nodes,omitempty"`
}

func NewTMRoot() *TMRoot {
	a := new(TMRoot)
	a.Xmlns = "urn:bbf:yang:bbf-qos-traffic-mngt"
	return a
}

//interfaces/interface/tm-root/scheduler-node
type SchedulerNode struct {
	Xmlns                     string                      `xml:"xmlns,attr"`
	XmlnsXc                   string                      `xml:"xmlns:xc,omitempty,attr"`
	Operation                 string                      `xml:"xc:operation,omitempty,attr"`
	Name                      string                      `xml:"name"`
	SchedulingLevel           string                      `xml:"scheduling-level,omitempty"`
	ContainsQueues            string                      `xml:"contains-queues,omitempty"`
	ChildSchedulerNodes4SNode []ChildSchedulerNodes4SNode `xml:"child-scheduler-nodes"`
	Queue                     []Queue                     `xml:"queue"`
	QueueMonitoring           *QueueMonitoring            `xml:"queue-monitoring,omitempty"`
}

func NewSchedulerNode(action string) *SchedulerNode {
	a := new(SchedulerNode)
	a.Xmlns = "urn:bbf:yang:bbf-qos-enhanced-scheduling"
	if action == "remove" {
		a.XmlnsXc = constants.XMLNS_VERSION
		a.Operation = constants.OPERATION_DELETE
	}
	return a
}

//interfaces/interface/tm-root/scheduler-node/child-scheduler-nodes
type ChildSchedulerNodes4SNode struct {
	Name     string `xml:"name"`
	Priority string `xml:"priority"`
	Weight   string `xml:"weight"`
}

func NewChildSchedulerNodes4SNode1() *ChildSchedulerNodes4SNode {
	a := new(ChildSchedulerNodes4SNode)
	a.Priority = "0"
	a.Weight = "1"
	return a
}

//interfaces/interface/tm-root/scheduler-node/queue
type Queue struct {
	LocalQueueId string `xml:"local-queue-id"`
	BacName      string `xml:"bac-name"`
	Priority     string `xml:"priority"`
	Weight       string `xml:"weight"`
}

func NewQueue() *Queue {
	a := new(Queue)
	return a
}

func NewQueue1() *Queue {
	a := new(Queue)
	a.BacName = "DEFAULTQ"
	a.Weight = "0"
	return a
}

//interfaces/interface/tm-root/scheduler-node/queue-monitoring
type QueueMonitoring struct {
	Xmlns             string `xml:"xmlns,attr"`
	EnableStatistics  string `xml:"enable-statistics"`
	EnablePerformance string `xml:"enable-performance"`
}

func NewQueueMonitoring() *QueueMonitoring {
	a := new(QueueMonitoring)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-qos-queue-monitoring-extension"
	a.EnableStatistics = "false"
	a.EnablePerformance = "false"
	return a
}

//interfaces/interface/tm-root/child-scheduler-nodes
type ChildSchedulerNodes struct {
	Xmlns     string `xml:"xmlns,attr"`
	XmlnsXc   string `xml:"xmlns:xc,attr"`
	Operation string `xml:"xc:operation,omitempty,attr"`
	Name      string `xml:"name"`
	Priority  string `xml:"priority,omitempty"`
	Weight    string `xml:"weight,omitempty"`
}

func NewChildSchedulerNodes(action string) *ChildSchedulerNodes {
	a := new(ChildSchedulerNodes)
	a.Xmlns = "urn:bbf:yang:bbf-qos-enhanced-scheduling"
	a.XmlnsXc = constants.XMLNS_VERSION
	if action == "remove" {
		a.Operation = constants.OPERATION_DELETE
	}
	return a
}

//interfaces/interface/v-ani
type VAni struct {
	Xmlns                         string `xml:"xmlns,attr"`
	OnuId                         string `xml:"onu-id"`
	ChannelPartition              string `xml:"channel-partition"`
	ExpectedSerialNumber          string `xml:"expected-serial-number"`
	PreferredChannelPair          string `xml:"preferred-channel-pair"`
	ManagementGemportAesIndicator string `xml:"management-gemport-aes-indicator"`
	UpstreamFec                   string `xml:"upstream-fec"`
	// OnuName                       interface{} `xml:",innerxml,omitempty"` // <onu-name xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-if-xponvani-aug">C12_ONT2</onu-name> it is for xgspon
	OnuName *C_TagWithAttrAndValue `xml:"onu-name,omitempty"`
}

func NewVAni() *VAni {
	a := new(VAni)
	return a
}
func NewVAniXgpon() *VAni {
	a := new(VAni)
	a.Xmlns = "urn:bbf:yang:bbf-xponvani"
	a.ManagementGemportAesIndicator = "false"
	a.UpstreamFec = "true"
	return a
}

//interfaces/interface/olt-v-enet
type OltVEnet struct {
	Xmlns               string                 `xml:"xmlns,attr"`
	LowerLayerInterface string                 `xml:"lower-layer-interface"`
	UniName             *C_TagWithAttrAndValue `xml:"uni-name,omitempty"`
}

func NewOltVEnet() *OltVEnet {
	a := new(OltVEnet)
	return a
}

//interfaces/interface/subif-lower-layer
type SubifLowerLayer struct {
	Xmlns     string `xml:"xmlns,attr"`
	Interface string `xml:"interface"`
}

func NewSubifLowerLayer() *SubifLowerLayer {
	a := new(SubifLowerLayer)
	a.Xmlns = "urn:bbf:yang:bbf-sub-interfaces"
	return a
}

//interfaces/interface/inline-frame-processing
type InlineFrameProcessing struct {
	Xmlns         string         `xml:"xmlns,attr"`
	EgressRewrite *EgressRewrite `xml:"egress-rewrite,omitempty"`
}

func NewInlineFrameProcessing() *InlineFrameProcessing {
	a := new(InlineFrameProcessing)
	a.Xmlns = "urn:bbf:yang:bbf-sub-interfaces"
	return a
}

//interfaces/interface/inline-frame-processing/egress-rewrite
type EgressRewrite struct {
	PushTag *PushTag `xml:"push-tag,omitempty"`
}

func NewEgressRewrite() *EgressRewrite {
	a := new(EgressRewrite)
	return a
}

//interfaces/interface/inline-frame-processing/egress-rewrite/push-tag
type PushTag struct {
	Xmlns    string    `xml:"xmlns,attr,omitempty"`
	Index    string    `xml:"index,omitempty"`
	Dot1qTag *Dot1qTag `xml:"dot1q-tag,omitempty"`
}

func NewPushTag() *PushTag {
	a := new(PushTag)
	return a
}

//interfaces/interface/inline-frame-processing/egress-rewrite/push-tag/dot1q-tag
type Dot1qTag struct {
	VlanId string `xml:"vlan-id,omitempty"`
}

func NewDot1qTag() *Dot1qTag {
	a := new(Dot1qTag)
	return a
}

//interfaces/interface/channel-temination
type ChannelTermination struct {
	Xmlns                  string `xml:"xmlns,attr"`
	ChannelPairRef         string `xml:"channel-pair-ref,omitempty"`
	ChannelTerminationType string `xml:"channel-termination-type,omitempty"`
	BerCalcPeriod          string `xml:"ber-calc-period,omitempty"`
	Location               string `xml:"location,omitempty"`
}

//interfaces/interface/channel-pair
type ChannelPair struct {
	Xmlns                             string `xml:"xmlns,attr"`
	ChannelGroupRef                   string `xml:"channel-group-ref,omitempty"`
	ChannelPartitionRef               string `xml:"channel-partition-ref,omitempty"`
	ChannelPairType                   string `xml:"channel-pair-type,omitempty"`
	GponPonIdInterval                 string `xml:"gpon-pon-id-interval,omitempty"`
	OnuPloamAuthenticationFailControl string `xml:"onu-ploam-authentication-fail-control,omitempty"`
}

//interfaces/interface/channel-partition
type ChannelPartition struct {
	Xmlns                           string `xml:"xmlns,attr"`
	ChannelGroupRef                 string `xml:"channel-group-ref,omitempty"`
	ChannelPartitionIndex           string `xml:"channel-partition-index,omitempty"`
	DownstreamFec                   string `xml:"downstream-fec,omitempty"`
	ClosestOnuDistance              string `xml:"closest-onu-distance,omitempty"`
	MaximumDifferentialXponDistance string `xml:"maximum-differential-xpon-distance,omitempty"`
	AuthenticationMethod            string `xml:"authentication-method,omitempty"`
	MulticastAesIndicator           string `xml:"multicast-aes-indicator,omitempty"`
}

//interfaces/interface/pae
type Pae struct {
	Xmlns            string            `xml:"xmlns,attr"`
	Operation        string            `xml:"operation,omitempty,attr"`
	EapolStatistics  *EapolStatistics  `xml:"eapol-statistics,omitempty"`
	PortCapabilities *PortCapabilities `xml:"port-capabilities,omitempty"`
	Authenticator    *Authenticator    `xml:"authenticator,omitempty"`
	ExtAuthenticator *ExtAuthenticator `xml:"ext-authenticator,omitempty"`
}

func NewPae() *Pae {
	a := new(Pae)
	a.Xmlns = "urn:ieee:std:802.1X:yang:ieee802-dot1x"
	return a
}

//interfaces/interface/pae/eapol-statistics
type EapolStatistics struct {
	InvalidEapolFrameRx    string `xml:"invalid-eapol-frame-rx,omitempty"`
	EapLengthErrorFramesRx string `xml:"eap-length-error-frames-rx,omitempty"`
	EapolStartFramesRx     string `xml:"eapol-start-frames-rx,omitempty"`
	EapolEapFramesRx       string `xml:"eapol-eap-frames-rx,omitempty"`
	EapolLogoffFramesRx    string `xml:"eapol-logoff-frames-rx,omitempty"`
	EapolAuthEapFramesTx   string `xml:"eapol-auth-eap-frames-tx,omitempty"`
	LastEapolFrameSource   string `xml:"last-eapol-frame-source,omitempty"`
}

func NewEapolStatistics() *EapolStatistics {
	a := new(EapolStatistics)
	return a
}

//interfaces/interface/pae/port-capabilities
type PortCapabilities struct {
	Auth      string `xml:"auth,omitempty"`
	Operation string `xml:"operation,omitempty,attr"`
}

func NewPortCapabilities() *PortCapabilities {
	a := new(PortCapabilities)
	return a
}

//interfaces/interface/pae/authenticator
type Authenticator struct {
	QuietPeriod string `xml:"quiet-period,omitempty"`
}

func NewAuthenticator() *Authenticator {
	a := new(Authenticator)
	return a
}

//interfaces/interface/pae/ext-authenticator
type ExtAuthenticator struct {
	Xmlns                         string `xml:"xmlns,attr,omitempty"`
	InitiateAuthenticationRequest string `xml:"initiate-authentication-request,omitempty"`
	TxPeriod                      string `xml:"tx-period,omitempty"`
	InterfaceSupplicantTimeout    string `xml:"interface-supplicant-timeout,omitempty"`
	HandshakeEnable               string `xml:"handshake-enable,omitempty"`
	HandshakePeriod               string `xml:"handshake-period,omitempty"`
	AuthenticationMode            string `xml:"authentication-mode,omitempty"`
	SignallingChannel             string `xml:"signalling-channel,omitempty"`
}

func NewExtAuthenticator() *ExtAuthenticator {
	a := new(ExtAuthenticator)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-802-dot1x-ext"
	return a
}

//interfaces/interface/aggregator/aggregator-lacp
type AggregatorLacp struct {
	ActorAdminKey string `xml:"actor-admin-key,omitempty"`
}

func NewAggregatorLacp() *AggregatorLacp {
	a := new(AggregatorLacp)
	return a
}

// interfaces/interface/aggregation-port/aggregation-port-lacp
type AggregationPortLacp struct {
	ActorAdminKey       string `xml:"actor-admin-key,omitempty"`
	ActorSystemPriority string `xml:"actor-system-priority,omitempty"`
	ActorPortPriority   string `xml:"actor-port-priority,omitempty"`
	ActorAdminState     string `xml:"actor-admin-state,omitempty"`
}

func NewAggregationPortLacp() *AggregationPortLacp {
	a := new(AggregationPortLacp)
	return a
}
