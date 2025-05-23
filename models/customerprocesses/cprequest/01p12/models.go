// Models for http://www.ebutilities.at/schemata/customerprocesses/cprequest/01p12
package cprequest

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type CPRequest struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `xml:"CPRequest"`

	MarketParticipantDirectory MarketParticipantDirectory `xml:"MarketParticipantDirectory"`

	ProcessDirectory ProcessDirectory `xml:"ProcessDirectory"`
}

// XSD ComplexType declarations

type Extension struct {
	XMLName xml.Name

	GridInvoiceRecipient *GridInvoiceRecipient `xml:"GridInvoiceRecipient"`

	ConsumptionBillingCycle *ConsumptionBillingCycle `xml:"ConsumptionBillingCycle"`

	TransmissionCycle *TransmissionCycle `xml:"TransmissionCycle"`

	MeteringIntervall *MeteringIntervall `xml:"MeteringIntervall"`

	LoadProfileType *LoadProfileType `xml:"LoadProfileType"`

	DateTimeFrom *ct.DateTimeU `xml:"DateTimeFrom"`

	DateTimeTo *ct.DateTimeU `xml:"DateTimeTo"`

	DisconnectionReason *DisconnectionReason `xml:"DisconnectionReason"`

	EmailCustomer *Email `xml:"EmailCustomer"`

	AssumptionOfCosts bool `xml:"AssumptionOfCosts"`
}

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

	ProcessDate string `xml:"ct:ProcessDate"`

	MeteringPoint ct.MeteringPoint `xml:"ct:MeteringPoint"`

	Extension *Extension `xml:"Extension"`

	AdditionalData []ct.AdditionalData `xml:"ct:AdditionalData"`

	VerificationDocument *ct.VerificationDocument `xml:"ct:VerificationDocument"`
}

// XSD SimpleType declarations

type ConsumptionBillingCycle string

const ConsumptionBillingCycle01 ConsumptionBillingCycle = "01"

const ConsumptionBillingCycle02 ConsumptionBillingCycle = "02"

const ConsumptionBillingCycle03 ConsumptionBillingCycle = "03"

const ConsumptionBillingCycle04 ConsumptionBillingCycle = "04"

const ConsumptionBillingCycle06 ConsumptionBillingCycle = "06"

const ConsumptionBillingCycle12 ConsumptionBillingCycle = "12"

type MeteringIntervall string

const MeteringIntervallQh MeteringIntervall = "QH"

const MeteringIntervallH MeteringIntervall = "H"

const MeteringIntervallD MeteringIntervall = "D"

const MeteringIntervallV MeteringIntervall = "V"

type DisconnectionReason string

const DisconnectionReason01 DisconnectionReason = "01"

const DisconnectionReason02 DisconnectionReason = "02"

type Email string

type GridInvoiceRecipient string

const GridInvoiceRecipientCustomer GridInvoiceRecipient = "CUSTOMER"

const GridInvoiceRecipientSupplier GridInvoiceRecipient = "SUPPLIER"

type LoadProfileType string

type TransmissionCycle string

const TransmissionCycleD TransmissionCycle = "D"

const TransmissionCycleM TransmissionCycle = "M"
