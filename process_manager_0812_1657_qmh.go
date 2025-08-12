// 代码生成时间: 2025-08-12 16:57:08
// process_manager.go
// 进程管理器，使用Go语言和GORM框架实现

package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Process 定义进程信息的结构体
type Process struct {
    ID        uint   "gorm:"primaryKey""
    Name      string
    Status    string // 可以是 `running`, `paused`, `stopped`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// DBConfig 数据库配置
type DBConfig struct {
    DbName string
}

// DB 初始化数据库连接
var DB *gorm.DB
var dbConfig = DBConfig{DbName: "process_manager.db"}

// InitDB 初始化数据库和表
func InitDB() error {
    var err error
    // 连接数据库
    DB, err = gorm.Open(sqlite.Open(dbConfig.DbName), &gorm.Config{})
    if err != nil {
        return err
    }
    // 自动迁移模式
    DB.AutoMigrate(&Process{})
    return nil
}

// AddProcess 添加新的进程
func AddProcess(name string) (*Process, error) {
    process := Process{Name: name, Status: "stopped"}
    result := DB.Create(&process)
    return &process, result.Error
}

// StartProcess 启动进程
func StartProcess(id uint) error {
    var process Process
    result := DB.First(&process, id)
    if result.Error != nil {
        return result.Error
    }
    process.Status = "running"
    result = DB.Save(&process)
    return result.Error
}

// StopProcess 停止进程
func StopProcess(id uint) error {
    var process Process
    result := DB.First(&process, id)
    if result.Error != nil {
        return result.Error
    }
    process.Status = "stopped"
    result = DB.Save(&process)
    return result.Error
}

// main 函数是程序的入口点
func main() {
    err := InitDB()
    if err != nil {
        fmt.Printf("Failed to connect database: %s
", err)
        return
    }
    defer DB.Close()
    
    // 示例：添加一个进程
    process, err := AddProcess("ExampleProcess")
    if err != nil {
        fmt.Printf("Failed to add process: %s
", err)
        return
    }
    fmt.Printf("Process added with ID: %d
", process.ID)
    
    // 示例：启动进程
    err = StartProcess(process.ID)
    if err != nil {
        fmt.Printf("Failed to start process: %s
", err)
        return
    }
    fmt.Println("Process started successfully.")
    
    // 示例：停止进程
    err = StopProcess(process.ID)
    if err != nil {
        fmt.Printf("Failed to stop process: %s
", err)
        return
    }
    fmt.Println("Process stopped successfully.")
}