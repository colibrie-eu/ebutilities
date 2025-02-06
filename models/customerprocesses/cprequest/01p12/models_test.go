package cprequest_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	cprequest "github.com/WattWiseAt/ebutilities/models/customerprocesses/cprequest/01p12"
	"github.com/WattWiseAt/ebutilities/utils"
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

	var cpreq cprequest.CPRequest

	err = xmlutils.Unmarshal(read, &cpreq)
	if err != nil {
		t.Fatal(err)
	}

	assertRequestPT(t, cpreq)

	marshal, err := xml.MarshalIndent(cpreq, "", "\t")
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

func assertRequestPT(t *testing.T, cpreq cprequest.CPRequest) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", cpreq.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/cprequest/01p12", cpreq.XMLNs2)
	assert.Equal(t, "PROD", string(cpreq.MarketParticipantDirectory.DocumentMode))
	assert.False(t, cpreq.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.12", cpreq.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(cpreq.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "EP1XXXXX", string(cpreq.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(cpreq.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "ATXXXXXX", string(cpreq.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-06-20T16:55:19", cpreq.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(cpreq.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ANFORDERUNG_PT", string(cpreq.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cpreq.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cpreq.ProcessDirectory.ConversationID))
	assert.Equal(t, "2024-06-20", cpreq.ProcessDirectory.ProcessDate)
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cpreq.ProcessDirectory.MeteringPoint))
	assert.Nil(t, cpreq.ProcessDirectory.Extension.GridInvoiceRecipient)
	assert.Nil(t, cpreq.ProcessDirectory.Extension.ConsumptionBillingCycle)
	assert.Nil(t, cpreq.ProcessDirectory.Extension.TransmissionCycle)
	assert.Nil(t, cpreq.ProcessDirectory.Extension.MeteringIntervall)
	assert.Nil(t, cpreq.ProcessDirectory.Extension.LoadProfileType)
	assert.Equal(t, "2024-05-01T00:00:00+02:00", string(cpreq.ProcessDirectory.Extension.DateTimeFrom))
	assert.Equal(t, "2024-05-31T00:00:00+02:00", string(cpreq.ProcessDirectory.Extension.DateTimeTo))
	assert.Nil(t, cpreq.ProcessDirectory.Extension.DisconnectionReason)
	assert.Nil(t, cpreq.ProcessDirectory.Extension.EmailCustomer)
	assert.False(t, cpreq.ProcessDirectory.Extension.AssumptionOfCosts)
}
