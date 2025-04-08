package consumptionrecord_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	consumptionrecord "github.com/colibrie-eu/ebutilities/models/customerprocesses/consumptionrecord/01p40"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath  = "../../../../testdata/examples/EP/"
	GeneratedPath = "../../../../testdata/generated/EP/"
)

func TestDataCRMSG(t *testing.T) {
	t.Parallel()

	filename := "4_DATEN_CRMSG.xml"

	read, err := utils.ReadFile(ExamplesPath + filename)
	if err != nil {
		t.Fatal(err)
	}

	var result consumptionrecord.ConsumptionRecord

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertDataCRMSG(t, result)

	marshal, err := xml.MarshalIndent(result, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	read, err = utils.ReadFile(GeneratedPath + filename)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(read, marshal) {
		t.Fatalf("expected: %s, got: %s", read, marshal)
	}
}

func assertDataCRMSG(t *testing.T, result consumptionrecord.ConsumptionRecord) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/consumptionrecord/01p40", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNsCt)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.False(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.40", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "EP1XXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-04-18T12:53:55.0459790Z", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "DATEN_CRMSG", string(result.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "2024-04-18", result.ProcessDirectory.ProcessDate)
	assert.Nil(t, result.ProcessDirectory.ECID)
	assert.Nil(t, result.ProcessDirectory.DeliveryPoint)
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MeteringPoint))
	assert.Equal(t, "00", string(*result.ProcessDirectory.Energy[0].MeteringReason))
	assert.Equal(t, "2024-03-01T00:00:00+01:00", string(result.ProcessDirectory.Energy[0].MeteringPeriodStart))
	assert.Equal(t, "2024-04-01T00:00:00+02:00", string(result.ProcessDirectory.Energy[0].MeteringPeriodEnd))
	assert.Equal(t, 2972, int(result.ProcessDirectory.Energy[0].NumberOfMeteringIntervall))
	assert.Equal(t, "QH", string(result.ProcessDirectory.Energy[0].MeteringIntervall))
	assert.Equal(t, "2024-03-01T00:00:00+01:00", string(result.ProcessDirectory.Energy[0].EnergyData[0].EP[0].DTF))
	assert.Equal(t, "2024-03-01T00:15:00+01:00", string(result.ProcessDirectory.Energy[0].EnergyData[0].EP[0].DTT))
	assert.Equal(t, "L1", string(result.ProcessDirectory.Energy[0].EnergyData[0].EP[0].MM))
	assert.InEpsilon(t, 0.003, float64(result.ProcessDirectory.Energy[0].EnergyData[0].EP[0].BQ), 0.0001)
}
