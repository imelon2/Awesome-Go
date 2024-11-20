package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	// WebSocket 클라이언트 연결 생성
	client, err := rpc.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer client.Close() // 종료 시 연결 닫기

	// Ethereum의 클라이언트 정보 조회
	var latestBlockNumber string
	err = client.CallContext(context.Background(), &latestBlockNumber, "eth_blockNumber")
	if err != nil {
		log.Fatalf("Failed to get latest block number: %v", err)
	}

	fmt.Printf("Latest Block Number (Hex): %s\n", latestBlockNumber)

	// 웹소켓 연결 종료
	if err != nil {
		log.Fatalf("Failed to close WebSocket: %v", err)
	}

	fmt.Println("WebSocket connection closed.")
	time.Sleep(5 * time.Second)
}
