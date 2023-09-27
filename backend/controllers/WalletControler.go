package controllers

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"golang.org/x/crypto/ripemd160"
)

func generateRandom32Bytes() []byte {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return bytes
}

func encodePrivateKey(privateKey []byte) string {
	return base58.CheckEncode(privateKey, 128)
}

func getPublicKey(x *big.Int, y *big.Int) string {
	return fmt.Sprintf("04%064x%064x", x, y)
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

func CreateEthWallet() (wif string, publickey string, ethaddress string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Panicln(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	wif = hexutil.Encode(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	PPublickey := hexutil.Encode(publicKeyBytes)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return wif, PPublickey, address
}

func CreateTonWallet() (wif string, publickey string, tonaddress string) {
	words := wallet.NewSeed()
	client := liteclient.NewConnectionPool()
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), "https://ton.org/global.config.json")
	if err != nil {
		log.Fatalln("get config err: ", err.Error())
		return
	}
	api := ton.NewAPIClient(client, ton.ProofCheckPolicySecure).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)
	w, err := wallet.FromSeed(api, words, wallet.V3)
	if err != nil {
		panic(err)
	}
	wif = strings.Join(words, " ")
	tonaddress = w.Address().String()
	publickey = ""
	return wif, publickey, tonaddress
}
