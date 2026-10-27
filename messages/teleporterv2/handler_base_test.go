// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/mocks"
	"github.com/ava-labs/icm-services/relayer/config"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	mock_vms "github.com/ava-labs/icm-services/vms/mocks"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	testBlockGasLimit = uint64(15_000_000)
	testNetworkID     = uint32(1)
)

var (
	testAdapterAddress    = common.HexToAddress("0xd81545385803bCD83bd59f58Ba2d2c0562387F83")
	testTeleporterAddress = common.HexToAddress("0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf")
	testRegistryAddress   = common.HexToAddress("0x1234567890AbcdEF1234567890aBcdef12345678")
	testRewardAddress     = common.HexToAddress("0x27aE10273D17Cd7e80de8580A51f476960626e5f")
	testRelayerAddress    = common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")

	testSourceBlockchainID      = ids.ID{1}
	testDestinationBlockchainID = ids.ID{2}
)

// testTeleporterMessage returns a message the test relayer is allowed to deliver to
// testTeleporterAddress on testDestinationBlockchainID.
func testTeleporterMessage() *teleportermessengerv2.TeleporterMessageV2 {
	return &teleportermessengerv2.TeleporterMessageV2{
		MessageNonce:            big.NewInt(1),
		OriginSenderAddress:     common.HexToAddress("0x1111111111111111111111111111111111111111"),
		OriginTeleporterAddress: testTeleporterAddress,
		DestinationBlockchainID: testDestinationBlockchainID,
		DestinationAddress:      common.HexToAddress("0x2222222222222222222222222222222222222222"),
		RequiredGasLimit:        big.NewInt(100_000),
		AllowedRelayerAddresses: []common.Address{testRelayerAddress},
		Receipts: []teleportermessengerv2.TeleporterMessageReceipt{
			{
				ReceivedMessageNonce: big.NewInt(1),
				RelayerRewardAddress: testRewardAddress,
			},
		},
		Message: []byte{1, 2, 3, 4},
	}
}

// rpcError mirrors the error libevm's RPC client returns for a JSON-RPC error object: the message
// plus whatever the node put in the "data" field, which is the revert output for a reverted call.
type rpcError struct {
	msg  string
	data any
}

func (e *rpcError) Error() string {
	return e.msg
}

func (e *rpcError) ErrorData() any {
	return e.data
}

// revertWithReason is the error a node returns for a call that reverted with Error(string).
func revertWithReason(t *testing.T, reason string) error {
	t.Helper()
	stringType, err := abi.NewType("string", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: stringType}}.Pack(reason)
	require.NoError(t, err)
	data := append(crypto.Keccak256([]byte("Error(string)"))[:4], packed...)
	return &rpcError{msg: "execution reverted: " + reason, data: hexutil.Encode(data)}
}

// panicRevert is the error a node returns for a call that reverted with Panic(code), as checked
// arithmetic does on overflow (code 0x11).
func panicRevert(t *testing.T, code uint64) error {
	t.Helper()
	uintType, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: uintType}}.Pack(new(big.Int).SetUint64(code))
	require.NoError(t, err)
	data := append(crypto.Keccak256([]byte("Panic(uint256)"))[:4], packed...)
	return &rpcError{msg: "execution reverted", data: hexutil.Encode(data)}
}

// handlerMocks are the mocked destination clients a test handler talks to, with what tests need
// to recognize and answer its calls.
type handlerMocks struct {
	destinationClient *mock_vms.MockDestinationClient
	ethClient         *mock_evm.MockClient
	// messageReceivedInput is the calldata of the messageReceived query for the handler's message.
	messageReceivedInput []byte
	// receiveSelector is the receiveCrossChainMessage function selector, which every delivery
	// (simulated or sent) starts with.
	receiveSelector []byte
}

