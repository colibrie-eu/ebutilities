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

	var result cmrevoke.CMRevoke

	err = xmlutils.Unmarshal(read, &result)
	if err != nil {
		t.Fatal(err)
	}

	assertRevokeCCMC(t, result)

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

func assertRevokeCCMC(t *testing.T, result cmrevoke.CMRevoke) {
	t.Helper()

	assert.Equal(t, "http://www.ebutilities.at/schemata/customerprocesses/common/types/01p20", result.XMLNs)
	assert.Equal(t, "http://www.ebutilities.at/schemata/customerconsent/cmrevoke/01p00", result.XMLNs2)
	assert.Equal(t, "PROD", string(result.MarketParticipantDirectory.DocumentMode))
	assert.False(t, result.MarketParticipantDirectory.Duplicate)
	assert.Equal(t, "01.00", result.MarketParticipantDirectory.SchemaVersion)
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Sender.AddressType))
	assert.Equal(t, "ATXXXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Sender.MessageAddress))
	assert.Equal(t, "ECNumber", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.AddressType))
	assert.Equal(t, "EP1XXXXX", string(result.MarketParticipantDirectory.RoutingHeader.Receiver.MessageAddress))
	assert.Equal(t, "2023-08-22T12:53:20", result.MarketParticipantDirectory.RoutingHeader.DocumentCreationDateTime)
	assert.Equal(t, "01", string(result.MarketParticipantDirectory.Sector))
	assert.Equal(t, "AUFHEBUNG_CCMC", result.MarketParticipantDirectory.MessageCode)
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MessageID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConversationID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.ConsentID))
	assert.Equal(t, "EPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", string(result.ProcessDirectory.MeteringPoint))
	assert.Equal(t, "2023-08-22", result.ProcessDirectory.ConsentEnd)
	assert.Equal(t, "Aufhebung durch Kunden", string(result.ProcessDirectory.Reason))
}
