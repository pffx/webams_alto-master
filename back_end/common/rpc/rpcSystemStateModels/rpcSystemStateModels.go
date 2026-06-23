package rpcSystemStateModels

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

//system-state/platform
type Platform struct {
	SoftwareRelease *C_TagWithAttrAndValue `xml:"software-release,omitempty"`
}

func NewPlatform() *Platform {
	a := new(Platform)
	return a
}

//system-state/clock
type Clock struct {
	CurrentDatetime string `xml:"current-datetime,omitempty"`
	BootDatetime    string `xml:"boot-datetime,omitempty"`
	SysUpTime       string `xml:"sys-up-time,omitempty"`
}

func NewClock() *Clock {
	a := new(Clock)
	return a
}

func NewSoftwareRelease() *C_TagWithAttrAndValue {
	a := new(C_TagWithAttrAndValue)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ietf-system-aug"
	return a
}

//system-state/radius-statistics
type RadiusStatistics struct {
	AccountingServer     []AccountingServer     `xml:"accounting-server,omitempty"`
	AuthenticationServer []AuthenticationServer `xml:"authentication-server,omitempty"`
}

func NewRadiusStatistics() *RadiusStatistics {
	a := new(RadiusStatistics)
	return a
}

//system-state/radius-statistics/accounting-server
type AccountingServer struct {
	Name                        string `xml:"name,omitempty"`
	AuthState                   string `xml:"auth-state,omitempty"`
	ServerReplyTime             string `xml:"server-reply-time,omitempty"`
	AccountingRequestTx         string `xml:"accounting-request-tx,omitempty"`
	AccountingRetransmissionTx  string `xml:"accounting-retransmission-tx,omitempty"`
	AccountingResponseRx        string `xml:"accounting-response-rx,omitempty"`
	InvalidAccountingResponseRx string `xml:"invalid-accounting-response-rx,omitempty"`
	BadAuthenticatorRx          string `xml:"bad-authenticator-rx,omitempty"`
	PendingRequestRx            string `xml:"pending-request-rx,omitempty"`
	Timeouts                    string `xml:"timeouts,omitempty"`
	UnknownPacketRx             string `xml:"unknown-packet-rx,omitempty"`
	PacketDiscardsRx            string `xml:"packet-discards-rx,omitempty"`
}

func NewAccountingServer() *AccountingServer {
	a := new(AccountingServer)
	return a
}

//system-state/radius-statistics/authentication-server
type AuthenticationServer struct {
	Name                    string `xml:"name,omitempty"`
	AuthState               string `xml:"auth-state,omitempty"`
	ServerReplyTime         string `xml:"server-reply-time,omitempty"`
	AccessRequestTx         string `xml:"access-request-tx,omitempty"`
	AccessRetransmissionTx  string `xml:"access-retransmission-tx,omitempty"`
	AccessAcceptRx          string `xml:"access-accept-rx,omitempty"`
	AccessRejectRx          string `xml:"access-reject-rx,omitempty"`
	AccessChallengeRx       string `xml:"access-challenge-rx,omitempty"`
	InvalidAccessResponseRx string `xml:"invalid-access-response-rx,omitempty"`
	BadAuthenticatorRx      string `xml:"bad-authenticator-rx,omitempty"`
	PendingRequestRx        string `xml:"pending-request-rx,omitempty"`
	Timeouts                string `xml:"timeouts,omitempty"`
	UnknownPacketRx         string `xml:"unknown-packet-rx,omitempty"`
	PacketDiscardsRx        string `xml:"packet-discards-rx,omitempty"`
}

func NewAuthenticationServerr() *AuthenticationServer {
	a := new(AuthenticationServer)
	return a
}
