package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get_pv() *ecdsa.PrivateKey {
	privateKey := os.Getenv("PRIVATE_KEY")
	pv, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return pv
}

func Get_fromaddr() common.Address {
	fromaddr := os.Getenv("fromAddress")
	return common.HexToAddress(fromaddr)
}

func Get_toaddr() common.Address {
	toaddr := os.Getenv("toAddress")
	return common.HexToAddress(toaddr)
}

func Get_contractaddr() common.Address {
	tokenaddr := os.Getenv("contractAddress")
	return common.HexToAddress(tokenaddr)

}
