// 代码生成时间: 2025-08-16 23:19:07
package main

import (
    "fmt"
    "log"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/sqlite"
# 添加错误处理
    "gorm.io/gorm"
)

// User represents a user with fields to be validated
type User struct {
    Name     string `gorm:"column:name;not null" json:"name"`
    Email    string `gorm:"column:email;uniqueIndex;not null" json:"email"`
    Password string `gorm:"column:password;not null" json:"password"`
}

// ValidateForm validates the user form data
func ValidateForm(user *User) error {
    // Validate name
    if user.Name == "" {
        return fmt.Errorf("name is required")
    }
# NOTE: 重要实现细节
    // Validate email
    if user.Email == "" {
        return fmt.Errorf("email is required")
    }
    // Validate email format
    if !containsOnly(user.Email, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@._+-") {
# TODO: 优化性能
        return fmt.Errorf("invalid email format")
    }
    // Validate password length
    if len(user.Password) < 8 {
        return fmt.Errorf("password must be at least 8 characters long")
    }
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
# 优化算法效率
        return fmt.Errorf("failed to hash password: %w", err)
    }
    user.Password = string(hashedPassword)
    return nil
}

// containsOnly checks if a string contains only specific characters
func containsOnly(s, valid string) bool {
    for i := range s {
        if !strings.ContainsRune(valid, rune(s[i])) {
            return false
