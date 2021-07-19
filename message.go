package pfcp

import (
	"github.com/mimetrix/pfcp/pfcpType"
)

type Message struct {
	Header Header      `json:"header,omitempty"`
	Body   interface{} `json:"body,omitempty"`
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
	RecoveryTimeStamp *pfcpType.RecoveryTimeStamp `tlv:"96" json:"recoveryTimeStamp,omitempty"`
}

type HeartbeatResponse struct {
	RecoveryTimeStamp *pfcpType.RecoveryTimeStamp `tlv:"96" json:"recoveryTimeStamp,omitempty"`
}

type PFCPPFDManagementRequest struct {
	ApplicationIDsPFDs []ApplicationIDsPFDs `tlv:"58" json:"applicationIDsPFDs,omitempty"`
}

type ApplicationIDsPFDs struct {
	ApplicationID pfcpType.ApplicationID `tlv:"24" json:"applicationID,omitempty"`
	PFD           *PFD                   `tlv:"59" json:"pfd,omitempty"`
}

type PFD struct {
	PFDContents []pfcpType.PFDContents `tlv:"61" json:"pfdContents,omitempty"`
}

type PFCPPFDManagementResponse struct {
	Cause       *pfcpType.Cause       `tlv:"19" json:"cause,omitempty"`
	OffendingIE *pfcpType.OffendingIE `tlv:"40" json:"offendingIE,omitempty"`
}

type PFCPAssociationSetupRequest struct {
	NodeID                         *pfcpType.NodeID                         `tlv:"60" json:"nodeID,omitempty"`
	RecoveryTimeStamp              *pfcpType.RecoveryTimeStamp              `tlv:"96" json:"recoveryTimeStamp,omitempty"`
	UPFunctionFeatures             *pfcpType.UPFunctionFeatures             `tlv:"43" json:"upFunctionFeatures,omitempty"`
	CPFunctionFeatures             *pfcpType.CPFunctionFeatures             `tlv:"89" json:"cpFunctionFeatures,omitempty"`
	UserPlaneIPResourceInformation *pfcpType.UserPlaneIPResourceInformation `tlv:"116" json:"userPlaneIPResourceInformation,omitempty"`
}

type PFCPAssociationSetupResponse struct {
	NodeID                         *pfcpType.NodeID                         `tlv:"60" json:"nodeID,omitempty"`
	Cause                          *pfcpType.Cause                          `tlv:"19" json:"cause,omitempty"`
	RecoveryTimeStamp              *pfcpType.RecoveryTimeStamp              `tlv:"96" json:"recoveryTimeStamp,omitempty"`
	UPFunctionFeatures             *pfcpType.UPFunctionFeatures             `tlv:"43" json:"upFunctionFeatures,omitempty"`
	CPFunctionFeatures             *pfcpType.CPFunctionFeatures             `tlv:"89" json:"cpFunctionFeatures,omitempty"`
	UserPlaneIPResourceInformation *pfcpType.UserPlaneIPResourceInformation `tlv:"116" json:"userPlaneIPResourceInformation,omitempty"`
}

type PFCPAssociationUpdateRequest struct {
	NodeID                         *pfcpType.NodeID                         `tlv:"60" json:"nodeID,omitempty"`
	UPFunctionFeatures             *pfcpType.UPFunctionFeatures             `tlv:"43" json:"upFunctionFeatures,omitempty"`
	CPFunctionFeatures             *pfcpType.CPFunctionFeatures             `tlv:"89" json:"cpFunctionFeatures,omitempty"`
	PFCPAssociationReleaseRequest  *PFCPAssociationReleaseRequest           `tlv:"111" json:"pfcpAssociationReleaseRequest,omitempty"`
	GracefulReleasePeriod          *pfcpType.GracefulReleasePeriod          `tlv:"112" json:"gracefulReleasePeriod,omitempty"`
	UserPlaneIPResourceInformation *pfcpType.UserPlaneIPResourceInformation `tlv:"116" json:"userPlaneIPResourceInformation,omitempty"`
}

type PFCPAssociationUpdateResponse struct {
	NodeID             *pfcpType.NodeID             `tlv:"60" json:"nodeID,omitempty"`
	Cause              *pfcpType.Cause              `tlv:"19" json:"cause,omitempty"`
	UPFunctionFeatures *pfcpType.UPFunctionFeatures `tlv:"43" json:"upFunctionFeatures,omitempty"`
	CPFunctionFeatures *pfcpType.CPFunctionFeatures `tlv:"89" json:"cpFunctionFeatures,omitempty"`
}

type PFCPAssociationReleaseRequest struct {
	NodeID *pfcpType.NodeID `tlv:"60" json:"nodeID,omitempty"`
}

type PFCPAssociationReleaseResponse struct {
	NodeID *pfcpType.NodeID `tlv:"60" json:"nodeID,omitempty"`
	Cause  *pfcpType.Cause  `tlv:"19" json:"cause,omitempty"`
}

type PFCPNodeReportRequest struct {
	NodeID                     *pfcpType.NodeID                     `tlv:"60" json:"nodeID,omitempty"`
	NodeReportType             *pfcpType.NodeReportType             `tlv:"101" json:"nodeReportType,omitempty"`
	UserPlanePathFailureReport *pfcpType.UserPlanePathFailureReport `tlv:"102" json:"userPlanePathFailureReport,omitempty"`
}

type UserPlanePathFailure struct {
	RemoteGTPUPeer *pfcpType.RemoteGTPUPeer `tlv:"103" json:"remoteGTPUPeer,omitempty"`
}

type PFCPNodeReportResponse struct {
	NodeID      *pfcpType.NodeID      `tlv:"60" json:"nodeID,omitempty"`
	Cause       *pfcpType.Cause       `tlv:"19" json:"cause,omitempty"`
	OffendingIE *pfcpType.OffendingIE `tlv:"40" json:"offendingIE,omitempty"`
}

