package cprequest_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	cprequest "github.com/colibrie-eu/ebutilities/models/customerprocesses/cprequest/01p12"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath  = "../../../../testdata/examples/EP/"
	GeneratedPath = "../../../../testdata/generated/EP/"
)

func TestRequestPT(t *testing.T) {
	t.Parallel()

	filename := "5_ANFORDERUNG_PT.xml"

	read, err := utils.ReadFile(ExamplesPath + filename)
	if err != nil {
		t.Fatal(err)
	}

	var result cprequest.CPRequest

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertRequestPT(t, result)

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

func assertRequestPT(t *testing.T, result cprequest.CPRequest) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/cprequest/01p12", result.XMLNs2)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.False(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.12", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "EP1XXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "ATXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-06-20T16:55:19", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ANFORDERUNG_PT", string(result.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "2024-06-20", result.ProcessDirectory.ProcessDate)
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MeteringPoint))
	assert.Nil(t, result.ProcessDirectory.Extension.GridInvoiceRecipient)
	assert.Nil(t, result.ProcessDirectory.Extension.ConsumptionBillingCycle)
	assert.Nil(t, result.ProcessDirectory.Extension.TransmissionCycle)
	assert.Nil(t, result.ProcessDirectory.Extension.MeteringIntervall)
	assert.Nil(t, result.ProcessDirectory.Extension.LoadProfileType)
	assert.Equal(t, "2024-05-01T00:00:00+02:00", string(result.ProcessDirectory.Extension.DateTimeFrom))
	assert.Equal(t, "2024-05-31T00:00:00+02:00", string(result.ProcessDirectory.Extension.DateTimeTo))
	assert.Nil(t, result.ProcessDirectory.Extension.DisconnectionReason)
	assert.Nil(t, result.ProcessDirectory.Extension.EmailCustomer)
	assert.False(t, result.ProcessDirectory.Extension.AssumptionOfCosts)
}
