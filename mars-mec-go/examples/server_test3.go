package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	//"mars-mec-go/examples/abigo/mecMaster"
	"math/big"
	//"strings"
)

func main() {

	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x62FA42a2b932D74f30Dec2248d0A4ccb1c690C88")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(18748968),
		ToBlock:   big.NewInt(18772001),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("数据总长度：", len(logs))

	logTransferSig := []byte("Deposited(address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	//contractAbi, err := abi.JSON(strings.NewReader(mecMaster.MecMasterMetaData.ABI))
	//contractAbi, err := abi.JSON(strings.NewReader(mecMaster.MecMasterMetaData.ABI))

	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		fmt.Println(logTransferSigHash.Hex())

		fmt.Printf("Log Name: Transfer\n")
		fmt.Println(common.HexToAddress(vLog.Topics[1].String()))

		fmt.Println(vLog.Topics[2].Big())

		//var a mecMaster.MecMasterDeposited

		//err := contractAbi.UnpackIntoInterface(&a, "Deposited", vLog.Data)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println(a)
		//transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
		//transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
		//
		//fmt.Printf("From: %s\n", transferEvent.From.Hex())
		//fmt.Printf("To: %s\n", transferEvent.To.Hex())
		//fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())
		//}
	}

}
