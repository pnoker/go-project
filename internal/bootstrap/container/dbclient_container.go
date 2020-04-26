/*
 * Copyright Pnoker. All Rights Reserved.
 */

package container

import (
	"emulator/internal/bootstrap/interfaces"
	"emulator/internal/config"
)

// DBClientInterfaceName contains the name of the interfaces.DBClient implementation in the DIC.
var DBClientInterfaceName = config.TypeInstanceToName((*interfaces.DBClient)(nil))

// DBClientFrom helper function queries the DIC and returns the interfaces.DBClient implementation.
func DBClientFrom(get config.Get) interfaces.DBClient {
	return get(DBClientInterfaceName).(interfaces.DBClient)
}
