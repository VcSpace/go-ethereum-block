package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"goweb3/Contract/IERC"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math/big"
)

func main() {
	client := connect.Connect_proxy_eth()
	contractAddress := config.Get_contractaddr()

	fromAddress := config.Get_fromaddr()
	toAddress := config.Get_toaddr()
	instance, err := IERC.NewIERC(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance

	bal, err := instance.BalanceOf(&bind.CallOpts{}, fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wei: %s\n", bal)

	pv := config.Get_pv()

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

	var amount *big.Int = big.NewInt(500000000000000000)
	tx, err := instance.Transfer(auth, toAddress, amount)
	//tx, err := instance.Claim(auth, toAddress, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("https://sepolia.etherscan.io/tx/%s \n", tx.Hash().Hex())
}
