package cmnotification_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	cmnotification "github.com/colibrie-eu/ebutilities/models/customerconsent/cmnotification/01p11"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath  = "../../../../testdata/examples/EP/"
	GeneratedPath = "../../../../testdata/generated/EP/"
)

func TestResponseCCMO(t *testing.T) {
	t.Parallel()

	filename := "2_ANTWORT_CCMO.xml"

	read, err := utils.ReadFile(ExamplesPath + filename)
	if err != nil {
		t.Fatal(err)
	}

	var result cmnotification.CMNotification

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertResponseCCMO(t, result)

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

func assertResponseCCMO(t *testing.T, result cmnotification.CMNotification) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerconsent/cmnotification/01p11", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNsCt)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.False(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.11", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "EPXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-04-18T12:20:33.7215520Z", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ANTWORT_CCMO", string(result.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "NYMV3L3I", string(result.ProcessDirectory.CMRequestID))
	assert.Nil(t, result.ProcessDirectory.ResponseData[0].ConsentID)
	assert.Equal(t, "ATXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(*result.ProcessDirectory.ResponseData[0].MeteringPoint))
	assert.Equal(t, 99, int(result.ProcessDirectory.ResponseData[0].ResponseCode[0]))
	assert.Nil(t, result.ProcessDirectory.ResponseData[0].ParamHist)
}
