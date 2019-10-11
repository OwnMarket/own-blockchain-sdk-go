package ownSdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// Types
////////////////////////////////////////////////////////////////////////////////////////////////////

type WalletInfo struct {
	PrivateKey string
	Address    string
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Encryption
////////////////////////////////////////////////////////////////////////////////////////////////////

func Encrypt(text []byte, passwordHash [32]byte) []byte {
	cypher, err := aes.NewCipher(passwordHash[:])
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0)
	}

	gcm, err := cipher.NewGCM(cypher)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	return gcm.Seal(nonce, nonce, text, nil)
}

func Decrypt(encryptedText []byte, passwordHash [32]byte) []byte {
	cypher, err := aes.NewCipher(passwordHash[:])
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(cypher)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedText) < nonceSize {
		fmt.Println(err)
	}
	nonce, encryptedText := encryptedText[:nonceSize], encryptedText[nonceSize:]
	text, err := gcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		panic(err.Error())
	}

	return text
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Encoding
////////////////////////////////////////////////////////////////////////////////////////////////////

func Encode64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Decode64(src string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		fmt.Println(err.Error())
	}
	return decoded
}

func Encode58(src []byte) string {
	return base58.Encode(src)
}

func Decode58(src string) []byte {
	decoded, err := base58.Decode(src)
	if err != nil {
		fmt.Println(err.Error())
		return make([]byte, 0)
	}
	return decoded
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Hashing
////////////////////////////////////////////////////////////////////////////////////////////////////

func xsha256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func xsha512(data []byte) [64]byte {
	return sha512.Sum512(data)
}

func xsha160(data []byte) []byte {
	_sha512 := xsha512(data)
	return _sha512[:20]
}

func Hash(data []byte) string {
	_sha256 := xsha256(data)
	return Encode58(_sha256[:])
}

func DeriveHash(address string, nonce int64, txActionNumber int16) string {
	addressBytes := Decode58(address)
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, uint64(nonce))
	txActionNumberBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(txActionNumberBytes, uint16(txActionNumber))
	concatBytes := make([]byte, 0)
	concatBytes = append(concatBytes, addressBytes[:]...)
	concatBytes = append(concatBytes, nonceBytes[:]...)
	concatBytes = append(concatBytes, txActionNumberBytes[:]...)
	return Hash(concatBytes)
}

func blockchainAddress(publicKey []byte) string {
	addressPrefix := []byte{6, 90} //CH
	_xsha256 := xsha256(publicKey)
	_xsha160_256 := xsha160(_xsha256[:])
	publicKeyHashWithPrefix := append(addressPrefix, _xsha160_256...)
	_xsha256 = xsha256(publicKeyHashWithPrefix)
	_xsha256_256 := xsha256(_xsha256[:])
	checksum := _xsha256_256[:4]
	return Encode58(append(publicKeyHashWithPrefix, checksum...))
}

