package ownsdk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// TX
////////////////////////////////////////////////////////////////////////////////////////////////////

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

func TestCreateTxExpirationTimeZeroIfNotProvided(t *testing.T) {
	senderWallet := GenerateWallet()
	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": []
}`, senderWallet.Address)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Actions: Network Management
////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAddTransferChxAction(t *testing.T) {
	senderWallet := GenerateWallet()
	recipientWallet := GenerateWallet()
	var amount float64 = 1000

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "TransferChx",
            "actionData": {
                "recipientAddress": "%s",
                "amount": %3.0f
            }
        }
    ]
}`, senderWallet.Address, recipientWallet.Address, amount)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddTransferChxAction(recipientWallet.Address, amount)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddDelegateStakeAction(t *testing.T) {
	senderWallet := GenerateWallet()
	validatorWallet := GenerateWallet()
	var amount float64 = 100000

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "DelegateStake",
            "actionData": {
                "validatorAddress": "%s",
                "amount": %3.0f
            }
        }
    ]
}`, senderWallet.Address, validatorWallet.Address, amount)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddDelegateStakeAction(validatorWallet.Address, amount)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddConfigureValidatorAction(t *testing.T) {
	senderWallet := GenerateWallet()
	networkAddress := "val01.some.domain.com:25718"
	var sharedRewardPercent float64 = 100000
	isEnabled := true

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "ConfigureValidator",
            "actionData": {
                "networkAddress": "%s",
                "sharedRewardPercent": %3.0f,
                "isEnabled": %t
            }
        }
    ]
}`, senderWallet.Address, networkAddress, sharedRewardPercent, isEnabled)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddConfigureValidatorAction(networkAddress, sharedRewardPercent, isEnabled)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}
