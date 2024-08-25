package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Printf("Private Key         : %s \n", hexutil.Encode(privateKeyBytes))     // private key
	fmt.Printf("Private Key (no 0x) : %s \n", hexutil.Encode(privateKeyBytes)[2:]) // strip 0x private key

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public Key          : %s \n", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("Address             : %s \n", address)

	/* ====================================================================================================== */

	importPrivateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19") // example PK
	if err != nil {
		log.Fatal(err)
	}
	importPublicKey := importPrivateKey.Public()
	importPublicKeyECDSA, ok := importPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	importAddress := crypto.PubkeyToAddress(*importPublicKeyECDSA).Hex()
	fmt.Printf("Import Address      : %s \n", importAddress) // 0x96216849c49358B10257cb55b28eA603c874b05E
}
