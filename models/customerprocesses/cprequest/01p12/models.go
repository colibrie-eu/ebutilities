// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://www.ebutilities.at/schemata/customerprocesses/cprequest/01p12
package cp

import (
	"encoding/xml"
	"github.com/WattWiseAt/ebutilities/models/customerprocesses/cprequest/ct"
)

// Element
type Cprequest struct {
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

	SchemaVersion string `xml:"SchemaVersion,attr"`

	DocumentMode ct.DocumentMode `xml:"DocumentMode,attr"`

	Duplicate bool `xml:"Duplicate,attr"`

	MessageCode ct.MessageCode `xml:"MessageCode"`

	RoutingHeader ct.RoutingHeader `xml:"RoutingHeader"`

	Sector ct.Sector `xml:"Sector"`
}

type ProcessDirectory struct {
	XMLName xml.Name

	Extension *Extension `xml:"Extension"`

	AdditionalData []ct.AdditionalData `xml:"AdditionalData"`

	VerificationDocument *ct.VerificationDocument `xml:"VerificationDocument"`

	MessageId ct.GroupingId `xml:"MessageId"`

	ConversationId ct.GroupingId `xml:"ConversationId"`

	ProcessDate string `xml:"ProcessDate"`

	MeteringPoint ct.MeteringPoint `xml:"MeteringPoint"`
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
