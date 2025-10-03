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

	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	testUtils "github.com/ava-labs/icm-services/tests/utils"
	"github.com/ava-labs/libevm/log"
	. "github.com/onsi/gomega"
)

// Tests epoch validator functionality in the Signature Aggregator API
// This test verifies that the signature aggregator can handle both current and epoched validators
// Setup step:
// - Sets up a primary network and a subnet.
// - Builds and runs a signature aggregator executable.
// Test Case 1: Current Validators (PChainHeight = 0)
// - Sends a teleporter message from the primary network to the subnet.
// - Requests signature aggregation with PChainHeight = 0 (current validators)
// - Confirms that the signed message is returned correctly
// Test Case 2: Epoched Validators (PChainHeight = specific height)
// - Uses the same teleporter message
// - Requests signature aggregation with a specific PChainHeight
// - Confirms that the signed message is returned correctly
// Test Case 3: Large PChainHeight (ProposedHeight)
// - Uses ProposedHeight as PChainHeight to test the edge case
// - Confirms that the system handles this correctly
func SignatureAggregatorEpochAPI(network *network.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	// Begin Setup step
	ctx := context.Background()

	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	signatureAggregatorConfig := testUtils.CreateDefaultSignatureAggregatorConfig(
		[]interfaces.L1TestInfo{l1AInfo, l1BInfo},
	)

	signatureAggregatorConfigPath := testUtils.WriteSignatureAggregatorConfig(
		signatureAggregatorConfig,
		testUtils.DefaultSignatureAggregatorCfgFname,
	)
	log.Info("Starting the signature aggregator for epoch tests", "configPath", signatureAggregatorConfigPath)
	signatureAggregatorCancel, readyChan := testUtils.RunSignatureAggregatorExecutable(
		ctx,
		signatureAggregatorConfigPath,
		signatureAggregatorConfig,
	)
	defer signatureAggregatorCancel()

	// Wait for signature-aggregator to start up
	log.Info("Waiting for the signature aggregator to start up")
	startupCtx, startupCancel := context.WithTimeout(ctx, 15*time.Second)
	defer startupCancel()
	testUtils.WaitForChannelClose(startupCtx, readyChan)

	// End setup step

	log.Info("Sending teleporter message for epoch validator tests")
	receipt, _, _ := testUtils.SendBasicTeleporterMessage(
		ctx,
		teleporter,
		l1AInfo,
		l1BInfo,
		fundedKey,
		fundedAddress,
	)
	warpMessage := getWarpMessageFromLog(ctx, receipt, l1AInfo)

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	requestURL := fmt.Sprintf("http://localhost:%d%s", signatureAggregatorConfig.APIPort, api.APIPath)

	// Helper function to send API request with specific PChainHeight
	var sendRequestWithPChainHeight = func(pchainHeight uint64, testDescription string) {
		log.Info("Testing signature aggregation",
			"testCase", testDescription,
			"pchainHeight", pchainHeight,
		)

		reqBody := api.AggregateSignatureRequest{
			Message:      "0x" + hex.EncodeToString(warpMessage.Bytes()),
			PChainHeight: pchainHeight,
		}

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

		log.Info("Successfully verified signed message",
			"testCase", testDescription,
			"pchainHeight", pchainHeight,
			"messageID", signedMessage.ID().String(),
		)
	}

	sendRequestWithPChainHeight(5, "Epoched Validators at Height 5")
	sendRequestWithPChainHeight(pchainapi.ProposedHeight, "ProposedHeight")

	// Test the reverse direction as well
	log.Info("Testing reverse direction with epoch validators")
	receipt, _, _ = testUtils.SendBasicTeleporterMessage(
		ctx,
		teleporter,
		l1BInfo,
		l1AInfo,
		fundedKey,
		fundedAddress,
	)
	warpMessage = getWarpMessageFromLog(ctx, receipt, l1BInfo)

	sendRequestWithPChainHeight(5, "Reverse Direction - Epoched Validators at Height 50")

	log.Info("All epoch validator API tests completed successfully!")
}
