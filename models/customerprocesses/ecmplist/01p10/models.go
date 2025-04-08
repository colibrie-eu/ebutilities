// Models for http://www.ebutilities.at/schemata/customerprocesses/ecmplist/01p10
package ecmplist

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type ECMPList struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `xml:"ECMPList"`

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

	MessageID ct.GroupingID `xml:"MessageId"`

	ConversationID ct.GroupingID `xml:"ConversationId"`

	ProcessDate string `xml:"ProcessDate"`

	ECID ct.MeteringPoint `xml:"ECID"`

	ECType ECType `xml:"ECType"`

	ECDisModel ECDisModel `xml:"ECDisModel"`

	MPListData []MPListData `xml:"MPListData"`
}

type MPListData struct {
	XMLName xml.Name

	MeteringPoint ct.MeteringPoint `xml:"MeteringPoint"`

	ConsentID *ct.GroupingID `xml:"ConsentId"`

	MPTimeData []MPTimeData `xml:"MPTimeData"`
}

type MPTimeData struct {
	XMLName xml.Name

	DateFrom string `xml:"DateFrom"`

	DateTo string `xml:"DateTo"`

	EnergyDirection EnergyDirection `xml:"EnergyDirection"`

	ECPartFact ECPartFact `xml:"ECPartFact"`

	PlantCategory *PlantCategory `xml:"PlantCategory"`

	DateActivate string `xml:"DateActivate"`

	DateDeactivate *string `xml:"DateDeactivate"`

	ECShare *ECShare `xml:"ECShare"`

	ECShC []ShareCalc `xml:"ECShC"`
}

type ShareCalc struct {
	XMLName xml.Name

	DateFrom string `xml:"DateFrom"`

	DateTo string `xml:"DateTo"`

	ECShareCalc ECShare `xml:"ECShareCalc"`
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
