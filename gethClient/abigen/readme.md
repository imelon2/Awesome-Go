solc --abi ./abigen/contracts/Store.sol -o ./abigen/build 

abigen --abi=./abigen/build/Store.abi --pkg=store --out=./abigen/factory/Store.go

solc --bin ./abigen/contracts/Store.sol -o ./abigen/build
abigen --bin=./abigen/build/Store.bin --abi=./abigen/build/Store.abi --pkg=store --out=./abigen/factory/Store.go
