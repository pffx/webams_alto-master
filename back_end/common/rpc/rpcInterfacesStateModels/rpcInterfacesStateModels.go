package rpcInterfacesStateModels

import "encoding/xml"

// interfaces-state/interface
type Interface struct {
	Name               string                `xml:"name,omitempty"`
	Type               C_TagWithAttrAndValue `xml:"type,omitempty"`
	AdminStatus        string                `xml:"admin-status,omitempty"`
	OperStatus         string                `xml:"oper-status,omitempty"`
	LastChange         string                `xml:"last-change,omitempty"`
	IfIndex            string                `xml:"if-index,omitempty"`
	HigherLayerIf      string                `xml:"higher-layer-if,omitempty"`
	LowerLayerIf       string                `xml:"lower-layer-if,omitempty"`
	PortLayerIf        string                `xml:"port-layer-if,omitempty"`
	ChannelTermination *ChannelTermination   `xml:"channel-termination,omitempty"`
	VAni               *VAni                 `xml:"v-ani,omitempty"`
	Pae                *Pae                  `xml:"pae,omitempty"`
	Ipv4Security       *Ipv4Security         `xml:"ipv4-security,omitempty"`
	Ipv6Security       *Ipv6Security         `xml:"ipv6-security,omitempty"`
	Performance        *Performance          `xml:"performance,omitempty"`
	Statistics         *Statistics           `xml:"statistics,omitempty"`
	SpeedMonitoring    *SpeedMonitoring      `xml:"speed-monitoring,omitempty"`
}

func NewInterface() *Interface {
	a := new(Interface)
	return a
}

type C_TagWithAttrAndValue struct {
	Value           string `xml:",chardata"`
	Xmlns           string `xml:"xmlns,attr,omitempty"`
	XmlnsBbfXPonift string `xml:"xmlns:bbf-xponift,attr,omitempty"`
}

// interfaces-state/interface/channel-termination
type ChannelTermination struct {
	Xmlns                                string                                `xml:"xmlns,attr,omitempty"`
	PonIdDisplay                         string                                `xml:"pon-id-display,omitempty"`
	Location                             string                                `xml:"location,omitempty"`
	OnusPresentOnLocalChannelTermination *OnusPresentOnLocalChannelTermination `xml:"onus-present-on-local-channel-termination,omitempty"`
}

func NewChannelTermination() *ChannelTermination {
	a := new(ChannelTermination)
	return a
}
func NewChannelTermination2() *ChannelTermination {
	a := new(ChannelTermination)
	a.Xmlns = "urn:bbf:yang:bbf-xpon"
	return a
}

// interfaces-state/interface/channel-termination/onus-present-on-local-channel-termination
type OnusPresentOnLocalChannelTermination struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	Onu   []Onu  `xml:"onu,omitempty"`
}

func NewOnusPresentOnLocalChannelTermination() *OnusPresentOnLocalChannelTermination {
	a := new(OnusPresentOnLocalChannelTermination)
	a.Xmlns = "urn:bbf:yang:bbf-xpon-onu-state"
	return a
}

// interfaces-state/interface/v-ani
type VAni struct {
	Xmlns                  string               `xml:"xmlns,attr,omitempty"`
	OnuId                  string               `xml:"onu-id,omitempty"`
	ManagementTcontAllocId string               `xml:"management-tcont-alloc-id,omitempty"`
	ManagementGemportId    string               `xml:"management-gemport-id,omitempty"`
	OnuPresentOnThisOlt    *OnuPresentOnThisOlt `xml:"onu-present-on-this-olt,omitempty"`
}

func NewVAni() *VAni {
	a := new(VAni)
	a.Xmlns = "urn:bbf:yang:bbf-xponvani"
	return a
}

// interfaces-state/interface/v-ani/onu-present-on-this-olt
type OnuPresentOnThisOlt struct {
	DetectedSerialNumber string `xml:"detected-serial-number,omitempty"`
	OnuFiberDistance     string `xml:"onu-fiber-distance,omitempty"`
}

// interfaces-state/interface/channel-termination/onus-present-on-local-channel-termination/onu/
type Onu struct {
	XMLName                xml.Name `xml:"onu"`
	DetectedSerialNumber   string   `xml:"detected-serial-number,omitempty"`
	OnuPresenceState       string   `xml:"onu-presence-state,omitempty"`
	OnuId                  string   `xml:"onu-id,omitempty"`
	DetectedRegistrationId string   `xml:"detected-registration-id,omitempty"`
	VAniRef                string   `xml:"v-ani-ref,omitempty"`
	OnuDetectedDatetime    string   `xml:"onu-detected-datetime,omitempty"`
	OnuStateLastChange     string   `xml:"onu-state-last-change,omitempty"`
	DetectedUpstreamRate   string   `xml:"detected-upstream-rate,omitempty"`
}

func NewOnu() *Onu {
	a := new(Onu)
	return a
}

// interfaces-state/interface/pae
type Pae struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	Port  *Port  `xml:"port,omitempty"`
}

