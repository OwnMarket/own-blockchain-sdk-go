package ownsdk

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// Types
////////////////////////////////////////////////////////////////////////////////////////////////////

type TxAction struct {
	ActionType string      `json:"actionType"`
	ActionData interface{} `json:"actionData"`
}

type Tx struct {
	SenderAddress  string     `json:"senderAddress"`
	Nonce          int64      `json:"nonce"`
	ExpirationTime int64      `json:"expirationTime"`
	ActionFee      float64    `json:"actionFee"`
	Actions        []TxAction `json:"actions"`
}

type SignedTx struct {
	Tx        string `json:"tx"`
	Signature string `json:"signature"`
}

type TransferChxTxActionDto struct {
	RecipientAddress string  `json:"recipientAddress"`
	Amount           float64 `json:"amount"`
}

type DelegateStakeTxActionDto struct {
	ValidatorAddress string  `json:"validatorAddress"`
	Amount           float64 `json:"amount"`
}

type ConfigureValidatorTxActionDto struct {
	NetworkAddress      string  `json:"networkAddress"`
	SharedRewardPercent float64 `json:"sharedRewardPercent"`
	IsEnabled           bool    `json:"isEnabled"`
}

type RemoveValidatorTxActionDto struct{}

type TransferAssetTxActionDto struct {
	FromAccountHash string  `json:"fromAccountHash"`
	ToAccountHash   string  `json:"toAccountHash"`
	AssetHash       string  `json:"assetHash"`
	Amount          float64 `json:"amount"`
}

type CreateAssetEmissionTxActionDto struct {
	EmissionAccountHash string  `json:"emissionAccountHash"`
	AssetHash           string  `json:"assetHash"`
	Amount              float64 `json:"amount"`
}

type CreateAssetTxActionDto struct{}

type SetAssetCodeTxActionDto struct {
	AssetHash string `json:"assetHash"`
	AssetCode string `json:"assetCode"`
}

type SetAssetControllerTxActionDto struct {
	AssetHash         string `json:"assetHash"`
	ControllerAddress string `json:"controllerAddress"`
}

type CreateAccountTxActionDto struct{}

