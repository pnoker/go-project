/*
 * Copyright Pnoker. All Rights Reserved.
 */

package interfaces

import (
	"emulator/internal/config"
)

// Database interface provides an abstraction for obtaining the database configuration information.
type Database interface {
	// GetDatabaseInfo returns a database information.
	GetDatabaseInfo() config.DatabaseInfo
}
