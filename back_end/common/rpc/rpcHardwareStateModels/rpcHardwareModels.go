package rpcHardwareStateModels

//Demo   <onu-name xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-if-xponvani-aug">C12_ONT2</onu-name>

type C_TagWithAttrAndValue struct {
	Name            string `xml:",chardata"`
	Xmlns           string `xml:"xmlns,attr,omitempty"`
	XmlnsBbfXPonift string `xml:"xmlns:bbf-xponift,attr,omitempty"`
	XmlnsBbfHwt     string `xml:"xmlns:bbf-hwt,attr,omitempty"`
	XmlnsNokiaHwi   string `xml:"xmlns:nokia-hwi,attr,omitempty"`
	XmlnsIanahw     string `xml:"xmlns:ianahw,attr,omitempty"`
}

func NewC_TagWithAttrAndValue() *C_TagWithAttrAndValue {
	a := new(C_TagWithAttrAndValue)
	return a
}

// hardware-state/component
type Component struct {
	Name               string                 `xml:"name" json:"Name"`
	Parent             string                 `xml:"parent,omitempty"`
	Description        string                 `xml:"description,omitempty"`
	SerialNumber       string                 `xml:"serial-num,omitempty"`
	ModelName          string                 `xml:"model-name,omitempty"`
	IsFru              bool                   `xml:"is-fru,omitempty"`
	AdminState         string                 `xml:"admin-state,omitempty"`
	ManufacturerName   string                 `xml:"mfg-name,omitempty"`
	ManufacturerDate   string                 `xml:"mfg-date,omitempty"`
	HardwareRev        string                 `xml:"hardware-rev,omitempty"`
	CleiCode           string                 `xml:"clei-code,omitempty"`
	ContainsChild      []string               `xml:"contains-child,omitempty"`
	Class              *C_TagWithAttrAndValue `xml:"class,omitempty"`
	State              *StateComponent        `xml:"state,omitempty"`
	SoftwaresOlt       *SoftwaresOlt          `xml:"software,omitempty"`
	Transceiver        *Transceiver           `xml:"transceiver,omitempty"`
	TransceiverLink    *TransceiverLink       `xml:"transceiver-link,omitempty"`
	Diagnostics        *Diagnostics           `xml:"diagnostics,omitempty"`
	SensorData         *SensorData            `xml:"sensor-data,omitempty"`
	CpuProcessorData   *CpuProcessorData      `xml:"cpu-processor-data,omitempty"`
	VolumeResourceData *VolumeResourceData    `xml:"volume-resource-data,omitempty"`
}

func NewComponent() *Component {
	a := new(Component)
	return a
}

// /hardware-state/component/sensor-data
type SensorData struct {
	Value string `xml:"value,omitempty"`
}

// /hardware-state/component/volume-resource-data
type VolumeResourceData struct {
	VolumeResourceList []VolumeResourceList `xml:"volume-resource-list,omitempty"`
}

// /hardware-state/component/volume-resource-data/volume-resource-list
type VolumeResourceList struct {
	Name string `xml:"name,omitempty"`
	Size string `xml:"size,omitempty"`
	Free string `xml:"free,omitempty"`
}

// /hardware-state/component/cpu-processor-data
type CpuProcessorData struct {
	Xmlns                  string           `xml:"xmlns,attr,omitempty"`
	NumberOfActiveSessions string           `xml:"number-of-active-sessions,omitempty"`
	SystemLoad             *SystemLoad      `xml:"system-load,omitempty"`
	TaskCounts             *TaskCounts      `xml:"task-counts,omitempty"`
	PercentCpuUsage        *PercentCpuUsage `xml:"percent-cpu-usage,omitempty"`
	MemoryUsage            *MemoryUsage     `xml:"memory-usage,omitempty"`
}

func NewCpuProcessorData() *CpuProcessorData {
	a := new(CpuProcessorData)
	a.Xmlns = "urn:bbf:yang:bbf-hardware-cpu-resource"
	return a
}

// /hardware-state/component/cpu-processor-data/system-load
type SystemLoad struct {
	AverageSystemLoad1Min  string `xml:"average-system-load-1-min,omitempty"`
	AverageSystemLoad5Min  string `xml:"average-system-load-5-min,omitempty"`
	AverageSystemLoad15Min string `xml:"average-system-load-15-min,omitempty"`
}

// /hardware-state/component/cpu-processor-data/task-counts
type TaskCounts struct {
	TotalTasks    string `xml:"total-tasks,omitempty"`
	RunningTasks  string `xml:"running-tasks,omitempty"`
	SleepingTasks string `xml:"sleeping-tasks,omitempty"`
	StoppedTasks  string `xml:"stopped-tasks,omitempty"`
	ZombieTasks   string `xml:"zombie-tasks,omitempty"`
}

// /hardware-state/component/cpu-processor-data/percent-cpu-usage
type PercentCpuUsage struct {
	PercentCpuCoreProcesses string `xml:"percent-cpu-core-processes,omitempty"`
	PercentCpuUser          string `xml:"percent-cpu-user,omitempty"`
	PercentCpuIdle          string `xml:"percent-cpu-idle,omitempty"`
	PercentCpuHwio          string `xml:"percent-cpu-hwio,omitempty"`
	PercentCpuIo            string `xml:"percent-cpu-io,omitempty"`
	PercentCpuNice          string `xml:"percent-cpu-nice,omitempty"`
	PercentCpuSwint         string `xml:"percent-cpu-swint,omitempty"`
}

