# Go-Ethereum-Web3

Ethereum Golang API, JSON-RPC client, Smart contract transaction API.

---

```
├── LICENSE
├── README.md
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
├── ethaccount
│   ├── account.go -- Account
│   └── checkaddr.go -- Check address is contract
└── go.mod

```