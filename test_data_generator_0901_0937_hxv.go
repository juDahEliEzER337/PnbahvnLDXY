// 代码生成时间: 2025-09-01 09:37:04
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "time"
)

// User 定义用户模型
type User struct {
    gorm.Model
    Name          string
    Email         string `gorm:"type:varchar(100);uniqueIndex"`
    BirthDate     time.Time
    Age           int
}

// 初始化数据库并自动迁移User模型
func initDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // 自动迁移
    db.AutoMigrate(&User{})
    return db
}

// GenerateTestData 生成测试数据
func GenerateTestData(db *gorm.DB, count int) error {
    for i := 1; i <= count; i++ {
        user := User{
            Name:  fmt.Sprintf("User%d", i),
            Email: fmt.Sprintf("user%d@example.com", i),
            Age:   i * 10,
            BirthDate: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, i),
        }
        if err := db.Create(&user).Error; err != nil {
            return fmt.Errorf("failed to create user: %w", err)
        }
    }
    return nil
}

func main() {
    db := initDB()
    defer db.Migrator.Close()

    // 生成100个测试用户
    if err := GenerateTestData(db, 100); err != nil {
        log.Fatalf("An error occurred while generating test data: %v", err)
    } else {
        fmt.Println("Test data generated successfully.")
    }
}
