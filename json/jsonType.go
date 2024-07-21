package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const (
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 20
)

type Address [AddressLength]byte
type Hash [HashLength]byte

type JsonType struct {
	Kind        uint8   `json:"kind"`
	Poster      Address `json:"sender"`
	BlockNumber uint64  `json:"blockNumber"`
	IsEmpty     bool    `json:"IsEmpty"`
	IsOmitempty bool    `json:"IsOmitempty,omitempty"` // omitempty : JSON 인코딩 시 값이 nil일 경우 해당 키가 생략됩니다
}

func main() {
	// JSON 데이터 정의
	ex := JsonType{Kind: 5, Poster: Address{19: 0x01}, BlockNumber: 20}

	/*
	 json.Marshal을 통해 json형태로 변환되는데 이때 []byte 형태로 변환되고 이는 ASCII CODE 형태이다.
	 이걸 string으로 변환하면 그대로 JSON이 출력된다.
	*/
	jsonString, err := json.Marshal(ex)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON To String : ", string(jsonString))
	fmt.Println("====================================")

	var jsonData JsonType
	err1 := json.Unmarshal([]byte(jsonString), &jsonData)
	if err1 != nil {
		panic(err1)
	}

	fmt.Printf("kind : %d\nPoster : %s\nBlockNumber : %d\nIsEmpty : %t\nIsOmitempty : %t\n", jsonData.Kind, hex.EncodeToString(jsonData.Poster[:]), jsonData.BlockNumber, jsonData.IsEmpty, jsonData.IsOmitempty)

}
