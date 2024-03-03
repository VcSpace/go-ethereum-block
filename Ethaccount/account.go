package main

import (
	"context"
	"fmt"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math"
	"math/big"
)

func main() {
	client := connect.Connect_proxy_eth()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())

	account := config.Get_fromaddr()
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	blockNumber := big.NewInt(5361150)
	balance2, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance2) // 16454338331596005302

	//bigint
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue) // 16.454338331596005302

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}
