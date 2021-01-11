package pfcp

import (
	"github.com/free5gc/pfcp/pfcpType"
)

type Message struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

func (message *Message) IsRequest() (IsRequest bool) {

	switch message.Header.MessageType {
	case PFCP_HEARTBEAT_REQUEST:
		IsRequest = true
	case PFCP_PFD_MANAGEMENT_REQUEST:
		IsRequest = true
	case PFCP_ASSOCIATION_SETUP_REQUEST:
		IsRequest = true
	case PFCP_ASSOCIATION_UPDATE_REQUEST:
		IsRequest = true
	case PFCP_ASSOCIATION_RELEASE_REQUEST:
		IsRequest = true
	case PFCP_NODE_REPORT_REQUEST:
		IsRequest = true
	case PFCP_SESSION_SET_DELETION_REQUEST:
		IsRequest = true
	case PFCP_SESSION_ESTABLISHMENT_REQUEST:
		IsRequest = true
	case PFCP_SESSION_MODIFICATION_REQUEST:
		IsRequest = true
	case PFCP_SESSION_DELETION_REQUEST:
		IsRequest = true
	case PFCP_SESSION_REPORT_REQUEST:
		IsRequest = true
	default:
		IsRequest = false
	}

	return
}

func (message *Message) IsResponse() (IsResponse bool) {
	IsResponse = false
	switch message.Header.MessageType {
	case PFCP_HEARTBEAT_RESPONSE:
		IsResponse = true
	case PFCP_PFD_MANAGEMENT_RESPONSE:
		IsResponse = true
	case PFCP_ASSOCIATION_SETUP_RESPONSE:
		IsResponse = true
	case PFCP_ASSOCIATION_UPDATE_RESPONSE:
		IsResponse = true
	case PFCP_ASSOCIATION_RELEASE_RESPONSE:
		IsResponse = true
	case PFCP_NODE_REPORT_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_SET_DELETION_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_ESTABLISHMENT_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_MODIFICATION_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_DELETION_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_REPORT_RESPONSE:
		IsResponse = true
	default:
		IsResponse = false
	}

	return
}

type HeartbeatRequest struct {
	RecoveryTimeStamp *pfcpType.RecoveryTimeStamp `tlv:"96" json:"recoveryTimeStamp"`
}

type HeartbeatResponse struct {
	RecoveryTimeStamp *pfcpType.RecoveryTimeStamp `tlv:"96" json:"recoveryTimeStamp"`
}

type PFCPPFDManagementRequest struct {
	ApplicationIDsPFDs []ApplicationIDsPFDs `tlv:"58" json:"applicationIDsPFDs"`
}

type ApplicationIDsPFDs struct {
	ApplicationID pfcpType.ApplicationID `tlv:"24" json:"applicationID"`
	PFD           *PFD                   `tlv:"59" json:"pfd"`
}

type PFD struct {
	PFDContents []pfcpType.PFDContents `tlv:"61" json:"pfdContents"`
}

type PFCPPFDManagementResponse struct {
	Cause       *pfcpType.Cause       `tlv:"19" json:"cause"`
	OffendingIE *pfcpType.OffendingIE `tlv:"40" json:"offendingIE"`
}

type PFCPAssociationSetupRequest struct {
	NodeID                         *pfcpType.NodeID                         `tlv:"60" json:"nodeID"`
	RecoveryTimeStamp              *pfcpType.RecoveryTimeStamp              `tlv:"96" json:"recoveryTimeStamp"`
	UPFunctionFeatures             *pfcpType.UPFunctionFeatures             `tlv:"43" json:"upFunctionFeatures"`
	CPFunctionFeatures             *pfcpType.CPFunctionFeatures             `tlv:"89" json:"cpFunctionFeatures"`
	UserPlaneIPResourceInformation *pfcpType.UserPlaneIPResourceInformation `tlv:"116" json:"userPlaneIPResourceInformation"`
}

type PFCPAssociationSetupResponse struct {
	NodeID                         *pfcpType.NodeID                         `tlv:"60" json:"nodeID"`
	Cause                          *pfcpType.Cause                          `tlv:"19" json:"cause"`
	RecoveryTimeStamp              *pfcpType.RecoveryTimeStamp              `tlv:"96" json:"recoveryTimeStamp"`
	UPFunctionFeatures             *pfcpType.UPFunctionFeatures             `tlv:"43" json:"upFunctionFeatures"`
	CPFunctionFeatures             *pfcpType.CPFunctionFeatures             `tlv:"89" json:"cpFunctionFeatures"`
	UserPlaneIPResourceInformation *pfcpType.UserPlaneIPResourceInformation `tlv:"116" json:"userPlaneIPResourceInformation"`
}

type PFCPAssociationUpdateRequest struct {
	NodeID                         *pfcpType.NodeID                         `tlv:"60" json:"nodeID"`
	UPFunctionFeatures             *pfcpType.UPFunctionFeatures             `tlv:"43" json:"upFunctionFeatures"`
	CPFunctionFeatures             *pfcpType.CPFunctionFeatures             `tlv:"89" json:"cpFunctionFeatures"`
	PFCPAssociationReleaseRequest  *PFCPAssociationReleaseRequest           `tlv:"111" json:"pfcpAssociationReleaseRequest"`
	GracefulReleasePeriod          *pfcpType.GracefulReleasePeriod          `tlv:"112" json:"gracefulReleasePeriod"`
	UserPlaneIPResourceInformation *pfcpType.UserPlaneIPResourceInformation `tlv:"116" json:"userPlaneIPResourceInformation"`
}

