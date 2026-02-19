// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/testinfo"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// Tests basic functionality of the Signature Aggregator API
// Setup step:
// - Sets up a primary network and a subnet.
// - Builds and runs a signature aggregator executable.
// Test Case 1:
// - Sends a teleporter message from the primary network to the subnet.
// - Reads the warp message unsigned bytes from the log
// - Sends the unsigned message to the signature aggregator API
// - Confirms that the signed message is returned and matches the originally sent message
func SignatureAggregatorAPI(
	ctx context.Context,
	log logging.Logger,
	network *network.LocalAvalancheNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	// Begin Setup step
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	signatureAggregatorConfig := utils.CreateDefaultSignatureAggregatorConfig(
		log,
		[]testinfo.L1TestInfo{l1AInfo, l1BInfo},
	)

	signatureAggregatorConfigPath := utils.WriteSignatureAggregatorConfig(
		log,
		signatureAggregatorConfig,
		utils.DefaultSignatureAggregatorCfgFname,
	)
	log.Info("Starting the signature aggregator",
		zap.String("configPath", signatureAggregatorConfigPath),
	)
	signatureAggregatorCancel, readyChan := utils.RunSignatureAggregatorExecutable(
		ctx,
		log,
		signatureAggregatorConfigPath,
		signatureAggregatorConfig,
	)
	defer signatureAggregatorCancel()

	// Wait for signature-aggregator to start up
	log.Info("Waiting for the relayer to start up")
	startupCtx, startupCancel := context.WithTimeout(ctx, 15*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	// End setup step
	// Begin Test Case 1

	log.Info("Sending teleporter message from A -> B")
	receipt, _, _ := utils.SendBasicTeleporterMessage(
		ctx,
		log,
		teleporter,
		l1AInfo,
		l1BInfo,
		fundedKey,
		fundedAddress,
	)
	warpMessage := getWarpMessageFromLog(ctx, log, receipt, l1AInfo)

	reqBody := api.AggregateSignatureRequest{
		Message: "0x" + hex.EncodeToString(warpMessage.Bytes()),
	}

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	requestURL := fmt.Sprintf("http://localhost:%d%s", signatureAggregatorConfig.APIPort, api.APIPath)

	var sendRequestToAPI = func() {
		b, err := json.Marshal(reqBody)
		Expect(err).Should(BeNil())
		bodyReader := bytes.NewReader(b)

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		Expect(err).Should(BeNil())
		req.Header.Set("Content-Type", "application/json")

		res, err := client.Do(req)
		Expect(err).Should(BeNil())
		Expect(res.Status).Should(Equal("200 OK"))
		Expect(res.Header.Get("Content-Type")).Should(Equal("application/json"))

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		Expect(err).Should(BeNil())

		var response api.AggregateSignatureResponse
		err = json.Unmarshal(body, &response)
		Expect(err).Should(BeNil())

		decodedMessage, err := hex.DecodeString(response.SignedMessage)
		Expect(err).Should(BeNil())

		signedMessage, err := avalancheWarp.ParseMessage(decodedMessage)
		Expect(err).Should(BeNil())
		Expect(signedMessage.ID()).Should(Equal(warpMessage.ID()))
	}

	sendRequestToAPI()

	// Try in the other direction
	log.Info("Sending teleporter message from B -> A")
	receipt, _, _ = utils.SendBasicTeleporterMessage(
		ctx,
		log,
		teleporter,
		l1BInfo,
		l1AInfo,
		fundedKey,
		fundedAddress,
	)
	warpMessage = getWarpMessageFromLog(ctx, log, receipt, l1BInfo)

	reqBody = api.AggregateSignatureRequest{
		Message: "0x" + hex.EncodeToString(warpMessage.Bytes()),
	}
	sendRequestToAPI()
}
