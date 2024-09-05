이더리움의 `github.com/ethereum/go-ethereum`는 트랜잭션 Type의 문제로 아비트럼 네트워크에서 사용할 수 없다.
때문에, 아비트럼은 `github.com/OffchainLabs/go-ethereum`를 사용해야 한다.
하지만, 아비트럼은 Go Package를 제공하지 않는다. 때문에, 직접 Repository를 Clone해서 사용해야한다.

go get github.com/ethereum/go-ethereum@v1.13.13


```
git clone https://github.com/OffchainLabs/go-ethereum.git
```

`github.com/OffchainLabs/go-ethereum`는 내부적으로 `github.com/ethereum/go-ethereum v1.13.13`를 설치하여 사용한다.
때문에 replace를 v1.13.13으로 받아 사용해야한다.