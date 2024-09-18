package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Dir(filename)

	// JSON 파일을 읽음
	fileContent, err := os.ReadFile(filepath.Join(root, "abi", "aggregateAbi.json"))
	if err != nil {
		log.Fatalf("Failed to read ABI JSON file: %v", err)
	}

	// 파일 내용을 ABI 파싱에 전달
	parsedABI, err := abi.JSON(strings.NewReader(string(fileContent)))
	if err != nil {
		log.Fatalf("Failed to get method from parsedABI: %v", err)
	}

	client, err := ethclient.Dial("https://ethereum-sepolia.g.allthatnode.com/full/evm/5da65013a9004d8da1983f17cae83366")
	if err != nil {
		log.Fatal(err)
	}
	txHash := common.HexToHash("0xa93c42dde60802530502af25a487d451018bc229fbf4a25f6838c12b2d355b0b")

	tx, err := client.TransactionReceipt(context.Background(), txHash)

	// fmt.Print(tx.Bloom., "\n\n")
	topic := tx.Logs[0].Topics[0]
	// _topic := tx.Logs[0].Data
	event, err := parsedABI.EventByID(topic) /* function selector */

	fmt.Print(len(tx.Logs[0].Topics), "\n\n")

	jsonCalldata := make(map[string]interface{})
	for _, data := range tx.Logs {
		var eventHashTopic = data.Topics[0]
		event, _ := parsedABI.EventByID(eventHashTopic)
		data.Topics = data.Topics[1:]

		_jsonCalldata := make(map[string]interface{})
		for j, topic := range data.Topics {
			_jsonCalldata[event.Inputs[j].Name] = topic
		}
		inter, err := event.Inputs.Unpack(data.Data)
		if err != nil {
			log.Fatalf("Failed to unpack calldata: %v", err)
		}

		for _, data := range inter {
			_jsonCalldata[event.Inputs[2].Name] = data // utils.ConvertBytesToHex(data)
		}
		jsonCalldata[event.Name] = _jsonCalldata
	}

	jsonData, err := json.MarshalIndent(jsonCalldata, "", "  ")
	if err != nil {
		log.Fatalf("Failed to MarshalIndent calldata: %v", err)
	}
	fmt.Printf("Function : %s\n", parsedABI.Events[event.Name])
	fmt.Println(string(jsonData))

}
