package main

import (
	"context"
	"fmt"
	"goweb3/connect"
	"log"
)

func main() {
	client := connect.Connect_proxy_eth()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744
}
