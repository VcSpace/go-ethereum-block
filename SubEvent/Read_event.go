package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	atk "goweb3/Contract/ATK"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math/big"
	"strings"
)

//See read_erc20_event.go

func main() {
	client := connect.Wss_proxy_eth()
	contractAddress := config.Get_contractaddr()

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(5335710),
		ToBlock:   big.NewInt(5436266),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(atk.AtkMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		//event := struct {
		//	address []common.Address
		//	amount  []*big.Int
		//}{}
		event, err := contractAbi.Unpack("Transfer", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(event) // foo
	}
}
