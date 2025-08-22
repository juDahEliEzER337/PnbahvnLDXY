// 代码生成时间: 2025-08-23 05:07:56
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "net/http"
    "encoding/json"
    "bytes"
)

// Message 代表消息实体
type Message struct {
    gorm.Model
    Title   string
    Content string
    User    string
}

// MessageService 处理消息通知的业务逻辑
type MessageService struct {
    DB *gorm.DB
}

// NewMessageService 创建一个新的 MessageService 实例
func NewMessageService(db *gorm.DB) *MessageService {
    return &MessageService{DB: db}
}

// SendMessage 发送消息
func (s *MessageService) SendMessage(title, content, user string) error {
    message := Message{
        Title:   title,
        Content: content,
        User:    user,
    }
    if err := s.DB.Create(&message).Error; err != nil {
        return err
    }
    return nil
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("message.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: &", err)
    }
    // 迁移模式
    db.AutoMigrate(&Message{})

    service := NewMessageService(db)

    // 创建 HTTP 服务器
    http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var message Message
        if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer r.Body.Close()

        if err := service.SendMessage(message.Title, message.Content, message.User); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Message sent successfully",
        })
    })

    log.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}