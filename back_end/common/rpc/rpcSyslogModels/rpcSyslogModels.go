package rpcSyslogModels

import "alto_server/constants"

//syslog/actions
type Actions struct {
	Remote *Remote `xml:"remote,omitempty"`
}

func NewActions() *Actions {
	a := new(Actions)
	return a
}

//syslog/actions/remote
type Remote struct {
	Destination []Destination `xml:"destination,omitempty"`
}

func NewRemote() *Remote {
	a := new(Remote)
	return a
}

//syslog/actions/remote/destination
type Destination struct {
	Operation string `xml:"operation,omitempty,attr"` // may be should be added automaticlly
	Name      string `xml:"name,omitempty"`
	Udp       *Udp   `xml:"udp,omitempty"`
}

func NewDestination() *Destination {
	a := new(Destination)
	return a
}
func NewDelDestination() *Destination {
	a := new(Destination)
	a.Operation = constants.OPERATION_DELETE
	return a
}

//syslog/actions/remote/destination/udp
type Udp struct {
	Address string `xml:"address,omitempty"`
	Port    string `xml:"port,omitempty"`
}

func NewUdp() *Udp {
	a := new(Udp)
	return a
}