type PFCPAssociationUpdateResponse struct {
	NodeID             *pfcpType.NodeID             `tlv:"60" json:"nodeID"`
	Cause              *pfcpType.Cause              `tlv:"19" json:"cause"`
	UPFunctionFeatures *pfcpType.UPFunctionFeatures `tlv:"43" json:"upFunctionFeatures"`
	CPFunctionFeatures *pfcpType.CPFunctionFeatures `tlv:"89" json:"cpFunctionFeatures"`
}

type PFCPAssociationReleaseRequest struct {
	NodeID *pfcpType.NodeID `tlv:"60" json:"nodeID"`
}

type PFCPAssociationReleaseResponse struct {
	NodeID *pfcpType.NodeID `tlv:"60" json:"nodeID"`
	Cause  *pfcpType.Cause  `tlv:"19" json:"cause"`
}

type PFCPNodeReportRequest struct {
	NodeID                     *pfcpType.NodeID                     `tlv:"60" json:"nodeID"`
	NodeReportType             *pfcpType.NodeReportType             `tlv:"101" json:"nodeReportType"`
	UserPlanePathFailureReport *pfcpType.UserPlanePathFailureReport `tlv:"102" json:"userPlanePathFailureReport"`
}

type UserPlanePathFailure struct {
	RemoteGTPUPeer *pfcpType.RemoteGTPUPeer `tlv:"103" json:"remoteGTPUPeer"`
}

type PFCPNodeReportResponse struct {
	NodeID      *pfcpType.NodeID      `tlv:"60" json:"nodeID"`
	Cause       *pfcpType.Cause       `tlv:"19" json:"cause"`
	OffendingIE *pfcpType.OffendingIE `tlv:"40" json:"offendingIE"`
}

type PFCPSessionSetDeletionRequest struct {
	NodeID     *pfcpType.NodeID `tlv:"60" json:"nodeID"`
	SGWCFQCSID *pfcpType.FQCSID `tlv:"65" json:"sgwcfqcsid"`
	PGWCFQCSID *pfcpType.FQCSID `tlv:"65" json:"pgwcfqcsid"`
	SGWUFQCSID *pfcpType.FQCSID `tlv:"65" json:"sgwufqcsid"`
	PGWUFQCSID *pfcpType.FQCSID `tlv:"65" json:"pgwufqcsid"`
	TWANFQCSID *pfcpType.FQCSID `tlv:"65" json:"twanfqcsid"`
	EPDGFQCSID *pfcpType.FQCSID `tlv:"65" json:"epdgfqcsid"`
	MMEFQCSID  *pfcpType.FQCSID `tlv:"65" json:"mmefqcsid"`
}

type PFCPSessionSetDeletionResponse struct {
	NodeID      *pfcpType.NodeID      `tlv:"60" json:"nodeID"`
	Cause       *pfcpType.Cause       `tlv:"19" json:"cause"`
	OffendingIE *pfcpType.OffendingIE `tlv:"40" json:"offendingIE"`
}

