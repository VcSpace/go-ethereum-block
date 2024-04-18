package main

import (
	"context"
	"fmt"
	"github.com/VcSpace/go-ethereum-block/connect"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math"
	"math/big"
)

func main() {
	wallet := "0xa9785f170f872a312a31aa30149e22c38ea151da"
	client := connect.Connect_proxy_eth()

	account := common.HexToAddress(wallet)
	//balance, err := client.BalanceAt(context.Background(), account, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fbalance := new(big.Float)
	//fbalance.SetString(balance.String())
	//ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//
	//fmt.Println(ethValue)

	blockNumber := big.NewInt(19647344)
	endNumber := big.NewInt(19647346)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	var from common.Address
	for {
		block, err = client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal(err)
		}
		for _, tx := range block.Transactions() {
			if count > 2 {
				break
			}

			var err error
			if from, err = types.Sender(types.LatestSignerForChainID(chainID), tx); err == nil {
				//fmt.Println(from.Hex()) //
			}
			if tx.To() != nil && (*tx.To() == account || from == account) {
				count++
				fmt.Printf("Tx: https://etherscan.io/tx/%s\n", tx.Hash().Hex()) // hash
				ethValue := new(big.Float).Quo(new(big.Float).SetInt(tx.Value()), new(big.Float).SetFloat64(math.Pow10(18)))
				fmt.Printf("Transfer_ETH: %f\n", ethValue)
				fmt.Println(tx.Gas())               // 121000
				fmt.Println(tx.GasPrice().Uint64()) // 150000000000
				fmt.Println(tx.Nonce())             // 49

				//https://ethereum.stackexchange.com/questions/149220/type-types-transaction-has-no-field-or-method-asmessage
				if from, err := types.Sender(types.LatestSignerForChainID(chainID), tx); err == nil {
					fmt.Println(from) //
				}

				fmt.Println(tx.To().Hex()) //
				//fmt.Println(tx.Data())
			}
		}

		blockNumber.Add(blockNumber, big.NewInt(1))
		if blockNumber.Cmp(endNumber) > 0 {
			break
		}

	}
}