type PFCPSessionSetDeletionRequest struct {
	NodeID     *pfcpType.NodeID `tlv:"60" json:"nodeID,omitempty"`
	SGWCFQCSID *pfcpType.FQCSID `tlv:"65" json:"sgwcfqcsid,omitempty"`
	PGWCFQCSID *pfcpType.FQCSID `tlv:"65" json:"pgwcfqcsid,omitempty"`
	SGWUFQCSID *pfcpType.FQCSID `tlv:"65" json:"sgwufqcsid,omitempty"`
	PGWUFQCSID *pfcpType.FQCSID `tlv:"65" json:"pgwufqcsid,omitempty"`
	TWANFQCSID *pfcpType.FQCSID `tlv:"65" json:"twanfqcsid,omitempty"`
	EPDGFQCSID *pfcpType.FQCSID `tlv:"65" json:"epdgfqcsid,omitempty"`
	MMEFQCSID  *pfcpType.FQCSID `tlv:"65" json:"mmefqcsid,omitempty"`
}

type PFCPSessionSetDeletionResponse struct {
	NodeID      *pfcpType.NodeID      `tlv:"60" json:"nodeID,omitempty"`
	Cause       *pfcpType.Cause       `tlv:"19" json:"cause,omitempty"`
	OffendingIE *pfcpType.OffendingIE `tlv:"40" json:"offendingIE,omitempty"`
}

type PFCPSessionEstablishmentRequest struct {
	NodeID                   *pfcpType.NodeID                   `tlv:"60" json:"nodeID,omitempty"`
	CPFSEID                  *pfcpType.FSEID                    `tlv:"57" json:"cpfseid,omitempty"`
	CreatePDR                []*CreatePDR                       `tlv:"1" json:"createPDR,omitempty"`
	CreateFAR                []*CreateFAR                       `tlv:"3" json:"createFAR,omitempty"`
	CreateURR                []*CreateURR                       `tlv:"6" json:"createURR,omitempty"`
	CreateQER                []*CreateQER                       `tlv:"7" json:"createQER,omitempty"`
	CreateBAR                []*CreateBAR                       `tlv:"85" json:"createBAR,omitempty"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint             `tlv:"127" json:"createTrafficEndpoint,omitempty"`
	PDNType                  *pfcpType.PDNType                  `tlv:"113" json:"pdnType,omitempty"`
	SGWCFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"sgwcfqcsid,omitempty"`
	MMEFQCSID                *pfcpType.FQCSID                   `tlv:"65" json:"mmefqcsid,omitempty"`
	PGWCFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"pgwcfqcsid,omitempty"`
	EPDGFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"epdgfqcsid,omitempty"`
	TWANFQCSID               *pfcpType.FQCSID                   `tlv:"65" json:"twanfqcsid,omitempty"`
	UserPlaneInactivityTimer *pfcpType.UserPlaneInactivityTimer `tlv:"117" json:"userPlaneInactivityTimer,omitempty"`
	UserID                   *pfcpType.UserID                   `tlv:"141" json:"userID,omitempty"`
	TraceInformation         *pfcpType.TraceInformation         `tlv:"152" json:"traceInformation,omitempty"`
}

type CreatePDR struct {
	PDRID                   *pfcpType.PacketDetectionRuleID   `tlv:"56" json:"pdrid,omitempty"`
	Precedence              *pfcpType.Precedence              `tlv:"29" json:"precedence,omitempty"`
	PDI                     *PDI                              `tlv:"2" json:"pdi,omitempty"`
	OuterHeaderRemoval      *pfcpType.OuterHeaderRemoval      `tlv:"95" json:"outerHeaderRemoval,omitempty"`
	FARID                   *pfcpType.FARID                   `tlv:"108" json:"farid,omitempty"`
	URRID                   *pfcpType.URRID                   `tlv:"81" json:"urrid,omitempty"`
	QERID                   *pfcpType.QERID                   `tlv:"109" json:"qerid,omitempty"`
	ActivatePredefinedRules *pfcpType.ActivatePredefinedRules `tlv:"106" json:"activatePredefinedRules,omitempty"`
}

type PDI struct {
	SourceInterface               *pfcpType.SourceInterface               `tlv:"20" json:"sourceInterface,omitempty"`
	LocalFTEID                    *pfcpType.FTEID                         `tlv:"21" json:"localFTEID,omitempty"`
	NetworkInstance               *pfcpType.Dnn                           `tlv:"22" json:"networkInstance,omitempty"`
	UEIPAddress                   *pfcpType.UEIPAddress                   `tlv:"93" json:"ueipAddress,omitempty"`
	TrafficEndpointID             *pfcpType.TrafficEndpointID             `tlv:"131" json:"trafficEndpointID,omitempty"`
	SDFFilter                     *pfcpType.SDFFilter                     `tlv:"23" json:"sdfFilter,omitempty"`
	ApplicationID                 *pfcpType.ApplicationID                 `tlv:"24" json:"applicationID,omitempty"`
	EthernetPDUSessionInformation *pfcpType.EthernetPDUSessionInformation `tlv:"142" json:"ethernetPDUSessionInformation,omitempty"`
	EthernetPacketFilter          *EthernetPacketFilter                   `tlv:"132" json:"ethernetPacketFilter,omitempty"`
	QFI                           *pfcpType.QFI                           `tlv:"124" json:"qfi,omitempty"`
	FramedRoute                   *pfcpType.FramedRoute                   `tlv:"153" json:"framedRoute,omitempty"`
	FramedRouting                 *pfcpType.FramedRouting                 `tlv:"154" json:"framedRouting,omitempty"`
	FramedIPv6Route               *pfcpType.FramedIPv6Route               `tlv:"155" json:"framedIPv6Route,omitempty"`
}

