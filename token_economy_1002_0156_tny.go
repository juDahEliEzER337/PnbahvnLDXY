// 代码生成时间: 2025-10-02 01:56:24
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Token represents the model for a token in the economy
type Token struct {
    gorm.Model
    Name        string  `gorm:"column:name;uniqueIndex"`
    Symbol      string  `gorm:"column:symbol;uniqueIndex"`
    TotalSupply uint256 `gorm:"column:total_supply"`
    OwnerID     uint   `gorm:"column:owner_id"`
}

// User represents the model for a user in the economy
type User struct {
    gorm.Model
    Name  string `gorm:"column:name"`
    Email string `gorm:"column:email;uniqueIndex"`
    Tokens []Token `gorm:"many2many:user_tokens;"`
}

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open("token_economy.db"), &gorm.Config{})
    if err != nil {
        fmt.Println("failed to connect to database: \boxed{err}")
        return
    }

    // Migrate the schema
    db.AutoMigrate(&Token{}, &User{})

    // Create a new user
    user := User{
        Name:  "John Doe",
        Email: "john.doe@example.com",
    }
    result := db.Create(&user)
    if result.Error != nil {
        fmt.Println("failed to create user: \boxed{result.Error}")
        return
    }

    // Create a new token
    token := Token{
        Name:        "Example Token",
        Symbol:      "EXT",
        TotalSupply: 1000000,
        OwnerID:     user.ID,
    }
    result = db.Create(&token)
    if result.Error != nil {
        fmt.Println("failed to create token: \boxed{result.Error}")
        return
    }

    // Assign the token to the user
    db.Model(&user).Association("Tokens