// /hardware-state/component/cpu-processor-data/memory-usage
type MemoryUsage struct {
	UsedMem         string `xml:"used-mem,omitempty"`
	TotalMemory     string `xml:"total-memory,omitempty"`
	FreeMemory      string `xml:"free-memory,omitempty"`
	BufferMemory    string `xml:"buffer-memory,omitempty"`
	TotalSwapMemory string `xml:"total-swap-memory,omitempty"`
	UsedSwap        string `xml:"used-swap,omitempty"`
	FreeSwap        string `xml:"free-swap,omitempty"`
	AvailableMem    string `xml:"available-mem,omitempty"`
}

// /hardware-state/component/state
type StateComponent struct {
	AdminState       string `xml:"admin-state,omitempty"`
	OperState        string `xml:"oper-state,omitempty"`
	StandbyState     string `xml:"standby-state,omitempty"`
	StateLastChanged string `xml:"state-last-changed,omitempty"`
}

func NewStateComponent() *StateComponent {
	a := new(StateComponent)
	return a
}

// /hardware-state/component/transceiver
type Transceiver struct {
	Xmlns       string       `xml:"xmlns,attr,omitempty"`
	Diagnostics *Diagnostics `xml:"diagnostics,omitempty"`
}

func NewTransceiver() *Transceiver {
	a := new(Transceiver)
	a.Xmlns = "urn:bbf:yang:bbf-hardware-transceivers"
	return a
}

// /hardware-state/component/transceiver-link
type TransceiverLink struct {
	Xmlns       string       `xml:"xmlns,attr,omitempty"`
	WaveLength  string       `xml:"wavelength,omitempty"`
	Diagnostics *Diagnostics `xml:"diagnostics"`
}

func NewTransceiverLink() *TransceiverLink {
	a := new(TransceiverLink)
	a.Xmlns = "urn:bbf:yang:bbf-hardware-transceivers"
	return a
}

// /hardware-state/component/transceiver/diagnostics
// /hardware-state/component/transceiver-link/diagnostics
type Diagnostics struct {
	TxBias        string    `xml:"tx-bias,omitempty"`
	TxPowerDbm    string    `xml:"tx-power-dbm,omitempty"`
	RxPowerDbm    string    `xml:"rx-power-dbm,omitempty"`
	Temperature   int       `xml:"temperature,omitempty"`
	SupplyVoltage int       `xml:"supply-voltage,omitempty"`
	RssiOnu       []RssiOnu `xml:"rssi-onu,omitempty"`
}

func NewDiagnostics() *Diagnostics {
	a := new(Diagnostics)
	return a
}

// /hardware-state/component/transceiver-link/diagnostics/rssi-onu
type RssiOnu struct {
	DetectedSerialNumber string `xml:"detected-serial-number,omitempty"`
	Rssi                 string `xml:"rssi,omitempty"`
	VAniRef              string `xml:"v-ani-ref,omitempty"`
}

// hardware-state/component/software
//used by OLT software and ont software
type SoftwaresOlt struct {
	Xmlns       string        `xml:"xmlns,attr,omitempty"`
	SoftwareOlt []SoftwareOlt `xml:"software,omitempty"`
}

func NewSoftwaresOlt() *SoftwaresOlt {
	a := new(SoftwaresOlt)
	a.Xmlns = "urn:bbf:yang:bbf-software-image-management-one-dot-one"
	return a
}

// hardware-state/component/software/software
type SoftwareOlt struct {
	Name           string          `xml:"name,omitempty"`
	Revisions      *Revisions      `xml:"revisions,omitempty"`
	Download       *Download       `xml:"download,omitempty"`
	ConfigDownload *ConfigDownload `xml:"config-download,omitempty"`
}

func NewSoftwareOlt() *SoftwareOlt {
	a := new(SoftwareOlt)
	return a
}

// hardware-state/component/software/software/download
type Download struct {
	CurrentState      *State `xml:"current-state,omitempty"`
	LastDownloadState *State `xml:"last-download-state,omitempty"`
}
type ConfigDownload struct {
	CurrentState      *State `xml:"current-state,omitempty"`
	LastDownloadState *State `xml:"last-download-state,omitempty"`
}

// hardware-state/component/software/software/download/current-state
// hardware-state/component/software/software/download/last-download-state
type State struct {
	State        string   `xml:"state,omitempty"`
	Timestamp    string   `xml:"timestamp,omitempty"`
	SoftwareName string   `xml:"software-name,omitempty"`
	Failure      *Failure `xml:"failure,omitempty"`
}

// hardware-state/component/software/software/download/current-state/failure
// hardware-state/component/software/software/download/last-download-state/failure
type Failure struct {
	FailureReason string `xml:"failure-reason,omitempty"`
	FailureString string `xml:"failure-string,omitempty"`
}

// hardware-state/component/software/software/revisions
type Revisions struct {
	Revision []Revision `xml:"revision,omitempty"`
}

// hardware-state/component/software/software/revisions/revision
type Revision struct {
	Name              string `xml:"name,omitempty"`
	DownloadTimestamp string `xml:"download-timestamp,omitempty"`
	Version           string `xml:"version,omitempty"`
	IsValid           string `xml:"is-valid,omitempty"`
	IsCommitted       string `xml:"is-committed,omitempty"`
	IsActive          string `xml:"is-active,omitempty"`
	Source            string `xml:"source,omitempty"`
	ProductCode       string `xml:"product-code,omitempty"`
}
