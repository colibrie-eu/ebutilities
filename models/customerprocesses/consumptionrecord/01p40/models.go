// Models for http://www.ebutilities.at/schemata/customerprocesses/consumptionrecord/01p40
package consumptionrecord

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type ConsumptionRecord struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `xml:"ConsumptionRecord"`

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

	ProcessDate string `xml:"ct:ProcessDate"`

	ECID *ct.MeteringPoint `xml:"ct:ECID"`

	DeliveryPoint *ct.MessageAddress `xml:"ct:DeliveryPoint"`

	MeteringPoint ct.MeteringPoint `xml:"ct:MeteringPoint"`

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

	EP []EnergyPosition `xml:",any"`
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

type ConsumptionRecordMessageCode string

const ConsumptionRecordMessageCodeDatenMsg ConsumptionRecordMessageCode = "DATEN_MSG"

type ConsumptionRecordVersion string

const ConsumptionRecordVersion0140 ConsumptionRecordVersion = "01.40"

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

const UOMTypeProz UOMType = "PROZ"

const UOMTypeCels UOMType = "CELS"

const UOMTypePce UOMType = "PCE"

const UOMTypeEur UOMType = "EUR"

const UOMTypeMb UOMType = "MB"

const UOMTypeGb UOMType = "GB"

const UOMTypeTb UOMType = "TB"

const UOMTypeH UOMType = "H"

const UOMTypeTag UOMType = "TAG"

const UOMTypeMin UOMType = "MIN"

const UOMTypeMon UOMType = "MON"

const UOMTypeKvarh UOMType = "KVARH"

const UOMTypeMvarh UOMType = "MVARH"

const UOMTypeKwt UOMType = "KWT"

const UOMTypeMwt UOMType = "MWT"

const UOMTypeGwt UOMType = "GWT"

const UOMTypeKwh UOMType = "KWH"

const UOMTypeMwh UOMType = "MWH"

const UOMTypeGwh UOMType = "GWH"

const UOMTypeLe UOMType = "LE"

const UOMTypeM2 UOMType = "M2"

const UOMTypeM3 UOMType = "M3"

const UOMTypeBm3 UOMType = "BM3"

const UOMTypeNm3 UOMType = "NM3"

const UOMTypeBm3H UOMType = "BM3H"

const UOMTypeNm3H UOMType = "NM3H"

const UOMTypeKwhh UOMType = "KWHH"

const UOMTypePau UOMType = "PAU"

type DateTimeS string

type DateTimeU string

type DocumentMode string

const DocumentModeProd DocumentMode = "PROD"

const DocumentModeSimu DocumentMode = "SIMU"

type GroupingID string

type MessageAddress string

type MessageCode string

type MeteringPoint string
