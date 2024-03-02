package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math/big"
)

func main() {
	client := connect.Connect_proxy_eth()
	from := config.Get_fromaddr()
	pv := config.Get_pv()
	toaddr := config.Get_toaddr()
	Contractaddr := config.Get_cotractaddr()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)

	var data []byte

	paddedAddress := common.LeftPadBytes(toaddr.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	transferFnSignature := []byte("transfer(address,uint256)")
	methodID := crypto.Keccak256(transferFnSignature)[:4]

	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: from,
		To:   &Contractaddr,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gasLimit)

	gasFeeCap, err := client.SuggestGasPrice(context.Background())
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        &Contractaddr,
		Value:     value,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), pv)
	err = client.SendTransaction(context.Background(), signedTx)
	fmt.Printf("https://sepolia.etherscan.io/tx/%s", signedTx.Hash().Hex())
}
