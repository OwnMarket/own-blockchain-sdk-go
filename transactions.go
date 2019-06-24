package ownsdk

import (
	"encoding/json"
	"fmt"

	. "github.com/shopspring/decimal"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// Types
////////////////////////////////////////////////////////////////////////////////////////////////////

type TxAction struct {
	ActionType string
	ActionData json.RawMessage
}

type Tx struct {
	SenderAddress  string
	Nonce          int64
	ActionFee      Decimal
	ExpirationTime int64
	Actions        []TxAction
}

type SignedTx struct {
	Tx        string
	Signature string
}

func CreateTx(senderAddress string, nonce int64, actionFee Decimal, expirationTime int64) *Tx {
	tx := &Tx{
		SenderAddress:  senderAddress,
		Nonce:          nonce,
		ActionFee:      actionFee,
		ExpirationTime: expirationTime,
		Actions:        make([]TxAction, 0),
	}

	return tx
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Actions
////////////////////////////////////////////////////////////////////////////////////////////////////

func (tx Tx) addAction(actionType string, actionData json.RawMessage) {
	txAction := TxAction{ActionType: actionType, ActionData: actionData}
	tx.Actions = append(tx.Actions, txAction)
}

func (tx Tx) AddTransferChxAction(recipientAddress string, amount Decimal) {
	jsonString := fmt.Sprintf("{recipientAddress: %s, amount: %s}", recipientAddress, amount.String())
	tx.addAction("TransferChx", json.RawMessage(jsonString))
}

func (tx Tx) AddDelegateStakeAction(validatorAddress string, amount Decimal) {
	jsonString := fmt.Sprintf("{validatorAddress: %s, amount: %s}", validatorAddress, amount.String())
	tx.addAction("DelegateStake", json.RawMessage(jsonString))
}

func (tx Tx) AddConfigureValidatorAction(networkAddress string, sharedRewardPercent Decimal, isEnabled bool) {
	jsonString := fmt.Sprintf("{networkAddress: %s, sharedRewardPercent: %s, isEnabled: %t}", networkAddress, sharedRewardPercent.String(), isEnabled)
	tx.addAction("ConfigureValidator", json.RawMessage(jsonString))
}

func (tx Tx) AddRemoveValidatorAction() {
	tx.addAction("DelegateStake", json.RawMessage("{}"))
}

func (tx Tx) AddTransferAssetAction(fromAccountHash string, toAccountHash string, assetHash string, amount Decimal) {
	jsonString := fmt.Sprintf("{fromAccountHash: %s, toAccountHash: %s, assetHash: %s, amount: %s}", fromAccountHash, toAccountHash, assetHash, amount.String())
	tx.addAction("TransferAsset", json.RawMessage(jsonString))
}

func (tx Tx) AddCreateAssetEmissionAction(emissionAccountHash string, assetHash string, amount Decimal) {
	jsonString := fmt.Sprintf("{emissionAccountHash: %s, assetHash: %s, amount: %s}", emissionAccountHash, assetHash, amount.String())
	tx.addAction("CreateAssetEmission", json.RawMessage(jsonString))
}

func (tx Tx) AddCreateAssetAction() {
	tx.addAction("CreateAsset", json.RawMessage("{}"))
	// TODO: return derive hash
}

func (tx Tx) AddSetAssetCodeAction(assetHash string, assetCode string) {
	jsonString := fmt.Sprintf("{assetHash: %s, assetCode: %s}", assetHash, assetCode)
	tx.addAction("SetAssetCode", json.RawMessage(jsonString))
}

func (tx Tx) AddSetAssetControllerAction(assetHash string, controllerAddress string) {
	jsonString := fmt.Sprintf("{assetHash: %s, controllerAddress: %s}", assetHash, controllerAddress)
	tx.addAction("SetAssetController", json.RawMessage(jsonString))
}

func (tx Tx) AddCreateAccountAction() {
	tx.addAction("CreateAccount", json.RawMessage("{}"))
	// TODO: return derive hash
}

func (tx Tx) AddSetAccountControllerAction(accountHash string, controllerAddress string) {
	jsonString := fmt.Sprintf("{accountHash: %s, controllerAddress: %s}", accountHash, controllerAddress)
	tx.addAction("SetAccountController", json.RawMessage(jsonString))
}

func (tx Tx) AddSubmitVoteAction(accountHash string, assetHash string, resolutionHash string, voteHash string) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, resolutionHash: %s, voteHash: %s}", accountHash, assetHash, resolutionHash, voteHash)
	tx.addAction("SubmitVote", json.RawMessage(jsonString))
}

func (tx Tx) AddSubmitVoteWeightAction(accountHash string, assetHash string, resolutionHash string, voteWeight Decimal) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, resolutionHash: %s, voteWeight: %s}", accountHash, assetHash, resolutionHash, voteWeight.String())
	tx.addAction("SubmitVoteWeight", json.RawMessage(jsonString))
}

func (tx Tx) AddSetAccountEligibilityAction(accountHash string, assetHash string, isPrimaryEligible bool, isSecondaryEligible bool) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, isPrimaryEligible: %t, isSecondaryEligible: %t}", accountHash, assetHash, isPrimaryEligible, isSecondaryEligible)
	tx.addAction("SetAccountEligibility", json.RawMessage(jsonString))
}

func (tx Tx) AddSetAssetEligibilityAction(assetHash string, isEligibilityRequired bool) {
	jsonString := fmt.Sprintf("{assetHash: %s, isEligibilityRequired: %t}", assetHash, isEligibilityRequired)
	tx.addAction("SetAssetEligibility", json.RawMessage(jsonString))
}

func (tx Tx) AddChangeKycControllerAddressAction(accountHash string, assetHash string, kycControllerAddress string) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, kycControllerAddress: %s}", accountHash, assetHash, kycControllerAddress)
	tx.addAction("ChangeKycControllerAddress", json.RawMessage(jsonString))
}

func (tx Tx) AddAddKycProviderAction(assetHash string, providerAddress string) {
	jsonString := fmt.Sprintf("{assetHash: %s, providerAddress: %s}", assetHash, providerAddress)
	tx.addAction("AddKycProvider", json.RawMessage(jsonString))
}

func (tx Tx) AddRemoveKycProviderAction(assetHash string, providerAddress string) {
	jsonString := fmt.Sprintf("{assetHash: %s, providerAddress: %s}", assetHash, providerAddress)
	tx.addAction("RemoveKycProvider", json.RawMessage(jsonString))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Signing
////////////////////////////////////////////////////////////////////////////////////////////////////

func (tx Tx) ToJson(indentation bool) string {
	var b []byte
	var err error
	if indentation {
		b, err = json.MarshalIndent(tx, "", "    ")
	} else {
		b, err = json.Marshal(tx)
	}

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(b)
}

func (tx Tx) sign(networkCode []byte, privateKey string) *SignedTx {
	json := tx.ToJson(false)
	signature := SignMessage(networkCode, privateKey, []byte(json))
	signedTx := &SignedTx{
		Tx:        Encode64([]byte(json)),
		Signature: signature,
	}

	return signedTx
}
