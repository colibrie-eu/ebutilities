// Models for http://www.ebutilities.at/schemata/customerconsent/cmrequest/01p21
package cmrequest

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMRequest struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `json:"-" xml:"CMRequest"`

	MarketParticipantDirectory MarketParticipantDirectory `db:"market_participant_directory" json:"marketParticipantDirectory" xml:"MarketParticipantDirectory"`

	ProcessDirectory ProcessDirectory `db:"process_directory" json:"processDirectory" xml:"ProcessDirectory"`
}

// XSD ComplexType declarations

type MarketParticipantDirectory struct {
	XMLName xml.Name `json:"-"`

	DocumentMode ct.DocumentMode `db:"document_mode" json:"documentMode,omitempty" xml:"DocumentMode,attr"`

	Duplicate bool `db:"duplicate" json:"duplicate,omitempty" xml:"Duplicate,attr"`

	SchemaVersion string `db:"schema_version" json:"schemaVersion,omitempty" xml:"SchemaVersion,attr"`

	RoutingHeader ct.RoutingHeader `db:"routing_header" json:"routingHeader" xml:"ct:RoutingHeader"`

	Sector ct.Sector `db:"sector" json:"sector,omitempty" xml:"ct:Sector"`

	MessageCode ct.MessageCode `db:"message_code" json:"messageCode,omitempty" xml:"MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name `json:"-"`

	MessageID ct.GroupingID `db:"message_id" json:"messageId,omitempty" xml:"ct:MessageId"`

	ConversationID ct.GroupingID `db:"conversation_id" json:"conversationId,omitempty" xml:"ct:ConversationId"`

	ProcessDate string `db:"process_date" json:"processDate,omitempty" xml:"ProcessDate"`

	MeteringPoint *ct.MeteringPoint `db:"metering_point" json:"meteringPoint,omitempty" xml:"MeteringPoint"`

	CMRequestID ct.GroupingID `db:"cm_request_id" json:"cmRequestId,omitempty" xml:"CMRequestId"`

	ConsentID *ct.GroupingID `db:"consent_id" json:"consentId,omitempty" xml:"ct:ConsentId"`

	CMRequest *ReqType `db:"cm_request" json:"cmRequest,omitempty" xml:"CMRequest"`

	AdditionalData []ct.AdditionalData `db:"additional_data" json:"additionalData,omitempty" xml:"AdditionalData"`
}

type ReqType struct {
	XMLName xml.Name `json:"-"`

	ReqDatType ReqDatType `db:"req_dat_type" json:"reqDatType,omitempty" xml:"ReqDatType"`

	DateFrom string `db:"date_from" json:"dateFrom,omitempty" xml:"DateFrom"`

	DateTo *string `db:"date_to" json:"dateTo,omitempty" xml:"DateTo"`

	MeteringIntervall *MeteringIntervallType `db:"metering_intervall" json:"meteringIntervall,omitempty" xml:"MeteringIntervall"`

	TransmissionCycle *TransmissionCycle `db:"transmission_cycle" json:"transmissionCycle,omitempty" xml:"TransmissionCycle"`

	ECID *ct.MeteringPoint `db:"ec_id" json:"ecId,omitempty" xml:"ECID"`

	ECPartFact *ECPartFact `db:"ec_part_fact" json:"ecPartFact,omitempty" xml:"ECPartFact"`

	ECShare *ECShare `db:"ec_share" json:"ecShare,omitempty" xml:"ECShare"`

	EnergyDirection *EnergyDirection `db:"energy_direction" json:"energyDirection,omitempty" xml:"EnergyDirection"`
}

type ReqDatParamType struct {
	XMLName xml.Name `json:"-"`

	ParamCyc *ParamCycType `db:"param_cyc" json:"paramCyc,omitempty" xml:"ParamCyc"`

	ParamHist *ParamHistType `db:"param_hist" json:"paramHist,omitempty" xml:"ParamHist"`
}

type ParamHistType struct {
	XMLName xml.Name `json:"-"`

	MeteringIntervall MeteringIntervallType `db:",any" json:"meteringIntervall,omitempty" xml:",any"`
}

type ParamCycType struct {
	XMLName xml.Name `json:"-"`

	MeteringIntervall MeteringIntervallType `db:"metering_intervall" json:"meteringIntervall,omitempty" xml:"MeteringIntervall"`

	TransmissionCycle TransmissionCycle `db:"transmission_cycle" json:"transmissionCycle,omitempty" xml:"TransmissionCycle"`
}

// XSD SimpleType declarations

type ECPartFact float64

type ECShare float64

type EnergyDirection string

const EnergyDirectionConsumption EnergyDirection = "CONSUMPTION"

const EnergyDirectionGeneration EnergyDirection = "GENERATION"

type MeteringIntervallType string

const MeteringIntervallTypeQh MeteringIntervallType = "QH"

const MeteringIntervallTypeH MeteringIntervallType = "H"

const MeteringIntervallTypeD MeteringIntervallType = "D"

const MeteringIntervallTypeV MeteringIntervallType = "V"

type ReqDatType string

type TransmissionCycle string

const TransmissionCycleD TransmissionCycle = "D"

const TransmissionCycleM TransmissionCycle = "M"

const TransmissionCycleV TransmissionCycle = "V"
