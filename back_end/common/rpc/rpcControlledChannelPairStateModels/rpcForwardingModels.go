package rpcControlledChannelPairStateModels

//controlled-channel-pair-state/controlled-channel-pair
type ControlledChannelPair struct {
	Name                         string                     `xml:"name,omitempty"`
	TypeBState                   string                     `xml:"type-b-state,omitempty"`
	LastSwitchoverReason         string                     `xml:"last-switchover-reason,omitempty"`
	LastSwitchoverDetailedReason string                     `xml:"last-switchover-detailed-reason,omitempty"`
	LastSwitchoverTimeStamp      string                     `xml:"last-switchover-time-stamp,omitempty"`
	NumberOfSwitchovers          string                     `xml:"number-of-switchovers,omitempty"`
	NumberOfSwitchoversFailures  string                     `xml:"number-of-switchovers-failures,omitempty"`
	MemberChannelTermination     []MemberChannelTermination `xml:"member-channel-termination,omitempty"`
}

//controlled-channel-pair-state/controlled-channel-pair/member-channel-termination
type MemberChannelTermination struct {
	OltDeviceName                   string `xml:"olt-device-name,omitempty"`
	ChannelTerminationYangNameInOlt string `xml:"channel-termination-yang-name-in-olt,omitempty"`
	ServiceState                    string `xml:"service-state,omitempty"`
	TypeBRole                       string `xml:"type-b-role,omitempty"`
}

func NewControlledChannelPair() *ControlledChannelPair {
	a := new(ControlledChannelPair)
	return a
}

func NewMemberChannelTermination() *MemberChannelTermination {
	a := new(MemberChannelTermination)
	return a
}
