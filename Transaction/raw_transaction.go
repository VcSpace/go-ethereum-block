package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math/big"
)

func main() {
	client := connect.Wss_eth()

	pv := config.Get_pv()
	fromAddress := config.Get_fromaddr()
	toaddr := config.Get_toaddr()

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	value := big.NewInt(1000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)             // in units
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

	ts := types.Transactions{signedTx}
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0]) //<-- ts.GetRlp undefined
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Printf("%s\n", rawTxHex) // f86...772

	send_rawTxBytes, err := hex.DecodeString(rawTxHex)
	sendtx := new(types.Transaction)
	rlp.DecodeBytes(send_rawTxBytes, &sendtx)
	err = client.SendTransaction(context.Background(), sendtx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("https://sepolia.etherscan.io/tx/%s \n", sendtx.Hash().Hex())
}
