package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// uint256 값을 디코딩하는 함수
func decodeUint256(data []byte) (*big.Int, error) {
	if len(data) != 32 {
		return nil, fmt.Errorf("expected 32 bytes, got %d", len(data))
	}
	return new(big.Int).SetBytes(data), nil
}

// https://github.com/OffchainLabs/arbitrum-sdk/blob/792a7ee3ccf09842653bc49b771671706894cbb4/src/lib/message/messageDataParser.ts#L13
func main() {
	// Solidity에서 인코딩된 데이터를 hex 문자열로 가정 (예제 값)
	encodedData := "00000000000000000000000065e1a5e8946e7e87d9774f5288f41c30a99fd30200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d614c3b5c8c00000000000000000000000000000000000000000000000000000cf60e420bbc000000000000000000000000086d9a1bc823afe264bbb5638873c25b1091d877c00000000000000000000000086d9a1bc823afe264bbb5638873c25b1091d877c000000000000000000000000000000000000000000000000000000000001dfcb000000000000000000000000000000000000000000000000000000000393870000000000000000000000000000000000000000000000000000000000000000c42e567b36000000000000000000000000c944e90c64b2c07662a292be6244bdf05cda44a700000000000000000000000086d9a1bc823afe264bbb5638873c25b1091d877c00000000000000000000000086d9a1bc823afe264bbb5638873c25b1091d877c000000000000000000000000000000000000000000000211fd618b2e9f2613a100000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000"

	// 16진수로 인코딩된 데이터를 바이트 배열로 변환
	// data := common.Hex2Bytes(encodedData)
	data, err := hex.DecodeString(encodedData)
	if err != nil {
		fmt.Print("\nhere\n")
		log.Fatal(err)
	}

	// 바이트 데이터를 읽기 위한 버퍼
	buf := bytes.NewReader(data)

	// 각 32바이트 값을 읽어서 디코딩
	toBytes := make([]byte, 32)
	l2CallValueBytes := make([]byte, 32)
	amountBytes := make([]byte, 32)
	maxSubmissionCostBytes := make([]byte, 32)
	excessFeeRefundAddressBytes := make([]byte, 32)
	callValueRefundAddressBytes := make([]byte, 32)
	gasLimitBytes := make([]byte, 32)
	maxFeePerGasBytes := make([]byte, 32)
	dataLengthBytes := make([]byte, 32)

	// 바이트 배열에서 각 값을 읽어옴
	_, err = buf.Read(toBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(l2CallValueBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(amountBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(maxSubmissionCostBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(excessFeeRefundAddressBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(callValueRefundAddressBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(gasLimitBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(maxFeePerGasBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Read(dataLengthBytes)
	if err != nil {
		log.Fatal(err)
	}

	l2CallValue, err := decodeUint256(l2CallValueBytes)
	if err != nil {
		log.Fatal(err)
	}
	amount, err := decodeUint256(amountBytes)
	if err != nil {
		log.Fatal(err)
	}
	maxSubmissionCost, err := decodeUint256(maxSubmissionCostBytes)
	if err != nil {
		log.Fatal(err)
	}
	gasLimit, err := decodeUint256(gasLimitBytes)
	if err != nil {
		log.Fatal(err)
	}
	maxFeePerGas, err := decodeUint256(maxFeePerGasBytes)
	if err != nil {
		log.Fatal(err)
	}
	dataLength, err := decodeUint256(dataLengthBytes)
	if err != nil {
		log.Fatal(err)
	}

	remainingData := make([]byte, buf.Len()) // 남은 데이터 길이만큼 버퍼를 할당
	_, err = buf.Read(remainingData)
	if err != nil {
		log.Fatal(err)
	}

	// 결과 출력
	fmt.Printf("To: %s\n", common.HexToAddress(hex.EncodeToString(toBytes)))
	fmt.Printf("L2 Call Value: %s\n", l2CallValue.String())
	fmt.Printf("Amount: %s\n", amount.String())
	fmt.Printf("maxSubmissionCost: %s\n", maxSubmissionCost.String())
	fmt.Printf("excessFeeRefundAddress: %s\n", common.HexToAddress(hex.EncodeToString(excessFeeRefundAddressBytes)))
	fmt.Printf("callValueRefundAddress: %s\n", common.HexToAddress(hex.EncodeToString(callValueRefundAddressBytes)))
	fmt.Printf("gasLimit: %s\n", gasLimit.String())
	fmt.Printf("maxFeePerGas: %s\n", maxFeePerGas.String())
	fmt.Printf("dataLength: %s\n", dataLength.String())
	fmt.Printf("나머지 buf: 0x%s\n", hex.EncodeToString(remainingData))
}