type PFCPSessionEstablishmentRequest struct {
	NodeID                   *pfcpType.NodeID                   `tlv:"60" json:"nodeID"`
	CPFSEID                  *pfcpType.FSEID                    `tlv:"57" json:"cpfseid"`
	CreatePDR                []*CreatePDR                       `tlv:"1" json:"createPDR"`
	CreateFAR                []*CreateFAR                       `tlv:"3" json:"createFAR"`
	CreateURR                []*CreateURR                       `tlv:"6" json:"createURR"`
	CreateQER                []*CreateQER                       `tlv:"7" json:"createQER"`
	CreateBAR                []*CreateBAR                       `tlv:"85" json:"createBAR"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint             `tlv:"127" json:"createTrafficEndpoint"`
	PDNType                  *pfcpType.PDNType                  `tlv:"113" json:"pdnType"`
	SGWCFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"sgwcfqcsid"`
	MMEFQCSID                *pfcpType.FQCSID                   `tlv:"65" json:"mmefqcsid"`
	PGWCFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"pgwcfqcsid"`
	EPDGFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"epdgfqcsid"`
	TWANFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"twanfqcsid"`
	UserPlaneInactivityTimer *pfcpType.UserPlaneInactivityTimer `tlv:"117" json:"userPlaneInactivityTimer"`
	UserID                   *pfcpType.UserID                   `tlv:"141" json:"userID"`
	TraceInformation         *pfcpType.TraceInformation         `tlv:"152" json:"traceInformation"`
}

type CreatePDR struct {
	PDRID                   *pfcpType.PacketDetectionRuleID   `tlv:"56" json:"pdrid"`
	Precedence              *pfcpType.Precedence              `tlv:"29" json:"precedence"`
	PDI                     *PDI                              `tlv:"2" json:"pdi"`
	OuterHeaderRemoval      *pfcpType.OuterHeaderRemoval      `tlv:"95" json:"outerHeaderRemoval"`
	FARID                   *pfcpType.FARID                   `tlv:"108" json:"farid"`
	URRID                   *pfcpType.URRID                   `tlv:"81" json:"urrid"`
	QERID                   *pfcpType.QERID                   `tlv:"109" json:"qerid"`
	ActivatePredefinedRules *pfcpType.ActivatePredefinedRules `tlv:"106" json:"activatePredefinedRules"`
}

type PDI struct {
	SourceInterface               *pfcpType.SourceInterface               `tlv:"20" json:"sourceInterface"`
	LocalFTEID                    *pfcpType.FTEID                         `tlv:"21" json:"localFTEID"`
	NetworkInstance               *pfcpType.Dnn                           `tlv:"22" json:"networkInstance"`
	UEIPAddress                   *pfcpType.UEIPAddress                   `tlv:"93" json:"ueipAddress"`
	TrafficEndpointID             *pfcpType.TrafficEndpointID             `tlv:"131" json:"trafficEndpointID"`
	SDFFilter                     *pfcpType.SDFFilter                     `tlv:"23" json:"sdfFilter"`
	ApplicationID                 *pfcpType.ApplicationID                 `tlv:"24" json:"applicationID"`
	EthernetPDUSessionInformation *pfcpType.EthernetPDUSessionInformation `tlv:"142" json:"ethernetPDUSessionInformation"`
	EthernetPacketFilter          *EthernetPacketFilter                   `tlv:"132" json:"ethernetPacketFilter"`
	QFI                           *pfcpType.QFI                           `tlv:"124" json:"qfi"`
	FramedRoute                   *pfcpType.FramedRoute                   `tlv:"153" json:"framedRoute"`
	FramedRouting                 *pfcpType.FramedRouting                 `tlv:"154" json:"framedRouting"`
	FramedIPv6Route               *pfcpType.FramedIPv6Route               `tlv:"155" json:"framedIPv6Route"`
}

type EthernetPacketFilter struct {
	EthernetFilterID         *pfcpType.EthernetFilterID         `tlv:"138" json:"ethernetFilterID"`
	EthernetFilterProperties *pfcpType.EthernetFilterProperties `tlv:"139" json:"ethernetFilterProperties"`
	MACAddress               *pfcpType.MACAddress               `tlv:"133" json:"macAddress"`
	Ethertype                *pfcpType.Ethertype                `tlv:"136" json:"ethertype"`
	CTAG                     *pfcpType.CTAG                     `tlv:"134" json:"ctag"`
	STAG                     *pfcpType.STAG                     `tlv:"135" json:"stag"`
	SDFFilter                *pfcpType.SDFFilter                `tlv:"23" json:"sdfFilter"`
}

type CreateFAR struct {
	FARID                 *pfcpType.FARID                 `tlv:"108" json:"farid"`
	ApplyAction           *pfcpType.ApplyAction           `tlv:"44" json:"applyAction"`
	ForwardingParameters  *ForwardingParametersIEInFAR    `tlv:"4" json:"forwardingParameters"`
	DuplicatingParameters *pfcpType.DuplicatingParameters `tlv:"5" json:"duplicatingParameters"`
	BARID                 *pfcpType.BARID                 `tlv:"88" json:"barid"`
}

type ForwardingParametersIEInFAR struct {
	DestinationInterface    *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface"`
	NetworkInstance         *pfcpType.Dnn                   `tlv:"22" json:"networkInstance"`
	RedirectInformation     *pfcpType.RedirectInformation   `tlv:"38" json:"redirectInformation"`
	OuterHeaderCreation     *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation"`
	TransportLevelMarking   *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking"`
	ForwardingPolicy        *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy"`
	HeaderEnrichment        *pfcpType.HeaderEnrichment      `tlv:"98" json:"headerEnrichment"`
	LinkedTrafficEndpointID *pfcpType.TrafficEndpointID     `tlv:"131" json:"linkedTrafficEndpointID"`
	Proxying                *pfcpType.Proxying              `tlv:"137" json:"proxying"`
}

type DuplicatingParametersIEInFAR struct {
	DestinationInterface  *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface"`
	OuterHeaderCreation   *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation"`
	TransportLevelMarking *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking"`
	ForwardingPolicy      *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy"`
}

type CreateURR struct {
	URRID                     *pfcpType.URRID                     `tlv:"81" json:"urrid"`
	MeasurementMethod         *pfcpType.MeasurementMethod         `tlv:"62" json:"measurementMethod"`
	ReportingTriggers         *pfcpType.ReportingTriggers         `tlv:"37" json:"reportingTriggers"`
	MeasurementPeriod         *pfcpType.MeasurementPeriod         `tlv:"64" json:"measurementPeriod"`
	VolumeThreshold           *pfcpType.VolumeThreshold           `tlv:"31" json:"volumeThreshold"`
	VolumeQuota               *pfcpType.VolumeQuota               `tlv:"73" json:"volumeQuota"`
	TimeThreshold             *pfcpType.TimeThreshold             `tlv:"32" json:"timeThreshold"`
	TimeQuota                 *pfcpType.TimeQuota                 `tlv:"74" json:"timeQuota"`
	QuotaHoldingTime          *pfcpType.QuotaHoldingTime          `tlv:"71" json:"quotaHoldingTime"`
	DroppedDLTrafficThreshold *pfcpType.DroppedDLTrafficThreshold `tlv:"72" json:"droppedDLTrafficThreshold"`
	MonitoringTime            *pfcpType.MonitoringTime            `tlv:"33" json:"monitoringTime"`
	EventInformation          *EventInformation                   `tlv:"148" json:"eventInformation"`
	SubsequentVolumeThreshold *pfcpType.SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold"`
	SubsequentTimeThreshold   *pfcpType.SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold"`
	SubsequentVolumeQuota     *pfcpType.SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota"`
	SubsequentTimeQuota       *pfcpType.SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota"`
	InactivityDetectionTime   *pfcpType.InactivityDetectionTime   `tlv:"36" json:"inactivityDetectionTime"`
	LinkedURRID               *pfcpType.LinkedURRID               `tlv:"82" json:"linkedURRID"`
	MeasurementInformation    *pfcpType.MeasurementInformation    `tlv:"100" json:"measurementInformation"`
	TimeQuotaMechanism        *pfcpType.TimeQuotaMechanism        `tlv:"115" json:"timeQuotaMechanism"`
	AggregatedURRs            *AggregatedURRs                     `tlv:"118" json:"aggregatedURRs"`
	FARIDForQuotaAction       *pfcpType.FARID                     `tlv:"108" json:"faridForQuotaAction"`
	EthernetInactivityTimer   *pfcpType.EthernetInactivityTimer   `tlv:"146" json:"ethernetInactivityTimer"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime           `tlv:"147" json:"additionalMonitoringTime"`
}

type AggregatedURRs struct {
	AggregatedURRID *pfcpType.AggregatedURRID `tlv:"120" json:"aggregatedURRID"`
	Multiplier      *pfcpType.Multiplier      `tlv:"119" json:"multiplier"`
}

type AdditionalMonitoringTime struct {
	MonitoringTime            *pfcpType.MonitoringTime            `tlv:"33" json:"monitoringTime"`
	SubsequentVolumeThreshold *pfcpType.SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold"`
	SubsequentTimeThreshold   *pfcpType.SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold"`
	SubsequentVolumeQuota     *pfcpType.SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota"`
	SubsequentTimeQuota       *pfcpType.SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota"`
}