func NewPae() *Pae {
	a := new(Pae)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-802-dot1x-ext"
	return a
}

// interfaces-state/interface/performance
type Performance struct {
	Xmlns          string          `xml:"xmlns,attr,omitempty"`
	Intervals15min *Intervals15min `xml:"intervals-15min,omitempty"`
}

func NewPerformance() *Performance {
	a := new(Performance)
	a.Xmlns = "urn:bbf:yang:bbf-interfaces-performance-management"
	return a
}

// interfaces-state/interface/statistics
type Statistics struct {
	// In statistics
	InPkts             string `xml:"in-pkts,omitempty"`
	InUnicastPackets   string `xml:"in-unicast-pkts,omitempty"`
	InBroadcastPackets string `xml:"in-broadcast-pkts,omitempty"`
	InMulticastPackets string `xml:"in-multicast-pkts,omitempty"`
	InOctets           string `xml:"in-octets,omitempty"`
	InDiscards         string `xml:"in-discards,omitempty"`

	// Out statistics
	OutPkts             string `xml:"out-pkts,omitempty"`
	OutUnicastPackets   string `xml:"out-unicast-pkts,omitempty"`
	OutBroadcastPackets string `xml:"out-broadcast-pkts,omitempty"`
	OutMulticastPackets string `xml:"out-multicast-pkts,omitempty"`
	OutOctets           string `xml:"out-octets,omitempty"`
	OutDiscards         string `xml:"out-discards,omitempty"`
	OutDroppedBytes     string `xml:"out-dropped-bytes,omitempty"`
	DiscontinuityTime   string `xml:"discontinuity-time,omitempty"`
}

func NewStatistics() *Statistics {
	a := new(Statistics)
	return a
}

// interfaces-state/interface/performance/intervals-15min
type Intervals15min struct {
	Current *PerformanceStatics `xml:"current,omitempty"`
}

func NewIntervals15min() *Intervals15min {
	a := new(Intervals15min)
	return a
}

// interfaces-state/interface/performance/intervals-15min/current
type PerformanceStatics struct {
	Xpon *Xpon `xml:"xpon,omitempty"`
}

func NewPerformanceStatics() *PerformanceStatics {
	a := new(PerformanceStatics)
	return a
}

// interfaces-state/interface/performance/intervals-15min/current/xpon
type Xpon struct {
	Phy *Phy `xml:"phy,omitempty"`
}

func NewXpon() *Xpon {
	a := new(Xpon)
	return a
}

// interfaces-state/interface/performance/intervals-15min/current/xpon/phy
type Phy struct {
	InBipErrors         string `xml:"in-bip-errors,omitempty"`
	InBipProtectedWords string `xml:"in-bip-protected-words,omitempty"`
}

func NewPhy() *Phy {
	a := new(Phy)
	return a
}

// interfaces-state/interface/ipv4-security
type Ipv4Security struct {
	Xmlns   string    `xml:"xmlns,attr,omitempty"`
	Address []Address `xml:"address,omitempty"`
}

func NewIpv4Security() *Ipv4Security {
	a := new(Ipv4Security)
	a.Xmlns = "uri:http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ipv4-address-spoofing-prevention"
	return a
}

// interfaces-state/interface/ipv4-security/address
// interfaces-state/interface/ipv6-security/address
type Address struct {
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

// interfaces-state/interface/ipv6-security
type Ipv6Security struct {
	Xmlns   string    `xml:"xmlns,attr,omitempty"`
	Address []Address `xml:"address,omitempty"`
}

func NewIpv6Security() *Ipv6Security {
	a := new(Ipv6Security)
	a.Xmlns = "uri:http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ipv6-address-spoofing-prevention"
	return a
}

// interfaces-state/interface/pae/port
type Port struct {
	AuthenticationStatus        string `xml:"authentication-status,omitempty"`
	LastAuthenticationTimestamp string `xml:"last-authentication-timestamp,omitempty"`
	LastTerminateCause          string `xml:"last-terminate-cause,omitempty"`
}

func NewPort() *Port {
	a := new(Port)
	return a
}

// interfaces-state/interface/speed-monitoring
type SpeedMonitoring struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	Data  *Data  `xml:"data,omitempty"`
}

// interfaces-state/interface/speed-monitoring/data
type Data struct {
	History *History `xml:"history,omitempty"`
}

// interfaces-state/interface/speed-monitoring/data/history
type History struct {
	ReceiveDatarate           string `xml:"receive-datarate,omitempty"`
	TransmitDatarate          string `xml:"transmit-datarate,omitempty"`
	TransmitBroadcastDatarate string `xml:"transmit-broadcast-datarate,omitempty"`
	TransmitMulticastDatarate string `xml:"transmit-multicast-datarate,omitempty"`
	TransmitUnicastDatarate   string `xml:"transmit-unicast-datarate,omitempty"`
}

func NewSpeedMonitoring() *SpeedMonitoring {
	a := new(SpeedMonitoring)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-interface-speed-monitoring"
	return a
}

func NewData() *Data {
	return new(Data)
}

func NewHistory() *History {
	return new(History)
}
