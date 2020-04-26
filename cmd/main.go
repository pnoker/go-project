/*
 * Copyright Pnoker. All Rights Reserved.
 */

package main

import (
	"context"
	"emulator/internal/bootstrap"
	"emulator/internal/bootstrap/container"
	"emulator/internal/bootstrap/handlers/database"
	"emulator/internal/bootstrap/handlers/monitor"
	"emulator/internal/bootstrap/interfaces"
	"emulator/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	configuration := &config.ConfigurationStruct{}
	dic := config.NewContainer(config.ServiceConstructorMap{
		container.ConfigurationName: func(get config.Get) interface{} {
			return configuration
		},
	})

	bootstrap.Run(
		ctx,
		cancel,
		configuration,
		dic,
		[]interfaces.BootstrapHandler{
			database.NewDatabase(configuration).BootstrapHandler,
			monitor.NewBootstrap().BootstrapHandler,
		})
}
