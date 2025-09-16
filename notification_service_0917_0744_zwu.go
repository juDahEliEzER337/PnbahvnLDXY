// 代码生成时间: 2025-09-17 07:44:37
// notification_service.go
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Notification 定义通知信息结构
type Notification struct {
    gorm.Model
    Title   string `gorm:"type:varchar(255)"`
    Content string `gorm:"type:text"`
}

// NotificationService 定义通知服务
type NotificationService struct {
    db *gorm.DB
}

// NewNotificationService 初始化通知服务
func NewNotificationService() (*NotificationService, error) {
    var err error
    db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect database: %w", err)
    }

    // 自动迁移模式
    db.AutoMigrate(&Notification{})

    return &NotificationService{db: db}, nil
}

// CreateNotification 创建一条通知
func (s *NotificationService) CreateNotification(title, content string) (*Notification, error) {
    notification := &Notification{Title: title, Content: content}
    if result := s.db.Create(notification); result.Error != nil {
        return nil, result.Error
    }
    return notification, nil
}

// GetAllNotifications 获取所有通知
func (s *NotificationService) GetAllNotifications() ([]Notification, error) {
    var notifications []Notification
    if result := s.db.Find(&notifications); result.Error != nil {
        return nil, result.Error
    }
    return notifications, nil
}

// main 函数入口
func main() {
    service, err := NewNotificationService()
    if err != nil {
        panic(err)
    }

    // 创建通知示例
    _, err = service.CreateNotification("Welcome", "This is a welcome notification.")
    if err != nil {
        panic(err)
    }

    // 获取所有通知示例
    notifications, err := service.GetAllNotifications()
    if err != nil {
        panic(err)
    }

    // 打印获取到的通知
    for _, notification := range notifications {
        fmt.Printf("ID: %d, Title: %s, Content: %s
", notification.ID, notification.Title, notification.Content)
    }
}