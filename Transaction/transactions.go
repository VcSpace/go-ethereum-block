package main

import (
	"context"
	"fmt"
	"github.com/VcSpace/go-ethereum-block/connect"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math"
	"math/big"
)

func main() {
	client := connect.Connect_proxy_eth()

	blockNumber := big.NewInt(19644570)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, tx := range block.Transactions() {
		if count > 2 {
			break
		}
		count++
		fmt.Printf("Tx: https://sepolia.etherscan.io/tx/%s\n", tx.Hash().Hex()) // hash
		ethValue := new(big.Float).Quo(new(big.Float).SetInt(tx.Value()), new(big.Float).SetFloat64(math.Pow10(18)))
		fmt.Printf("Transfer_ETH: %f\n", ethValue)
		fmt.Println(tx.Gas())               // 121000
		fmt.Println(tx.GasPrice().Uint64()) // 150000000000
		fmt.Println(tx.Nonce())             // 49

		//https://ethereum.stackexchange.com/questions/149220/type-types-transaction-has-no-field-or-method-asmessage
		if from, err := types.Sender(types.LatestSignerForChainID(chainID), tx); err == nil {
			fmt.Println(from.Hex()) //
		}

		fmt.Println(tx.To().Hex()) //
		fmt.Println(tx.Data())

	}

}
