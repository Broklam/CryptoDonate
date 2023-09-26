package controllers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/ripemd160"
)

func generateRandom32Bytes() []byte {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return bytes
}

func getCompressedPrivateKey(privateKey []byte) []byte {
	decoded, _ := hex.DecodeString(fmt.Sprintf("%x01", privateKey))
	return decoded
}

func encodePrivateKey(privateKey []byte) string {
	return base58.CheckEncode(privateKey, 128)
}

func getPublicKey(x *big.Int, y *big.Int) string {
	return fmt.Sprintf("04%064x%064x", x, y)
}

func getCompressedPublicKey(x *big.Int, y *big.Int) string {
	z := new(big.Int)
	z.Mod(y, big.NewInt(2))
	if z.Cmp(big.NewInt(1)) == 0 {
		return fmt.Sprintf("03%064x", x)
	} else {
		return fmt.Sprintf("02%064x", x)
	}
}

func publicKeyToAddress(publicKey string) string {
	decoded, _ := hex.DecodeString(publicKey)

	sha := sha256.New()
	sha.Write(decoded)
	intermed := sha.Sum(nil)

	ripemd := ripemd160.New()
	ripemd.Write(intermed)
	digest := ripemd.Sum(nil)

	return base58.CheckEncode(digest, 0)
}
func CreateBtcWallet() (wif string, publickey string, btcaddress string) {
	privateKey := generateRandom32Bytes()
	WifPrivateKey := encodePrivateKey(privateKey)
	s256 := secp256k1.S256()
	x, y := s256.ScalarBaseMult(privateKey)
	publicKey := getPublicKey(x, y)
	BtcAddress := publicKeyToAddress(publicKey)
	return WifPrivateKey, publicKey, BtcAddress
}
