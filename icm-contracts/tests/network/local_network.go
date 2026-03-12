package network

import (
	"crypto/ecdsa"

	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/common"
)

type LocalNetwork interface {
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)
	GetNetworkInfo() []testinfo.NetworkTestInfo
	TearDownNetwork()
}

// NewTeleporterTestInfo Get a map of teleporter info for all networks
func NewTeleporterTestInfo(networks ...LocalNetwork) utils.TeleporterTestInfo {
	t := make(utils.TeleporterTestInfo)
	for _, nw := range networks {
		for _, info := range nw.GetNetworkInfo() {
			t[info.ChainID()] = &utils.ChainTeleporterInfo{}
		}
	}
	return t
}
