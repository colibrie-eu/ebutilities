package cmrevoke_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	cmrevoke "github.com/colibrie-eu/ebutilities/models/customerconsent/cmrevoke/01p00"
	"github.com/colibrie-eu/ebutilities/utils"
	"github.com/mantyr/xmlutils"
	"github.com/stretchr/testify/assert"
)

const (
	ExamplesPath  = "../../../../testdata/examples/EP/"
	GeneratedPath = "../../../../testdata/generated/EP/"
)

func TestRevokeCCMC(t *testing.T) {
	t.Parallel()

	filename := "6_AUFHEBUNG_CCMC.xml"

	read, err := utils.ReadFile(ExamplesPath + filename)
	if err != nil {
		t.Fatal(err)
	}

	var cmrev cmrevoke.CMRevoke

	err = xmlutils.Unmarshal(read, &cmrev)
	if err != nil {
		t.Fatal(err)
	}

	assertRevokeCCMC(t, cmrev)

	marshal, err := xml.MarshalIndent(cmrev, "", "\t")
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

func assertRevokeCCMC(t *testing.T, cmrev cmrevoke.CMRevoke) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", cmrev.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerconsent/cmrevoke/01p00", cmrev.XMLNs2)
	assert.Equal(t, "PROD", string(cmrev.MarketParticipantDirectory.DocumentMode))
	assert.False(t, cmrev.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.00", cmrev.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(cmrev.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(cmrev.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(cmrev.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "EP1XXXXX", string(cmrev.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2023-08-22T12:53:20", cmrev.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(cmrev.MarketParticipantDirectory.Sector))
	assert.Equal(t, "AUFHEBUNG_CCMC", cmrev.MarketParticipantDirectory.MessageCode)
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmrev.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmrev.ProcessDirectory.ConversationID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmrev.ProcessDirectory.ConsentID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(cmrev.ProcessDirectory.MeteringPoint))
	assert.Equal(t, "2023-08-22", cmrev.ProcessDirectory.ConsentEnd)
	assert.Equal(t, "Aufhebung durch Kunden", string(cmrev.ProcessDirectory.Reason))
}
