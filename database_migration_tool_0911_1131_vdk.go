// 代码生成时间: 2025-09-11 11:31:07
package main

import (
    "fmt"
    "gorm.io/driver/sqlite" // You can change this to your preferred database driver
    "gorm.io/gorm"
)

// DBConfig contains the configuration for the database connection.
type DBConfig struct {
    DSN string
}

// Migrate performs database migration using GORM.
func Migrate(dbConfig *DBConfig) error {
    // Initialize GORM DB instance
    db, err := gorm.Open(sqlite.Open(dbConfig.DSN), &gorm.Config{})
    if err != nil {
        return fmt.Errorf("failed to connect to database: %w", err)
    }

    // Migrate the schema
    err = db.AutoMigrate()
    if err != nil {
        return fmt.Errorf("failed to migrate database: %w", err)
    }

    return nil
}

// main function to run the migration tool.
func main() {
    // Define database configuration
    dbConfig := &DBConfig{
        DSN: "gorm.db", // Replace with your database file path or connection string
    }

    // Run migration
    if err := Migrate(dbConfig); err != nil {
        fmt.Println("Migration failed: ", err)
    } else {
        fmt.Println("Migration successful")
    }
}
