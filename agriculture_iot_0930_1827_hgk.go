// 代码生成时间: 2025-09-30 18:27:49
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// SensorReading represents a reading from an IoT sensor in agriculture
type SensorReading struct {
    gorm.Model
    SensorID    string
    Temperature float64
    Humidity     float64
    CO2Level    float64
    SoilMoisture float64
}

// DatabaseClient represents a database client for connecting to a database
type DatabaseClient struct {
    DB *gorm.DB
}

// NewDatabaseClient creates a new database client
func NewDatabaseClient() *DatabaseClient {
    db, err := gorm.Open(sqlite.Open("agriculture.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&SensorReading{})
    return &DatabaseClient{DB: db}
}

// AddSensorReading adds a new sensor reading to the database
func (client *DatabaseClient) AddSensorReading(reading SensorReading) error {
    result := client.DB.Create(&reading)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    dbClient := NewDatabaseClient()
    defer dbClient.DB.Close()

    // Example sensor reading
    sensorReading := SensorReading{
        SensorID:    "sensor1",
        Temperature: 22.5,
        Humidity:     45.0,
        CO2Level:     400,
        SoilMoisture: 30.0,
    }

    // Add sensor reading to the database
    if err := dbClient.AddSensorReading(sensorReading); err != nil {
        fmt.Println("Error adding sensor reading:", err)
    } else {
        fmt.Println("Sensor reading added successfully")
    }
}
