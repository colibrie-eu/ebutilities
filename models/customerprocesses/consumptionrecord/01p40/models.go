// Models for http://www.ebutilities.at/schemata/customerprocesses/consumptionrecord/01p40
package consumptionrecord

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type ConsumptionRecord struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNs2 string `xml:"xmlns:ns2,attr"`

	XMLName xml.Name `xml:"ns2:ConsumptionRecord"`

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

	ProcessDate string `xml:"ProcessDate"`

	ECID *ct.MeteringPoint `xml:"ECID"`

	DeliveryPoint *ct.MessageAddress `xml:"DeliveryPoint"`

	MeteringPoint ct.MeteringPoint `xml:"MeteringPoint"`

	Energy []Energy `xml:"Energy"`
}

type Energy struct {
	XMLName xml.Name

	MeteringReason *MeteringReason `xml:"MeteringReason"`

	MeteringPeriodStart DateTimeU `xml:"MeteringPeriodStart"`

	MeteringPeriodEnd DateTimeU `xml:"MeteringPeriodEnd"`

	MeteringIntervall MeteringIntervall `xml:"MeteringIntervall"`

	NumberOfMeteringIntervall int64 `xml:"NumberOfMeteringIntervall"`

	EnergyData []EnergyData `xml:"EnergyData"`
}

type EnergyData struct {
	XMLName xml.Name

	MeterCode string `xml:"MeterCode,attr"`

	UOM UOMType `xml:"UOM,attr"`

	Ep []EnergyPosition `xml:",any"`
}

type EnergyPosition struct {
	XMLName xml.Name

	DTF DateTimeU `xml:"DTF"`

	DTT DateTimeU `xml:"DTT"`

	MM MeteringMethod `xml:"MM"`

	BQ BillingQuantity `xml:"BQ"`
}

// XSD SimpleType declarations

type BillingQuantity float64

type Competence string

type ConRecMessageCode string

const ConRecMessageCodeDatenMsg ConRecMessageCode = "DATEN_MSG"

type ConRecVersion string

const ConRecVersion0140 ConRecVersion = "01.40"

type Email string

type Fax string

type MeteringIntervall string

const MeteringIntervallQh MeteringIntervall = "QH"

const MeteringIntervallH MeteringIntervall = "H"

const MeteringIntervallD MeteringIntervall = "D"

const MeteringIntervallV MeteringIntervall = "V"

type MeteringMethod string

const MeteringMethod01 MeteringMethod = "01"

const MeteringMethod02 MeteringMethod = "02"

const MeteringMethod03 MeteringMethod = "03"

const MeteringMethod04 MeteringMethod = "04"

const MeteringMethod05 MeteringMethod = "05"

const MeteringMethodL1 MeteringMethod = "L1"

const MeteringMethodL2 MeteringMethod = "L2"

const MeteringMethodL3 MeteringMethod = "L3"

type MeteringReason string

const MeteringReason00 MeteringReason = "00"

const MeteringReason01 MeteringReason = "01"

const MeteringReason02 MeteringReason = "02"

const MeteringReason03 MeteringReason = "03"

const MeteringReason04 MeteringReason = "04"

const MeteringReason05 MeteringReason = "05"

type Name string

type Phone string

type ReferenceNumber string

type UOMType string

const UomtypeProz UOMType = "PROZ"

const UomtypeCels UOMType = "CELS"

const UomtypePce UOMType = "PCE"

const UomtypeEur UOMType = "EUR"

const UomtypeMb UOMType = "MB"

const UomtypeGb UOMType = "GB"

const UomtypeTb UOMType = "TB"

const UomtypeH UOMType = "H"

const UomtypeTag UOMType = "TAG"

const UomtypeMin UOMType = "MIN"

const UomtypeMon UOMType = "MON"

const UomtypeKvarh UOMType = "KVARH"

const UomtypeMvarh UOMType = "MVARH"

const UomtypeKwt UOMType = "KWT"

const UomtypeMwt UOMType = "MWT"

const UomtypeGwt UOMType = "GWT"

const UomtypeKwh UOMType = "KWH"

const UomtypeMwh UOMType = "MWH"

const UomtypeGwh UOMType = "GWH"

const UomtypeLe UOMType = "LE"

const UomtypeM2 UOMType = "M2"

const UomtypeM3 UOMType = "M3"

const UomtypeBm3 UOMType = "BM3"

const UomtypeNm3 UOMType = "NM3"

const UomtypeBm3H UOMType = "BM3H"

const UomtypeNm3H UOMType = "NM3H"

const UomtypeKwhh UOMType = "KWHH"

const UomtypePau UOMType = "PAU"

type DateTimeS string

type DateTimeU string

type DocumentMode string

const DocumentModeProd DocumentMode = "PROD"

const DocumentModeSimu DocumentMode = "SIMU"

type GroupingID string

type MessageAddress string

type MessageCode string

type MeteringPoint string
