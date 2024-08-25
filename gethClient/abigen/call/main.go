package main

import (
	"fmt"
	"log"

	store "my-geth/abigen/factory"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x41e646D31eF6d96Ee6901938A299b4549721597e")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("string public version : %s \n", version)
}
