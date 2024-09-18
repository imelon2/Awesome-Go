package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// JSON 파일을 읽음
	fileContent, err := os.ReadFile("./abi/aggregateAbi.json")
	if err != nil {
		log.Fatalf("Failed to read ABI JSON file: %v", err)
	}

	// 파일 내용을 ABI 파싱에 전달
	parsedABI, err := abi.JSON(strings.NewReader(string(fileContent)))
	if err != nil {
		log.Fatalf("Failed to get method from parsedABI: %v", err)
	}
	// 예시 calldata
	calldata := common.FromHex("0x6bf6a42d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013abeaf000000000000000000000000000000000000000000000000000000000ec1fa3a0000000000000000000000000000000000000000000000000000000000000001")
	arg := calldata[4:]
	// 함수 시그니처 추출
	method, err := parsedABI.MethodById(common.FromHex("0x6bf6a42d"))
	if err != nil {
		log.Fatalf("Failed to get method from calldata: %v", err)
	}

	inter, err := method.Inputs.Unpack(arg)

	if err != nil {
		log.Fatalf("Failed to unpack calldata: %v", err)
	}

	fmt.Printf("Decoded Function : %s\n", parsedABI.Methods["transfer"])

	for i, data := range inter {
		fmt.Printf("Decoded %s : %s\n", method.Inputs[i].Name, data)
	}
	// result, err := method.Inputs.Pack(inter)
	// err = parsedABI.UnpackIntoInterface(inter, method.Name, arg)
}
