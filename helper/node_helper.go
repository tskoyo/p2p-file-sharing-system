package helper

import "p2p-file-sharing-system/types"

func BuildNodeConfig(id string, port int) types.NodeConfig {
	return types.NodeConfig{
		ID:   id,
		Port: port,
	}
}
