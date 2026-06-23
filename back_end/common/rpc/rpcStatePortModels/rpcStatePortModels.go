package rpcStatePortModels

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
// state/port
type Port struct {
	Name          string       `xml:"name,omitempty" json:"Name,omitempty"`
	PortId        string       `xml:"port-id,omitempty" json:"PortId,omitempty"`
	PortName      string       `xml:"port-name,omitempty" json:"PortName,omitempty"`
	DownReason    string       `xml:"down-reason,omitempty" json:"DownReason,omitempty"`
	OperState     string       `xml:"oper-state,omitempty" json:"OperState,omitempty"`
	PhysicalLink  bool         `xml:"physical-link,omitempty" json:"PhysicalLink,omitempty"`
	PortClass     string       `xml:"port-class,omitempty" json:"PortClass,omitempty"`
	PortState     string       `xml:"port-state,omitempty" json:"PortState,omitempty"`
	PreviousState string       `xml:"previous-state,omitempty" json:"PreviousState,omitempty"`
	Type          string       `xml:"type,omitempty" json:"Type,omitempty"`
	Statistics    *Statistics  `xml:"statistics,omitempty" json:"Statistics,omitempty"`
	Ethernet      *Ethernet    `xml:"ethernet,omitempty" json:"Ethernet,omitempty"`
	Transceiver   *Transceiver `xml:"transceiver"`
}

type Statistics struct {
	InDiscards                int64 `xml:"in-discards,omitempty" json:"InDiscards,omitempty"`
	InPackets                 int64 `xml:"in-packets,omitempty" json:"InPackets,omitempty"`
	InUnknownProtocolDiscards int64 `xml:"in-unknown-protocol-discards,omitempty" json:"InUnknownProtocolDiscards,omitempty"`
	OutDiscards               int64 `xml:"out-discards,omitempty" json:"OutDiscards,omitempty"`
	OutPackets                int64 `xml:"out-packets,omitempty" json:"OutPackets,omitempty"`
}

type Ethernet struct {
	OperSpeed            int64          `xml:"oper-speed,omitempty" json:"OperSpeed,omitempty"`
	OperStateChangeCount int64          `xml:"oper-state-change-count,omitempty" json:"OperStateChangeCount,omitempty"`
	Statistics           *EthernetStats `xml:"statistics,omitempty" json:"Statistics,omitempty"`
	Backplane            *Backplane     `xml:"backplane,omitempty" json:"Backplane,omitempty"`
	Lldp                 *Lldp          `xml:"lldp,omitempty" json:"Lldp,omitempty"`
	Performance          *Performance   `xml:"performance,omitempty" json:"Performance,omitempty"`
}

type EthernetStats struct {
	InBroadcastPackets    int64 `xml:"in-broadcast-packets,omitempty" json:"InBroadcastPackets,omitempty"`
	InMulticastPackets    int64 `xml:"in-multicast-packets,omitempty" json:"InMulticastPackets,omitempty"`
	InUnicastPackets      int64 `xml:"in-unicast-packets,omitempty" json:"InUnicastPackets,omitempty"`
	InErrors              int64 `xml:"in-errors,omitempty" json:"InErrors,omitempty"`
	InOctets              int64 `xml:"in-octets,omitempty" json:"InOctets,omitempty"`
	OutBroadcastPackets   int64 `xml:"out-broadcast-packets,omitempty" json:"OutBroadcastPackets,omitempty"`
	OutMulticastPackets   int64 `xml:"out-multicast-packets,omitempty" json:"OutMulticastPackets,omitempty"`
	OutUnicastPackets     int64 `xml:"out-unicast-packets,omitempty" json:"OutUnicastPackets,omitempty"`
	OutErrors             int64 `xml:"out-errors,omitempty" json:"OutErrors,omitempty"`
	OutOctets             int64 `xml:"out-octets,omitempty" json:"OutOctets,omitempty"`
	OversizePackets       int64 `xml:"oversize-packets,omitempty" json:"OversizePackets,omitempty"`
	UndersizePackets      int64 `xml:"undersize-packets,omitempty" json:"UndersizePackets,omitempty"`
	TotalBroadcastPackets int64 `xml:"total-broadcast-packets,omitempty" json:"TotalBroadcastPackets,omitempty"`
	TotalMulticastPackets int64 `xml:"total-multicast-packets,omitempty" json:"TotalMulticastPackets,omitempty"`
	TotalOctets           int64 `xml:"total-octets,omitempty" json:"TotalOctets,omitempty"`
	TotalPackets          int64 `xml:"total-packets,omitempty" json:"TotalPackets,omitempty"`
}

