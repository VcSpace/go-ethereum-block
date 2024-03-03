package main

import (
	"context"
	"fmt"
	"goweb3/config"
	"goweb3/connect"
	"log"
)

func main() {
	// smart contract address
	client := connect.Connect_proxy_eth()
	address := config.Get_contractaddr()
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

}
