package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// ERC20 ABI 정의
	erc20ABI := `[{
		"constant": false,
		"inputs": [{"name": "to", "type": "address"},{"name": "value", "type": "uint256"}],
		"name": "transfer",
		"outputs": [{"name": "", "type": "bool"}],
		"type": "function"
	}]`

	// ABI 파싱
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// 예시 calldata
	calldata := common.FromHex("0xa9059cbb0000000000000000000000001c6c54f1e18e5c47dfae928aaf8779de7e09e44800000000000000000000000000000000000000000000000000000000000000a")
	arg := calldata[4:]
	// 함수 시그니처 추출
	method, err := parsedABI.MethodById(common.FromHex("0xa9059cbc"))
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
