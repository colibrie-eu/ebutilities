// Models for http://www.ebutilities.at/schemata/customerconsent/cmrequest/01p20
package cmrequest

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMRequest struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `xml:"CMRequest"`

	MarketParticipantDirectory MarketParticipantDirectory `xml:"MarketParticipantDirectory"`

	ProcessDirectory ProcessDirectory `xml:"ProcessDirectory"`
}

// XSD ComplexType declarations

type MarketParticipantDirectory struct {
	XMLName xml.Name

	DocumentMode ct.DocumentMode `xml:"DocumentMode,attr"`

	Duplicate bool `xml:"Duplicate,attr"`

	SchemaVersion string `xml:"SchemaVersion,attr"`

	RoutingHeader ct.RoutingHeader `xml:"ct:RoutingHeader"`

	Sector ct.Sector `xml:"ct:Sector"`

	MessageCode ct.MessageCode `xml:"MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name

	MessageID ct.GroupingID `xml:"ct:MessageId"`

	ConversationID ct.GroupingID `xml:"ct:ConversationId"`

	ProcessDate string `xml:"ProcessDate"`

	MeteringPoint *ct.MeteringPoint `xml:"MeteringPoint"`

	CMRequestID ct.GroupingID `xml:"CMRequestId"`

	ConsentID *ct.GroupingID `xml:"ct:ConsentId"`

	CMRequest ReqType `xml:"CMRequest"`
}

type ReqType struct {
	XMLName xml.Name

	ReqDatType ReqDatType `xml:"ReqDatType"`

	DateFrom string `xml:"DateFrom"`

	DateTo string `xml:"DateTo"`

	MeteringIntervall *MeteringIntervallType `xml:"MeteringIntervall"`

	TransmissionCycle *TransmissionCycle `xml:"TransmissionCycle"`

	ECID *ct.MeteringPoint `xml:"ECID"`

	ECPartFact *ECPartFact `xml:"ECPartFact"`

	ECShare *ECShare `xml:"ECShare"`

	EnergyDirection *EnergyDirection `xml:"EnergyDirection"`
}

type ReqDatParamType struct {
	XMLName xml.Name

	ParamCyc *ParamCycType `xml:"ParamCyc"`

	ParamHist *ParamHistType `xml:"ParamHist"`
}

type ParamHistType struct {
	XMLName xml.Name

	MeteringIntervall MeteringIntervallType `xml:",any"`
}

type ParamCycType struct {
	XMLName xml.Name

	MeteringIntervall MeteringIntervallType `xml:"MeteringIntervall"`

	TransmissionCycle TransmissionCycle `xml:"TransmissionCycle"`
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
