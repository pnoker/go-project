/*
 * Copyright Pnoker. All Rights Reserved.
 */

package bootstrap

import (
	"context"
	"emulator/internal/bootstrap/interfaces"
	"emulator/internal/config"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Deferred defines the signature of a function returned by RunAndReturnWaitGroup that should be executed via defer.
type Deferred func()

type Processor struct {
	ctx context.Context
	wg  *sync.WaitGroup
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// fatalError logs an error and exits the application.  It's intended to be used only within the bootstrap prior to
// any go routines being spawned.
func fatalError(err error) {
	os.Exit(1)
}

// loadConfiguration attempts to read and unmarshal toml-based configuration into a configuration struct.
func (cp *Processor) loadConfiguration(config config.Configuration) error {
	configDir := "./res"

	fileName := configDir + "/application.toml"

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("could not load configuration file (%s): %s", fileName, err.Error())
	}
	if err = toml.Unmarshal(contents, config); err != nil {
		return fmt.Errorf("could not load configuration file (%s): %s", fileName, err.Error())
	}

	return nil
}

// NewProcessor creates a new configuration Processor
func NewProcessor(ctx context.Context, wg *sync.WaitGroup) *Processor {
	return &Processor{
		ctx: ctx,
		wg:  wg,
	}
}

func (cp *Processor) Process(serviceConfig config.Configuration) error {
	// Local configuration must be loaded first in case need registry config info and/or
	// need to push it to the Configuration Provider.
	if err := cp.loadConfiguration(serviceConfig); err != nil {
		return err
	}
	return nil
}

func RunAndReturnWaitGroup(
	ctx context.Context,
	cancel context.CancelFunc,
	serviceConfig config.Configuration,
	dic *config.Container,
	handlers []interfaces.BootstrapHandler) (*sync.WaitGroup, Deferred, bool) {
	var wg sync.WaitGroup
	deferred := func() {}

	configProcessor := NewProcessor(ctx, &wg)
	if err := configProcessor.Process(serviceConfig); err != nil {
		fatalError(err)
	}

	// call individual bootstrap handlers.
	startedSuccessfully := true
	for i := range handlers {
		if handlers[i](ctx, &wg, dic) == false {
			cancel()
			startedSuccessfully = false
			break
		}
	}

	return &wg, deferred, startedSuccessfully
}

func Run(
	ctx context.Context,
	cancel context.CancelFunc,
	serviceConfig config.Configuration,
	dic *config.Container,
	handlers []interfaces.BootstrapHandler) {
	wg, deferred, _ := RunAndReturnWaitGroup(
		ctx,
		cancel,
		serviceConfig,
		dic,
		handlers)

	defer deferred()

	// wait for go routines to stop executing.
	wg.Wait()
}
