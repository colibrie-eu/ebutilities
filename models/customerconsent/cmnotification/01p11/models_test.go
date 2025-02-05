package cmnotification_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	cmnotification "github.com/WattWiseAt/ebutilities/models/customerconsent/cmnotification/01p11"
	"github.com/WattWiseAt/ebutilities/utils"
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

	var cmnotif cmnotification.CMNotification

	err = xmlutils.Unmarshal(read, &cmnotif)
	if err != nil {
		t.Fatal(err)
	}

	assertResponseCCMO(t, cmnotif)

	marshal, err := xml.MarshalIndent(cmnotif, "", "\t")
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

func assertResponseCCMO(t *testing.T, cmnotif cmnotification.CMNotification) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", cmnotif.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerconsent/cmnotification/01p11", cmnotif.XMLNs2)
	assert.Equal(t, "PROD", string(cmnotif.MarketParticipantDirectory.DocumentMode))
	assert.False(t, cmnotif.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.11", cmnotif.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(cmnotif.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(cmnotif.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(cmnotif.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "EPXXXXXX", string(cmnotif.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-04-18T12:20:33.7215520Z", cmnotif.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(cmnotif.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ANTWORT_CCMO", string(cmnotif.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmnotif.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmnotif.ProcessDirectory.ConversationID))
	assert.Equal(t, "NYMV3L3I", string(cmnotif.ProcessDirectory.CMRequestID))
	assert.Nil(t, cmnotif.ProcessDirectory.ResponseData[0].ConsentID)
	assert.Equal(t, "ATXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(*cmnotif.ProcessDirectory.ResponseData[0].MeteringPoint))
	assert.Equal(t, 99, int(cmnotif.ProcessDirectory.ResponseData[0].ResponseCode[0]))
	assert.Nil(t, cmnotif.ProcessDirectory.ResponseData[0].ParamHist)
}
