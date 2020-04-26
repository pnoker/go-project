/*
 * Copyright Pnoker. All Rights Reserved.
 */

package interfaces

import (
	"context"
	"emulator/internal/config"
	"sync"
)

// BootstrapHandler defines the contract each bootstrap handler must fulfill.  Implementation returns true if the
// handler completed successfully, false if it did not.
type BootstrapHandler func(ctx context.Context, wg *sync.WaitGroup, dic *config.Container) (success bool)