func IsValidBlockchainAddress(address string) bool {
	addressPrefix := []byte{6, 90} //CH
	if address == "" || !strings.HasPrefix(address, "CH") {
		return false
	}
	addressBytes := Decode58(address)
	if len(addressBytes) != 26 || bytes.Compare(addressBytes[:2], addressPrefix[:]) != 0 {
		return false
	}

	publicKeyHashWithPrefix := addressBytes[:22]
	checksum := addressBytes[22:]
	_xsha256 := xsha256(publicKeyHashWithPrefix)
	_xsha256_256 := xsha256(_xsha256[:])
	calculatedChecksum := _xsha256_256[:4]

	return bytes.Compare(checksum, calculatedChecksum) == 0
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Signing
////////////////////////////////////////////////////////////////////////////////////////////////////

func GenerateWallet() *WalletInfo {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKey := elliptic.Marshal(secp256k1.S256(), key.X, key.Y)

	privateKeyBytes := make([]byte, 32)
	blob := key.D.Bytes()
	copy(privateKeyBytes[32-len(blob):], blob)

	wallet := &WalletInfo{
		PrivateKey: Encode58(privateKeyBytes),
		Address:    blockchainAddress(publicKey),
	}
	return wallet
}

func AddressFromPrivateKey(privateKey string) string {
	bytes := Decode58(privateKey)
	key, err := crypto.ToECDSA(bytes)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	publicKey := elliptic.Marshal(secp256k1.S256(), key.X, key.Y)
	return blockchainAddress(publicKey)
}

func WalletFromPrivateKey(privateKey string) *WalletInfo {
	address := AddressFromPrivateKey(privateKey)
	wallet := &WalletInfo{
		PrivateKey: privateKey,
		Address:    address,
	}
	return wallet
}

func sign(privateKey string, dataHash [32]byte) string {
	privateKeyBytes := Decode58(privateKey)
	signatureBytes, err := secp256k1.Sign(dataHash[:], privateKeyBytes)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return Encode58(signatureBytes)
}

func SignMessage(networkCode string, privateKey string, message string) string {
	messageHash := xsha256([]byte(message))
	networkIdBytes := xsha256([]byte(networkCode))
	dataToSign := xsha256(append(messageHash[:], networkIdBytes[:]...))
	return sign(privateKey, dataToSign)
}

func SignPlainText(privateKey string, text string) string {
	dataToSign := xsha256([]byte(text))
	return sign(privateKey, dataToSign)
}

func VerifyPlainTextSignature(signature string, text string) string {
	dataToVerify := xsha256([]byte(text))
	signatureBytes := Decode58(signature)
	publicKey, err := secp256k1.RecoverPubkey(dataToVerify[:], signatureBytes)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return blockchainAddress(publicKey)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Hierarchical Deterministic Cryptography
////////////////////////////////////////////////////////////////////////////////////////////////////

func GenerateMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

func GenerateSeedFromMnemonic(mnemonic string, passphrase string) []byte {
	if !bip39.IsMnemonicValid(mnemonic) {
		panic("Invalid mnemonic")
	}
	return bip39.NewSeed(mnemonic, passphrase)
}

func GenerateSeedFromKeystore(keyStoreEncrypted []byte, passwordHash [32]byte) []byte {
	return Decrypt(keyStoreEncrypted, passwordHash)
}

func generateMasterKeyFromSeed(seed []byte) *bip32.Key {
	masterKey, _ := bip32.NewMasterKey(seed)
	return masterKey
}

func GenerateKeystore(mnemonic string, passwordHash [32]byte) []byte {
	seed := GenerateSeedFromMnemonic(mnemonic, "")
	return Encrypt(seed, passwordHash)
}

// m/44'/{coin}'/0'/0/${address}
func newKeyFromMasterKey(masterKey *bip32.Key, coin, address uint32) (*bip32.Key, error) {
	const _purpose uint32 = 44
	const _account = 0
	const _chain = 0

	child, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44) // m/44'
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(bip32.FirstHardenedChild + coin) // m/44'/${coin}'
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(bip32.FirstHardenedChild + _account) // m/44'/${coin}'/0'
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(_chain) // m/44'/${coin}'/0'/0
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(address) // m/44'/${coin}'/0'/0/${address}
	if err != nil {
		return nil, err
	}

	return child, nil
}

func generateWalletFromSeedWithExplicitCoinIndex(seed []byte, coin uint32, keyIndex uint32) *WalletInfo {
	masterKey := generateMasterKeyFromSeed(seed)
	childKey, err := newKeyFromMasterKey(masterKey, coin, keyIndex)
	if err != nil {
		fmt.Println(err)
	}
	privateKeyBytes := childKey.Key
	privateKey := Encode58(privateKeyBytes)
	wallet := &WalletInfo{
		PrivateKey: privateKey,
		Address:    AddressFromPrivateKey(privateKey),
	}
	return wallet
}

func GenerateWalletFromSeed(seed []byte, keyIndex uint32) *WalletInfo {
	return generateWalletFromSeedWithExplicitCoinIndex(seed, 25718, keyIndex)
}

func RestoreWalletsFromSeed(seed []byte, walletCount uint32) [](*WalletInfo) {
	var wallets [](*WalletInfo)
	var i uint32 = 0
	for ; i < walletCount; i++ {
		wallet := GenerateWalletFromSeed(seed, i)
		wallets = append(wallets, wallet)
	}
	return wallets
}

func GenerateWalletFromKeystore(keyStoreEncrypted []byte, passwordHash [32]byte, keyIndex uint32) *WalletInfo {
	seed := GenerateSeedFromKeystore(keyStoreEncrypted, passwordHash)
	return GenerateWalletFromSeed(seed, keyIndex)
}

func RestoreWalletsFromKeystore(keyStoreEncrypted []byte, passwordHash [32]byte, walletCount uint32) [](*WalletInfo) {
	seed := GenerateSeedFromKeystore(keyStoreEncrypted, passwordHash)
	return RestoreWalletsFromSeed(seed, walletCount)
}
