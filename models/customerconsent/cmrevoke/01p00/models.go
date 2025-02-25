// Models for http://www.ebutilities.at/schemata/customerconsent/cmrevoke/01p00
package cmrevoke

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMRevoke struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNs2 string `xml:"xmlns:ns2,attr"`

	XMLName xml.Name `xml:"ns2:CMRevoke"`

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

	MessageCode string `xml:"ns2:MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name `xml:"ns2:ProcessDirectory"`

	MessageID ct.GroupingID `xml:"MessageId"`

	ConversationID ct.GroupingID `xml:"ConversationId"`

	ConsentID ct.GroupingID `xml:"ns2:ConsentId"`

	MeteringPoint ct.MeteringPoint `xml:"ns2:MeteringPoint"`

	ConsentEnd string `xml:"ns2:ConsentEnd"`

	Reason ReasonType `xml:"ns2:Reason"`
}

// XSD SimpleType declarations

type ReasonType string
