package ownSdk

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

func TestAddRemoveValidatorAction(t *testing.T) {
	senderWallet := GenerateWallet()

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "RemoveValidator",
            "actionData": {}
        }
    ]
}`, senderWallet.Address)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddRemoveValidatorAction()
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Actions: Asset Management
////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAddTransferAssetAction(t *testing.T) {
	senderWallet := GenerateWallet()
	fromAccountHash := "FAccH1"
	toAccountHash := "TAccH1"
	assetHash := "AssetH1"
	var amount float64 = 100

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "TransferAsset",
            "actionData": {
                "fromAccountHash": "%s",
                "toAccountHash": "%s",
                "assetHash": "%s",
                "amount": %3.0f
            }
        }
    ]
}`, senderWallet.Address, fromAccountHash, toAccountHash, assetHash, amount)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddTransferAssetAction(fromAccountHash, toAccountHash, assetHash, amount)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddCreateAssetEmissionAction(t *testing.T) {
	senderWallet := GenerateWallet()
	emissionAccountHash := "EAccH1"
	assetHash := "AssetH1"
	var amount float64 = 10000

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "CreateAssetEmission",
            "actionData": {
                "emissionAccountHash": "%s",
                "assetHash": "%s",
                "amount": %3.0f
            }
        }
    ]
}`, senderWallet.Address, emissionAccountHash, assetHash, amount)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddCreateAssetEmissionAction(emissionAccountHash, assetHash, amount)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddCreateAssetAction(t *testing.T) {
	senderWallet := GenerateWallet()

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "CreateAsset",
            "actionData": {}
        }
    ]
}`, senderWallet.Address)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddCreateAssetAction()
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddCreateAssetActionReturnsAssetHash(t *testing.T) {
	senderWallet := GenerateWallet()
	var nonce int64 = 1
	expectedHash := DeriveHash(senderWallet.Address, nonce, 1)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	actualHash := tx.AddCreateAssetAction()
	assert.Equal(t, expectedHash, actualHash)
}

func TestAddSetAssetCodeAction(t *testing.T) {
	senderWallet := GenerateWallet()
	assetHash := "AssetH1"
	const assetCode = "AST1"

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "SetAssetCode",
            "actionData": {
                "assetHash": "%s",
                "assetCode": "%s"
            }
        }
    ]
}`, senderWallet.Address, assetHash, assetCode)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddSetAssetCodeAction(assetHash, assetCode)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddSetAssetControllerAction(t *testing.T) {
	senderWallet := GenerateWallet()
	assetHash := "AssetH1"
	controllerWallet := GenerateWallet()

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "SetAssetController",
            "actionData": {
                "assetHash": "%s",
                "controllerAddress": "%s"
            }
        }
    ]
}`, senderWallet.Address, assetHash, controllerWallet.Address)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddSetAssetControllerAction(assetHash, controllerWallet.Address)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Actions: Account Management
////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAddCreateAccountActionReturnsAccountHash(t *testing.T) {
	senderWallet := GenerateWallet()
	var nonce int64 = 1
	expectedHash := DeriveHash(senderWallet.Address, nonce, 1)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	actualHash := tx.AddCreateAccountAction()
	assert.Equal(t, expectedHash, actualHash)
}

