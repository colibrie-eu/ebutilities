package ecmplist_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	ecmplist "github.com/colibrie-eu/ebutilities/models/customerprocesses/ecmplist/01p10"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath  = "../../../../testdata/examples/EP/"
	GeneratedPath = "../../../../testdata/generated/EP/"
)

func TestResult(t *testing.T) {
	t.Parallel()

	filename := "SENDEN_ECP.xml"

	read, err := utils.ReadFile(ExamplesPath+filename, 1)
	if err != nil {
		t.Fatal(err)
	}

	var result ecmplist.ECMPList

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertResult(t, result)

	marshal, err := xml.MarshalIndent(result, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	read, err = utils.ReadFile(GeneratedPath+filename, 1)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(read, marshal) {
		t.Fatalf("expected: %s, got: %s", read, marshal)
	}
}

func assertResult(t *testing.T, result ecmplist.ECMPList) {
	t.Helper()

	dateDeactivate := "2022-12-17"

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/ecmplist/01p10", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNsCt)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.True(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.10", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "AT001000", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "RC100123", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2022-12-17T09:30:47Z", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "SENDEN_ECP", string(result.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "123456789", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "0ASDF", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "2022-12-17", result.ProcessDirectory.ProcessDate)
	assert.Equal(t, "AT00100000000RC100123000000123456", string(result.ProcessDirectory.ECID))
	assert.Equal(t, ecmplist.ECTypeRcR, result.ProcessDirectory.ECType)
	assert.Equal(t, ecmplist.ECDisModelS, result.ProcessDirectory.ECDisModel)
	assert.Equal(t, "AT0010000103600000000123456123456", string(result.ProcessDirectory.MPListData[0].MeteringPoint))
	assert.Equal(t, "IWRN74PW", string(*result.ProcessDirectory.MPListData[0].ConsentID))
	assert.Equal(t, "2022-11-01", result.ProcessDirectory.MPListData[0].MPTimeData[0].DateFrom)
	assert.Equal(t, "2022-12-17", result.ProcessDirectory.MPListData[0].MPTimeData[0].DateTo)
	assert.Equal(t, ecmplist.EnergyDirectionGeneration, result.ProcessDirectory.MPListData[0].MPTimeData[0].EnergyDirection)
	assert.InEpsilon(t, 100, float64(result.ProcessDirectory.MPListData[0].MPTimeData[0].ECPartFact), 0.0001)
	assert.Equal(t, ecmplist.PlantCategory("SONNE"), *result.ProcessDirectory.MPListData[0].MPTimeData[0].PlantCategory)
	assert.Equal(t, "2022-11-01", result.ProcessDirectory.MPListData[0].MPTimeData[0].DateActivate)
	assert.Equal(t, &dateDeactivate, result.ProcessDirectory.MPListData[0].MPTimeData[0].DateDeactivate)
	assert.InEpsilon(t, 80, float64(*result.ProcessDirectory.MPListData[1].MPTimeData[0].ECShare), 0.0001)
	assert.Equal(t, "2022-11-15", result.ProcessDirectory.MPListData[1].MPTimeData[0].ECShC[0].DateFrom)
	assert.Equal(t, "2022-11-30", result.ProcessDirectory.MPListData[1].MPTimeData[0].ECShC[0].DateTo)
	assert.InEpsilon(t, 66.6666, float64(result.ProcessDirectory.MPListData[1].MPTimeData[0].ECShC[0].ECShareCalc), 0.0001)
}
