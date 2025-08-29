// 代码生成时间: 2025-08-29 21:35:33
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "golang.org/x/crypto/pbkdf2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// User represents a user with password
type User struct {
    gorm.Model
    Username string
    PasswordHash string
}

// AuthService handles user authentication
type AuthService struct {
    DB *gorm.DB
}

// NewAuthService initializes a new AuthService struct
func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{DB: db}
}

// HashPassword hashes a password using PBKDF2
func HashPassword(password string) (string, error) {
    passwordBytes := []byte(password)
    salt := []byte("salt")
    key := pbkdf2.Key(passwordBytes, salt, 4096, 32, sha256.New)
    return hex.EncodeToString(key), nil
}

// VerifyPassword checks if the provided password matches the stored hash
func VerifyPassword(password, hash string) bool {
    bytes, err := hex.DecodeString(hash)
    if err != nil {
        return false
    }
    newHash, err := HashPassword(password)
    if err != nil {
        return false
    }
    return newHash == hash
}

// AuthenticateUser checks if a username and password combination is valid
func (s *AuthService) AuthenticateUser(username, password string) error {
    var user User
    if result := s.DB.Where(&User{Username: username}).First(&user); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return fmt.Errorf("user not found")
        }
        return result.Error
    }
    if !VerifyPassword(password, user.PasswordHash) {
        return fmt.Errorf("invalid password")
    }
    return nil
}

func main() {
    // Initialize a new GORM DB connection
    db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }
    db.AutoMigrate(&User{})

    // Initialize AuthService
    authService := NewAuthService(db)

    // Example usage: Authenticate a user
    if err := authService.AuthenticateUser("exampleUser", "examplePassword"); err != nil {
        fmt.Println("Authentication failed: ", err)
    } else {
        fmt.Println("User authenticated successfully")
    }
}