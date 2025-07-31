// 代码生成时间: 2025-08-01 01:38:57
package main

import (
    "fmt"
    "log"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Theme 定义主题数据库模型
type Theme struct {
    gorm.Model
    Name string
}

// Database 数据库连接
var db *gorm.DB

func main() {
    // 连接数据库
    var err error
    db, err = gorm.Open(sqlite.Open("theme.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // 迁移模式
    db.AutoMigrate(&Theme{})

    // 设置路由
    http.HandleFunc("/", themeHandler)
    http.HandleFunc("/switch/", switchThemeHandler)

    // 启动服务器
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

// themeHandler 显示当前主题
func themeHandler(w http.ResponseWriter, r *http.Request) {
    var theme Theme
    if result := db.First(&theme, 1); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            fmt.Fprintln(w, "No theme found")
        } else {
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        }
    } else {
        fmt.Fprintln(w, theme.Name)
    }
}

// switchThemeHandler 切换主题
func switchThemeHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    themeName := r.FormValue("theme")
    if themeName == "" {
        http.Error(w, "Theme name is required", http.StatusBadRequest)
        return
    }

    // 查找当前主题
    var currentTheme Theme
    if result := db.First(&currentTheme, 1); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            // 如果没有主题，创建新主题
            db.Create(&Theme{Name: themeName})
        } else {
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        // 更新当前主题
        currentTheme.Name = themeName
        if result := db.Save(&currentTheme); result.Error != nil {
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
            return
        }
    }

    fmt.Fprintln(w, "Theme switched to: ", themeName)
}
