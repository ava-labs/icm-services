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
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/staking"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	testUtils "github.com/ava-labs/icm-services/tests/utils"
	. "github.com/onsi/gomega"
)

// Tests signature aggregation with a private network
// Steps:
// - Sets up a primary network and a subnet.
// - Generates a config with temporary paths set for TLS cert and key
// - Starts the signature aggregator with the generated config once to populate the TLS cert and key
// - Reads the nodeID from the TLS cert and stops the signature aggregator
// - Sends the teleporter message from B -> A
// - Restarts the subnet B nodes with the validatorOnly flag set to true and nodeID added to allowedNodes
// - Restarts the signature aggregator with the same config which should re-use
// now populated TLS cert and key and result in same nodeID
// - Requests an aggregated signature from the signature aggregator API which
// will only be returned successfully if the nodeID is explicitly allowed by the subnet
func ValidatorsOnlyNetwork(network *network.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	// Begin Setup step
	ctx := context.Background()

	l1AInfo := network.GetPrimaryNetworkInfo()
	_, l1BInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	// Start the signature-aggregator for the first time to generate the
	// TLS cert key pair
	dir, err := os.MkdirTemp(os.TempDir(), "sig-agg-tls-cert")
	Expect(err).Should(BeNil())

	// Create a config without TLS cert and key
	baseConfig := testUtils.CreateDefaultSignatureAggregatorConfig(
		[]interfaces.L1TestInfo{l1AInfo, l1BInfo},
	)
	baseConfigPath := testUtils.WriteSignatureAggregatorConfig(
		baseConfig,
		testUtils.DefaultSignatureAggregatorCfgFname,
	)

	keyPath := dir + "/key.pem"
	certPath := dir + "/cert.pem"

	signatureAggregatorConfig := baseConfig
	signatureAggregatorConfig.TLSCertPath = certPath
	signatureAggregatorConfig.TLSKeyPath = keyPath

	signatureAggregatorConfigPath := testUtils.WriteSignatureAggregatorConfig(
		signatureAggregatorConfig,
		testUtils.DefaultSignatureAggregatorCfgFname,
	)
	log.Println("Starting the signature aggregator", "configPath", signatureAggregatorConfigPath)
	signatureAggregatorCancel, readyChan := testUtils.RunSignatureAggregatorExecutable(
		ctx,
		signatureAggregatorConfigPath,
		signatureAggregatorConfig,
	)

	// Wait for signature-aggregator to start up
	log.Println("Waiting for the signature-aggregator to start up")
	startupCtx, startupCancel := context.WithTimeout(ctx, 15*time.Second)
	defer startupCancel()
	testUtils.WaitForChannelClose(startupCtx, readyChan)

	cert, err := staking.LoadTLSCertFromFiles(keyPath, certPath)
	Expect(err).Should(BeNil())
	peerCert, err := staking.ParseCertificate(cert.Leaf.Raw)
	Expect(err).Should(BeNil())
	nodeID := ids.NodeIDFromCert(peerCert)

	signatureAggregatorCancel()
	log.Println("Retrieved nodeID", "nodeID", nodeID)

	// We have to send the message before making the network private.

	log.Println("Sending teleporter message from B -> A")
	receipt, _, _ := testUtils.SendBasicTeleporterMessage(
		ctx,
		teleporter,
		l1BInfo,
		l1AInfo,
		fundedKey,
		fundedAddress,
	)
	warpMessage := getWarpMessageFromLog(ctx, receipt, l1BInfo)

	// Restart l1B and make it private
	relayerNodeIDSet := set.NewSet[ids.NodeID](1)
	relayerNodeIDSet.Add(nodeID)

	l1BNodes := set.NewSet[ids.NodeID](1)

	// Make l1BInfo a validator only network
	for _, subnet := range network.Subnets {
		if subnet.SubnetID == l1BInfo.SubnetID {
			subnet.Config = make(map[string]interface{})
			subnet.Config["validatorOnly"] = true
			subnet.Config["allowedNodes"] = relayerNodeIDSet
			err := subnet.Write(network.GetSubnetDir(), network.GetChainConfigDir())
			Expect(err).Should(BeNil())
			l1BNodes.Add(subnet.ValidatorIDs...)
		}
	}
	// Restart l1B nodes
	for _, tmpnetNode := range network.Nodes {
		if l1BNodes.Contains(tmpnetNode.NodeID) {
			tmpnetNode.RuntimeConfig.ReuseDynamicPorts = true
			// Restart the network to apply the new chain configs
			ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
			defer cancel()
			err := network.RestartNode(ctx, logging.NoLog{}, tmpnetNode)
			Expect(err).Should(BeNil())
		}
	}

	// End setup step

	requestURL := fmt.Sprintf("http://localhost:%d%s", signatureAggregatorConfig.APIPort, api.APIPath)

	reqBody := api.AggregateSignatureRequest{
		Message: "0x" + hex.EncodeToString(warpMessage.Bytes()),
	}
	client := http.Client{
		Timeout: 20 * time.Second,
	}

	var sendRequestToAPI = func(success bool) {
		b, err := json.Marshal(reqBody)
		Expect(err).Should(BeNil())
		bodyReader := bytes.NewReader(b)

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		Expect(err).Should(BeNil())
		req.Header.Set("Content-Type", "application/json")

		res, err := client.Do(req)
		Expect(err).Should(BeNil())

		if !success {
			Expect(res.Status).ShouldNot(Equal("200 OK"))
			return
		}

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

	// start sig-agg again with a floating TLS cert - this should fail
	log.Println("Starting the signature aggregator with a floating TLS cert")
	signatureAggregatorCancel, readyChan = testUtils.RunSignatureAggregatorExecutable(
		ctx,
		baseConfigPath,
		baseConfig,
	)

	// Wait for signature-aggregator to start up
	log.Println("Waiting for the signature-aggregator to start up")
	startupCtx, startupCancel = context.WithTimeout(ctx, 15*time.Second)
	defer startupCancel()
	testUtils.WaitForChannelClose(startupCtx, readyChan)

	sendRequestToAPI(false)
	signatureAggregatorCancel()

	// start sig-agg again with the same TLS cert
	log.Println("Starting the signature aggregator with the same TLS cert")
	signatureAggregatorCancel, readyChan = testUtils.RunSignatureAggregatorExecutable(
		ctx,
		signatureAggregatorConfigPath,
		signatureAggregatorConfig,
	)
	defer signatureAggregatorCancel()

	// Wait for signature-aggregator to start up
	log.Println("Waiting for the signature-aggregator to start up")
	startupCtx, startupCancel = context.WithTimeout(ctx, 15*time.Second)
	defer startupCancel()
	testUtils.WaitForChannelClose(startupCtx, readyChan)

	sendRequestToAPI(true)
}