type EthernetPacketFilter struct {
	EthernetFilterID         *pfcpType.EthernetFilterID         `tlv:"138" json:"ethernetFilterID,omitempty"`
	EthernetFilterProperties *pfcpType.EthernetFilterProperties `tlv:"139" json:"ethernetFilterProperties,omitempty"`
	MACAddress               *pfcpType.MACAddress               `tlv:"133" json:"macAddress,omitempty"`
	Ethertype                *pfcpType.Ethertype                `tlv:"136" json:"ethertype,omitempty"`
	CTAG                     *pfcpType.CTAG                     `tlv:"134" json:"ctag,omitempty"`
	STAG                     *pfcpType.STAG                     `tlv:"135" json:"stag,omitempty"`
	SDFFilter                *pfcpType.SDFFilter                `tlv:"23" json:"sdfFilter,omitempty"`
}

type CreateFAR struct {
	FARID                 *pfcpType.FARID                 `tlv:"108" json:"farid,omitempty"`
	ApplyAction           *pfcpType.ApplyAction           `tlv:"44" json:"applyAction,omitempty"`
	ForwardingParameters  *ForwardingParametersIEInFAR    `tlv:"4" json:"forwardingParameters,omitempty"`
	DuplicatingParameters *pfcpType.DuplicatingParameters `tlv:"5" json:"duplicatingParameters,omitempty"`
	BARID                 *pfcpType.BARID                 `tlv:"88" json:"barid,omitempty"`
}

type ForwardingParametersIEInFAR struct {
	DestinationInterface    *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface,omitempty"`
	NetworkInstance         *pfcpType.Dnn                   `tlv:"22" json:"networkInstance,omitempty"`
	RedirectInformation     *pfcpType.RedirectInformation   `tlv:"38" json:"redirectInformation,omitempty"`
	OuterHeaderCreation     *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking   *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy        *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
	HeaderEnrichment        *pfcpType.HeaderEnrichment      `tlv:"98" json:"headerEnrichment,omitempty"`
	LinkedTrafficEndpointID *pfcpType.TrafficEndpointID     `tlv:"131" json:"linkedTrafficEndpointID,omitempty"`
	Proxying                *pfcpType.Proxying              `tlv:"137" json:"proxying,omitempty"`
}

type DuplicatingParametersIEInFAR struct {
	DestinationInterface  *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface,omitempty"`
	OuterHeaderCreation   *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy      *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
}

type CreateURR struct {
	URRID                     *pfcpType.URRID                     `tlv:"81" json:"urrid,omitempty"`
	MeasurementMethod         *pfcpType.MeasurementMethod         `tlv:"62" json:"measurementMethod,omitempty"`
	ReportingTriggers         *pfcpType.ReportingTriggers         `tlv:"37" json:"reportingTriggers,omitempty"`
	MeasurementPeriod         *pfcpType.MeasurementPeriod         `tlv:"64" json:"measurementPeriod,omitempty"`
	VolumeThreshold           *pfcpType.VolumeThreshold           `tlv:"31" json:"volumeThreshold,omitempty"`
	VolumeQuota               *pfcpType.VolumeQuota               `tlv:"73" json:"volumeQuota,omitempty"`
	TimeThreshold             *pfcpType.TimeThreshold             `tlv:"32" json:"timeThreshold,omitempty"`
	TimeQuota                 *pfcpType.TimeQuota                 `tlv:"74" json:"timeQuota,omitempty"`
	QuotaHoldingTime          *pfcpType.QuotaHoldingTime          `tlv:"71" json:"quotaHoldingTime,omitempty"`
	DroppedDLTrafficThreshold *pfcpType.DroppedDLTrafficThreshold `tlv:"72" json:"droppedDLTrafficThreshold,omitempty"`
	MonitoringTime            *pfcpType.MonitoringTime            `tlv:"33" json:"monitoringTime,omitempty"`
	EventInformation          *EventInformation                   `tlv:"148" json:"eventInformation,omitempty"`
	SubsequentVolumeThreshold *pfcpType.SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold,omitempty"`
	SubsequentTimeThreshold   *pfcpType.SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold,omitempty"`
	SubsequentVolumeQuota     *pfcpType.SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota,omitempty"`
	SubsequentTimeQuota       *pfcpType.SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota,omitempty"`
	InactivityDetectionTime   *pfcpType.InactivityDetectionTime   `tlv:"36" json:"inactivityDetectionTime,omitempty"`
	LinkedURRID               *pfcpType.LinkedURRID               `tlv:"82" json:"linkedURRID,omitempty"`
	MeasurementInformation    *pfcpType.MeasurementInformation    `tlv:"100" json:"measurementInformation,omitempty"`
	TimeQuotaMechanism        *pfcpType.TimeQuotaMechanism        `tlv:"115" json:"timeQuotaMechanism,omitempty"`
	AggregatedURRs            *AggregatedURRs                     `tlv:"118" json:"aggregatedURRs,omitempty"`
	FARIDForQuotaAction       *pfcpType.FARID                     `tlv:"108" json:"faridForQuotaAction,omitempty"`
	EthernetInactivityTimer   *pfcpType.EthernetInactivityTimer   `tlv:"146" json:"ethernetInactivityTimer,omitempty"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime           `tlv:"147" json:"additionalMonitoringTime,omitempty"`
}

type AggregatedURRs struct {
	AggregatedURRID *pfcpType.AggregatedURRID `tlv:"120" json:"aggregatedURRID,omitempty"`
	Multiplier      *pfcpType.Multiplier      `tlv:"119" json:"multiplier,omitempty"`
}

type AdditionalMonitoringTime struct {
	MonitoringTime            *pfcpType.MonitoringTime            `tlv:"33" json:"monitoringTime,omitempty"`
	SubsequentVolumeThreshold *pfcpType.SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold,omitempty"`
	SubsequentTimeThreshold   *pfcpType.SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold,omitempty"`
	SubsequentVolumeQuota     *pfcpType.SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota,omitempty"`
	SubsequentTimeQuota       *pfcpType.SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota,omitempty"`
}

