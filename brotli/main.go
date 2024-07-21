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
	// data := []byte("0x02f87383066eee50839896808405f5e100826b4b94572b7a9650d7c58e3ceafe4a6bbedbfaf391d3eb872386f26fc1000080c001a02a3b6dd2f3ce1f514b68e4632595fbcfe8007738781c4454f4ce8e802eb855aca0313fe54d6adb2e62423be620f8e7a005838e8ed5b8f2a3d42760f74de004fefe")
	_hexData := "f87383066eee50839896808405f5e100826b4b94572b7a9650d7c58e3ceafe4a6bbedbfaf391d3eb872386f26fc1000080c001a02a3b6dd2f3ce1f514b68e4632595fbcfe8007738781c4454f4ce8e802eb855aca0313fe54d6adb2e62423be620f8e7a005838e8ed5b8f2a3d42760f74de004fefe"
	data, err := hex.DecodeString(_hexData)

	// Brotli 압축 수준 설정 (0~11 사이의 값, 기본값은 11)
	// options := brotli.WriterOptions{Quality: 1}

	// 설정된 옵션으로 Brotli writer 생성
	var compressedBuffer bytes.Buffer

	writer := brotli.NewWriterLevel(&compressedBuffer, 1)
	writer.Write(data)
	writer.Close()

	println(brotli.DefaultCompression)
	compressedData := compressedBuffer.Bytes()
	fmt.Printf("Compressed data: %x\n", compressedData)
	fmt.Printf("Compressed length: %d\n", compressedBuffer.Len())

	// Brotli로 압축 해제
	reader := brotli.NewReader(bytes.NewReader(compressedData))
	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error decompressing data:", err)
		return
	}

	fmt.Printf("Decompressed data: %s\n", decompressedData)
}
