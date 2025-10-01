// 代码生成时间: 2025-10-01 20:46:58
// learning_resource_service.go

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// LearningResource represents a learning resource entity
type LearningResource struct {
    gorm.Model
    Title       string `gorm:"type:varchar(255)" json:"title"`
    Author      string `gorm:"type:varchar(255)" json:"author"`
    Description string `gorm:"type:text" json:"description"`
    URL         string `gorm:"type:varchar(255)" json:"url"`
}

// DatabaseConfig contains the configuration for the database
type DatabaseConfig struct {
    DBName string
    DBUser string
    DBPass string
}

// DB is a global variable to hold the GORM database connection
var DB *gorm.DB

// InitializeDB initializes the database connection
func InitializeDB(config DatabaseConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
        config.DBUser, config.DBPass, config.DBName)
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&LearningResource{})

    return db, nil
}

// CreateLearningResource creates a new learning resource
func CreateLearningResource(db *gorm.DB, resource *LearningResource) error {
    if err := db.Create(&resource).Error; err != nil {
        return err
    }
    return nil
}

// GetLearningResources retrieves all learning resources
func GetLearningResources(db *gorm.DB) ([]LearningResource, error) {
    var resources []LearningResource
    if err := db.Find(&resources).Error; err != nil {
        return nil, err
    }
    return resources, nil
}

// UpdateLearningResource updates an existing learning resource
func UpdateLearningResource(db *gorm.DB, id uint, updates map[string]interface{}) error {
    resource := LearningResource{}
    if err := db.First(&resource, id).Error; err != nil {
        return err
    }
    for key, value := range updates {
        resource[key] = value
    }
    if err := db.Save(&resource).Error; err != nil {
        return err
    }
    return nil
}

// DeleteLearningResource deletes a learning resource
func DeleteLearningResource(db *gorm.DB, id uint) error {
    if err := db.Delete(&LearningResource{}, id).Error; err != nil {
        return err
    }
    return nil
}

func main() {
    // Define the database configuration
    config := DatabaseConfig{DBName: "learning_resources.db", DBUser: "admin", DBPass: "password"}

    // Initialize the database
    db, err := InitializeDB(config)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    DB = db
    defer DB.Close()

    // Create a new learning resource
    newResource := LearningResource{
        Title:       "Learn Go",
        Author:      "John Doe",
        Description: "A comprehensive guide to learning Go programming.",
        URL:         "https://example.com/learn-go",
    }
    if err := CreateLearningResource(DB, &newResource); err != nil {
        log.Fatalf("Failed to create learning resource: %v", err)
    }

    // Get all learning resources
    resources, err := GetLearningResources(DB)
    if err != nil {
        log.Fatalf("Failed to retrieve learning resources: %v", err)
    }
    data, _ := json.Marshal(resources)
    fmt.Println(string(data))

    // Update a learning resource
    if err := UpdateLearningResource(DB, newResource.ID, map[string]interface{}{
        "Title":       "Advanced Go",
        "Description": "A guide to advanced Go programming concepts.",
    }); err != nil {
        log.Fatalf("Failed to update learning resource: %v", err)
    }

    // Delete a learning resource
    if err := DeleteLearningResource(DB, newResource.ID); err != nil {
        log.Fatalf("Failed to delete learning resource: %v", err)
    }
}