type EventInformation struct {
	EventID        *pfcpType.EventID        `tlv:"150" json:"eventID,omitempty"`
	EventThreshold *pfcpType.EventThreshold `tlv:"151" json:"eventThreshold,omitempty"`
}

type CreateQER struct {
	QERID              *pfcpType.QERID              `tlv:"109" json:"qerid,omitempty"`
	QERCorrelationID   *pfcpType.QERCorrelationID   `tlv:"28" json:"qerCorrelationID,omitempty"`
	GateStatus         *pfcpType.GateStatus         `tlv:"25" json:"gateStatus,omitempty"`
	MaximumBitrate     *pfcpType.MBR                `tlv:"26" json:"maximumBitrate,omitempty"`
	GuaranteedBitrate  *pfcpType.GBR                `tlv:"27" json:"guaranteedBitrate,omitempty"`
	PacketRate         *pfcpType.PacketRate         `tlv:"94" json:"packetRate,omitempty"`
	DLFlowLevelMarking *pfcpType.DLFlowLevelMarking `tlv:"97" json:"dlFlowLevelMarking,omitempty"`
	QoSFlowIdentifier  *pfcpType.QFI                `tlv:"124" json:"qoSFlowIdentifier,omitempty"`
	ReflectiveQoS      *pfcpType.RQI                `tlv:"123" json:"reflectiveQoS,omitempty"`
}

type CreateBAR struct {
	BARID                          *pfcpType.BARID                          `tlv:"88" json:"barid,omitempty"`
	DownlinkDataNotificationDelay  *pfcpType.DownlinkDataNotificationDelay  `tlv:"46" json:"downlinkDataNotificationDelay,omitempty"`
	SuggestedBufferingPacketsCount *pfcpType.SuggestedBufferingPacketsCount `tlv:"140" json:"suggestedBufferingPacketsCount,omitempty"`
}

type CreateTrafficEndpoint struct {
	TrafficEndpointID             *pfcpType.TrafficEndpointID             `tlv:"131" json:"trafficEndpointID,omitempty"`
	LocalFTEID                    *pfcpType.FTEID                         `tlv:"21" json:"localFTEID,omitempty"`
	NetworkInstance               *pfcpType.Dnn                           `tlv:"22" json:"networkInstance,omitempty"`
	UEIPAddress                   *pfcpType.UEIPAddress                   `tlv:"93" json:"ueipAddress,omitempty"`
	EthernetPDUSessionInformation *pfcpType.EthernetPDUSessionInformation `tlv:"142" json:"ethernetPDUSessionInformation,omitempty"`
	FramedRoute                   *pfcpType.FramedRoute                   `tlv:"153" json:"framedRoute,omitempty"`
	FramedRouting                 *pfcpType.FramedRouting                 `tlv:"154" json:"framedRouting,omitempty"`
	FramedIPv6Route               *pfcpType.FramedIPv6Route               `tlv:"155" json:"framedIPv6Route,omitempty"`
}

type PFCPSessionEstablishmentResponse struct {
	NodeID                     *pfcpType.NodeID            `tlv:"60" json:"nodeID,omitempty"`
	Cause                      *pfcpType.Cause             `tlv:"19" json:"cause,omitempty"`
	OffendingIE                *pfcpType.OffendingIE       `tlv:"40" json:"offendingIE,omitempty"`
	UPFSEID                    *pfcpType.FSEID             `tlv:"57" json:"upfseid,omitempty"`
	CreatedPDR                 *CreatedPDR                 `tlv:"8" json:"createdPDR,omitempty"`
	LoadControlInformation     *LoadControlInformation     `tlv:"51" json:"loadControlInformation,omitempty"`
	OverloadControlInformation *OverloadControlInformation `tlv:"54" json:"overloadControlInformation,omitempty"`
	SGWUFQCSID                 *pfcpType.FQCSID            `tlv:"65" json:"sgwufqcsid,omitempty"`
	PGWUFQCSID                 *pfcpType.FQCSID            `tlv:"65" json:"pgwufqcsid,omitempty"`
	FailedRuleID               *pfcpType.FailedRuleID      `tlv:"114" json:"failedRuleID,omitempty"`
	CreatedTrafficEndpoint     *CreatedTrafficEndpoint     `tlv:"128" json:"createdTrafficEndpoint,omitempty"`
}

type CreatedPDR struct {
	PDRID      *pfcpType.PacketDetectionRuleID `tlv:"56" json:"pdrid,omitempty"`
	LocalFTEID *pfcpType.FTEID                 `tlv:"21" json:"localFTEID,omitempty"`
}

type LoadControlInformation struct {
	LoadControlSequenceNumber *pfcpType.SequenceNumber `tlv:"52" json:"loadControlSequenceNumber,omitempty"`
	LoadMetric                *pfcpType.Metric         `tlv:"53" json:"loadMetric,omitempty"`
}

type OverloadControlInformation struct {
	OverloadControlSequenceNumber   *pfcpType.SequenceNumber `tlv:"52" json:"overloadControlSequenceNumber,omitempty"`
	OverloadReductionMetric         *pfcpType.Metric         `tlv:"53" json:"overloadReductionMetric,omitempty"`
	PeriodOfValidity                *pfcpType.Timer          `tlv:"55" json:"periodOfValidity,omitempty"`
	OverloadControlInformationFlags *pfcpType.OCIFlags       `tlv:"110" json:"overloadControlInformationFlags,omitempty"`
}

type CreatedTrafficEndpoint struct {
	TrafficEndpointID *pfcpType.TrafficEndpointID `tlv:"131" json:"trafficEndpointID,omitempty"`
	LocalFTEID        *pfcpType.FTEID             `tlv:"21" json:"localFTEID,omitempty"`
}

