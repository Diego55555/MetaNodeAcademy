package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 文件，使用系统环境变量")
	}

	rpcURL := os.Getenv("SEPOLIA_HTTPS_RPC_URL")

	ctx, cancel := context.WithCancel(context.Background()) // 上下文对象，用于控制 client 的生命周期
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL) //以太坊客户端连接
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}
	log.Printf("chain id: %v\n", chainID)

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	exchangeTokenContract, err := NewExchangeToken(contractAddress, client)
	if err != nil {
		log.Fatalf("failed to bind ExchangeToken contract: %v", err)
	}

	privateKeyHex := os.Getenv("SEPOLIA_PRIVATE_KEY1")
	privateKeyECDSA, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		log.Fatalf("failed to create transactor: %v", err)
	}
	opt.Value = big.NewInt(1e15)

	//使用0.001ETH兑换1ETK代币
	_, err = exchangeTokenContract.ExchangeETHToETK(opt)
	if err != nil {
		log.Fatalf("failed to call exchangeETHToETK: %v", err)
	}
}
