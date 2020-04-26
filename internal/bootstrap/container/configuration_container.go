/*
 * Copyright Pnoker. All Rights Reserved.
 */

package container

import (
	"emulator/internal/config"
)

// ConfigurationName contains the name of the metadata's config.ConfigurationStruct implementation in the DIC.
var ConfigurationName = config.TypeInstanceToName((*config.ConfigurationStruct)(nil))

// ConfigurationFrom helper function queries the DIC and returns metadata's config.ConfigurationStruct implementation.
func ConfigurationFrom(get config.Get) *config.ConfigurationStruct {
	return get(ConfigurationName).(*config.ConfigurationStruct)
}