type Backplane struct {
	BackplaneKr *BackplaneKr `xml:"backplane-kr,omitempty" json:"BackplaneKr,omitempty"`
}

type BackplaneKr struct {
	OperKrMode string `xml:"oper-kr-mode,omitempty" json:"OperKrMode,omitempty"`
}

type Lldp struct {
	DestMac []DestMac `xml:"dest-mac,omitempty" json:"DestMac,omitempty"`
}

type DestMac struct {
	MacType       string          `xml:"mac-type,omitempty" json:"MacType,omitempty"`
	Statistics    *LldpStatistics `xml:"statistics,omitempty" json:"Statistics,omitempty"`
	TxMgmtAddress []TxMgmtAddress `xml:"tx-mgmt-address,omitempty" json:"TxMgmtAddress,omitempty"`
}

type LldpStatistics struct {
	Transmit *TransmitStats `xml:"transmit,omitempty" json:"Transmit,omitempty"`
	Receive  *ReceiveStats  `xml:"receive,omitempty" json:"Receive,omitempty"`
}

type TransmitStats struct {
	Frames            int64 `xml:"frames,omitempty" json:"Frames,omitempty"`
	LengthErrorFrames int64 `xml:"length-error-frames,omitempty" json:"LengthErrorFrames,omitempty"`
}

type ReceiveStats struct {
	AgeOuts       int64 `xml:"age-outs,omitempty" json:"AgeOuts,omitempty"`
	Frames        int64 `xml:"frames,omitempty" json:"Frames,omitempty"`
	FrameDiscards int64 `xml:"frame-discards,omitempty" json:"FrameDiscards,omitempty"`
	FrameErrors   int64 `xml:"frame-errors,omitempty" json:"FrameErrors,omitempty"`
	TlvDiscards   int64 `xml:"tlv-discards,omitempty" json:"TlvDiscards,omitempty"`
	TlvUnknown    int64 `xml:"tlv-unknown,omitempty" json:"TlvUnknown,omitempty"`
}

type TxMgmtAddress struct {
	MgmtAddressSystemType string `xml:"mgmt-address-system-type,omitempty" json:"MgmtAddressSystemType,omitempty"`
	MgmtAddress           string `xml:"mgmt-address,omitempty" json:"MgmtAddress,omitempty"`
	MgmtAddressSubtype    string `xml:"mgmt-address-subtype,omitempty" json:"MgmtAddressSubtype,omitempty"`
}

type Performance struct {
	Intervals15min *Intervals15min `xml:"intervals-15min,omitempty" json:"Intervals15min,omitempty"`
}

type Intervals15min struct {
	Current *Current `xml:"current,omitempty" json:"Current,omitempty"`
	History *History `xml:"history,omitempty" json:"History,omitempty"`
}

type Current struct {
	ValidIntervals             int64 `xml:"valid-intervals,omitempty" json:"ValidIntervals,omitempty"`
	InBandwidth                int64 `xml:"in-bandwidth,omitempty" json:"InBandwidth,omitempty"`
	OutBandwidth               int64 `xml:"out-bandwidth,omitempty" json:"OutBandwidth,omitempty"`
	RxCrcAlignErrors           int64 `xml:"rx-crc-align-errors,omitempty" json:"RxCrcAlignErrors,omitempty"`
	TxCrcAlignErrors           int64 `xml:"tx-crc-align-errors,omitempty" json:"TxCrcAlignErrors,omitempty"`
	TxCollision                int64 `xml:"tx-collision,omitempty" json:"TxCollision,omitempty"`
	InOctets                   int64 `xml:"in-octets,omitempty" json:"InOctets,omitempty"`
	InPkts                     int64 `xml:"in-pkts,omitempty" json:"InPkts,omitempty"`
	InPktsDrop                 int64 `xml:"in-pkts-drop,omitempty" json:"InPktsDrop,omitempty"`
	InDiscards                 int64 `xml:"in-discards,omitempty" json:"InDiscards,omitempty"`
	InErrors                   int64 `xml:"in-errors,omitempty" json:"InErrors,omitempty"`
	InUnknownProtocolDiscards  int64 `xml:"in-unknown-protocol-discards,omitempty" json:"InUnknownProtocolDiscards,omitempty"`
	OutOctets                  int64 `xml:"out-octets,omitempty" json:"OutOctets,omitempty"`
	OutPkts                    int64 `xml:"out-pkts,omitempty" json:"OutPkts,omitempty"`
	OutPktsDrop                int64 `xml:"out-pkts-drop,omitempty" json:"OutPktsDrop,omitempty"`
	OutDiscards                int64 `xml:"out-discards,omitempty" json:"OutDiscards,omitempty"`
	OutUnknownProtocolDiscards int64 `xml:"out-unknown-protocol-discards,omitempty" json:"OutUnknownProtocolDiscards,omitempty"`
	OutErrors                  int64 `xml:"out-errors,omitempty" json:"OutErrors,omitempty"`
}

