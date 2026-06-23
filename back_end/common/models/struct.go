package models

import (
	"encoding/xml"
	"net"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

type Olt struct {
	Id       int
	Ip       string
	UserName string
	Passwd   string
	Type     string
	LtNum    int
	Lt1      string
	Lt2      string
	Lt3      string
	Lt4      string
	Lt5      string
	Lt6      string
	Lt7      string
	Lt8      string
	Lt9      string
	Lt10     string
	Lt11     string
	Lt12     string
	Lt13     string
	Lt14     string
}

//export by default
type UserInfo struct {
	Account    string
	Password   string
	Role       string
	Department string
}
type OltParamInfo struct {
	OltId   string `json:"oltId" form:"oltId"`
	OltPort string `json:"oltPort" form:"oltPort"`
}

type OltInfo struct {
	IP       string `form:"IP" json:"IP" binding:"required"`
	Username string `form:"Username" json:"Username" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	// Controlled bool
	Status   string
	HostName string
	Software string
	OltType  string
	OltLtNum int
}
type LTInfoStruct struct {
	Planned string
}
type OltInfoWithLTCard struct {
	IP       string `form:"IP" json:"IP" binding:"required"`
	Username string `form:"Username" json:"Username" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	// Controlled bool
	Status   string
	HostName string
	Software string
	OltType  string
	OltLtNum int
	Lt1Info  string
	Lt2Info  string
	Lt3Info  string
	Lt4Info  string
	Lt5Info  string
	Lt6Info  string
	Lt7Info  string
	Lt8Info  string
	Lt9Info  string
	Lt10Info string
	Lt11Info string
	Lt12Info string
	Lt13Info string
	Lt14Info string
}
type IpInfo struct {
	IpAddr  string
	NetMask string
	GateWay string
}
type ComponentForPlanLT struct {
	XMLName   xml.Name `xml:"component"`
	Name      string   `xml:"name" json:"Name"`
	ModelName string   `xml:"model-name" json:"ModelName"`
}
type HardwareStateForPlanLT struct {
	XMLName   xml.Name           `xml:"hardware-state"`
	Component ComponentForPlanLT `json:"Component"`
}
type LtBoardInfoData struct {
	XMLName       xml.Name               `xml:"data"`
	HardwareState HardwareStateForPlanLT `json:"HardwareState"`
}
type SoftwareDownloadInfo struct {
	SoftwareName string
	SoftwareUrl  string
}
type ltPlannedInfo struct {
	ltName    string
	ltParent  string
	modelName string
}
type SoftwareDownloadName struct {
	SoftwareName string
}
type LtIndex struct {
	LtIndex string
}
type PlanLtInfo struct {
	LtIndex     string
	LtIndexAdd8 string
	ModelName   string
}
type RpcError struct {
	ErrorType     string `xml:"error-type"`
	ErrorTag      string `xml:"error-tag"`
	ErrorSeverity string `xml:"error-severity"`
	ErrorAppTag   string `xml:"error-app-tag"`
	ErrorMessage  string `xml:"error-message"`
	ErrorInfo     string `xml:"error-info"`
}

type OltServiceInfo struct {
	Name        string
	Version     string
	Description string
	Status      string
	ID          string
}

type Cli struct {
	User       string
	Pwd        string
	Addr       string
	Client     *gossh.Client
	Session    *gossh.Session
	LastResult string
}

// 连接对象
func (c *Cli) Connect() (*Cli, error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = c.User
	config.Auth = []gossh.AuthMethod{gossh.Password(c.Pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key gossh.PublicKey) error { return nil }
	config.Timeout = 3 * time.Second
	client, err := gossh.Dial("tcp", c.Addr, config)
	if nil != err {
		return c, err
	}
	c.Client = client
	return c, nil
}

// 执行shell
func (c Cli) Run(shell string) (string, error) {
	if c.Client == nil {
		if _, err := c.Connect(); err != nil {
			return "", err
		}
	}
	session, err := c.Client.NewSession()
	if err != nil {
		return "", err
	}
	// 关闭会话
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

func NewOltInfo() *OltInfo {
	a := new(OltInfo)
	return a
}

func NewIpInfo() *IpInfo {
	a := new(IpInfo)
	return a
}

func NewRpcError() *RpcError {
	a := new(RpcError)
	return a
}

func NewOltServiceInfo() *OltServiceInfo {
	a := new(OltServiceInfo)
	return a
}

func NewUserInfo() *UserInfo {
	a := new(UserInfo)
	return a
}
