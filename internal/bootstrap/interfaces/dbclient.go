/*
 * Copyright Pnoker. All Rights Reserved.
 */

package interfaces

import (
	"emulator/internal/models"
)

type DBClient interface {
	CloseSession()

	//Node
	AddNodeInfo(nodeInfo models.NodeInfo) error
	UpdateNodeInfo(nodeInfo models.NodeInfo) error
	GetAllNodeInfos() ([]models.NodeInfo, error)
	GetNodeInfoByKey(key string) (models.NodeInfo, error)
}
