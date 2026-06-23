package rpcAlarmsModels

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
//alarms/summary
type Summary struct {
	AlarmSummary []AlarmSummary `xml:"alarm-summary,omitempty"`
}

func NewSummary() *Summary {
	a := new(Summary)
	return a
}

//alarms/alarm-list
type AlarmList struct {
	Alarm          []Alarm `xml:"alarm,omitempty"`
	NumberOfAlarms string  `xml:"number-of-alarms,omitempty"`
	LastChanged    string  `xml:"last-changed,omitempty"`
}

func NewAlarmList() *AlarmList {
	a := new(AlarmList)
	return a
}

//alarms/alarm-list/alarm
type Alarm struct {
	Resource          string `xml:"resource,omitempty"`
	AlarmTypeId       string `xml:"alarm-type-id,omitempty"`
	TimeCreated       string `xml:"time-created,omitempty"`
	IsCleared         string `xml:"is-cleared,omitempty"`
	LastChanged       string `xml:"last-changed,omitempty"`
	PerceivedSeverity string `xml:"perceived-severity,omitempty"`
	AlarmText         string `xml:"alarm-text,omitempty"`
	// used for ihub yang
	LastStatusChanged     string `xml:"last-status-change,omitempty"`
	AlarmIdentity         string `xml:"alarm-identity,omitempty"`
	LastPerceivedSeverity string `xml:"last-perceived-severity,omitempty"`
	LastAlarmText         string `xml:"last-alarm-text,omitempty"`
}

func NewAlarm() *Alarm {
	a := new(Alarm)
	a.IsCleared = "false"
	return a
}
func NewAlarmIhub() *Alarm {
	a := new(Alarm)
	return a
}

// used for ihub yang
//alarms/summary/alarm-summary
type AlarmSummary struct {
	Sseverity string `xml:"severity,omitempty"`
	Total     int    `xml:"total,omitempty"`
}

func NewAlarmSummary() *AlarmSummary {
	a := new(AlarmSummary)
	return a
}

// ********************************forwarding part end****************************************
//****************************************************************************************