func TestAddSetAccountControllerAction(t *testing.T) {
	senderWallet := GenerateWallet()
	accountHash := "AccountH1"
	controllerWallet := GenerateWallet()

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "SetAccountController",
            "actionData": {
                "accountHash": "%s",
                "controllerAddress": "%s"
            }
        }
    ]
}`, senderWallet.Address, accountHash, controllerWallet.Address)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddSetAccountControllerAction(accountHash, controllerWallet.Address)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Actions: Voting
////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAddSubmitVoteAction(t *testing.T) {
	senderWallet := GenerateWallet()
	accountHash := "AccountH1"
	assetHash := "AssetH1"
	resolutionHash := "ResolutionH1"
	voteHash := "VoteH1"

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "SubmitVote",
            "actionData": {
                "accountHash": "%s",
                "assetHash": "%s",
                "resolutionHash": "%s",
                "voteHash": "%s"
            }
        }
    ]
}`, senderWallet.Address, accountHash, assetHash, resolutionHash, voteHash)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddSubmitVoteAction(accountHash, assetHash, resolutionHash, voteHash)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddSubmitVoteWeightAction(t *testing.T) {
	senderWallet := GenerateWallet()
	accountHash := "AccountH1"
	assetHash := "AssetH1"
	resolutionHash := "ResolutionH1"
	var voteWeight float64 = 12345

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "SubmitVoteWeight",
            "actionData": {
                "accountHash": "%s",
                "assetHash": "%s",
                "resolutionHash": "%s",
                "voteWeight": %5.0f
            }
        }
    ]
}`, senderWallet.Address, accountHash, assetHash, resolutionHash, voteWeight)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddSubmitVoteWeightAction(accountHash, assetHash, resolutionHash, voteWeight)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Actions: Eligibility
////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAddSetAccountEligibilityAction(t *testing.T) {
	senderWallet := GenerateWallet()
	accountHash := "FAccH1"
	assetHash := "AssetH1"
	isPrimaryEligible := false
	isSecondaryEligible := true

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "SetAccountEligibility",
            "actionData": {
                "accountHash": "%s",
                "assetHash": "%s",
                "isPrimaryEligible": %t,
                "isSecondaryEligible": %t
            }
        }
    ]
}`, senderWallet.Address, accountHash, assetHash, isPrimaryEligible, isSecondaryEligible)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddSetAccountEligibilityAction(accountHash, assetHash, isPrimaryEligible, isSecondaryEligible)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddSetAssetEligibilityAction(t *testing.T) {
	senderWallet := GenerateWallet()
	assetHash := "AssetH1"
	isEligibilityRequired := true

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "SetAssetEligibility",
            "actionData": {
                "assetHash": "%s",
                "isEligibilityRequired": %t
            }
        }
    ]
}`, senderWallet.Address, assetHash, isEligibilityRequired)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddSetAssetEligibilityAction(assetHash, isEligibilityRequired)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddChangeKycControllerAddressAction(t *testing.T) {
	senderWallet := GenerateWallet()
	accountHash := "FAccH1"
	assetHash := "AssetH1"
	kycControllerAddress := "KycCtrlAddr1"

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "ChangeKycControllerAddress",
            "actionData": {
                "accountHash": "%s",
                "assetHash": "%s",
                "kycControllerAddress": "%s"
            }
        }
    ]
}`, senderWallet.Address, accountHash, assetHash, kycControllerAddress)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddChangeKycControllerAddressAction(accountHash, assetHash, kycControllerAddress)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddAddKycProviderAction(t *testing.T) {
	senderWallet := GenerateWallet()
	assetHash := "AssetH1"
	providerAddress := GenerateWallet().Address

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "AddKycProvider",
            "actionData": {
                "assetHash": "%s",
                "providerAddress": "%s"
            }
        }
    ]
}`, senderWallet.Address, assetHash, providerAddress)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddAddKycProviderAction(assetHash, providerAddress)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

func TestAddRemoveKycProviderAction(t *testing.T) {
	senderWallet := GenerateWallet()
	assetHash := "AssetH1"
	providerAddress := GenerateWallet().Address

	expectedJson :=
		fmt.Sprintf(
			`{
    "senderAddress": "%s",
    "nonce": 1,
    "expirationTime": 0,
    "actionFee": 0.01,
    "actions": [
        {
            "actionType": "RemoveKycProvider",
            "actionData": {
                "assetHash": "%s",
                "providerAddress": "%s"
            }
        }
    ]
}`, senderWallet.Address, assetHash, providerAddress)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddRemoveKycProviderAction(assetHash, providerAddress)
	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Actions: Multiple
////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAddTransferChxActionMultiple(t *testing.T) {
	senderWallet := GenerateWallet()
	recipientWallet1 := GenerateWallet()
	recipientWallet2 := GenerateWallet()
	var amount1 float64 = 200
	var amount2 float64 = 300

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
        },
        {
            "actionType": "TransferChx",
            "actionData": {
                "recipientAddress": "%s",
                "amount": %3.0f
            }
        }
    ]
}`, senderWallet.Address, recipientWallet1.Address, amount1, recipientWallet2.Address, amount2)

	tx := CreateTx(senderWallet.Address, 1, 0.01, 0)
	tx.AddTransferChxAction(recipientWallet1.Address, amount1)
	tx.AddTransferChxAction(recipientWallet2.Address, amount2)

	actualJson := tx.ToJson(true)
	assert.Equal(t, expectedJson, actualJson)
}
