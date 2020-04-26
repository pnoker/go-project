/*
 * Copyright Pnoker. All Rights Reserved.
 */

package mongo

import (
	"errors"
)

const (

	//Node
	NodeInfo = "node_info"
)

var (
	ErrNotFound            = errors.New("Item not found")
	ErrUnsupportedDatabase = errors.New("Unsupported database type")
	ErrInvalidObjectId     = errors.New("Invalid object ID")
	ErrNotUnique           = errors.New("Resource already exists")
	ErrCommandStillInUse   = errors.New("Command is still in use by device profiles")
	ErrSlugEmpty           = errors.New("Slug is nil or empty")
	ErrNameEmpty           = errors.New("Name is required")
)

type Configuration struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Timeout  int
}
