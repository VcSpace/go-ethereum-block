package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"goweb3/connect"
	"log"
)

func main() {
	client := connect.Wss_eth()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) //

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        //
			fmt.Println(block.Number().Uint64())   //
			fmt.Println(block.Time())              //
			fmt.Println(block.Nonce())             //
			fmt.Println(len(block.Transactions())) //
		}
	}
}
