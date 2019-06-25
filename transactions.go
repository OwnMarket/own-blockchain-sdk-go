package ownSdk

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

type SubmitVoteTxActionDto struct {
	AccountHash    string `json:"accountHash"`
	AssetHash      string `json:"assetHash"`
	ResolutionHash string `json:"resolutionHash"`
	VoteHash       string `json:"voteHash"`
}

type SubmitVoteWeightTxActionDto struct {
	AccountHash    string  `json:"accountHash"`
	AssetHash      string  `json:"assetHash"`
	ResolutionHash string  `json:"resolutionHash"`
	VoteWeight     float64 `json:"voteWeight"`
}

type SetAccountEligibilityTxActionDto struct {
	AccountHash         string `json:"accountHash"`
	AssetHash           string `json:"assetHash"`
	IsPrimaryEligible   bool   `json:"isPrimaryEligible"`
	IsSecondaryEligible bool   `json:"isSecondaryEligible"`
}

type SetAssetEligibilityTxActionDto struct {
	AssetHash             string `json:"assetHash"`
	IsEligibilityRequired bool   `json:"isEligibilityRequired"`
}

type ChangeKycControllerAddressTxActionDto struct {
	AccountHash          string `json:"accountHash"`
	AssetHash            string `json:"assetHash"`
	KycControllerAddress string `json:"kycControllerAddress"`
}

type AddKycProviderTxActionDto struct {
	AssetHash       string `json:"assetHash"`
	ProviderAddress string `json:"providerAddress"`
}

type RemoveKycProviderTxActionDto struct {
	AssetHash       string `json:"assetHash"`
	ProviderAddress string `json:"providerAddress"`
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
	dto := SubmitVoteTxActionDto{
		AccountHash:    accountHash,
		AssetHash:      assetHash,
		ResolutionHash: resolutionHash,
		VoteHash:       voteHash,
	}
	tx.addAction("SubmitVote", dto)
}

func (tx *Tx) AddSubmitVoteWeightAction(accountHash string, assetHash string, resolutionHash string, voteWeight float64) {
	dto := SubmitVoteWeightTxActionDto{
		AccountHash:    accountHash,
		AssetHash:      assetHash,
		ResolutionHash: resolutionHash,
		VoteWeight:     voteWeight,
	}
	tx.addAction("SubmitVoteWeight", dto)
}

func (tx *Tx) AddSetAccountEligibilityAction(accountHash string, assetHash string, isPrimaryEligible bool, isSecondaryEligible bool) {
	dto := SetAccountEligibilityTxActionDto{
		AccountHash:         accountHash,
		AssetHash:           assetHash,
		IsPrimaryEligible:   isPrimaryEligible,
		IsSecondaryEligible: isSecondaryEligible,
	}
	tx.addAction("SetAccountEligibility", dto)
}

func (tx *Tx) AddSetAssetEligibilityAction(assetHash string, isEligibilityRequired bool) {
	dto := SetAssetEligibilityTxActionDto{
		AssetHash:             assetHash,
		IsEligibilityRequired: isEligibilityRequired,
	}
	tx.addAction("SetAssetEligibility", dto)
}

func (tx *Tx) AddChangeKycControllerAddressAction(accountHash string, assetHash string, kycControllerAddress string) {
	dto := ChangeKycControllerAddressTxActionDto{
		AccountHash:          accountHash,
		AssetHash:            assetHash,
		KycControllerAddress: kycControllerAddress,
	}
	tx.addAction("ChangeKycControllerAddress", dto)
}

func (tx *Tx) AddAddKycProviderAction(assetHash string, providerAddress string) {
	dto := AddKycProviderTxActionDto{
		AssetHash:       assetHash,
		ProviderAddress: providerAddress,
	}
	tx.addAction("AddKycProvider", dto)
}

func (tx *Tx) AddRemoveKycProviderAction(assetHash string, providerAddress string) {
	dto := RemoveKycProviderTxActionDto{
		AssetHash:       assetHash,
		ProviderAddress: providerAddress,
	}
	tx.addAction("RemoveKycProvider", dto)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Signing
////////////////////////////////////////////////////////////////////////////////////////////////////

func toJson(data interface{}, indentation bool) string {
	var b []byte
	var err error
	if indentation {
		b, err = json.MarshalIndent(data, "", "    ")
	} else {
		b, err = json.Marshal(data)
	}

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(b)
}
func (tx *Tx) ToJson(indentation bool) string {
	return toJson(tx, indentation)
}

func (signedTx *SignedTx) ToJson(indentation bool) string {
	return toJson(signedTx, indentation)
}

func (tx *Tx) Sign(networkCode []byte, privateKey string) *SignedTx {
	json := tx.ToJson(false)
	signature := SignMessage(networkCode, privateKey, []byte(json))
	signedTx := &SignedTx{
		Tx:        Encode64([]byte(json)),
		Signature: signature,
	}

	return signedTx
}
