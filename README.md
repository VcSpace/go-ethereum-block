# Go-Ethereum-Web3

Ethereum Golang API, JSON-RPC client, Smart contract transaction API.

---

```
├── Contract
│   ├── CallERC20.go -- ERC20Client
│   ├── CallStorage.go -- StorageClient
│   ├── Deploy.go
│   ├── ERC20
│   │   ├── ERC20.abi
│   │   ├── ERC20.sol
│   │   └── EToken.go
│   ├── IERC
│   │   ├── IERC.go
│   │   ├── IERC20.abi
│   │   └── IERC20.sol
│   └── Storage
│       ├── Storage.abi
│       ├── Storage.go
│       └── Storage.sol
├── LICENSE -- GPL3.0
├── README.md
├── SubEvent
│   ├── Read_erc20_event.go
│   ├── Read_event.go
│   └── Sub_event.go
├── Transaction
│   ├── blocks.go -- Get block
│   ├── raw_transaction.go -- Create RawTransaction and SendTransaction
│   ├── send_erc20.go -- Contract transfer
│   ├── send_ether.go -- Transfer ETH
│   ├── sub_block.go -- Subscribe block
│   └── transactions.go -- Output tx data
├── config
│   └── config.go -- File .env config
├── connect
│   └── connect.go -- Connect RPC/WSS
├── Ethaccount
│   ├── account.go -- Account
│   └── checkaddr.go -- Check address is contract
└── go.mod

```

---

## Deploy contract

```
solc --abi --bin -o ./ Storage.sol
abigen --bin=Storage.bin --abi=Storage.abi --pkg=storage --out=Storage.go
```

## interface

```
solc --abi -o ./ IERC20.sol
abigen --abi=IERC20.abi --pkg=IERC --out=IERC.go
```

---

## Reference:

https://mhxw.life/eth-dev-with-go/docs/en/

https://www.yuhenm.com/archives/529.html

https://mirror.xyz/rbtree.eth/B2OZSszjxD3BfI07WOuAFzzrACilxvZcgb09GYdMgng

---

## Infrastructure:

https://chainide.com/

https://www.infura.io/