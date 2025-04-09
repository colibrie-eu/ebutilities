// Models for http://www.ebutilities.at/schemata/customerprocesses/ecmplist/01p10
package ecmplist

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type ECMPList struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `json:"-" xml:"ECMPList"`

	MarketParticipantDirectory MarketParticipantDirectory `json:"marketParticipantDirectory,omitempty" xml:"MarketParticipantDirectory"`

	ProcessDirectory ProcessDirectory `json:"processDirectory,omitempty" xml:"ProcessDirectory"`
}

// XSD ComplexType declarations

type MarketParticipantDirectory struct {
	XMLName xml.Name `json:"-"`

	DocumentMode ct.DocumentMode `json:"documentMode,omitempty" xml:"DocumentMode,attr"`

	Duplicate bool `json:"duplicate,omitempty" xml:"Duplicate,attr"`

	SchemaVersion string `json:"schemaVersion,omitempty" xml:"SchemaVersion,attr"`

	RoutingHeader ct.RoutingHeader `json:"routingHeader,omitempty" xml:"ct:RoutingHeader"`

	Sector ct.Sector `json:"sector,omitempty" xml:"ct:Sector"`

	MessageCode ct.MessageCode `json:"messageCode,omitempty" xml:"MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name `json:"-"`

	MessageID ct.GroupingID `json:"messageId,omitempty" xml:"MessageId"`

	ConversationID ct.GroupingID `json:"conversationId,omitempty" xml:"ConversationId"`

	ProcessDate string `json:"processDate,omitempty" xml:"ProcessDate"`

	ECID ct.MeteringPoint `json:"ecId,omitempty" xml:"ECID"`

	ECType ECType `json:"ecType,omitempty" xml:"ECType"`

	ECDisModel ECDisModel `json:"ecDisModel,omitempty" xml:"ECDisModel"`

	MPListData []MPListData `json:"mpListData,omitempty" xml:"MPListData"`
}

type MPListData struct {
	XMLName xml.Name `json:"-"`

	MeteringPoint ct.MeteringPoint `json:"meteringPoint,omitempty" xml:"MeteringPoint"`

	ConsentID *ct.GroupingID `json:"consentId,omitempty" xml:"ConsentId"`

	MPTimeData []MPTimeData `json:"mpTimeData,omitempty" xml:"MPTimeData"`
}

type MPTimeData struct {
	XMLName xml.Name `json:"-"`

	DateFrom string `json:"dateFrom,omitempty" xml:"DateFrom"`

	DateTo string `json:"dateTo,omitempty" xml:"DateTo"`

	EnergyDirection EnergyDirection `json:"energyDirection,omitempty" xml:"EnergyDirection"`

	ECPartFact ECPartFact `json:"ecPartFact,omitempty" xml:"ECPartFact"`

	PlantCategory *PlantCategory `json:"plantCategory,omitempty" xml:"PlantCategory"`

	DateActivate string `json:"dateActivate,omitempty" xml:"DateActivate"`

	DateDeactivate *string `json:"dateDeactivate,omitempty" xml:"DateDeactivate"`

	ECShare *ECShare `json:"ecShare,omitempty" xml:"ECShare"`

	ECShC []ShareCalc `json:"ecShC,omitempty" xml:"ECShC"`
}

type ShareCalc struct {
	XMLName xml.Name `json:"-"`

	DateFrom string `json:"dateFrom,omitempty" xml:"DateFrom"`

	DateTo string `json:"dateTo,omitempty" xml:"DateTo"`

	ECShareCalc ECShare `json:"ecShareCalc,omitempty" xml:"ECShareCalc"`
}

// XSD SimpleType declarations

type ECDisModel string

const ECDisModelD ECDisModel = "D"

const ECDisModelS ECDisModel = "S"

type ECShare float64

type ECType string

const ECTypeGc ECType = "GC"

const ECTypeRcR ECType = "RC_R"

const ECTypeRcL ECType = "RC_L"

const ECTypeCc ECType = "CC"

type EnergyDirection string

const EnergyDirectionConsumption EnergyDirection = "CONSUMPTION"

const EnergyDirectionGeneration EnergyDirection = "GENERATION"

type ECPartFact float64

type PlantCategory string
