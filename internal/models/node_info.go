/*
 * Copyright Pnoker. All Rights Reserved.
 */

package models

import "time"

type NodeInfo struct {
	Key        string    `bson:"key"`
	ExpireTime time.Time `bson:"expire_time"`
}
