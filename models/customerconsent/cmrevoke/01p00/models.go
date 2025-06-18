// Models for http://www.ebutilities.at/schemata/customerconsent/cmrevoke/01p00
package cmrevoke

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMRevoke struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `json:"-" xml:"CMRevoke"`

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

	MessageCode string `db:"message_code" json:"messageCode,omitempty" xml:"MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name `json:"-"`

	MessageID ct.GroupingID `db:"message_id" json:"messageId,omitempty" xml:"ct:MessageId"`

	ConversationID ct.GroupingID `db:"conversation_id" json:"conversationId,omitempty" xml:"ct:ConversationId"`

	ConsentID ct.GroupingID `db:"consent_id" json:"consentId,omitempty" xml:"ConsentId"`

	MeteringPoint ct.MeteringPoint `db:"metering_point" json:"meteringPoint,omitempty" xml:"MeteringPoint"`

	ConsentEnd string `db:"consent_end" json:"consentEnd,omitempty" xml:"ConsentEnd"`

	Reason *ReasonType `db:"reason" json:"reason,omitempty" xml:"Reason"`
}

// XSD SimpleType declarations

type ReasonType string
