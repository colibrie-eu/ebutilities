package cmnotification_test

import (
	"testing"

	cmnotification "github.com/colibrie-eu/ebutilities/models/customerconsent/cmnotification/01p12"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath = "../../../../testdata/examples/EP/"
)

func TestResponseCCMO(t *testing.T) {
	t.Parallel()

	filename := "ABLEHNUNG_ECON.xml"

	read, err := utils.ReadFile(ExamplesPath+filename, 1)
	if err != nil {
		t.Fatal(err)
	}

	var result cmnotification.CMNotification

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertResponseCCMO(t, result)
}

func assertResponseCCMO(t *testing.T, result cmnotification.CMNotification) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerconsent/cmnotification/01p12", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNsCt)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.False(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.12", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "RCXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2025-07-16T14:45:30.3610520Z", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ABLEHNUNG_ECON", string(result.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "ATXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "RCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "RCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.CMRequestID))
	assert.Nil(t, result.ProcessDirectory.ResponseData[0].ConsentID)
	assert.Equal(t, "ATXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(*result.ProcessDirectory.ResponseData[0].MeteringPoint))
	assert.Equal(t, 86, int(result.ProcessDirectory.ResponseData[0].ResponseCode[0]))
	assert.Equal(t, 188, int(result.ProcessDirectory.ResponseData[0].ResponseCode[1]))
}
