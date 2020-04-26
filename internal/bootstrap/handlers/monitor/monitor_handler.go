/*
 * Copyright Pnoker. All Rights Reserved.
 */

package monitor

import (
	"context"
	"emulator/internal/bootstrap/container"
	"emulator/internal/config"
	"emulator/internal/utils"
	"log"
	"sync"
)

// Monitor contains references to dependencies required by the start Monitor handler.
type Monitor struct{}

// NewBootstrap is a factory method that returns an initialized Monitor Monitor struct.
func NewBootstrap() *Monitor {
	return &Monitor{}
}

// BootstrapHandler fulfills the BootstrapHandler contract.  It creates no go routines.  It logs a "standard" set of
// Monitor when the service first starts up successfully.
func (h Monitor) BootstrapHandler(ctx context.Context, wg *sync.WaitGroup, dic *config.Container) bool {
	nodeName, dbClient := utils.RandomString(16), container.DBClientFrom(dic.Get)

	distributeTicker := utils.Timer(func() {
		dbClient.GetAllNodeInfos()
		log.Println(nodeName)
	}, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		distributeTicker.Stop()
		log.Println("info.monitor closed")
	}()

	return true
}
