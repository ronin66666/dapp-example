package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"mars-mec-go/examples/abigo/mgf"
	_ "mars-mec-go/examples/admin"
	_ "mars-mec-go/examples/gateway"
	_ "mars-mec-go/examples/login"
	_ "mars-mec-go/examples/user"
	"mars-mec-go/log"
	"math/big"
	"strconv"
)

func main2() {
	//初始化数据库链接
	//err := db.InitMysql()
	//if err != nil {
	//	log.Fatal(err.Error())
	//	panic(err)
	//}
	//defer db.Close()
	//node.Start()

	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	if err != nil {
		log.Fatal(err.Error())
	}

	address := common.HexToAddress("0xb4D9C3F5A51bE88Dfc536710AdE66C47C8aEf7e8")
	instance, err := mgf.NewToken(address, client)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKey, err := crypto.HexToECDSA("4d1220d865ef25b3d6ef20e8f957f9263492ae9d2cb05a500fa828c87e2f5c94")
	if err != nil {
		log.Fatal(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err.Error())
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	valueF, err := strconv.ParseFloat("1000", 64)
	valueWei, _ := new(big.Int).SetString(fmt.Sprintf("%.0f", valueF*1000000000000000000), 10)

	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &htaToAddress,
	//	Data: data,
	//})

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	transfer, err := instance.Transfer(auth, common.HexToAddress("0x3a0E2634e7CA96e9Ed58465CF8730E8b7B842A64"), valueWei)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(transfer)

	//balance, err := instance.BalanceOf(nil, common.HexToAddress("0x41338701219B24A94AfDcEB6bD18Fcd43676bD3c"))
	//fmt.Println()
	//fmt.Println(balance)
}