type History struct {
	IntervalNumber             int64  `xml:"interval-number,omitempty" json:"IntervalNumber,omitempty"`
	InBandwidth                int64  `xml:"in-bandwidth,omitempty" json:"InBandwidth,omitempty"`
	OutBandwidth               int64  `xml:"out-bandwidth,omitempty" json:"OutBandwidth,omitempty"`
	RxCrcAlignErrors           int64  `xml:"rx-crc-align-errors,omitempty" json:"RxCrcAlignErrors,omitempty"`
	TxCrcAlignErrors           int64  `xml:"tx-crc-align-errors,omitempty" json:"TxCrcAlignErrors,omitempty"`
	TxCollision                int64  `xml:"tx-collision,omitempty" json:"TxCollision,omitempty"`
	MeasuredTime               int64  `xml:"measured-time,omitempty" json:"MeasuredTime,omitempty"`
	InvalidDataFlag            bool   `xml:"invalid-data-flag,omitempty" json:"InvalidDataFlag,omitempty"`
	TimeStamp                  string `xml:"time-stamp,omitempty" json:"TimeStamp,omitempty"`
	InOctets                   int64  `xml:"in-octets,omitempty" json:"InOctets,omitempty"`
	InPkts                     int64  `xml:"in-pkts,omitempty" json:"InPkts,omitempty"`
	InPktsDrop                 int64  `xml:"in-pkts-drop,omitempty" json:"InPktsDrop,omitempty"`
	InDiscards                 int64  `xml:"in-discards,omitempty" json:"InDiscards,omitempty"`
	InErrors                   int64  `xml:"in-errors,omitempty" json:"InErrors,omitempty"`
	InUnknownProtocolDiscards  int64  `xml:"in-unknown-protocol-discards,omitempty" json:"InUnknownProtocolDiscards,omitempty"`
	OutOctets                  int64  `xml:"out-octets,omitempty" json:"OutOctets,omitempty"`
	OutPkts                    int64  `xml:"out-pkts,omitempty" json:"OutPkts,omitempty"`
	OutPktsDrop                int64  `xml:"out-pkts-drop,omitempty" json:"OutPktsDrop,omitempty"`
	OutDiscards                int64  `xml:"out-discards,omitempty" json:"OutDiscards,omitempty"`
	OutUnknownProtocolDiscards int64  `xml:"out-unknown-protocol-discards,omitempty" json:"OutUnknownProtocolDiscards,omitempty"`
	OutErrors                  int64  `xml:"out-errors,omitempty" json:"OutErrors,omitempty"`
}

type Transceiver struct {
	ModelNumber                 string                       `xml:"model-number"`
	OperState                   string                       `xml:"oper-state"`
	SffEquipped                 string                       `xml:"sff-equipped"`
	VendorManufactureDate       string                       `xml:"vendor-manufacture-date"`
	VendorPartNumber            string                       `xml:"vendor-part-number"`
	VendorSerialNumber          string                       `xml:"vendor-serial-number"`
	DigitalDiagnosticMonitoring *DigitalDiagnosticMonitoring `xml:"digital-diagnostic-monitoring"`
}

type DigitalDiagnosticMonitoring struct {
	Temperature          *MonitoringValue `xml:"temperature"`
	TransmitBiasCurrent  *MonitoringValue `xml:"transmit-bias-current"`
	TransmitOutputPower  *MonitoringValue `xml:"transmit-output-power"`
	ReceivedOpticalPower *MonitoringValue `xml:"received-optical-power"`
	SupplyVoltage        *MonitoringValue `xml:"supply-voltage"`
}

type MonitoringValue struct {
	Current string `xml:"current"`
}

func NewPort() *Port {
	return &Port{}
}

func NewPorts() []Port {
	return make([]Port, 0)
}
