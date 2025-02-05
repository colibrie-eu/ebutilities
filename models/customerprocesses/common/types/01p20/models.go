// Models for http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20
package ct

import (
	"encoding/xml"
)

// XSD ComplexType declarations

type AdditionalData struct {
	XMLName xml.Name

	Name AdditionalDataName `xml:"Name,attr"`

	Text string `xml:",chardata"`
}

type MarketParticipantDirectory struct {
	XMLName xml.Name

	DocumentMode DocumentMode `xml:"DocumentMode,attr"`

	Duplicate bool `xml:"Duplicate,attr"`

	RoutingHeader RoutingHeader `xml:"RoutingHeader"`

	Sector Sector `xml:"Sector"`
}

type ProcessDirectory struct {
	XMLName xml.Name

	MessageID GroupingID `xml:"MessageId"`

	ConversationID GroupingID `xml:"ConversationId"`

	ProcessDate string `xml:"ProcessDate"`

	MeteringPoint MeteringPoint `xml:"MeteringPoint"`
}

type ProcessDirectoryS struct {
	XMLName xml.Name

	MessageID GroupingID `xml:"MessageId"`

	ConversationID GroupingID `xml:"ConversationId"`
}

type RoutingAddress struct {
	XMLName xml.Name

	AddressType AddressType `xml:"AddressType,attr"`

	MessageAddress MessageAddress `xml:",any"`
}

type RoutingHeader struct {
	XMLName xml.Name

	Sender RoutingAddress `xml:"Sender"`

	Receiver RoutingAddress `xml:"Receiver"`

	DocumentCreationDateTime string `xml:"DocumentCreationDateTime"`
}

type VerificationDocument struct {
	XMLName xml.Name

	Docnumber Docnumber `xml:",any"`
}

// XSD SimpleType declarations

type AdditionalDataName string

type AdditionalDataValue string

type AddressType string

const AddressTypeECNumber AddressType = "ECNumber"

const AddressTypeOther AddressType = "Other"

type DateTimeS string

type DateTimeU string

type Docnumber string

type DocumentMode string

const DocumentModeProd DocumentMode = "PROD"

const DocumentModeSimu DocumentMode = "SIMU"

type GroupingID string

type MessageAddress string

type MessageCode string

type MeteringPoint string

type Sector string

const Sector01 Sector = "01"

const Sector02 Sector = "02"

const Sector03 Sector = "03"

const Sector04 Sector = "04"

const Sector05 Sector = "05"

const Sector06 Sector = "06"

const Sector07 Sector = "07"

const Sector08 Sector = "08"

const Sector09 Sector = "09"

const Sector99 Sector = "99"
