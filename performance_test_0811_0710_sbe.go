// 代码生成时间: 2025-08-11 07:10:46
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "os"
    "time"
)

// 定义一个与数据库模型对应的结构体
type Performance struct {
    ID        uint      `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// 初始化数据库连接
var db *gorm.DB
var err error

func initDB() {
    // 使用SQLite内存数据库进行性能测试
    db, err = gorm.Open(sqlite.Open("file:performance_test.db?mode=memory&cache=shared&_foreign_keys=on"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // 迁移模式，确保数据库模式匹配
    db.AutoMigrate(&Performance{})
}

func main() {
    initDB()
    defer db.Close()

    // 执行性能测试
    testPerformance()
}

// testPerformance 函数用于执行性能测试
func testPerformance() {
    const batchSize = 10000
    const repeatCount = 10

    for i := 0; i < repeatCount; i++ {
        start := time.Now()

        // 批量插入
        for j := 0; j < batchSize; j++ {
            db.Create(&Performance{CreatedAt: time.Now(), UpdatedAt: time.Now()})
        }

        // 测量时间并打印结果
        fmt.Printf("Batch %d: %v
", i+1, time.Since(start))
    }
}