type EventInformation struct {
	EventID        *pfcpType.EventID        `tlv:"150" json:"eventID"`
	EventThreshold *pfcpType.EventThreshold `tlv:"151" json:"eventThreshold"`
}

type CreateQER struct {
	QERID              *pfcpType.QERID              `tlv:"109" json:"qerid"`
	QERCorrelationID   *pfcpType.QERCorrelationID   `tlv:"28" json:"qerCorrelationID"`
	GateStatus         *pfcpType.GateStatus         `tlv:"25" json:"gateStatus"`
	MaximumBitrate     *pfcpType.MBR                `tlv:"26" json:"maximumBitrate"`
	GuaranteedBitrate  *pfcpType.GBR                `tlv:"27" json:"guaranteedBitrate"`
	PacketRate         *pfcpType.PacketRate         `tlv:"94" json:"packetRate"`
	DLFlowLevelMarking *pfcpType.DLFlowLevelMarking `tlv:"97" json:"dlFlowLevelMarking"`
	QoSFlowIdentifier  *pfcpType.QFI                `tlv:"124" json:"qoSFlowIdentifier"`
	ReflectiveQoS      *pfcpType.RQI                `tlv:"123" json:"reflectiveQoS"`
}

type CreateBAR struct {
	BARID                          *pfcpType.BARID                          `tlv:"88" json:"barid"`
	DownlinkDataNotificationDelay  *pfcpType.DownlinkDataNotificationDelay  `tlv:"46" json:"downlinkDataNotificationDelay"`
	SuggestedBufferingPacketsCount *pfcpType.SuggestedBufferingPacketsCount `tlv:"140" json:"suggestedBufferingPacketsCount"`
}

type CreateTrafficEndpoint struct {
	TrafficEndpointID             *pfcpType.TrafficEndpointID             `tlv:"131" json:"trafficEndpointID"`
	LocalFTEID                    *pfcpType.FTEID                         `tlv:"21" json:"localFTEID"`
	NetworkInstance               *pfcpType.Dnn                           `tlv:"22" json:"networkInstance"`
	UEIPAddress                   *pfcpType.UEIPAddress                   `tlv:"93" json:"ueipAddress"`
	EthernetPDUSessionInformation *pfcpType.EthernetPDUSessionInformation `tlv:"142" json:"ethernetPDUSessionInformation"`
	FramedRoute                   *pfcpType.FramedRoute                   `tlv:"153" json:"framedRoute"`
	FramedRouting                 *pfcpType.FramedRouting                 `tlv:"154" json:"framedRouting"`
	FramedIPv6Route               *pfcpType.FramedIPv6Route               `tlv:"155" json:"framedIPv6Route"`
}

type PFCPSessionEstablishmentResponse struct {
	NodeID                     *pfcpType.NodeID            `tlv:"60" json:"nodeID"`
	Cause                      *pfcpType.Cause             `tlv:"19" json:"cause"`
	OffendingIE                *pfcpType.OffendingIE       `tlv:"40" json:"offendingIE"`
	UPFSEID                    *pfcpType.FSEID             `tlv:"57" json:"upfseid"`
	CreatedPDR                 *CreatedPDR                 `tlv:"8" json:"createdPDR"`
	LoadControlInformation     *LoadControlInformation     `tlv:"51" json:"loadControlInformation"`
	OverloadControlInformation *OverloadControlInformation `tlv:"54" json:"overloadControlInformation"`
	SGWUFQCSID                 *pfcpType.FQCSID            `tlv:"65" json:"sgwufqcsid"`
	PGWUFQCSID                 *pfcpType.FQCSID            `tlv:"65" json:"pgwufqcsid"`
	FailedRuleID               *pfcpType.FailedRuleID      `tlv:"114" json:"failedRuleID"`
	CreatedTrafficEndpoint     *CreatedTrafficEndpoint     `tlv:"128" json:"createdTrafficEndpoint"`
}

type CreatedPDR struct {
	PDRID      *pfcpType.PacketDetectionRuleID `tlv:"56" json:"pdrid"`
	LocalFTEID *pfcpType.FTEID                 `tlv:"21" json:"localFTEID"`
}

