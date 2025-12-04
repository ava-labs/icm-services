// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"os"
	"strconv"

	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/log"
)

func main() {
	if len(os.Args) < 2 {
		panic("Invalid argument count. Must provide at least one argument to specify command type.")
	}
	commandType := os.Args[1]

	switch commandType {
	case "constructKeylessTx":
		// Get the byte code of the teleporter contract to be deployed.
		if len(os.Args) != 3 {
			panic("Invalid argument count. Must provide JSON file containing contract bytecode.")
		}
		_, _, _, _, err := deploymentUtils.ConstructKeylessTransaction(
			os.Args[2],
			true,
			deploymentUtils.GetDefaultContractCreationGasPrice(),
		)
		if err != nil {
			panic("Failed to construct keyless transaction.")
		}
	case "deriveContractAddress":
		// Get the byte code of the teleporter contract to be deployed.
		if len(os.Args) != 4 {
			panic("Invalid argument count. Must provide address and nonce.")
		}

		deployerAddress := common.HexToAddress(os.Args[2])
		nonce, err := strconv.ParseUint(os.Args[3], 10, 64)
		if err != nil {
			panic("Failed to parse nonce as uint")
		}

		resultAddress := crypto.CreateAddress(deployerAddress, nonce)
		log.Info(resultAddress.Hex())
	default:
		panic("Invalid command type. Supported options are \"constructKeylessTx\" and \"deriveContractAddress\".")
	}
}