type PFCPSessionModificationRequest struct {
	CPFSEID                  *pfcpType.FSEID                          `tlv:"57" json:"cpfseid,omitempty"`
	RemovePDR                []*RemovePDR                             `tlv:"15" json:"removePDR,omitempty"`
	RemoveFAR                []*RemoveFAR                             `tlv:"16" json:"removeFAR,omitempty"`
	RemoveURR                []*RemoveURR                             `tlv:"17" json:"removeURR,omitempty"`
	RemoveQER                []*pfcpType.RemoveQER                    `tlv:"18" json:"removeQER,omitempty"`
	RemoveBAR                []*RemoveBAR                             `tlv:"87" json:"removeBAR,omitempty"`
	RemoveTrafficEndpoint    *RemoveTrafficEndpoint                   `tlv:"130" json:"removeTrafficEndpoint,omitempty"`
	CreatePDR                []*CreatePDR                             `tlv:"1" json:"createPDR,omitempty"`
	CreateFAR                []*CreateFAR                             `tlv:"3" json:"createFAR,omitempty"`
	CreateURR                []*CreateURR                             `tlv:"6" json:"createURR,omitempty"`
	CreateQER                []*CreateQER                             `tlv:"7" json:"createQER,omitempty"`
	CreateBAR                []*CreateBAR                             `tlv:"85" json:"createBAR,omitempty"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint                   `tlv:"127" json:"createTrafficEndpoint,omitempty"`
	UpdatePDR                []*UpdatePDR                             `tlv:"9" json:"updatePDR,omitempty"`
	UpdateFAR                []*UpdateFAR                             `tlv:"10" json:"updateFAR,omitempty"`
	UpdateURR                []*UpdateURR                             `tlv:"13" json:"updateURR,omitempty"`
	UpdateQER                []*UpdateQER                             `tlv:"14" json:"updateQER,omitempty"`
	UpdateBAR                *UpdateBARPFCPSessionModificationRequest `tlv:"86" json:"updateBAR,omitempty"`
	UpdateTrafficEndpoint    *UpdateTrafficEndpoint                   `tlv:"129" json:"updateTrafficEndpoint,omitempty"`
	PFCPSMReqFlags           *pfcpType.PFCPSMReqFlags                 `tlv:"49" json:"pfcpsmReqFlags,omitempty"`
	QueryURR                 *QueryURR                                `tlv:"77" json:"queryURR,omitempty"`
	PGWCFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"pgwcfqcsid,omitempty"`
	SGWCFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"sgwcfqcsid,omitempty"`
	MMEFQCSID                *pfcpType.FQCSID                         `tlv:"65" json:"mmefqcsid,omitempty"`
	EPDGFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"epdgfqcsid,omitempty"`
	TWANFQCSID               *pfcpType.FQCSID                         `tlv:"65" json:"twanfqcsid,omitempty"`
	UserPlaneInactivityTimer *pfcpType.UserPlaneInactivityTimer       `tlv:"117" json:"userPlaneInactivityTimer,omitempty"`
	QueryURRReference        *pfcpType.QueryURRReference              `tlv:"125" json:"queryURRReference,omitempty"`
	TraceInformation         *pfcpType.TraceInformation               `tlv:"152" json:"traceInformation,omitempty"`
}

type UpdatePDR struct {
	PDRID                     *pfcpType.PacketDetectionRuleID     `tlv:"56" json:"pdrid,omitempty"`
	OuterHeaderRemoval        *pfcpType.OuterHeaderRemoval        `tlv:"95" json:"outerHeaderRemoval,omitempty"`
	Precedence                *pfcpType.Precedence                `tlv:"29" json:"precedence,omitempty"`
	PDI                       *PDI                                `tlv:"2" json:"pdi,omitempty"`
	FARID                     *pfcpType.FARID                     `tlv:"108" json:"farid,omitempty"`
	URRID                     *pfcpType.URRID                     `tlv:"81" json:"urrid,omitempty"`
	QERID                     *pfcpType.QERID                     `tlv:"109" json:"qerid,omitempty"`
	ActivatePredefinedRules   *pfcpType.ActivatePredefinedRules   `tlv:"106" json:"activatePredefinedRules,omitempty"`
	DeactivatePredefinedRules *pfcpType.DeactivatePredefinedRules `tlv:"107" json:"deactivatePredefinedRules,omitempty"`
}

type UpdateFAR struct {
	FARID                       *pfcpType.FARID                       `tlv:"108" json:"farid,omitempty"`
	ApplyAction                 *pfcpType.ApplyAction                 `tlv:"44" json:"applyAction,omitempty"`
	UpdateForwardingParameters  *UpdateForwardingParametersIEInFAR    `tlv:"11" json:"updateForwardingParameters,omitempty"`
	UpdateDuplicatingParameters *pfcpType.UpdateDuplicatingParameters `tlv:"105" json:"updateDuplicatingParameters,omitempty"`
	BARID                       *pfcpType.BARID                       `tlv:"88" json:"barid,omitempty"`
}

type UpdateForwardingParametersIEInFAR struct {
	DestinationInterface    *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface,omitempty"`
	NetworkInstance         *pfcpType.Dnn                   `tlv:"22" json:"networkInstance,omitempty"`
	RedirectInformation     *pfcpType.RedirectInformation   `tlv:"38" json:"redirectInformation,omitempty"`
	OuterHeaderCreation     *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking   *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy        *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
	HeaderEnrichment        *pfcpType.HeaderEnrichment      `tlv:"98" json:"headerEnrichment,omitempty"`
	PFCPSMReqFlags          *pfcpType.PFCPSMReqFlags        `tlv:"49" json:"pfcpsmReqFlags,omitempty"`
	LinkedTrafficEndpointID *pfcpType.TrafficEndpointID     `tlv:"131" json:"linkedTrafficEndpointID,omitempty"`
}