type LoadControlInformation struct {
	LoadControlSequenceNumber *pfcpType.SequenceNumber `tlv:"52" json:"loadControlSequenceNumber"`
	LoadMetric                *pfcpType.Metric         `tlv:"53" json:"loadMetric"`
}

type OverloadControlInformation struct {
	OverloadControlSequenceNumber   *pfcpType.SequenceNumber `tlv:"52" json:"overloadControlSequenceNumber"`
	OverloadReductionMetric         *pfcpType.Metric         `tlv:"53" json:"overloadReductionMetric"`
	PeriodOfValidity                *pfcpType.Timer          `tlv:"55" json:"periodOfValidity"`
	OverloadControlInformationFlags *pfcpType.OCIFlags       `tlv:"110" json:"overloadControlInformationFlags"`
}

type CreatedTrafficEndpoint struct {
	TrafficEndpointID *pfcpType.TrafficEndpointID `tlv:"131" json:"trafficEndpointID"`
	LocalFTEID        *pfcpType.FTEID             `tlv:"21" json:"localFTEID"`
}

type PFCPSessionModificationRequest struct {
	CPFSEID                  *pfcpType.FSEID                          `tlv:"57" json:"cpfseid"`
	RemovePDR                []*RemovePDR                             `tlv:"15" json:"removePDR"`
	RemoveFAR                []*RemoveFAR                             `tlv:"16" json:"removeFAR"`
	RemoveURR                []*RemoveURR                             `tlv:"17" json:"removeURR"`
	RemoveQER                []*pfcpType.RemoveQER                    `tlv:"18" json:"removeQER"`
	RemoveBAR                []*RemoveBAR                             `tlv:"87" json:"removeBAR"`
	RemoveTrafficEndpoint    *RemoveTrafficEndpoint                   `tlv:"130" json:"removeTrafficEndpoint"`
	CreatePDR                []*CreatePDR                             `tlv:"1" json:"createPDR"`
	CreateFAR                []*CreateFAR                             `tlv:"3" json:"createFAR"`
	CreateURR                []*CreateURR                             `tlv:"6" json:"createURR"`
	CreateQER                []*CreateQER                             `tlv:"7" json:"createQER"`
	CreateBAR                []*CreateBAR                             `tlv:"85" json:"createBAR"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint                   `tlv:"127" json:"createTrafficEndpoint"`
	UpdatePDR                []*UpdatePDR                             `tlv:"9" json:"updatePDR"`
	UpdateFAR                []*UpdateFAR                             `tlv:"10" json:"updateFAR"`
	UpdateURR                []*UpdateURR                             `tlv:"13" json:"updateURR"`
	UpdateQER                []*UpdateQER                             `tlv:"14" json:"updateQER"`
	UpdateBAR                *UpdateBARPFCPSessionModificationRequest `tlv:"86" json:"updateBAR"`
	UpdateTrafficEndpoint    *UpdateTrafficEndpoint                   `tlv:"129" json:"updateTrafficEndpoint"`
	PFCPSMReqFlags           *pfcpType.PFCPSMReqFlags                 `tlv:"49" json:"pfcpsmReqFlags"`
	QueryURR                 *QueryURR                                `tlv:"77" json:"queryURR"`
	PGWCFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"pgwcfqcsid"`
	SGWCFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"sgwcfqcsid"`
	MMEFQCSID                *pfcpType.FQCSID                         `tlv:"65" json:"mmefqcsid"`
	EPDGFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"epdgfqcsid"`
	TWANFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"twanfqcsid"`
	UserPlaneInactivityTimer *pfcpType.UserPlaneInactivityTimer       `tlv:"117" json:"userPlaneInactivityTimer"`
	QueryURRReference        *pfcpType.QueryURRReference              `tlv:"125" json:"queryURRReference"`
	TraceInformation         *pfcpType.TraceInformation               `tlv:"152" json:"traceInformation"`
}

type UpdatePDR struct {
	PDRID                     *pfcpType.PacketDetectionRuleID     `tlv:"56" json:"pdrid"`
	OuterHeaderRemoval        *pfcpType.OuterHeaderRemoval        `tlv:"95" json:"outerHeaderRemoval"`
	Precedence                *pfcpType.Precedence                `tlv:"29" json:"precedence"`
	PDI                       *PDI                                `tlv:"2" json:"pdi"`
	FARID                     *pfcpType.FARID                     `tlv:"108" json:"farid"`
	URRID                     *pfcpType.URRID                     `tlv:"81" json:"urrid"`
	QERID                     *pfcpType.QERID                     `tlv:"109" json:"qerid"`
	ActivatePredefinedRules   *pfcpType.ActivatePredefinedRules   `tlv:"106" json:"activatePredefinedRules"`
	DeactivatePredefinedRules *pfcpType.DeactivatePredefinedRules `tlv:"107" json:"deactivatePredefinedRules"`
}

type UpdateFAR struct {
	FARID                       *pfcpType.FARID                       `tlv:"108" json:"farid"`
	ApplyAction                 *pfcpType.ApplyAction                 `tlv:"44" json:"applyAction"`
	UpdateForwardingParameters  *UpdateForwardingParametersIEInFAR    `tlv:"11" json:"updateForwardingParameters"`
	UpdateDuplicatingParameters *pfcpType.UpdateDuplicatingParameters `tlv:"105" json:"updateDuplicatingParameters"`
	BARID                       *pfcpType.BARID                       `tlv:"88" json:"barid"`
}

type UpdateForwardingParametersIEInFAR struct {
	DestinationInterface    *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface"`
	NetworkInstance         *pfcpType.Dnn                   `tlv:"22" json:"networkInstance"`
	RedirectInformation     *pfcpType.RedirectInformation   `tlv:"38" json:"redirectInformation"`
	OuterHeaderCreation     *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation"`
	TransportLevelMarking   *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking"`
	ForwardingPolicy        *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy"`
	HeaderEnrichment        *pfcpType.HeaderEnrichment      `tlv:"98" json:"headerEnrichment"`
	PFCPSMReqFlags          *pfcpType.PFCPSMReqFlags        `tlv:"49" json:"pfcpsmReqFlags"`
	LinkedTrafficEndpointID *pfcpType.TrafficEndpointID     `tlv:"131" json:"linkedTrafficEndpointID"`
}

