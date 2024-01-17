package main

import (
	"fmt"
	"sync"
)

type DatabaseConnection struct {
	// Assume some database connection details
	connected bool
}

var dbInstance *DatabaseConnection
var dbOnce sync.Once

func initializeDatabase() {
	// Simulate expensive database initialization
	fmt.Println("Initializing the database...")
	// Code for connecting to the database goes here
	// ...
	dbInstance = &DatabaseConnection{connected: true}
}

func getDatabaseInstance() *DatabaseConnection {
	dbOnce.Do(initializeDatabase)
	return dbInstance
}

func main() {
	// Multiple goroutines calling getDatabaseInstance, but initializeDatabase will be executed only once.
	for i := 0; i < 3; i++ {
		go func() {
			db := getDatabaseInstance()
			fmt.Printf("Database connected: %v\n", db.connected)
		}()
	}

	// Sleep to allow goroutines to complete before program exit
	select {}
}
