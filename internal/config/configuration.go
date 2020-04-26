/*
 * Copyright Pnoker. All Rights Reserved.
 */

package config

import (
	"emulator/internal/db/mongo"
	"reflect"
)

// Configuration interface provides an abstraction around a configuration struct.
type Configuration interface{}

// TypeInstanceToName converts an instance of a type to a unique name.
func TypeInstanceToName(v interface{}) string {
	t := reflect.TypeOf(v)

	if name := t.Name(); name != "" {
		// non-interface types
		return t.PkgPath() + "." + name
	}

	// interface types
	e := t.Elem()
	return e.PkgPath() + "." + e.Name()
}

type DatabaseInfo mongo.Configuration

// Struct used to parse the JSON configuration file
type ConfigurationStruct struct {
	Databases DatabaseInfo
}

// --- Database
// GetDatabaseInfo returns a database information map.
func (c *ConfigurationStruct) GetDatabaseInfo() DatabaseInfo {
	return c.Databases
}
