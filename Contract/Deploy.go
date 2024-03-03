package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"goweb3/Contract/Sol"
	"goweb3/config"
	"goweb3/connect"
	"log"
	"math/big"
)

func main() {
	client := connect.Connect_proxy_eth()
	pv := config.Get_pv()
	fromAddress := config.Get_fromaddr()

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasFeeCap, err := client.SuggestGasPrice(context.Background())
	//gasTipCap, err := client.SuggestGasTipCap(context.Background())

	auth, _ := bind.NewKeyedTransactorWithChainID(pv, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasFeeCap

	address, tx, instance, err := storage.DeployStorage(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}
