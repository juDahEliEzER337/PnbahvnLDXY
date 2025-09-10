// 代码生成时间: 2025-09-10 22:20:42
package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "time"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// 初始化数据库连接
func initDB() *gorm.DB {
    connStr := "file:performance_test.db?cache=shared&mode=memory&_fk=1" // 使用内存数据库进行测试
    db, err := gorm.Open(sqlite.Open(connStr), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }
    return db
}

// 模型定义
type PerformanceTest struct {
    ID        uint      `gorm:"primaryKey"`
    CreatedAt time.Time `gorm:"index"`
    Data      string    `gorm:"type:text"`
}

// 插入测试数据
func insertTestData(db *gorm.DB, count int) {
    for i := 0; i < count; i++ {
        err := db.Create(&PerformanceTest{
            CreatedAt: time.Now(),
            Data:      strconv.Itoa(i),
        }).Error
        if err != nil {
            log.Printf("failed to insert test data: %v", err)
        }
    }
}

// 读取测试数据
func readTestData(db *gorm.DB) {
    var testData PerformanceTest
    err := db.First(&testData).Error
    if err != nil {
        log.Printf("failed to read test data: %v", err)
    } else {
        fmt.Println("Read test data: ID =", testData.ID, ", Data =", testData.Data)
    }
}

func main() {
    db := initDB()
    defer db.Migrator().DropTable(&PerformanceTest{}) // 清理环境

    // 创建表
    err := db.AutoMigrate(&PerformanceTest{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    // 插入数据
    const testCount = 10000
    start := time.Now()
    insertTestData(db, testCount)
    fmt.Printf("Insert %d records took %s
", testCount, time.Since(start))

    // 读取数据
    start = time.Now()
    readTestData(db)
    fmt.Printf("Read one record took %s
", time.Since(start))
}
