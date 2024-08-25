package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	// contains filtered or unexported fields
}

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	header, err := client.HeaderByNumber(context.Background(), nil /* Latest */)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(header.Number.Int64())

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144
}
