package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"runtime"

	receiver "github.com/ava-labs/icm-services/icm-relayer-tests/Receiver"
	sender "github.com/ava-labs/icm-services/icm-relayer-tests/Sender"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func buildAndRunRelayer() {
	BuildAllExecutables(context.TODO())
	RunRelayer(context.TODO(), "/Users/anvii.mishra/Desktop/icm-services/sample-relayer-config.json")
}

func messageExchange(sender_addr string, receiver_addr string) {
	rpcClient, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	if err != nil {
		fmt.Println(err)
	}
	defer rpcClient.Close()
	contractAddress := common.HexToAddress(sender_addr)
	senderContract, error := sender.NewSender(contractAddress, rpcClient)
	privateKey, err := crypto.HexToECDSA("0523fee5412aa3a43468b851fbb970d5d5cc29b672bd1d1a467c8cfa07a5ed4e")
	if err != nil {
		fmt.Println(err)
	}
	chainID := big.NewInt(43113)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println(err)
	}
	destinationAddress := common.HexToAddress(receiver_addr)
	message := "123444"
	var destinationBlockchainID [32]byte
	copy(destinationBlockchainID[:], common.Hex2Bytes("9f3be606497285d0ffbb5ac9ba24aa60346a9b1812479ed66cb329f394a4b1c7"))
	allowedRelayer := common.HexToAddress("12559337e972F8DcE9d739231347d03544fCE91C")
	fmt.Printf("destination address: %s\n", destinationAddress.String())
	tx, err := senderContract.SendMessage(auth, destinationAddress, message, destinationBlockchainID, allowedRelayer)
	if err != nil {
		fmt.Println(error)
	}
	log.Println("Transaction sent:", tx.Hash().Hex())
	log.Println("Transaction id:", tx.Hash().Hex())
	receiveMessage(receiver_addr)
}

func receiveMessage(receiver_addr string) {
	client, err := ethclient.Dial("https://subnets.avax.network/dispatch/testnet/rpc")
	if err != nil {
		fmt.Println(err)
	}
	contractAddress := common.HexToAddress("0x530B7E1c84eE39f4B410FE58d945bB8cDD6E763A")
	receiverCaller, err := receiver.NewReceiverCaller(contractAddress, client)
	if err != nil {
		panic(err)
	}
	msg, err := receiverCaller.LastMessage(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Last received message:", msg)
}

func crossChainRelaying(sender string, relayer string) {
	fmt.Println("Extracting chains...")
	jsonFile, err := os.Open("icm-relayer-tests/sender_receiver_info.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]map[string]map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	sender_address := result["chain-info"]["c-chain"]["sender_contract"]
	receiver_address := result["chain-info"]["dispatch"]["receiver_contract"]
	fmt.Println("Extracted chains.")
	fmt.Println("Calling ABI send...")
	fmt.Println(receiver_address)
	messageExchange(sender_address, receiver_address)
	fmt.Println("Sent.")
}

func sendMessage(sender string, receiver string) {
	fmt.Println("Building relayer...")
	buildAndRunRelayer()
	fmt.Println("Built relayer.")
	fmt.Println("Sending message...")
	crossChainRelaying(sender, receiver)
}

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Sender chain:", args[1])
		fmt.Println("Receiver chain:", args[2])
	} else {
		fmt.Println("Please pass in a sender and receiver chain.")
		return
	}
	sender_chain := args[1]
	receiver_chain := args[2]
	fmt.Println("Beginning process...")
	sendMessage(sender_chain, receiver_chain)
	fmt.Println("Concluded.")
	runtime.Goexit()
}