// newTestHandler builds the handler for testTeleporterMessage through the factory, for the given
// verifier type, over a destination client whose only sender is testRelayerAddress.
func newTestHandler(
	t *testing.T,
	ctrl *gomock.Controller,
	verifierType string,
) (messages.MessageHandler, *handlerMocks) {
	t.Helper()
	teleporterMessage := testTeleporterMessage()

	// The two adapters encode the message differently on the source chain.
	var payload []byte
	if verifierType == VerifierTypeWarp {
		var err error
		payload, err = teleporterMessage.Pack()
		require.NoError(t, err)
	} else {
		payload = SerializeTeleporterMessageV2(teleporterMessage)
	}
	addressedCall, err := warpPayload.NewAddressedCall(testAdapterAddress.Bytes(), payload)
	require.NoError(t, err)
	unsignedMessage, err := warp.NewUnsignedMessage(testNetworkID, testSourceBlockchainID, addressedCall.Bytes())
	require.NoError(t, err)

	settings := map[string]any{
		"reward-address":     testRewardAddress.Hex(),
		"verifier-type":      verifierType,
		"teleporter-address": testTeleporterAddress.Hex(),
	}
	if verifierType == VerifierTypeMerkle {
		settings["registry-address"] = testRegistryAddress.Hex()
	}
	factory, err := NewMessageHandlerFactory(
		testAdapterAddress,
		config.MessageProtocolConfig{
			MessageFormat: config.TELEPORTER_V2.String(),
			Settings:      settings,
		},
		nil,
		ids.Empty,
	)
	require.NoError(t, err)

	destinationClient := mock_vms.NewMockDestinationClient(ctrl)
	ethClient := mock_evm.NewMockClient(ctrl)
	destinationClient.EXPECT().SenderAddresses().Return([]common.Address{testRelayerAddress}).AnyTimes()
	destinationClient.EXPECT().BlockGasLimit().Return(testBlockGasLimit).AnyTimes()
	destinationClient.EXPECT().Client().Return(ethClient).AnyTimes()

	handler, err := factory.NewMessageHandler(
		logging.NoLog{},
		&messages.SourceMessage{
			SourceBlockchainID: testSourceBlockchainID,
			ProtocolAddress:    testAdapterAddress,
			Payload:            unsignedMessage.Bytes(),
		},
		destinationClient,
		nil,
		mocks.NewMockMetrics(ctrl),
		ids.Empty,
		0,
	)
	require.NoError(t, err)

	messageID, err := teleporterUtils.CalculateMessageID(
		testTeleporterAddress,
		testSourceBlockchainID,
		testDestinationBlockchainID,
		teleporterMessage.MessageNonce,
	)
	require.NoError(t, err)
	messageReceivedInput, err := teleportermessengerv2.PackMessageReceived(messageID)
	require.NoError(t, err)
	teleporterABI, err := teleportermessengerv2.TeleporterMessengerV2MetaData.GetAbi()
	require.NoError(t, err)

	return handler, &handlerMocks{
		destinationClient:    destinationClient,
		ethClient:            ethClient,
		messageReceivedInput: messageReceivedInput,
		receiveSelector:      teleporterABI.Methods["receiveCrossChainMessage"].ID,
	}
}

// callRoutes answers the eth_call traffic a handler generates while sending a message.
type callRoutes struct {
	// delivered is what the messenger's messageReceived query reports.
	delivered bool
	// simulate answers the warp path's delivery simulation and receives the simulated call.
	simulate func(call ethereum.CallMsg) ([]byte, error)
	// commitment is what the merkle registry reports as the current validator set commitment.
	commitment *merkleregistry.ValidatorSetMerkleCommitment
}

func (h *handlerMocks) routeCalls(t *testing.T, routes callRoutes) {
	t.Helper()
	deliveredOutput, err := teleportermessengerv2.PackMessageReceivedOutput(routes.delivered)
	require.NoError(t, err)
	var commitmentOutput []byte
	if routes.commitment != nil {
		registryABI, err := merkleregistry.MerkleValidatorSetRegistryMetaData.GetAbi()
		require.NoError(t, err)
		commitmentOutput, err = registryABI.PackOutput("getValidatorSetCommitment", *routes.commitment)
		require.NoError(t, err)
	}

	h.ethClient.EXPECT().
		CallContract(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, call ethereum.CallMsg, _ *big.Int) ([]byte, error) {
			switch {
			case call.To != nil && *call.To == testRegistryAddress && commitmentOutput != nil:
				return commitmentOutput, nil
			case bytes.Equal(call.Data, h.messageReceivedInput):
				return deliveredOutput, nil
			case routes.simulate != nil && bytes.HasPrefix(call.Data, h.receiveSelector):
				return routes.simulate(call)
			default:
				t.Fatalf("unexpected eth_call to %v with data %x", call.To, call.Data)
				return nil, nil
			}
		}).
		AnyTimes()
}
