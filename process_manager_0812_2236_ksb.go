// 代码生成时间: 2025-08-12 22:36:45
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Process 代表进程信息的结构体
type Process struct {
    gorm.Model
    Name        string `gorm:"type:varchar(100);uniqueIndex"` // 进程名称
    Description string `gorm:"type:varchar(255)"`             // 进程描述
    PID         int    `gorm:"uniqueIndex"`                     // 进程ID
    Status      string `gorm:"type:varchar(50)"`                  // 进程状态
}

// ProcessManager 包含数据库连接的进程管理器
type ProcessManager struct {
    DB *gorm.DB
}

// NewProcessManager 创建一个新的进程管理器实例
func NewProcessManager() *ProcessManager {
    db, err := gorm.Open(sqlite.Open("process_manager.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }

    // Migrate the schema
    db.AutoMigrate(&Process{})
    return &ProcessManager{DB: db}
}

// AddProcess 添加一个新的进程到数据库
func (pm *ProcessManager) AddProcess(name, description string, pid int, status string) error {
    process := Process{Name: name, Description: description, PID: pid, Status: status}
    if result := pm.DB.Create(&process); result.Error != nil {
        return result.Error
    }
    return nil
}

// GetProcess 根据进程ID获取进程信息
func (pm *ProcessManager) GetProcess(pid int) (*Process, error) {
    var process Process
    if result := pm.DB.First(&process, pid); result.Error != nil {
        return nil, result.Error
    }
    return &process, nil
}

// UpdateProcess 更新进程的状态
func (pm *ProcessManager) UpdateProcess(pid int, status string) error {
    result := pm.DB.Model(&Process{}).Where("pid = ?", pid).Update("status", status)
    return result.Error
}

// DeleteProcess 根据进程ID删除进程信息
func (pm *ProcessManager) DeleteProcess(pid int) error {
    result := pm.DB.Delete(&Process{}, pid)
    return result.Error
}

// ListProcesses 列出所有进程信息
func (pm *ProcessManager) ListProcesses() ([]Process, error) {
    var processes []Process
    if result := pm.DB.Find(&processes); result.Error != nil {
        return nil, result.Error
    }
    return processes, nil
}

func main() {
    pm := NewProcessManager()
    defer pm.DB.Close()

    // 添加一个示例进程
    if err := pm.AddProcess("SampleProcess", "This is a sample process.", 1234, "running"); err != nil {
        log.Printf("Error adding process: %v", err)
    }

    // 获取进程信息
    process, err := pm.GetProcess(1234)
    if err != nil {
        log.Printf("Error getting process: %v", err)
    } else {
        fmt.Printf("Process Found: %+v
", process)
    }

    // 更新进程状态
    if err := pm.UpdateProcess(1234, "stopped"); err != nil {
        log.Printf("Error updating process: %v", err)
    }

    // 列出所有进程
    processes, err := pm.ListProcesses()
    if err != nil {
        log.Printf("Error listing processes: %v", err)
    } else {
        fmt.Printf("Processes: %+v
", processes)
    }
}
