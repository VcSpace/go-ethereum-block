package connect

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var connectUrl string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectUrl = os.Getenv("connectUrl")
}

func Connect_proxy_eth() *ethclient.Client {
	client, err := ethclient.Dial(connectUrl)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func Connect_eth() *ethclient.Client {
	proxyURL, err := url.Parse("http://localhost:8889") // 你的代理地址
	if err != nil {
		log.Fatal(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	//httpClient := &http.Client{
	//	Transport: transport,
	//}

	ctx := context.Background()
	httpClient := rpc.WithHTTPClient(&http.Client{
		Timeout:   15 * time.Second,
		Transport: transport,
	})
	rpcClient, err := rpc.DialOptions(ctx, connectUrl, httpClient)
	//rpcClient, err := rpc.DialHTTPWithClient("https://sepolia.infura.io/v3/88af0cf7642044f0bae9f75123d887c6", httpClient)
	if err != nil {
		log.Fatal(err)
	}

	client := ethclient.NewClient(rpcClient)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected")
	}

	return client
}

func Wss_eth() *ethclient.Client {
	proxyURL, err := url.Parse("http://localhost:8889") // 你的代理地址
	if err != nil {
		log.Fatal(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	ctx := context.Background()
	httpClient := rpc.WithHTTPClient(&http.Client{
		Timeout:   15 * time.Second,
		Transport: transport,
	})
	rpcClient, err := rpc.DialOptions(ctx, connectUrl, httpClient)
	if err != nil {
		log.Fatal(err)
	}

	client := ethclient.NewClient(rpcClient)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected")
	}

	return client
}
