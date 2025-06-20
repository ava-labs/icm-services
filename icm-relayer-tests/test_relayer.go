package main

import (
	"context"
	"crypto/ecdsa"
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
	opts                    *bind.TransactOpts
	destinationBlockchainID [32]byte
	allowedRelayer          common.Address
	senderContract          *sender.Sender
}

func configureWallets(sender_addr string, receiver_addr string, sender_rpc string, priv_key *ecdsa.PrivateKey, relayer_allowed common.Address) *SendingConfiguration {
	sendConfig := new(SendingConfiguration)
	sendConfig.rpcClient, _ = ethclient.Dial(sender_rpc)
	defer sendConfig.rpcClient.Close()
	contractAddress := common.HexToAddress(sender_addr)
	sendConfig.senderContract, _ = sender.NewSender(contractAddress, sendConfig.rpcClient)
	privateKey, err := crypto.HexToECDSA("priv_key")
	if err != nil {
		fmt.Println(err)
	}

	sendConfig.opts, _ = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(43113))

	sendConfig.destinationContractAddr = common.HexToAddress(receiver_addr)
	copy(sendConfig.destinationBlockchainID[:], common.Hex2Bytes("9f3be606497285d0ffbb5ac9ba24aa60346a9b1812479ed66cb329f394a4b1c7"))
	sendConfig.allowedRelayer = common.HexToAddress("wallet_addr")
	return sendConfig
}
func configureSenderAndSendMessage(sender_addr string, receiver_addr string, ch chan string, tps int, load int, sender_rpc string) {
	sendCofig := configureWallets(sender_addr, receiver_addr, sender_rpc, pr)
	sendMessagedAtLoadAndTPS(load, tps, ch, sendConfig)
}

func sendIndividualMessage(sendConfig *SendingConfiguration, ch chan string) bool {
	tx, err := sendConfig.senderContract.SendMessage(
		sendConfig.opts,
		sendConfig.destinationContractAddr,
		"Hello",
		sendConfig.destinationBlockchainID,
		sendConfig.allowedRelayer,
	)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	success, receipt, error := waitForTransactionSuccess(context.TODO(), sendConfig.rpcClient, tx.Hash())
	if !success {
		if receipt != nil {
			ch <- fmt.Sprintf("Receipt for tx %s : %s", tx.Hash(), receipt.TxHash.String())
		}
		if error != nil {
			ch <- fmt.Sprintf("Error for tx %s : %s", tx.Hash(), error.Error())
		} else {
			ch <- "Failure"
		}
		return false
	} else {
		fmt.Println("Sent at" + time.Now().String() + " for " + tx.Hash().Hex() + "\n")
	}
	ch <- tx.Hash().Hex()
	return true
}

func sendMessagedAtLoadAndTPS(load int, tps int, ch chan string, sendConfig *SendingConfiguration) {
	var wg sync.WaitGroup
	duration := time.Second / time.Duration(tps)
	ticker := time.NewTicker(duration)
	fmt.Println("Ticker duration:", duration)
	defer ticker.Stop()

	count := 0
	maxCount := load

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t, "for count", strconv.Itoa(count))
			wg.Add(1)
			go func() {
				defer wg.Done()
				sendIndividualMessage(sendConfig, ch)
			}()
			count++
			if count >= maxCount {
				wg.Wait()
				close(ch)
				return
			}
		}
	}
}

// func logSendMessageError(receipt *types.Receipt, issue error) {
// 	file, err := os.OpenFile("errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("Error opening/creating file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	if _, err := file.WriteString("Errors:\n"); err != nil {
// 		fmt.Println("Error writing to file:", err)
// 		return
// 	}
// 	if receipt == nil {
// 		file.WriteString(issue.Error())
// 	} else if issue == nil {
// 		file.WriteString(receipt.TxHash.String())

// 	}
// 	file.WriteString(receipt.TxHash.String() + ":" + issue.Error())
// }

func waitForTransactionSuccess(
	ctx context.Context,
	client ethclient.Client,
	txHash common.Hash,
) (bool, *types.Receipt, error) {
	cctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	receipt, err := WaitMined(cctx, client, txHash)
	if err != nil {
		return false, receipt, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return false, receipt, err
	}
	return true, receipt, err
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

func readConfigAndSendMessage(sender string, relayer string, ch chan string, tps int, load int, walletInfo string) {
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

// func buildRelayerWithConfig(file string, loggerToRelayerCh chan string) {
// 	BuildAllExecutables(context.TODO())
// 	RunRelayer(context.TODO(), file)
// }

func directLogsToFile(loadToLoggerCh chan string, file *os.File) {
	for {
		select {
		case msg, ok := <-loadToLoggerCh:
			if !ok {
				fmt.Println("Channel closed. Done receiving.")
				return
			}
			fmt.Println("Receive at: ", time.Now().String(), " for ", msg+"\n")
			if _, err := file.WriteString(msg + "+\n"); err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		case <-time.After(10000 * time.Second):
			fmt.Println("Operation timed out")
			return
		}
	}
}

func main() {
	args := os.Args
	loadToLoggerCh := make(chan string)
	loggerToRelayerCh := make(chan string)
	senderChain := args[1]
	receiverChain := args[2]
	load, err := strconv.ParseInt(args[3], 10, 64)
	load = load / 5
	if err != nil {
		fmt.Println("Please pass in a load.")
		return
	}
	tps, err := strconv.ParseInt(args[4], 10, 64)
	if err != nil {
		fmt.Println("Please pass in a tps.")
		return
	}
	walletInfo := args[5]
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		readConfigAndSendMessage(senderChain, receiverChain, loadToLoggerCh, int(tps), int(load), walletInfo)
	}()
	go func() {
		sum := 9 + 6
		fmt.Println(sum)
		// buildRelayerWithConfig("/Users/anvii.mishra/Desktop/icm-services/sample-relayer-config.json", loggerToRelayerCh)
		defer wg.Done()
	}()
	go func() {
		defer wg.Done()
		listenToAndCollectLogs(loadToLoggerCh, loggerToRelayerCh)
	}()
	wg.Wait()
}
