package cpnotification_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	cpnotification "github.com/colibrie-eu/ebutilities/models/customerprocesses/cpnotification/01p13"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath  = "../../../../testdata/examples/EP/"
	GeneratedPath = "../../../../testdata/generated/EP/"
)

func TestResponsePT(t *testing.T) {
	t.Parallel()

	filename := "ANTWORT_PT.xml"

	read, err := utils.ReadFile(ExamplesPath+filename, 1)
	if err != nil {
		t.Fatal(err)
	}

	var result cpnotification.CPNotification

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertResponsePT(t, result)

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

func assertResponsePT(t *testing.T, result cpnotification.CPNotification) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/cpnotification/01p13", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNsCt)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.False(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.13", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "RCXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2025-07-30T06:42:57.4683500Z", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ANTWORT_PT", string(result.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "ATXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "RCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "RCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ResponseData.OriginalMessageID))
	assert.Equal(t, 70, int(result.ProcessDirectory.ResponseData.ResponseCode[0]))
}
