package ownsdk

import (
	"testing"

	ownsdk "github.com/OwnMarket/own-blockchain-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	password := []byte("pass")
	decodedPassword := ownsdk.Decode58(ownsdk.Hash(password))
	var passwordHash [32]byte
	copy(passwordHash[:], decodedPassword)
	encryptedText := ownsdk.Encrypt([]byte("Chainium"), passwordHash)
	decryptedText := ownsdk.Decrypt(encryptedText, passwordHash)

	assert.Equal(t, "Chainium", string(decryptedText))
}

func TestEncoding(t *testing.T) {
	encoded := ownsdk.Encode58([]byte("Chainium"))
	decoded := ownsdk.Decode58(encoded)
	assert.Equal(t, "Chainium", string(decoded))
}

func TestWallet(t *testing.T) {
	wallet := ownsdk.GenerateWallet()
	addressFromPrivateKey := ownsdk.AddressFromPrivateKey(wallet.PrivateKey)
	assert.Equal(t, wallet.Address, addressFromPrivateKey)
}

func TestSigningMessage(t *testing.T) {
	msg := []byte("Chainium")
	networkCode := []byte("UNIT_TESTS") //TODO: replace with OWN_PUBLIC_BLOCKCHAIN_MAINNET for mainnet!
	sig := ownsdk.SignMessage(networkCode, "B6WNNx9oK8qRUU52PpzjXHZuv4NUb3Z33hdju3hhrceS", msg)
	expectedSig := "6Hhxz2eP3AagR56mP4AAaKViUxHi3gM9c5weLDR48x4X4ynRBDfxsHGjhX9cni1mtCkNxbnZ783YPgMwVYV52X1w5"
	assert.Equal(t, expectedSig, sig)
}

func TestSignPlainText(t *testing.T) {
	txt := []byte("Chainium")
	privateKey := "3rzY3EENhYrWXzUqNnMEbGUr3iEzzSZrjMwJ1CgQpJpq"
	expectedSig := "EzCsWgPozyVT9o6TycYV6q1n4YK4QWixa6Lk4GFvwrj6RU3K1wHcwNPZJUMBYcsGp5oFhytHiThon5zqE8uLk8naB"
	sig := ownsdk.SignPlainText(privateKey, txt)
	assert.Equal(t, expectedSig, sig)
}

func TestVerifyPlainTextSignature(t *testing.T) {
	txt := []byte("Chainium")
	wallet := ownsdk.GenerateWallet()
	expectedAddress := ownsdk.AddressFromPrivateKey(wallet.PrivateKey)
	sig := ownsdk.SignPlainText(wallet.PrivateKey, txt)
	address := ownsdk.VerifyPlainTextSignature(sig, txt)
	assert.Equal(t, expectedAddress, address)
}

func TestHdCrypto(t *testing.T) {
	mnemonic := "receive raccoon rocket donkey cherry garbage medal skirt random smoke young before scale leave hold insect foster blouse mail donkey regular vital hurt april"
	seed := ownsdk.GenerateSeedFromMnemonic(mnemonic, "")
	wallet := ownsdk.GenerateWalletFromSeed(seed, 0)
	wallets := ownsdk.RestoreWalletsFromSeed(seed, 1)
	expectedPrivateKey := "ECPVXjz78oMdmLKbHVAAo7X7evtTh4EfnaW5Yc1SHWaj"
	expectedAddress := "CHb5Z6Za34nv28Z3rLZ2Yd8LFikHaTqLhxB"
	assert.Equal(t, expectedPrivateKey, wallet.PrivateKey)
	assert.Equal(t, expectedPrivateKey, wallets[0].PrivateKey)
	assert.Equal(t, expectedAddress, wallet.Address)
	assert.Equal(t, expectedAddress, wallets[0].Address)
}
