package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
 * go get github.com/ethereum/go-ethereum/accounts/keystore@v1.14.8
 */

func createKs() {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	// 0xa526210406df1b40c6b176cfbca08cfe6da5f4442271976e8d96eb9d413c9a35

	privateKey, err := crypto.HexToECDSA("a526210406df1b40c6b176cfbca08cfe6da5f4442271976e8d96eb9d413c9a35")
	password := ""
	account, err := ks.ImportECDSA(privateKey, password)
	// account, err := ks.NewAccount(password)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // b06778606daddbdd1306e6e5604c6ac8684773d3
}

func importKs() {
	file := "./keystore/UTC--2024-08-25T00-57-16.305567000Z--dfb859857f05c24e201419d7b43df70f2dec7cf3"
	ks := keystore.NewKeyStore("./keystore/tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // b06778606daddbdd1306e6e5604c6ac8684773d3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func raedWallets() {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	// wallets := ks.Wallets()

	// a := wallets[0].Accounts()
	a := ks.Accounts()

	fmt.Println(a[0].Address)
	fmt.Println(a[0].URL)
	fmt.Println(a[0].URL.Path)
	fmt.Println(a[0].URL.Scheme)
	fmt.Println(len(a))
}

func main() {
	createKs()

	// importKs()

	// raedWallets()
}
