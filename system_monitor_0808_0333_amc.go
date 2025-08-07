// 代码生成时间: 2025-08-08 03:33:00
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// SystemPerformance represents the system's performance data.
type SystemPerformance struct {
    gorm.Model
    CPUUsage    float64
    MemoryUsage uint64
    DiskUsage   uint64
}

// DBClient is a wrapper around the *gorm.DB connection.
type DBClient struct {
    *gorm.DB
}

// NewDBClient creates a new database client with a SQLite database.
func NewDBClient(dataSourceName string) (*DBClient, error) {
    db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema.
    db.AutoMigrate(&SystemPerformance{})

    return &DBClient{db}, nil
}

// RecordPerformance records the system performance data into the database.
func (client *DBClient) RecordPerformance(performance SystemPerformance) error {
    if err := client.Create(&performance).Error; err != nil {
        return err
    }
    return nil
}

// main function to run the system performance monitoring tool.
func main() {
    dbClient, err := NewDBClient("system_monitor.db")
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Sample performance data.
    performance := SystemPerformance{
        CPUUsage:    75.4,
        MemoryUsage: 2048, // MB
        DiskUsage:   500,  // GB
    }

    // Record the performance data.
    if err := dbClient.RecordPerformance(performance); err != nil {
        log.Printf("Failed to record performance data: %v", err)
        return
    }

    fmt.Println("System performance data recorded successfully.")
}