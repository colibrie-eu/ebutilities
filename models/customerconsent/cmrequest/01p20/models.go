// Models for http://www.ebutilities.at/schemata/customerconsent/cmrequest/01p20
package cmrequest

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMRequest struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNs2 string `xml:"xmlns:ns2,attr"`

	XMLName xml.Name `xml:"ns2:CMRequest"`

	MarketParticipantDirectory MarketParticipantDirectory `xml:"ns2:MarketParticipantDirectory"`

	ProcessDirectory ProcessDirectory `xml:"ns2:ProcessDirectory"`
}

// XSD ComplexType declarations

type MarketParticipantDirectory struct {
	XMLName xml.Name `xml:"ns2:MarketParticipantDirectory"`

	DocumentMode ct.DocumentMode `xml:"DocumentMode,attr"`

	Duplicate bool `xml:"Duplicate,attr"`

	SchemaVersion string `xml:"SchemaVersion,attr"`

	RoutingHeader ct.RoutingHeader `xml:"RoutingHeader"`

	Sector ct.Sector `xml:"Sector"`

	MessageCode ct.MessageCode `xml:"ns2:MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name `xml:"ns2:ProcessDirectory"`

	MessageID ct.GroupingID `xml:"MessageId"`

	ConversationID ct.GroupingID `xml:"ConversationId"`

	ProcessDate string `xml:"ns2:ProcessDate"`

	MeteringPoint ct.MeteringPoint `xml:"ns2:MeteringPoint"`

	CMRequestID ct.GroupingID `xml:"ns2:CMRequestId"`

	ConsentID *ct.GroupingID `xml:"ns2:ConsentId"`

	CMRequest ReqType `xml:"ns2:CMRequest"`
}

type ReqType struct {
	XMLName xml.Name

	ReqDatType ReqDatType `xml:"ns2:ReqDatType"`

	DateFrom string `xml:"ns2:DateFrom"`

	DateTo string `xml:"ns2:DateTo"`

	MeteringIntervall MeteringIntervallType `xml:"ns2:MeteringIntervall"`

	TransmissionCycle TransmissionCycle `xml:"ns2:TransmissionCycle"`

	ECID *ct.MeteringPoint `xml:"ECID"`

	ECPartFact *ECPartFact `xml:"ECPartFact"`

	ECShare *ECShare `xml:"ECShare"`

	EnergyDirection EnergyDirection `xml:"ns2:EnergyDirection"`
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