type UpdateDuplicatingParametersIEInFAR struct {
	DestinationInterface  *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface"`
	OuterHeaderCreation   *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation"`
	TransportLevelMarking *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking"`
	ForwardingPolicy      *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy"`
}

type UpdateURR struct {
	URRID                     *pfcpType.URRID                     `tlv:"81" json:"urrid"`
	MeasurementMethod         *pfcpType.MeasurementMethod         `tlv:"62" json:"measurementMethod"`
	ReportingTriggers         *pfcpType.ReportingTriggers         `tlv:"37" json:"reportingTriggers"`
	MeasurementPeriod         *pfcpType.MeasurementPeriod         `tlv:"64" json:"measurementPeriod"`
	VolumeThreshold           *pfcpType.VolumeThreshold           `tlv:"31" json:"volumeThreshold"`
	VolumeQuota               *pfcpType.VolumeQuota               `tlv:"73" json:"volumeQuota"`
	TimeThreshold             *pfcpType.TimeThreshold             `tlv:"32" json:"timeThreshold"`
	TimeQuota                 *pfcpType.TimeQuota                 `tlv:"74" json:"timeQuota"`
	QuotaHoldingTime          *pfcpType.QuotaHoldingTime          `tlv:"71" json:"quotaHoldingTime"`
	DroppedDLTrafficThreshold *pfcpType.DroppedDLTrafficThreshold `tlv:"72" json:"droppedDLTrafficThreshold"`
	MonitoringTime            *pfcpType.MonitoringTime            `tlv:"33" json:"monitoringTime"`
	EventInformation          *EventInformation                   `tlv:"148" json:"eventInformation"`
	SubsequentVolumeThreshold *pfcpType.SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold"`
	SubsequentTimeThreshold   *pfcpType.SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold"`
	SubsequentVolumeQuota     *pfcpType.SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota"`
	SubsequentTimeQuota       *pfcpType.SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota"`
	InactivityDetectionTime   *pfcpType.InactivityDetectionTime   `tlv:"36" json:"inactivityDetectionTime"`
	LinkedURRID               *pfcpType.LinkedURRID               `tlv:"82" json:"linkedURRID"`
	MeasurementInformation    *pfcpType.MeasurementInformation    `tlv:"100" json:"measurementInformation"`
	TimeQuotaMechanism        *pfcpType.TimeQuotaMechanism        `tlv:"115" json:"timeQuotaMechanism"`
	AggregatedURRs            *AggregatedURRs                     `tlv:"118" json:"aggregatedURRs"`
	FARIDForQuotaAction       *pfcpType.FARID                     `tlv:"108" json:"faridForQuotaAction"`
	EthernetInactivityTimer   *pfcpType.EthernetInactivityTimer   `tlv:"146" json:"ethernetInactivityTimer"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime           `tlv:"147" json:"additionalMonitoringTime"`
}

type UpdateQER struct {
	QERID              *pfcpType.QERID              `tlv:"109" json:"qerid"`
	QERCorrelationID   *pfcpType.QERCorrelationID   `tlv:"28" json:"qerCorrelationID"`
	GateStatus         *pfcpType.GateStatus         `tlv:"25" json:"gateStatus"`
	MaximumBitrate     *pfcpType.MBR                `tlv:"26" json:"maximumBitrate"`
	GuaranteedBitrate  *pfcpType.GBR                `tlv:"27" json:"guaranteedBitrate"`
	PacketRate         *pfcpType.PacketRate         `tlv:"94" json:"packetRate"`
	DLFlowLevelMarking *pfcpType.DLFlowLevelMarking `tlv:"97" json:"dlFlowLevelMarking"`
	QoSFlowIdentifier  *pfcpType.QFI                `tlv:"124" json:"qoSFlowIdentifier"`
	ReflectiveQoS      *pfcpType.RQI                `tlv:"123" json:"reflectiveQoS"`
}

type RemovePDR struct {
	PDRID *pfcpType.PacketDetectionRuleID `tlv:"56" json:"pdrid"`
}

type RemoveFAR struct {
	FARID *pfcpType.FARID `tlv:"108" json:"farid"`
}

type RemoveURR struct {
	URRID *pfcpType.URRID `tlv:"81" json:"urrid"`
}

type RemoveQERIEPFCPSessionModificationRequest struct {
	QERID *pfcpType.QERID `tlv:"109" json:"qerid"`
}

type QueryURR struct {
	URRID *pfcpType.URRID `tlv:"81" json:"urrid"`
}

type UpdateBARPFCPSessionModificationRequest struct {
	BARID                          *pfcpType.BARID                          `tlv:"88" json:"barid"`
	DownlinkDataNotificationDelay  *pfcpType.DownlinkDataNotificationDelay  `tlv:"46" json:"downlinkDataNotificationDelay"`
	SuggestedBufferingPacketsCount *pfcpType.SuggestedBufferingPacketsCount `tlv:"140" json:"suggestedBufferingPacketsCount"`
}

type RemoveBAR struct {
	BARID *pfcpType.BARID `tlv:"88" json:"barid"`
}

type UpdateTrafficEndpoint struct {
	TrafficEndpointID *pfcpType.TrafficEndpointID `tlv:"131" json:"trafficEndpointID"`
	LocalFTEID        *pfcpType.FTEID             `tlv:"21" json:"localFTEID"`
	NetworkInstance   *pfcpType.Dnn               `tlv:"22" json:"networkInstance"`
	UEIPAddress       *pfcpType.UEIPAddress       `tlv:"93" json:"ueipAddress"`
	FramedRoute       *pfcpType.FramedRoute       `tlv:"153" json:"framedRoute"`
	FramedRouting     *pfcpType.FramedRouting     `tlv:"154" json:"framedRouting"`
	FramedIPv6Route   *pfcpType.FramedIPv6Route   `tlv:"155" json:"framedIPv6Route"`
}

type RemoveTrafficEndpoint struct {
	TrafficEndpointID *pfcpType.TrafficEndpointID `tlv:"131" json:"trafficEndpointID"`
}

type PFCPSessionModificationResponse struct {
	Cause                             *pfcpType.Cause                             `tlv:"19" json:"cause"`
	OffendingIE                       *pfcpType.OffendingIE                       `tlv:"40" json:"offendingIE"`
	CreatedPDR                        *CreatedPDR                                 `tlv:"8" json:"createdPDR"`
	LoadControlInformation            *LoadControlInformation                     `tlv:"51" json:"loadControlInformation"`
	OverloadControlInformation        *OverloadControlInformation                 `tlv:"54" json:"overloadControlInformation"`
	UsageReport                       *UsageReportPFCPSessionModificationResponse `tlv:"78" json:"usageReport"`
	FailedRuleID                      *pfcpType.FailedRuleID                      `tlv:"114" json:"failedRuleID"`
	AdditionalUsageReportsInformation *pfcpType.AdditionalUsageReportsInformation `tlv:"126" json:"additionalUsageReportsInformation"`
	CreatedUpdatedTrafficEndpoint     *CreatedTrafficEndpoint                     `tlv:"128" json:"createdUpdatedTrafficEndpoint"`
}

type UsageReportPFCPSessionModificationResponse struct {
	URRID                      *pfcpType.URRID               `tlv:"81" json:"urrid"`
	URSEQN                     *pfcpType.URSEQN              `tlv:"104" json:"urseqn"`
	UsageReportTrigger         *pfcpType.UsageReportTrigger  `tlv:"63" json:"usageReportTrigger"`
	StartTime                  *pfcpType.StartTime           `tlv:"75" json:"startTime"`
	EndTime                    *pfcpType.EndTime             `tlv:"76" json:"endTime"`
	VolumeMeasurement          *pfcpType.VolumeMeasurement   `tlv:"66" json:"volumeMeasurement"`
	DurationMeasurement        *pfcpType.DurationMeasurement `tlv:"67" json:"durationMeasurement"`
	TimeOfFirstPacket          *pfcpType.TimeOfFirstPacket   `tlv:"69" json:"timeOfFirstPacket"`
	TimeOfLastPacket           *pfcpType.TimeOfLastPacket    `tlv:"70" json:"timeOfLastPacket"`
	UsageInformation           *pfcpType.UsageInformation    `tlv:"90" json:"usageInformation"`
	QueryURRReference          *pfcpType.QueryURRReference   `tlv:"125" json:"queryURRReference"`
	EthernetTrafficInformation *EthernetTrafficInformation   `tlv:"143" json:"ethernetTrafficInformation"`
}

type PFCPSessionDeletionRequest struct {
}

type PFCPSessionDeletionResponse struct {
	Cause                      *pfcpType.Cause                         `tlv:"19" json:"cause"`
	OffendingIE                *pfcpType.OffendingIE                   `tlv:"40" json:"offendingIE"`
	LoadControlInformation     *LoadControlInformation                 `tlv:"51" json:"loadControlInformation"`
	OverloadControlInformation *OverloadControlInformation             `tlv:"54" json:"overloadControlInformation"`
	UsageReport                *UsageReportPFCPSessionDeletionResponse `tlv:"79" json:"usageReport"`
}

type UsageReportPFCPSessionDeletionResponse struct {
	URRID                      *pfcpType.URRID               `tlv:"81" json:"urrid"`
	URSEQN                     *pfcpType.URSEQN              `tlv:"104" json:"urseqn"`
	UsageReportTrigger         *pfcpType.UsageReportTrigger  `tlv:"63" json:"usageReportTrigger"`
	StartTime                  *pfcpType.StartTime           `tlv:"75" json:"startTime"`
	EndTime                    *pfcpType.EndTime             `tlv:"76" json:"endTime"`
	VolumeMeasurement          *pfcpType.VolumeMeasurement   `tlv:"66" json:"volumeMeasurement"`
	DurationMeasurement        *pfcpType.DurationMeasurement `tlv:"67" json:"durationMeasurement"`
	TimeOfFirstPacket          *pfcpType.TimeOfFirstPacket   `tlv:"69" json:"timeOfFirstPacket"`
	TimeOfLastPacket           *pfcpType.TimeOfLastPacket    `tlv:"70" json:"timeOfLastPacket"`
	UsageInformation           *pfcpType.UsageInformation    `tlv:"90" json:"usageInformation"`
	EthernetTrafficInformation *EthernetTrafficInformation   `tlv:"143" json:"ethernetTrafficInformation"`
}

type PFCPSessionReportRequest struct {
	ReportType                        *pfcpType.ReportType                        `tlv:"39" json:"reportType"`
	DownlinkDataReport                *DownlinkDataReport                         `tlv:"83" json:"downlinkDataReport"`
	UsageReport                       *UsageReportPFCPSessionReportRequest        `tlv:"80" json:"usageReport"`
	ErrorIndicationReport             *ErrorIndicationReport                      `tlv:"99" json:"errorIndicationReport"`
	LoadControlInformation            *LoadControlInformation                     `tlv:"51" json:"loadControlInformation"`
	OverloadControlInformation        *OverloadControlInformation                 `tlv:"54" json:"overloadControlInformation"`
	AdditionalUsageReportsInformation *pfcpType.AdditionalUsageReportsInformation `tlv:"126" json:"additionalUsageReportsInformation"`
}

type DownlinkDataReport struct {
	PDRID                          *pfcpType.PacketDetectionRuleID          `tlv:"56" json:"pdrid"`
	DownlinkDataServiceInformation *pfcpType.DownlinkDataServiceInformation `tlv:"45" json:"downlinkDataServiceInformation"`
}

type UsageReportPFCPSessionReportRequest struct {
	URRID                           *pfcpType.URRID                  `tlv:"81" json:"urrid"`
	URSEQN                          *pfcpType.URSEQN                 `tlv:"104" json:"urseqn"`
	UsageReportTrigger              *pfcpType.UsageReportTrigger     `tlv:"63" json:"usageReportTrigger"`
	StartTime                       *pfcpType.StartTime              `tlv:"75" json:"startTime"`
	EndTime                         *pfcpType.EndTime                `tlv:"76" json:"endTime"`
	VolumeMeasurement               *pfcpType.VolumeMeasurement      `tlv:"66" json:"volumeMeasurement"`
	DurationMeasurement             *pfcpType.DurationMeasurement    `tlv:"67" json:"durationMeasurement"`
	ApplicationDetectionInformation *ApplicationDetectionInformation `tlv:"68" json:"applicationDetectionInformation"`
	UEIPAddress                     *pfcpType.UEIPAddress            `tlv:"93" json:"ueipAddress"`
	NetworkInstance                 *pfcpType.Dnn                    `tlv:"22" json:"networkInstance"`
	TimeOfFirstPacket               *pfcpType.TimeOfFirstPacket      `tlv:"69" json:"timeOfFirstPacket"`
	TimeOfLastPacket                *pfcpType.TimeOfLastPacket       `tlv:"70" json:"timeOfLastPacket"`
	UsageInformation                *pfcpType.UsageInformation       `tlv:"90" json:"usageInformation"`
	QueryURRReference               *pfcpType.QueryURRReference      `tlv:"125" json:"queryURRReference"`
	EventReporting                  *EventReporting                  `tlv:"149" json:"eventReporting"`
	EthernetTrafficInformation      *EthernetTrafficInformation      `tlv:"143" json:"ethernetTrafficInformation"`
}

type ApplicationDetectionInformation struct {
	ApplicationID         *pfcpType.ApplicationID         `tlv:"24" json:"applicationID"`
	ApplicationInstanceID *pfcpType.ApplicationInstanceID `tlv:"91" json:"applicationInstanceID"`
	FlowInformation       *pfcpType.FlowInformation       `tlv:"92" json:"flowInformation"`
}

type EventReporting struct {
	EventID *pfcpType.EventID `tlv:"150" json:"eventID"`
}

type EthernetTrafficInformation struct {
	MACAddressesDetected *pfcpType.MACAddressesDetected `tlv:"144" json:"macAddressesDetected"`
	MACAddressesRemoved  *pfcpType.MACAddressesRemoved  `tlv:"145" json:"macAddressesRemoved"`
}

type ErrorIndicationReport struct {
	RemoteFTEID *pfcpType.FTEID `tlv:"21" json:"remoteFTEID"`
}

type PFCPSessionReportResponse struct {
	Cause        *pfcpType.Cause                              `tlv:"19" json:"cause"`
	OffendingIE  *pfcpType.OffendingIE                        `tlv:"40" json:"offendingIE"`
	UpdateBAR    *pfcpType.UpdateBARPFCPSessionReportResponse `tlv:"12" json:"updateBAR"`
	SxSRRspFlags *pfcpType.PFCPSRRspFlags                     `tlv:"50" json:"sxSRRspFlags"`
}

type UpdateBARIEInPFCPSessionReportResponse struct {
	BARID                           *pfcpType.BARID                           `tlv:"88" json:"barid"`
	DownlinkDataNotificationDelay   *pfcpType.DownlinkDataNotificationDelay   `tlv:"46" json:"downlinkDataNotificationDelay"`
	DLBufferingDuration             *pfcpType.DLBufferingDuration             `tlv:"47" json:"dlBufferingDuration"`
	DLBufferingSuggestedPacketCount *pfcpType.DLBufferingSuggestedPacketCount `tlv:"48" json:"dlBufferingSuggestedPacketCount"`
	SuggestedBufferingPacketsCount  *pfcpType.SuggestedBufferingPacketsCount  `tlv:"140" json:"suggestedBufferingPacketsCount"`
}
