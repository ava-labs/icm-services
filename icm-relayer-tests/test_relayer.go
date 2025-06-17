package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"

	sender "github.com/ava-labs/icm-services/icm-relayer-tests/Sender"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SendingConfiguration struct {
	rpcClient               ethclient.Client
	destinationContractAddr common.Address
	auth                    *bind.TransactOpts
	destinationBlockchainID [32]byte
	allowedRelayer          common.Address
	senderContract          *sender.Sender
}

func configureSenderAndSendMessage(sender_addr string, receiver_addr string, ch chan string, tps int, load int, sender_rpc string) {
	sendConfig := new(SendingConfiguration)
	sendConfig.rpcClient, _ = ethclient.Dial(sender_rpc)
	defer sendConfig.rpcClient.Close()
	contractAddress := common.HexToAddress(sender_addr)
	sendConfig.senderContract, _ = sender.NewSender(contractAddress, sendConfig.rpcClient)
	privateKey, err := crypto.HexToECDSA("0523fee5412aa3a43468b851fbb970d5d5cc29b672bd1d1a467c8cfa07a5ed4e")
	if err != nil {
		fmt.Println(err)
	}
	sendConfig.auth, _ = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(43113))
	sendConfig.destinationContractAddr = common.HexToAddress(receiver_addr)
	copy(sendConfig.destinationBlockchainID[:], common.Hex2Bytes("9f3be606497285d0ffbb5ac9ba24aa60346a9b1812479ed66cb329f394a4b1c7"))
	sendConfig.allowedRelayer = common.HexToAddress("12559337e972F8DcE9d739231347d03544fCE91C")
	sendMessagedAtLoadAndTPS(load, tps, ch, sendConfig)
}

func sendMessagedAtLoadAndTPS(load int, tps int, ch chan string, sendConfig *SendingConfiguration) {
	duration := time.Second / time.Duration(tps)
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	count := 0
	maxCount := load
	for {
		select {
		case <-ticker.C:
			tx, err := sendConfig.senderContract.SendMessage(
				sendConfig.auth,
				sendConfig.destinationContractAddr,
				"Hello",
				sendConfig.destinationBlockchainID,
				sendConfig.allowedRelayer,
			)
			if err != nil {
				fmt.Println("oh no")
				continue
			}
			success, receipt := waitForTransactionSuccess(context.TODO(), sendConfig.rpcClient, tx.Hash())
			if !success {
				logSendMessageError(receipt)
				continue
			}
			ch <- tx.Hash().Hex()
			count++
			if count >= maxCount {
				close(ch)
				return
			}
		}
	}
}

func logSendMessageError(receipt *types.Receipt) {
	file, err := os.OpenFile("errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
		return
	}
	defer file.Close()
	if _, err := file.WriteString("Errors:\n"); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func waitForTransactionSuccess(
	ctx context.Context,
	client ethclient.Client,
	txHash common.Hash,
) (bool, *types.Receipt) {
	cctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	receipt, err := WaitMined(cctx, client, txHash)
	if err != nil {
		return false, receipt
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return false, receipt
	}
	return true, receipt
}

func WaitMined(ctx context.Context, client ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func readConfigAndSendMessage(sender string, relayer string, ch chan string, tps int, load int) {
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
	sender_rpc := "https://api.avax-test.network/ext/bc/C/rpc"
	configureSenderAndSendMessage(sender_address, receiver_address, ch, tps, load, sender_rpc)
}

func listenToAndCollectLogs(loadToLoggerCh chan string, loggerToRelayer chan string) {
	file, err := os.OpenFile("received.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
		return
	}
	defer file.Close()
	if _, err := file.WriteString("Begining Receiving:\n"); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	directLogsToFile(loadToLoggerCh, file)
}

func buildRelayerWithConfig(file string, loggerToRelayerCh chan string) {
	BuildAllExecutables(context.TODO())
	RunRelayer(context.TODO(), file)
}

func directLogsToFile(loadToLoggerCh chan string, file *os.File) {
	for {
		select {
		case msg, ok := <-loadToLoggerCh:
			if !ok {
				fmt.Println("Channel closed. Done receiving.")
				return
			}
			if _, err := file.WriteString(msg + "+\n"); err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		case <-time.After(15 * time.Second):
			fmt.Println("Operation timed out")
			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) > 2 && len(args) < 5 {
		fmt.Println("Sender chain:", args[1])
		fmt.Println("Receiver chain:", args[2])
	} else if len(args) == 5 {
		fmt.Println("Load:", args[3])
		fmt.Println("TPS:", args[4])
	} else {
		fmt.Println("Please pass in a sender, receiver chain.")
		return
	}
	loadToLoggerCh := make(chan string)
	loggerToRelayerCh := make(chan string)
	sender_chain := args[1]
	receiver_chain := args[2]
	load, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		fmt.Println("Please pass in a load.")
		return
	}
	tps, err := strconv.ParseInt(args[4], 10, 64)
	if err != nil {
		fmt.Println("Please pass in a tps.")
		return
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		readConfigAndSendMessage(sender_chain, receiver_chain, loadToLoggerCh, int(tps), int(load))
	}()
	go func() {
		buildRelayerWithConfig("/Users/anvii.mishra/Desktop/icm-services/sample-relayer-config.json", loggerToRelayerCh)
		defer wg.Done()
	}()
	go func() {
		defer wg.Done()
		listenToAndCollectLogs(loadToLoggerCh, loggerToRelayerCh)
	}()
	wg.Wait()
}
