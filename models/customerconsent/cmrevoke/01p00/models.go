// Models for http://www.ebutilities.at/schemata/customerconsent/cmrevoke/01p00
package cmrevoke

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMRevoke struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `xml:"CMRevoke"`

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

	MessageCode string `xml:"MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name

	MessageID ct.GroupingID `xml:"ct:MessageId"`

	ConversationID ct.GroupingID `xml:"ct:ConversationId"`

	ConsentID ct.GroupingID `xml:"ConsentId"`

	MeteringPoint ct.MeteringPoint `xml:"MeteringPoint"`

	ConsentEnd string `xml:"ConsentEnd"`

	Reason *ReasonType `xml:"Reason"`
}

// XSD SimpleType declarations

type ReasonType string
