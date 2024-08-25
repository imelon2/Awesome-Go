package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x1001150aE8Ec8843BDcA3c7dE86A291B43a7F835")

	balance, err := client.BalanceAt(context.Background(), account, nil /* Block Number */)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		log.Fatal(err)
	}

	ethValue2 := new(big.Float).SetInt(pendingBalance)
	ethValue2.Quo(ethValue2, big.NewFloat(1e18))

	fmt.Printf("Balance: %.18f ETH\n", ethValue2)
}
