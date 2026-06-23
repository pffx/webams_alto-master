package rpcSystemModels

//Demo   <onu-name xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-if-xponvani-aug">C12_ONT2</onu-name>

type C_TagWithAttrAndValue struct {
	Name            string `xml:",chardata"`
	Xmlns           string `xml:"xmlns,attr,omitempty"`
	XmlnsBbfXPonift string `xml:"xmlns:bbf-xponift,attr,omitempty"`
	XmlnsBbfHwt     string `xml:"xmlns:bbf-hwt,attr,omitempty"`
	XmlnsNokiaHwi   string `xml:"xmlns:nokia-hwi,attr,omitempty"`
	Operation       string `xml:"operation,omitempty,attr"`
}

func NewC_TagWithAttrAndValue() *C_TagWithAttrAndValue {
	a := new(C_TagWithAttrAndValue)
	return a
}

//system/radius
type Radius struct {
	Server           []Server          `xml:"server,omitempty"`
	OperatorAuth     *OperatorAuth     `xml:"operator-auth,omitempty"`
	Policy           []Policy          `xml:"policy,omitempty"`
	Domain           *Domain           `xml:"domain,omitempty"`
	ConnectionPolicy *ConnectionPolicy `xml:"connection-policy,omitempty"`
}

func NewRadius() *Radius {
	a := new(Radius)
	return a
}

//system/radius/server
type Server struct {
	Name      string `xml:"name,omitempty"`
	Operation string `xml:"operation,omitempty,attr"`
	Udp       *Udp   `xml:"udp,omitempty"`
}

func NewServer() *Server {
	a := new(Server)
	return a
}

//system/radius/server/udp
type Udp struct {
	Address            string `xml:"address,omitempty"`
	AuthenticationPort string `xml:"authentication-port,omitempty"`
	SharedSecret       string `xml:"shared-secret,omitempty"`
}

func NewUdp() *Udp {
	a := new(Udp)
	return a
}

//only in NT
//system/radius/operator-auth
type OperatorAuth struct {
	Xmlns  string   `xml:"xmlns,attr,omitempty"`
	Policy []Policy `xml:"policy,omitempty"`
	// CliSessionPolicy string   `xml:"cli-session-policy,omitempty"`
	CliSessionPolicy *C_TagWithAttrAndValue `xml:"cli-session-policy,omitempty"`
}

func NewOperatorAuth() *OperatorAuth {
	a := new(OperatorAuth)
	a.Xmlns = "http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-radius-operator-auth"
	return a
}

//system/radius/operator-auth/policy
type Policy struct {
	Name                   string `xml:"name,omitempty"`
	Xmlns                  string `xml:"xmlns,attr,omitempty"`
	Operation              string `xml:"operation,omitempty,attr"`
	NasId                  string `xml:"nas-id,omitempty"`
	NasPortIdSyntax        string `xml:"nas-port-id-syntax,omitempty"`
	NasIpAddress           string `xml:"nas-ip-address,omitempty"`
	CallingStationIdFormat string `xml:"calling-station-id-format,omitempty"`
	AuthServerFirst        string `xml:"auth-server-first,omitempty"`
	AuthServerSecond       string `xml:"auth-server-second,omitempty"`
	AcctServerFirst        string `xml:"acct-server-first,omitempty"`
	AcctServerSecond       string `xml:"acct-server-second,omitempty"`
	KeepDomainName         string `xml:"keep-domain-name,omitempty"`
	DisableEap             string `xml:"disable-eap,omitempty"`
	AcctInterval           string `xml:"acct-interval,omitempty"`
	DisableAccounting      string `xml:"disable-accounting,omitempty"`
}

func NewPolicy() *Policy {
	a := new(Policy)
	return a
}

//system/radius/operator-auth/domain
type Domain struct {
	Xmlns         string `xml:"xmlns,attr,omitempty"`
	Name          string `xml:"name,omitempty"`
	Operation     string `xml:"operation,omitempty,attr"`
	Authenticator string `xml:"authenticator,omitempty"`
}

func NewDomain() *Domain {
	a := new(Domain)
	a.Xmlns = "urn:aul:params:xml:ns:yang:nokia-radius"
	return a
}

//system/radius/operator-auth/connection-policy
type ConnectionPolicy struct {
	Xmlns              string `xml:"xmlns,attr,omitempty"`
	DomainName         string `xml:"domain-name,omitempty"`
	Operation          string `xml:"operation,omitempty,attr"`
	RejectNoDomain     string `xml:"reject-no-domain,omitempty"`
	RejectInvDomain    string `xml:"reject-inv-domain,omitempty"`
	AcceptOrRejectAll  string `xml:"accept-or-reject-all,omitempty"`
	AccountingOnReboot string `xml:"accounting-on-reboot,omitempty"`
}

func NewConnectionPolicy() *ConnectionPolicy {
	a := new(ConnectionPolicy)
	a.Xmlns = "urn:aul:params:xml:ns:yang:nokia-radius"
	return a
}
