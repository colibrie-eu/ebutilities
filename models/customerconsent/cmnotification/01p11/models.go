// Models for http://www.ebutilities.at/schemata/customerconsent/cmnotification/01p11
package cmnotification

import (
	"encoding/xml"

	ct "github.com/WattWiseAt/ebutilities/models/customerprocesses/common/types/01p20"
)

type CMNotification struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNs2 string `xml:"xmlns:ns2,attr"`

	XMLName xml.Name `xml:"ns2:CMNotification"`

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

	CMRequestID ct.GroupingID `xml:"ns2:CMRequestId"`

	ResponseData []ResponseDataType `xml:"ResponseData"`
}

type ResponseDataType struct {
	XMLName xml.Name

	ConsentID *ct.GroupingID `xml:"ConsentId"`

	MeteringPoint *ct.MeteringPoint `xml:"MeteringPoint"`

	ResponseCode []ResponseCode `xml:"ResponseCode"`

	ParamHist []ParamHistType `xml:"ParamHist"`
}

type ParamHistType struct {
	XMLName xml.Name

	DateFrom string `xml:"DateFrom"`

	DateTo string `xml:"DateTo"`

	MeteringIntervall MeteringIntervallType `xml:"MeteringIntervall"`
}

// XSD SimpleType declarations

type MeteringIntervallType string

const MeteringIntervallTypeQh MeteringIntervallType = "QH"

const MeteringIntervallTypeH MeteringIntervallType = "H"

const MeteringIntervallTypeD MeteringIntervallType = "D"

const MeteringIntervallTypeV MeteringIntervallType = "V"

type ResponseCode int64
