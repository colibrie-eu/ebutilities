// Models for http://www.ebutilities.at/schemata/customerprocesses/ecmplist/01p10
package ecmplist

import (
	"encoding/xml"

	ct "github.com/colibrie-eu/ebutilities/models/customerprocesses/common/types/01p20"
)

type ECMPList struct {
	XMLNs string `xml:"xmlns,attr"`

	XMLNsCt string `xml:"xmlns:ct,attr"`

	XMLName xml.Name `json:"-" xml:"ECMPList"`

	MarketParticipantDirectory MarketParticipantDirectory `db:"market_participant_directory" json:"marketParticipantDirectory,omitempty" xml:"MarketParticipantDirectory"`

	ProcessDirectory ProcessDirectory `db:"process_directory" json:"processDirectory,omitempty" xml:"ProcessDirectory"`
}

// XSD ComplexType declarations

type MarketParticipantDirectory struct {
	XMLName xml.Name `json:"-"`

	DocumentMode ct.DocumentMode `db:"document_mode" json:"documentMode,omitempty" xml:"DocumentMode,attr"`

	Duplicate bool `db:"duplicate" json:"duplicate,omitempty" xml:"Duplicate,attr"`

	SchemaVersion string `db:"schema_version" json:"schemaVersion,omitempty" xml:"SchemaVersion,attr"`

	RoutingHeader ct.RoutingHeader `db:"routing_header" json:"routingHeader,omitempty" xml:"ct:RoutingHeader"`

	Sector ct.Sector `db:"sector" json:"sector,omitempty" xml:"ct:Sector"`

	MessageCode ct.MessageCode `db:"message_code" json:"messageCode,omitempty" xml:"MessageCode"`
}

type ProcessDirectory struct {
	XMLName xml.Name `json:"-"`

	MessageID ct.GroupingID `db:"message_id" json:"messageId,omitempty" xml:"MessageId"`

	ConversationID ct.GroupingID `db:"conversation_id" json:"conversationId,omitempty" xml:"ConversationId"`

	ProcessDate string `db:"process_date" json:"processDate,omitempty" xml:"ProcessDate"`

	ECID ct.MeteringPoint `db:"ec_id" json:"ecId,omitempty" xml:"ECID"`

	ECType ECType `db:"ec_type" json:"ecType,omitempty" xml:"ECType"`

	ECDisModel ECDisModel `db:"ec_dis_model" json:"ecDisModel,omitempty" xml:"ECDisModel"`

	MPListData []MPListData `db:"mp_list_data" json:"mpListData,omitempty" xml:"MPListData"`
}

type MPListData struct {
	XMLName xml.Name `json:"-"`

	MeteringPoint ct.MeteringPoint `db:"metering_point_id" json:"meteringPoint,omitempty" xml:"MeteringPoint"`

	ConsentID *ct.GroupingID `db:"consent_id" json:"consentId,omitempty" xml:"ConsentId"`

	MPTimeData []MPTimeData `db:"mp_time_data" json:"mpTimeData" xml:"MPTimeData"`
}

type MPTimeData struct {
	XMLName xml.Name `json:"-"`

	ID int `db:"id" json:"-" xml:"-"`

	DateFrom string `db:"date_from" json:"dateFrom,omitempty" xml:"DateFrom"`

	DateTo string `db:"date_to" json:"dateTo,omitempty" xml:"DateTo"`

	EnergyDirection EnergyDirection `db:"energy_direction" json:"energyDirection,omitempty" xml:"EnergyDirection"`

	ECPartFact ECPartFact `db:"ec_part_fact" json:"ecPartFact,omitempty" xml:"ECPartFact"`

	PlantCategory *PlantCategory `db:"plant_category" json:"plantCategory,omitempty" xml:"PlantCategory"`

	DateActivate string `db:"date_activate" json:"dateActivate,omitempty" xml:"DateActivate"`

	DateDeactivate *string `db:"date_deactivate" json:"dateDeactivate,omitempty" xml:"DateDeactivate"`

	ECShare *ECShare `db:"ec_share" json:"ecShare,omitempty" xml:"ECShare"`

	ECShC []ShareCalc `db:"ec_sh_c" json:"ecShC,omitempty" xml:"ECShC"`
}

type ShareCalc struct {
	XMLName xml.Name `json:"-"`

	DateFrom string `db:"date_from" json:"dateFrom,omitempty" xml:"DateFrom"`

	DateTo string `db:"date_to" json:"dateTo,omitempty" xml:"DateTo"`

	ECShareCalc ECShare `db:"ec_share_calc" json:"ecShareCalc,omitempty" xml:"ECShareCalc"`
}

// XSD SimpleType declarations

type ECDisModel string

const ECDisModelD ECDisModel = "D"

const ECDisModelS ECDisModel = "S"

type ECShare float64

type ECType string

const ECTypeGc ECType = "GC"

const ECTypeRcR ECType = "RC_R"

const ECTypeRcL ECType = "RC_L"

const ECTypeCc ECType = "CC"

type EnergyDirection string

const EnergyDirectionConsumption EnergyDirection = "CONSUMPTION"

const EnergyDirectionGeneration EnergyDirection = "GENERATION"

type ECPartFact float64

type PlantCategory string
