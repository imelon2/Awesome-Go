package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	_ "os"

	"github.com/andybalholm/brotli"
)

func main() {
	// data := []byte("0x02f87583066eee7e8459682f00846553f10083063f3094d644352a429f3ff3d21128820dcbc53e063685b1872386f26fc1000080c080a09937bb87e8f6005f927b6fa708c15f2a86f7a4b780a483989d667a798591b5caa013b09a720502ba3c11ee3d3e496683599bf4f5121ae062789f5073d7c94e685a")
	_hexData := "01f86e8208ae018405f5e1008252b694940e3cb4f37ae0259499e71f3a558b5de0471fa0872386f26fc1000080c080a085160014ba197f3472c82bbfaa1e1cb16b0859c7268d07e8c0acd9b9813f23cea0727a66d3f9dd1b363a7e0ce4d644bce65434d21724d8d21700484899a04b1ed7"

	data, err1 := hex.DecodeString(_hexData)

	if err1 != nil {
		fmt.Println("Error decompressing data:", err1)
		return
	}

	// Brotli 압축 수준 설정 (0~11 사이의 값, 기본값은 11)
	// options := brotli.WriterOptions{Quality: 1}

	// 설정된 옵션으로 Brotli writer 생성
	var compressedBuffer bytes.Buffer

	writer := brotli.NewWriterLevel(&compressedBuffer, 10)
	writer.Write(data)
	writer.Close()

	compressedData := compressedBuffer.Bytes()
	fmt.Printf("Compressed data: %x \n", compressedData)
	fmt.Printf("Compressed length: %d \n", compressedBuffer.Len())
	// return
	// Brotli로 압축 해제
	reader := brotli.NewReader(bytes.NewReader(compressedData))
	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error decompressing data:", err)
		return
	}

	// fmt.Printf("Decompressed data: %s\n", decompressedData)
	fmt.Printf("\n")
	fmt.Printf("Decompressed data: %s\n", hex.EncodeToString(decompressedData))

}
