package consumptionrecord_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	consumptionrecord "github.com/WattWiseAt/ebutilities/models/customerprocesses/consumptionrecord/01p40"
	"github.com/WattWiseAt/ebutilities/utils"
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

	var conrec consumptionrecord.ConsumptionRecord

	err = xmlutils.Unmarshal(read, &conrec)
	if err != nil {
		t.Fatal(err)
	}

	assertDataCRMSG(t, conrec)

	marshal, err := xml.MarshalIndent(conrec, "", "\t")
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

func assertDataCRMSG(t *testing.T, conrec consumptionrecord.ConsumptionRecord) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", conrec.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/consumptionrecord/01p40", conrec.XMLNs2)
	assert.Equal(t, "PROD", string(conrec.MarketParticipantDirectory.DocumentMode))
	assert.False(t, conrec.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.40", conrec.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(conrec.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(conrec.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(conrec.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "EP1XXXXX", string(conrec.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-04-18T12:53:55.0459790Z", conrec.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(conrec.MarketParticipantDirectory.Sector))
	assert.Equal(t, "DATEN_CRMSG", string(conrec.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(conrec.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(conrec.ProcessDirectory.ConversationID))
	assert.Equal(t, "2024-04-18", conrec.ProcessDirectory.ProcessDate)
	assert.Nil(t, conrec.ProcessDirectory.ECID)
	assert.Nil(t, conrec.ProcessDirectory.DeliveryPoint)
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(conrec.ProcessDirectory.MeteringPoint))
	assert.Equal(t, "00", string(*conrec.ProcessDirectory.Energy[0].MeteringReason))
	assert.Equal(t, "2024-03-01T00:00:00+01:00", string(conrec.ProcessDirectory.Energy[0].MeteringPeriodStart))
	assert.Equal(t, "2024-04-01T00:00:00+02:00", string(conrec.ProcessDirectory.Energy[0].MeteringPeriodEnd))
	assert.Equal(t, 2972, int(conrec.ProcessDirectory.Energy[0].NumberOfMeteringIntervall))
	assert.Equal(t, "QH", string(conrec.ProcessDirectory.Energy[0].MeteringIntervall))
	assert.Equal(t, "2024-03-01T00:00:00+01:00", string(conrec.ProcessDirectory.Energy[0].EnergyData[0].Ep[0].DTF))
	assert.Equal(t, "2024-03-01T00:15:00+01:00", string(conrec.ProcessDirectory.Energy[0].EnergyData[0].Ep[0].DTT))
	assert.Equal(t, "L1", string(conrec.ProcessDirectory.Energy[0].EnergyData[0].Ep[0].MM))
	assert.InEpsilon(t, 0.003, float64(conrec.ProcessDirectory.Energy[0].EnergyData[0].Ep[0].BQ), 0.0001)
}