type SetAccountControllerTxActionDto struct {
	AccountHash       string `json:"accountHash"`
	ControllerAddress string `json:"controllerAddress"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Constructor
////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateTx(senderAddress string, nonce int64, actionFee float64, expirationTime int64) *Tx {
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

func (tx *Tx) addAction(actionType string, actionData interface{}) {
	txAction := TxAction{ActionType: actionType, ActionData: actionData}
	tx.Actions = append(tx.Actions, txAction)
}

func (tx *Tx) AddTransferChxAction(recipientAddress string, amount float64) {
	dto := TransferChxTxActionDto{
		RecipientAddress: recipientAddress,
		Amount:           amount,
	}
	tx.addAction("TransferChx", dto)
}

func (tx *Tx) AddDelegateStakeAction(validatorAddress string, amount float64) {
	dto := DelegateStakeTxActionDto{
		ValidatorAddress: validatorAddress,
		Amount:           amount,
	}
	tx.addAction("DelegateStake", dto)
}

func (tx *Tx) AddConfigureValidatorAction(networkAddress string, sharedRewardPercent float64, isEnabled bool) {
	dto := ConfigureValidatorTxActionDto{
		NetworkAddress:      networkAddress,
		SharedRewardPercent: sharedRewardPercent,
		IsEnabled:           isEnabled,
	}
	tx.addAction("ConfigureValidator", dto)
}

func (tx *Tx) AddRemoveValidatorAction() {
	dto := RemoveValidatorTxActionDto{}
	tx.addAction("RemoveValidator", dto)
}

func (tx *Tx) AddTransferAssetAction(fromAccountHash string, toAccountHash string, assetHash string, amount float64) {
	dto := TransferAssetTxActionDto{
		FromAccountHash: fromAccountHash,
		ToAccountHash:   toAccountHash,
		AssetHash:       assetHash,
		Amount:          amount,
	}
	tx.addAction("TransferAsset", dto)
}

func (tx *Tx) AddCreateAssetEmissionAction(emissionAccountHash string, assetHash string, amount float64) {
	dto := CreateAssetEmissionTxActionDto{
		EmissionAccountHash: emissionAccountHash,
		AssetHash:           assetHash,
		Amount:              amount,
	}
	tx.addAction("CreateAssetEmission", dto)
}

func (tx *Tx) AddCreateAssetAction() string {
	dto := CreateAssetTxActionDto{}
	tx.addAction("CreateAsset", dto)
	return DeriveHash(tx.SenderAddress, tx.Nonce, int16(len(tx.Actions)))
}

func (tx *Tx) AddSetAssetCodeAction(assetHash string, assetCode string) {
	dto := SetAssetCodeTxActionDto{
		AssetHash: assetHash,
		AssetCode: assetCode,
	}
	tx.addAction("SetAssetCode", dto)
}

func (tx *Tx) AddSetAssetControllerAction(assetHash string, controllerAddress string) {
	dto := SetAssetControllerTxActionDto{
		AssetHash:         assetHash,
		ControllerAddress: controllerAddress,
	}
	tx.addAction("SetAssetController", dto)
}

func (tx *Tx) AddCreateAccountAction() string {
	dto := CreateAccountTxActionDto{}
	tx.addAction("CreateAccount", dto)
	return DeriveHash(tx.SenderAddress, tx.Nonce, int16(len(tx.Actions)))
}

func (tx *Tx) AddSetAccountControllerAction(accountHash string, controllerAddress string) {
	dto := SetAccountControllerTxActionDto{
		AccountHash:       accountHash,
		ControllerAddress: controllerAddress,
	}
	tx.addAction("SetAccountController", dto)
}

func (tx *Tx) AddSubmitVoteAction(accountHash string, assetHash string, resolutionHash string, voteHash string) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, resolutionHash: %s, voteHash: %s}", accountHash, assetHash, resolutionHash, voteHash)
	tx.addAction("SubmitVote", json.RawMessage(jsonString))
}

func (tx *Tx) AddSubmitVoteWeightAction(accountHash string, assetHash string, resolutionHash string, voteWeight float64) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, resolutionHash: %s, voteWeight: %f}", accountHash, assetHash, resolutionHash, voteWeight)
	tx.addAction("SubmitVoteWeight", json.RawMessage(jsonString))
}

func (tx *Tx) AddSetAccountEligibilityAction(accountHash string, assetHash string, isPrimaryEligible bool, isSecondaryEligible bool) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, isPrimaryEligible: %t, isSecondaryEligible: %t}", accountHash, assetHash, isPrimaryEligible, isSecondaryEligible)
	tx.addAction("SetAccountEligibility", json.RawMessage(jsonString))
}

func (tx *Tx) AddSetAssetEligibilityAction(assetHash string, isEligibilityRequired bool) {
	jsonString := fmt.Sprintf("{assetHash: %s, isEligibilityRequired: %t}", assetHash, isEligibilityRequired)
	tx.addAction("SetAssetEligibility", json.RawMessage(jsonString))
}

func (tx *Tx) AddChangeKycControllerAddressAction(accountHash string, assetHash string, kycControllerAddress string) {
	jsonString := fmt.Sprintf("{accountHash: %s, assetHash: %s, kycControllerAddress: %s}", accountHash, assetHash, kycControllerAddress)
	tx.addAction("ChangeKycControllerAddress", json.RawMessage(jsonString))
}

func (tx *Tx) AddAddKycProviderAction(assetHash string, providerAddress string) {
	jsonString := fmt.Sprintf("{assetHash: %s, providerAddress: %s}", assetHash, providerAddress)
	tx.addAction("AddKycProvider", json.RawMessage(jsonString))
}

func (tx *Tx) AddRemoveKycProviderAction(assetHash string, providerAddress string) {
	jsonString := fmt.Sprintf("{assetHash: %s, providerAddress: %s}", assetHash, providerAddress)
	tx.addAction("RemoveKycProvider", json.RawMessage(jsonString))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Signing
////////////////////////////////////////////////////////////////////////////////////////////////////

func (tx *Tx) ToJson(indentation bool) string {
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

func (tx *Tx) sign(networkCode []byte, privateKey string) *SignedTx {
	json := tx.ToJson(false)
	signature := SignMessage(networkCode, privateKey, []byte(json))
	signedTx := &SignedTx{
		Tx:        Encode64([]byte(json)),
		Signature: signature,
	}

	return signedTx
}