type UpdateDuplicatingParametersIEInFAR struct {
	DestinationInterface  *pfcpType.DestinationInterface  `tlv:"42" json:"destinationInterface,omitempty"`
	OuterHeaderCreation   *pfcpType.OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking *pfcpType.TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy      *pfcpType.ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
}

type UpdateURR struct {
	URRID                     *pfcpType.URRID                     `tlv:"81" json:"urrid,omitempty"`
	MeasurementMethod         *pfcpType.MeasurementMethod         `tlv:"62" json:"measurementMethod,omitempty"`
	ReportingTriggers         *pfcpType.ReportingTriggers         `tlv:"37" json:"reportingTriggers,omitempty"`
	MeasurementPeriod         *pfcpType.MeasurementPeriod         `tlv:"64" json:"measurementPeriod,omitempty"`
	VolumeThreshold           *pfcpType.VolumeThreshold           `tlv:"31" json:"volumeThreshold,omitempty"`
	VolumeQuota               *pfcpType.VolumeQuota               `tlv:"73" json:"volumeQuota,omitempty"`
	TimeThreshold             *pfcpType.TimeThreshold             `tlv:"32" json:"timeThreshold,omitempty"`
	TimeQuota                 *pfcpType.TimeQuota                 `tlv:"74" json:"timeQuota,omitempty"`
	QuotaHoldingTime          *pfcpType.QuotaHoldingTime          `tlv:"71" json:"quotaHoldingTime,omitempty"`
	DroppedDLTrafficThreshold *pfcpType.DroppedDLTrafficThreshold `tlv:"72" json:"droppedDLTrafficThreshold,omitempty"`
	MonitoringTime            *pfcpType.MonitoringTime            `tlv:"33" json:"monitoringTime,omitempty"`
	EventInformation          *EventInformation                   `tlv:"148" json:"eventInformation,omitempty"`
	SubsequentVolumeThreshold *pfcpType.SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold,omitempty"`
	SubsequentTimeThreshold   *pfcpType.SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold,omitempty"`
	SubsequentVolumeQuota     *pfcpType.SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota,omitempty"`
	SubsequentTimeQuota       *pfcpType.SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota,omitempty"`
	InactivityDetectionTime   *pfcpType.InactivityDetectionTime   `tlv:"36" json:"inactivityDetectionTime,omitempty"`
	LinkedURRID               *pfcpType.LinkedURRID               `tlv:"82" json:"linkedURRID,omitempty"`
	MeasurementInformation    *pfcpType.MeasurementInformation    `tlv:"100" json:"measurementInformation,omitempty"`
	TimeQuotaMechanism        *pfcpType.TimeQuotaMechanism        `tlv:"115" json:"timeQuotaMechanism,omitempty"`
	AggregatedURRs            *AggregatedURRs                     `tlv:"118" json:"aggregatedURRs,omitempty"`
	FARIDForQuotaAction       *pfcpType.FARID                     `tlv:"108" json:"faridForQuotaAction,omitempty"`
	EthernetInactivityTimer   *pfcpType.EthernetInactivityTimer   `tlv:"146" json:"ethernetInactivityTimer,omitempty"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime           `tlv:"147" json:"additionalMonitoringTime,omitempty"`
}

type UpdateQER struct {
	QERID              *pfcpType.QERID              `tlv:"109" json:"qerid,omitempty"`
	QERCorrelationID   *pfcpType.QERCorrelationID   `tlv:"28" json:"qerCorrelationID,omitempty"`
	GateStatus         *pfcpType.GateStatus         `tlv:"25" json:"gateStatus,omitempty"`
	MaximumBitrate     *pfcpType.MBR                `tlv:"26" json:"maximumBitrate,omitempty"`
	GuaranteedBitrate  *pfcpType.GBR                `tlv:"27" json:"guaranteedBitrate,omitempty"`
	PacketRate         *pfcpType.PacketRate         `tlv:"94" json:"packetRate,omitempty"`
	DLFlowLevelMarking *pfcpType.DLFlowLevelMarking `tlv:"97" json:"dlFlowLevelMarking,omitempty"`
	QoSFlowIdentifier  *pfcpType.QFI                `tlv:"124" json:"qoSFlowIdentifier,omitempty"`
	ReflectiveQoS      *pfcpType.RQI                `tlv:"123" json:"reflectiveQoS,omitempty"`
}

type RemovePDR struct {
	PDRID *pfcpType.PacketDetectionRuleID `tlv:"56" json:"pdrid,omitempty"`
}

type RemoveFAR struct {
	FARID *pfcpType.FARID `tlv:"108" json:"farid,omitempty"`
}

type RemoveURR struct {
	URRID *pfcpType.URRID `tlv:"81" json:"urrid,omitempty"`
}

type RemoveQERIEPFCPSessionModificationRequest struct {
	QERID *pfcpType.QERID `tlv:"109" json:"qerid,omitempty"`
}

type QueryURR struct {
	URRID *pfcpType.URRID `tlv:"81" json:"urrid,omitempty"`
}

type UpdateBARPFCPSessionModificationRequest struct {
	BARID                          *pfcpType.BARID                          `tlv:"88" json:"barid,omitempty"`
	DownlinkDataNotificationDelay  *pfcpType.DownlinkDataNotificationDelay  `tlv:"46" json:"downlinkDataNotificationDelay,omitempty"`
	SuggestedBufferingPacketsCount *pfcpType.SuggestedBufferingPacketsCount `tlv:"140" json:"suggestedBufferingPacketsCount,omitempty"`
}

type RemoveBAR struct {
	BARID *pfcpType.BARID `tlv:"88" json:"barid,omitempty"`
}

type UpdateTrafficEndpoint struct {
	TrafficEndpointID *pfcpType.TrafficEndpointID `tlv:"131" json:"trafficEndpointID,omitempty"`
	LocalFTEID        *pfcpType.FTEID             `tlv:"21" json:"localFTEID,omitempty"`
	NetworkInstance   *pfcpType.Dnn               `tlv:"22" json:"networkInstance,omitempty"`
	UEIPAddress       *pfcpType.UEIPAddress       `tlv:"93" json:"ueipAddress,omitempty"`
	FramedRoute       *pfcpType.FramedRoute       `tlv:"153" json:"framedRoute,omitempty"`
	FramedRouting     *pfcpType.FramedRouting     `tlv:"154" json:"framedRouting,omitempty"`
	FramedIPv6Route   *pfcpType.FramedIPv6Route   `tlv:"155" json:"framedIPv6Route,omitempty"`
}

type RemoveTrafficEndpoint struct {
	TrafficEndpointID *pfcpType.TrafficEndpointID `tlv:"131" json:"trafficEndpointID,omitempty"`
}

type PFCPSessionModificationResponse struct {
	Cause                             *pfcpType.Cause                             `tlv:"19" json:"cause,omitempty"`
	OffendingIE                       *pfcpType.OffendingIE                       `tlv:"40" json:"offendingIE,omitempty"`
	CreatedPDR                        *CreatedPDR                                 `tlv:"8" json:"createdPDR,omitempty"`
	LoadControlInformation            *LoadControlInformation                     `tlv:"51" json:"loadControlInformation,omitempty"`
	OverloadControlInformation        *OverloadControlInformation                 `tlv:"54" json:"overloadControlInformation,omitempty"`
	UsageReport                       *UsageReportPFCPSessionModificationResponse `tlv:"78" json:"usageReport,omitempty"`
	FailedRuleID                      *pfcpType.FailedRuleID                      `tlv:"114" json:"failedRuleID,omitempty"`
	AdditionalUsageReportsInformation *pfcpType.AdditionalUsageReportsInformation `tlv:"126" json:"additionalUsageReportsInformation,omitempty"`
	CreatedUpdatedTrafficEndpoint     *CreatedTrafficEndpoint                     `tlv:"128" json:"createdUpdatedTrafficEndpoint,omitempty"`
}

type UsageReportPFCPSessionModificationResponse struct {
	URRID                      *pfcpType.URRID               `tlv:"81" json:"urrid,omitempty"`
	URSEQN                     *pfcpType.URSEQN              `tlv:"104" json:"urseqn,omitempty"`
	UsageReportTrigger         *pfcpType.UsageReportTrigger  `tlv:"63" json:"usageReportTrigger,omitempty"`
	StartTime                  *pfcpType.StartTime           `tlv:"75" json:"startTime,omitempty"`
	EndTime                    *pfcpType.EndTime             `tlv:"76" json:"endTime,omitempty"`
	VolumeMeasurement          *pfcpType.VolumeMeasurement   `tlv:"66" json:"volumeMeasurement,omitempty"`
	DurationMeasurement        *pfcpType.DurationMeasurement `tlv:"67" json:"durationMeasurement,omitempty"`
	TimeOfFirstPacket          *pfcpType.TimeOfFirstPacket   `tlv:"69" json:"timeOfFirstPacket,omitempty"`
	TimeOfLastPacket           *pfcpType.TimeOfLastPacket    `tlv:"70" json:"timeOfLastPacket,omitempty"`
	UsageInformation           *pfcpType.UsageInformation    `tlv:"90" json:"usageInformation,omitempty"`
	QueryURRReference          *pfcpType.QueryURRReference   `tlv:"125" json:"queryURRReference,omitempty"`
	EthernetTrafficInformation *EthernetTrafficInformation   `tlv:"143" json:"ethernetTrafficInformation,omitempty"`
}

type PFCPSessionDeletionRequest struct{}

type PFCPSessionDeletionResponse struct {
	Cause                      *pfcpType.Cause                         `tlv:"19" json:"cause,omitempty"`
	OffendingIE                *pfcpType.OffendingIE                   `tlv:"40" json:"offendingIE,omitempty"`
	LoadControlInformation     *LoadControlInformation                 `tlv:"51" json:"loadControlInformation,omitempty"`
	OverloadControlInformation *OverloadControlInformation             `tlv:"54" json:"overloadControlInformation,omitempty"`
	UsageReport                *UsageReportPFCPSessionDeletionResponse `tlv:"79" json:"usageReport,omitempty"`
}

type UsageReportPFCPSessionDeletionResponse struct {
	URRID                      *pfcpType.URRID               `tlv:"81" json:"urrid,omitempty"`
	URSEQN                     *pfcpType.URSEQN              `tlv:"104" json:"urseqn,omitempty"`
	UsageReportTrigger         *pfcpType.UsageReportTrigger  `tlv:"63" json:"usageReportTrigger,omitempty"`
	StartTime                  *pfcpType.StartTime           `tlv:"75" json:"startTime,omitempty"`
	EndTime                    *pfcpType.EndTime             `tlv:"76" json:"endTime,omitempty"`
	VolumeMeasurement          *pfcpType.VolumeMeasurement   `tlv:"66" json:"volumeMeasurement,omitempty"`
	DurationMeasurement        *pfcpType.DurationMeasurement `tlv:"67" json:"durationMeasurement,omitempty"`
	TimeOfFirstPacket          *pfcpType.TimeOfFirstPacket   `tlv:"69" json:"timeOfFirstPacket,omitempty"`
	TimeOfLastPacket           *pfcpType.TimeOfLastPacket    `tlv:"70" json:"timeOfLastPacket,omitempty"`
	UsageInformation           *pfcpType.UsageInformation    `tlv:"90" json:"usageInformation,omitempty"`
	EthernetTrafficInformation *EthernetTrafficInformation   `tlv:"143" json:"ethernetTrafficInformation,omitempty"`
}

type PFCPSessionReportRequest struct {
	ReportType                        *pfcpType.ReportType                        `tlv:"39" json:"reportType,omitempty"`
	DownlinkDataReport                *DownlinkDataReport                         `tlv:"83" json:"downlinkDataReport,omitempty"`
	UsageReport                       *UsageReportPFCPSessionReportRequest        `tlv:"80" json:"usageReport,omitempty"`
	ErrorIndicationReport             *ErrorIndicationReport                      `tlv:"99" json:"errorIndicationReport,omitempty"`
	LoadControlInformation            *LoadControlInformation                     `tlv:"51" json:"loadControlInformation,omitempty"`
	OverloadControlInformation        *OverloadControlInformation                 `tlv:"54" json:"overloadControlInformation,omitempty"`
	AdditionalUsageReportsInformation *pfcpType.AdditionalUsageReportsInformation `tlv:"126" json:"additionalUsageReportsInformation,omitempty"`
}

type DownlinkDataReport struct {
	PDRID                          *pfcpType.PacketDetectionRuleID          `tlv:"56" json:"pdrid,omitempty"`
	DownlinkDataServiceInformation *pfcpType.DownlinkDataServiceInformation `tlv:"45" json:"downlinkDataServiceInformation,omitempty"`
}

type UsageReportPFCPSessionReportRequest struct {
	URRID                           *pfcpType.URRID                  `tlv:"81" json:"urrid,omitempty"`
	URSEQN                          *pfcpType.URSEQN                 `tlv:"104" json:"urseqn,omitempty"`
	UsageReportTrigger              *pfcpType.UsageReportTrigger     `tlv:"63" json:"usageReportTrigger,omitempty"`
	StartTime                       *pfcpType.StartTime              `tlv:"75" json:"startTime,omitempty"`
	EndTime                         *pfcpType.EndTime                `tlv:"76" json:"endTime,omitempty"`
	VolumeMeasurement               *pfcpType.VolumeMeasurement      `tlv:"66" json:"volumeMeasurement,omitempty"`
	DurationMeasurement             *pfcpType.DurationMeasurement    `tlv:"67" json:"durationMeasurement,omitempty"`
	ApplicationDetectionInformation *ApplicationDetectionInformation `tlv:"68" json:"applicationDetectionInformation,omitempty"`
	UEIPAddress                     *pfcpType.UEIPAddress            `tlv:"93" json:"ueipAddress,omitempty"`
	NetworkInstance                 *pfcpType.Dnn                    `tlv:"22" json:"networkInstance,omitempty"`
	TimeOfFirstPacket               *pfcpType.TimeOfFirstPacket      `tlv:"69" json:"timeOfFirstPacket,omitempty"`
	TimeOfLastPacket                *pfcpType.TimeOfLastPacket       `tlv:"70" json:"timeOfLastPacket,omitempty"`
	UsageInformation                *pfcpType.UsageInformation       `tlv:"90" json:"usageInformation,omitempty"`
	QueryURRReference               *pfcpType.QueryURRReference      `tlv:"125" json:"queryURRReference,omitempty"`
	EventReporting                  *EventReporting                  `tlv:"149" json:"eventReporting,omitempty"`
	EthernetTrafficInformation      *EthernetTrafficInformation      `tlv:"143" json:"ethernetTrafficInformation,omitempty"`
}

type ApplicationDetectionInformation struct {
	ApplicationID         *pfcpType.ApplicationID         `tlv:"24" json:"applicationID,omitempty"`
	ApplicationInstanceID *pfcpType.ApplicationInstanceID `tlv:"91" json:"applicationInstanceID,omitempty"`
	FlowInformation       *pfcpType.FlowInformation       `tlv:"92" json:"flowInformation,omitempty"`
}

type EventReporting struct {
	EventID *pfcpType.EventID `tlv:"150" json:"eventID,omitempty"`
}

type EthernetTrafficInformation struct {
	MACAddressesDetected *pfcpType.MACAddressesDetected `tlv:"144" json:"macAddressesDetected,omitempty"`
	MACAddressesRemoved  *pfcpType.MACAddressesRemoved  `tlv:"145" json:"macAddressesRemoved,omitempty"`
}

type ErrorIndicationReport struct {
	RemoteFTEID *pfcpType.FTEID `tlv:"21" json:"remoteFTEID,omitempty"`
}

type PFCPSessionReportResponse struct {
	Cause        *pfcpType.Cause                              `tlv:"19" json:"cause,omitempty"`
	OffendingIE  *pfcpType.OffendingIE                        `tlv:"40" json:"offendingIE,omitempty"`
	UpdateBAR    *pfcpType.UpdateBARPFCPSessionReportResponse `tlv:"12" json:"updateBAR,omitempty"`
	SxSRRspFlags *pfcpType.PFCPSRRspFlags                     `tlv:"50" json:"sxSRRspFlags,omitempty"`
}

type UpdateBARIEInPFCPSessionReportResponse struct {
	BARID                           *pfcpType.BARID                           `tlv:"88" json:"barid,omitempty"`
	DownlinkDataNotificationDelay   *pfcpType.DownlinkDataNotificationDelay   `tlv:"46" json:"downlinkDataNotificationDelay,omitempty"`
	DLBufferingDuration             *pfcpType.DLBufferingDuration             `tlv:"47" json:"dlBufferingDuration,omitempty"`
	DLBufferingSuggestedPacketCount *pfcpType.DLBufferingSuggestedPacketCount `tlv:"48" json:"dlBufferingSuggestedPacketCount,omitempty"`
	SuggestedBufferingPacketsCount  *pfcpType.SuggestedBufferingPacketsCount  `tlv:"140" json:"suggestedBufferingPacketsCount,omitempty"`
}
