# NewsChain

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

### Screens

##### User register

![https://github.com/s4kibs4mi/newschain/blob/master/resources/1_profile.png?raw=true]()

##### Write news

![https://github.com/s4kibs4mi/newschain/blob/master/resources/2_write.png?raw=true]()

##### Publish news

![https://github.com/s4kibs4mi/newschain/blob/master/resources/3_write_push.png?raw=true]()

##### View news

![https://github.com/s4kibs4mi/newschain/blob/master/resources/4_view.png?raw=true]()

##### List news

![https://github.com/s4kibs4mi/newschain/blob/master/resources/5_list.png?raw=true]()
