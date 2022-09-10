// This file contains the token and chain config to be used in the mainnet environment.
//
// This file is maintained by hand. Add / remove / update entries as appropriate.

package governor

import (
	"github.com/certusone/wormhole/node/pkg/vaa"
)

func chainList() []chainConfigEntry {
	return []chainConfigEntry{
		chainConfigEntry{emitterChainID: vaa.ChainIDSolana, dailyLimit: 50000000},
		chainConfigEntry{emitterChainID: vaa.ChainIDEthereum, dailyLimit: 50000000},
		chainConfigEntry{emitterChainID: vaa.ChainIDTerra, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDBSC, dailyLimit: 5000000},
		chainConfigEntry{emitterChainID: vaa.ChainIDPolygon, dailyLimit: 5000000},
		chainConfigEntry{emitterChainID: vaa.ChainIDAvalanche, dailyLimit: 5000000},
		chainConfigEntry{emitterChainID: vaa.ChainIDOasis, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDAlgorand, dailyLimit: 200000},
		chainConfigEntry{emitterChainID: vaa.ChainIDAurora, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDFantom, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDKarura, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDAcala, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDKlaytn, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDCelo, dailyLimit: 500000},
		chainConfigEntry{emitterChainID: vaa.ChainIDTerra2, dailyLimit: 500000},
	}
}