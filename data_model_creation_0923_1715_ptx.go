// 代码生成时间: 2025-09-23 17:15:31
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Define the data model structure.
// User represents a user with an ID, Name, and Age.
type User struct {
    gorm.Model
    Name string `gorm:"type:varchar(100)"`
    Age  int    `gorm:"type:integer"`
}

// DatabaseConfig holds the configuration for the database.
type DatabaseConfig struct {
    Driver   string
    DataSource string
}

func main() {
    // Define the database configuration.
    config := DatabaseConfig{
        Driver:   "sqlite",
        DataSource: "test.db",
    }

    // Connect to the database.
    db, err := gorm.Open(sqlite.Open(config.DataSource), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    // Migrate the schema.
    if err := db.AutoMigrate(&User{}); err != nil {
        panic("failed to migrate database: " + err.Error())
    }

    // Create a new user and insert it into the database.
    newUser := User{Name: "John Doe", Age: 30}
    if result := db.Create(&newUser); result.Error != nil {
        panic("failed to create user: " + result.Error.Error())
    } else {
        fmt.Printf("User created successfully: %+v
", newUser)
    }

    // Query the user from the database.
    var user User
    if result := db.First(&user, 1); result.Error != nil {
        panic("failed to query user: " + result.Error.Error())
    } else {
        fmt.Printf("User found: %+v
", user)
    }

    // Update the user's age.
    if result := db.Model(&user).Update("Age", 31); result.Error != nil {
        panic("failed to update user: " + result.Error.Error())
    } else {
        fmt.Printf("User updated successfully: %+v
", user)
    }

    // Delete the user from the database.
    if result := db.Delete(&user, 1); result.Error != nil {
        panic("failed to delete user: " + result.Error.Error())
    } else {
        fmt.Printf("User deleted successfully
")
    }
}
