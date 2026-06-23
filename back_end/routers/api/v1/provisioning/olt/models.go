package olt

import "encoding/xml"

type CurrentState struct {
	XMLName   xml.Name `xml:"current-state"`
	State     string   `xml:"state" json:"State"`
	Timestamp string   `xml:"timestamp" json:"Timestamp"`
}
type LastDownloadState struct {
	XMLName      xml.Name `xml:"last-download-state"`
	State        string   `xml:"state" json:"State"`
	Timestamp    string   `xml:"timestamp" json:"Timestamp"`
	SoftwareName string   `xml:"software-name" json:"SoftwareName"`
}

type Download struct {
	XMLName           xml.Name          `xml:"download"`
	CurrentState      CurrentState      `json:"CurrentState"`
	LastDownloadState LastDownloadState `json:"LastDownloadState"`
}

type Software2 struct {
	XMLName   xml.Name  `xml:"software"`
	Name      string    `xml:"name" json:"Name"`
	Download  Download  `json:"Download"`
	Revisions Revisions `json:"Revisions"`
}

type Software1 struct {
	XMLName  xml.Name  `xml:"software"`
	Software Software2 `json:"Software"`
}

type Component struct {
	XMLName  xml.Name  `xml:"component"`
	Name     string    `xml:"name" json:"Name"`
	Software Software1 `json:"Software"`
}
type HardwareState struct {
	XMLName   xml.Name  `xml:"hardware-state"`
	Component Component `json:"Component"`
}
type SoftwareVersionData struct {
	XMLName       xml.Name      `xml:"data"`
	HardwareState HardwareState `json:"HardwareState"`
}

type Revision struct {
	XMLName           xml.Name `xml:"revision"`
	Name              string   `xml:"name" json:""`
	DownloadTimestamp string   `xml:"download-timestamp" json:""`
	Version           string   `xml:"version" json:""`
	IsValid           string   `xml:"is-valid" json:""`
	IsCommitted       string   `xml:"is-committed" json:""`
	IsActive          string   `xml:"is-active" json:""`
}
type Revisions struct {
	XMLName  xml.Name   `xml:"revisions"`
	Revision []Revision `xml:"revision" json:""`
}
