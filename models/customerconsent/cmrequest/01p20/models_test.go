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

	var cmreq cmrequest.CMRequest

	err = xmlutils.Unmarshal(read, &cmreq)
	if err != nil {
		t.Fatal(err)
	}

	assertRequestCCMO(t, cmreq)

	marshal, err := xml.MarshalIndent(cmreq, "", "\t")
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

func assertRequestCCMO(t *testing.T, cmreq cmrequest.CMRequest) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", cmreq.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerconsent/cmrequest/01p20", cmreq.XMLNs2)
	assert.Equal(t, "PROD", string(cmreq.MarketParticipantDirectory.DocumentMode))
	assert.False(t, cmreq.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.20", cmreq.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(cmreq.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "EP1XXXXX", string(cmreq.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(cmreq.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "ATXXXXXX", string(cmreq.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2024-04-15T10:56:47", cmreq.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(cmreq.MarketParticipantDirectory.Sector))
	assert.Equal(t, "ANFORDERUNG_CCMO", string(cmreq.MarketParticipantDirectory.MessageCode))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmreq.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmreq.ProcessDirectory.ConversationID))
	assert.Equal(t, "2024-04-15", cmreq.ProcessDirectory.ProcessDate)
	assert.Equal(t, "ATXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmreq.ProcessDirectory.MeteringPoint))
	assert.Equal(t, "NYMV3L3I", string(cmreq.ProcessDirectory.CMRequestID))
	assert.Nil(t, cmreq.ProcessDirectory.ConsentID)
	assert.Equal(t, "ECProfileData", string(cmreq.ProcessDirectory.CMRequest.ReqDatType))
	assert.Equal(t, "2024-05-01", cmreq.ProcessDirectory.CMRequest.DateFrom)
	assert.Equal(t, "9999-12-31", cmreq.ProcessDirectory.CMRequest.DateTo)
	assert.Equal(t, "QH", string(cmreq.ProcessDirectory.CMRequest.MeteringIntervall))
	assert.Equal(t, "D", string(cmreq.ProcessDirectory.CMRequest.TransmissionCycle))
	assert.Nil(t, cmreq.ProcessDirectory.CMRequest.ECID)
	assert.Nil(t, cmreq.ProcessDirectory.CMRequest.ECPartFact)
	assert.Nil(t, cmreq.ProcessDirectory.CMRequest.ECShare)
	assert.Equal(t, "CONSUMPTION", string(cmreq.ProcessDirectory.CMRequest.EnergyDirection))
}
