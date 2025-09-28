// 代码生成时间: 2025-09-29 00:01:35
package main

import (
    "fmt"
    "log"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// APITestTool is the main struct that holds the database connection
type APITestTool struct {
    DB *gorm.DB
}

// NewAPITestTool creates a new APITestTool instance with a SQLite database connection
func NewAPITestTool() *APITestTool {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Migrate the schema
    db.AutoMigrate(&APITestResult{})

    return &APITestTool{DB: db}
}

// APITestResult represents the result of an API test
type APITestResult struct {
    ID        uint   "gorm:primary_key"
    URL       string
    Method    string
    StatusCode int
    Response  string
    Timestamp string
}

// RunTest performs an API test and records the result in the database
func (tool *APITestTool) RunTest(url, method string) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to perform test: %w", err)
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("failed to read response: %w", err)
    }

    result := APITestResult{
        URL:       url,
        Method:    method,
        StatusCode: resp.StatusCode,
        Response:  string(body),
        Timestamp: time.Now().Format(time.RFC3339),
    }

    // Save the result to the database
    if err := tool.DB.Create(&result).Error; err != nil {
        return fmt.Errorf("failed to save test result: %w", err)
    }

    return nil
}

func main() {
    tool := NewAPITestTool()
    defer tool.DB.Close()

    // Example test
    if err := tool.RunTest("https://api.example.com/test", "GET"); err != nil {
        log.Println("API test failed: ", err)
    } else {
        log.Println("API test successful")
    }
}
