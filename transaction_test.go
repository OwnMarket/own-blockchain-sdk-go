package ownsdk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTxCorrectJson(t *testing.T) {
	senderWallet := GenerateWallet()
	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 123,
    "actionFee": 0.01,
    "actions": []
}`, senderWallet.Address)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 123)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}
