package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math/big"
)

func main() {
	client := connect.Connect_proxy_eth()
	pv := config.Get_pv()
	from := config.Get_fromaddr()
	toaddr := config.Get_toaddr()
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)
	gasFeeCap, err := client.SuggestGasPrice(context.Background())
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	var data []byte

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        &toaddr,
		Value:     value,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), pv)
	err = client.SendTransaction(context.Background(), signedTx)
	fmt.Printf("https://sepolia.etherscan.io/tx/%s \n", signedTx.Hash().Hex())
}
