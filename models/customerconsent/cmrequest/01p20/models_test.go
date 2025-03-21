package cmrequest_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	cmrequest "github.com/colibrie-eu/ebutilities/models/customerconsent/cmrequest/01p20"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath  = "../../../../testdata/examples/EP/"
	GeneratedPath = "../../../../testdata/generated/EP/"
)

func TestRequestCCMO(t *testing.T) {
	t.Parallel()

	filename := "1_ANFORDERUNG_CCMO.xml"

	read, err := utils.ReadFile(ExamplesPath + filename)
	if err != nil {
		t.Fatal(err)
	}

	var result cmrequest.CMRequest

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertRequestCCMO(t, result)

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

func assertRequestCCMO(t *testing.T, result cmrequest.CMRequest) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerconsent/cmrequest/01p20", result.XMLNs2)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.False(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.20", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "EP1XXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "ATXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-04-15T10:56:47", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ANFORDERUNG_CCMO", string(result.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "2024-04-15", result.ProcessDirectory.ProcessDate)
	assert.Equal(t, "ATXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MeteringPoint))
	assert.Equal(t, "NYMV3L3I", string(result.ProcessDirectory.CMRequestID))
	assert.Nil(t, result.ProcessDirectory.ConsentID)
	assert.Equal(t, "ECProfileData", string(result.ProcessDirectory.CMRequest.ReqDatType))
	assert.Equal(t, "2024-05-01", result.ProcessDirectory.CMRequest.DateFrom)
	assert.Equal(t, "9999-12-31", result.ProcessDirectory.CMRequest.DateTo)
	assert.Equal(t, "QH", string(result.ProcessDirectory.CMRequest.MeteringIntervall))
	assert.Equal(t, "D", string(result.ProcessDirectory.CMRequest.TransmissionCycle))
	assert.Nil(t, result.ProcessDirectory.CMRequest.ECID)
	assert.Nil(t, result.ProcessDirectory.CMRequest.ECPartFact)
	assert.Nil(t, result.ProcessDirectory.CMRequest.ECShare)
	assert.Equal(t, "CONSUMPTION", string(result.ProcessDirectory.CMRequest.EnergyDirection))
}
