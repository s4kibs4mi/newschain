# NewsChain

Smart contract to publish news on ethereum blockchain.

### Install solc

```
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
```

### Install go-ethereum & abigen

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

### Generate smart contract abi

```
solc --abi contracts/NewsChain.sol -o bindings/
```

### Generate go code

```
$HOME/go/bin/abigen --abi ./bindings/NewsChain.abi --pkg contracts --out bindings/newschain.go
```

### Truffle Install

```
npm install -g truffle
```

### Compile solidity code

```
truffle compile
```

### Deploy smart contract

```
truffle migrate --reset
```
