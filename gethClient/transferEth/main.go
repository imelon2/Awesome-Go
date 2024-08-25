package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value, _ := new(big.Int).SetString("100000000000000000", 10) // 0.1 eth
	gasLimit := uint64(21000)
	// gasPrice := big.NewInt(30000000000) // 20 gwei
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x1001150aE8Ec8843BDcA3c7dE86A291B43a7F835")

	fmt.Printf("From      : %s \n", fromAddress.String())
	fmt.Printf("To        : %s \n", toAddress.String())
	fmt.Printf("Nonce     : %d \n", nonce)
	fmt.Printf("Value     : %s \n", value.String())
	fmt.Printf("Gas Limit : %d \n", gasLimit)
	fmt.Printf("Gas pirce : %s \n", gasPrice)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil /* calldata */)

	chainID, err := client.NetworkID(context.Background())
	fmt.Printf("Chain Id  : %s \n\n", chainID)
	if err != nil {
		fmt.Println("chainID")
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(4693)), privateKey)
	if err != nil {
		fmt.Println("signedTx")
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("SendTransaction")
		log.Fatal(err)
	}

	fmt.Printf("Transaction Hash: %s\n", signedTx.Hash().Hex())
	fmt.Print("Wait Mined Transaction ... \n\n")

	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the receipt to JSON
	receiptBytes, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction receipt in JSON format:")
	fmt.Println(string(receiptBytes))
}
