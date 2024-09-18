package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var homeDir string

	// 운영 체제에 따라 환경 변수를 읽음
	if runtime.GOOS == "windows" {
		homeDir = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	} else {
		homeDir = os.Getenv("HOME")
	}

	// 홈 디렉토리 출력
	fmt.Println("Home Directory:", homeDir)
}
