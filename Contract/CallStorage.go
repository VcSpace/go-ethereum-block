package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"goweb3/Contract/Storage"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math/big"
)

func main() {
	client := connect.Connect_proxy_eth()
	contractAddress := config.Get_contractaddr()

	instance, err := storage.NewStorage(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance

	num, err := instance.Retrieve(nil)
	fmt.Println(num) //return num

	//input data
	pv := config.Get_pv()
	fromAddress := config.Get_fromaddr()

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pv, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	var amount *big.Int = big.NewInt(5)
	tx, err := instance.Store(auth, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("https://sepolia.etherscan.io/tx/%s \n", tx.Hash().Hex())

	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode))
}
