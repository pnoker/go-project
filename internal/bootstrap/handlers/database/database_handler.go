/*
 * Copyright Pnoker. All Rights Reserved.
 */

package database

import (
	"context"
	"emulator/internal/bootstrap/container"
	"emulator/internal/bootstrap/interfaces"
	"emulator/internal/config"
	"emulator/internal/db/mongo"
	"log"
	"os"
	"sync"
)

// Database contains references to dependencies required by the database bootstrap implementation.
type Database struct {
	database interfaces.Database
}

// NewDatabase is a factory method that returns an initialized Database receiver struct.
func NewDatabase(database interfaces.Database) Database {
	return Database{
		database: database,
	}
}

// Return the dbClient interface
func (d Database) newDBClient() (interfaces.DBClient, error) {
	databaseInfo := d.database.GetDatabaseInfo()
	host := os.Getenv(databaseInfo.Host)
	if host == "" {
		host = "localhost"
	}
	return mongo.NewClient(
		mongo.Configuration{
			Username: databaseInfo.Username,
			Password: databaseInfo.Password,
			Host:     host,
			Port:     databaseInfo.Port,
			Database: databaseInfo.Database,
			Timeout:  databaseInfo.Timeout,
		})
}

func (d Database) BootstrapHandler(ctx context.Context, wg *sync.WaitGroup, dic *config.Container) bool {
	var dbClient interfaces.DBClient
	dbClient, err := d.newDBClient()
	if err != nil {
		log.Println("error.couldn't create database client:", err.Error())
		dbClient = nil
	}
	if dbClient == nil {
		return false
	}

	dic.Update(config.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get config.Get) interface{} {
			return dbClient
		},
	})

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		dbClient.CloseSession()
		log.Println("info.database disconnected")
	}()

	return true
}
