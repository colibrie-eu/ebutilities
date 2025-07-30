// Models for http://www.ebutilities.at/schemata/customerprocesses/cpnotification/01p13
package cpnotification

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CPNotification struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `json:"-" xml:"CPNotification"`

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

	ResponseData ResponseData `db:"response_data" json:"responseData" xml:"ResponseData"`

	AdditionalData []ct.AdditionalData `db:"additional_data" json:"additionalData,omitempty" xml:"ct:AdditionalData"`
}

type ResponseData struct {
	XMLName xml.Name `json:"-"`

	OriginalMessageID ct.GroupingID `db:"original_message_id" json:"originalMessageId,omitempty" xml:"OriginalMessageID"`

	ResponseCode []ResponseCode `db:"response_code" json:"responseCode,omitempty" xml:"ResponseCode"`
}

// XSD SimpleType declarations

type ResponseCode int64
