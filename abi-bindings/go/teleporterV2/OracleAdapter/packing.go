// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package oracleadapter

func PackReceiveOracleMessage(warpIndex uint32, oracleMsg OracleMessage) ([]byte, error) {
	abi, err := OracleAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return abi.Pack("receiveOracleMessage", warpIndex, oracleMsg)
}
