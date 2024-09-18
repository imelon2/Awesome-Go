package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 첫 번째 JSON 데이터
	erc20ABI := `[
		{
			"constant": false,
			"inputs": [
				{"name": "to", "type": "address"},
				{"name": "value", "type": "uint256"}
			],
			"name": "transfer",
			"outputs": [
				{"name": "", "type": "bool"}
			],
			"type": "function"
		}
	]`

	// 두 번째 JSON 데이터
	erc20ABI2 := `[
		{
			"constant": false,
			"inputs": [
				{"name": "to", "type": "address"},
				{"name": "value", "type": "uint256"}
			],
			"name": "transfer",
			"outputs": [
				{"name": "", "type": "bool"}
			],
			"type": "function"
		}
	]`

	// 첫 번째 JSON을 구조체로 디코딩
	var functions1 []interface{}
	err := json.Unmarshal([]byte(erc20ABI), &functions1)
	if err != nil {
		fmt.Println("Error decoding erc20ABI:", err)
		return
	}

	// 두 번째 JSON을 구조체로 디코딩
	var functions2 []interface{}
	err = json.Unmarshal([]byte(erc20ABI2), &functions2)
	if err != nil {
		fmt.Println("Error decoding erc20ABI2:", err)
		return
	}

	// 두 배열을 합침
	combinedFunctions := append(functions1, functions2...)

	// 합친 배열을 다시 JSON으로 인코딩
	combinedJSON, err := json.MarshalIndent(combinedFunctions, "", "  ")
	if err != nil {
		fmt.Println("Error encoding combined functions:", err)
		return
	}

	// 결과 출력
	fmt.Println(string(combinedJSON))
}
