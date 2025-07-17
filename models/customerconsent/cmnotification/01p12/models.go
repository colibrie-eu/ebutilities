// Models for http://www.ebutilities.at/schemata/customerconsent/cmnotification/01p12
package cmnotification

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMNotification struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `xml:"CMNotification"`

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

	CMRequestID ct.GroupingID `xml:"CMRequestId"`

	ResponseData []ResponseDataType `xml:"ResponseData"`

	CustomerData *CustomerDataType `xml:"CustomerData"`
}

type ResponseDataType struct {
	XMLName xml.Name

	ConsentID *ct.GroupingID `xml:"ct:ConsentId"`

	MeteringPoint *ct.MeteringPoint `xml:"MeteringPoint"`

	ResponseCode []ResponseCode `xml:"ResponseCode"`
}

type CustomerDataType struct {
	XMLName xml.Name

	Name1 Name `xml:"Name1"`

	Name2 *Name `xml:"Name2"`
}

// XSD SimpleType declarations

type Name string

type ResponseCode int64
