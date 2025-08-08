// 代码生成时间: 2025-08-08 14:35:57
package main

import (
    "fmt"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// 定义User模型
type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"type:varchar(100);uniqueIndex"`
}

// 定义HTTP请求处理器的结构体
type HttpRequestHandler struct {
    DB *gorm.DB
}

// NewHttpRequestHandler初始化HTTP请求处理器
func NewHttpRequestHandler() *HttpRequestHandler {
    // 连接数据库（这里以SQLite为例，实际项目中可以更换为其他数据库）
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 迁移schema
    db.AutoMigrate(&User{})

    return &HttpRequestHandler{DB: db}
}

// CreateUser处理创建用户的请求
func (h *HttpRequestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
        return
    }

    var user User
    if err := r.ParseForm(); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    user.Name = r.FormValue("name")
    user.Email = r.FormValue("email")

    if err := h.DB.Create(&user).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "User created successfully")
}

// 主函数启动HTTP服务器
func main() {
    handler := NewHttpRequestHandler()
    http.HandleFunc("/user", handler.CreateUser)
    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}