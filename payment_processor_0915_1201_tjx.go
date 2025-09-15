// 代码生成时间: 2025-09-15 12:01:06
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// Payment 定义支付实体
type Payment struct {
    gorm.Model
    UserID    uint   `gorm:"index;uniqueIndex:idx_user_id_status"`
    Amount   float64
    Status   string // 支付状态，例如：pending, completed, failed
}

// PaymentService 定义支付服务
type PaymentService struct {
    db *gorm.DB
}

// NewPaymentService 创建新的支付服务实例
func NewPaymentService(db *gorm.DB) *PaymentService {
    return &PaymentService{db: db}
}

// CreatePayment 创建支付记录
func (s *PaymentService) CreatePayment(userID uint, amount float64) (*Payment, error) {
    payment := Payment{UserID: userID, Amount: amount, Status: "pending"}
    result := s.db.Create(&payment)
    return &payment, result.Error
}

// ProcessPayment 处理支付流程
func (s *PaymentService) ProcessPayment(paymentID uint) error {
    var payment Payment
    // 查询支付记录
    if err := s.db.First(&payment, paymentID).Error; err != nil {
        return err
    }
    // 检查支付状态
    if payment.Status != "pending" {
        return ErrPaymentStatusInvalid
    }
    // 模拟支付处理逻辑
    // 这里可以添加真实的支付处理逻辑，如调用外部支付服务API等
    payment.Status = "completed"
    // 更新支付状态
    if err := s.db.Save(&payment).Error; err != nil {
        return err
    }
    return nil
}

// ErrPaymentStatusInvalid 定义支付状态无效错误
var ErrPaymentStatusInvalid = gorm.ErrInvalidData{
    Err:       "payment status is invalid",
    Statement: "WHERE status != 'pending'",
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }
    // 自动迁移
    db.AutoMigrate(&Payment{})

    // 创建支付服务实例
    paymentService := NewPaymentService(db)

    // 创建支付记录
    payment, err := paymentService.CreatePayment(1, 100.0)
    if err != nil {
        log.Fatal("failed to create payment:", err)
    }
    log.Printf("Created payment: %+v", payment)

    // 处理支付
    if err := paymentService.ProcessPayment(payment.ID); err != nil {
        log.Fatal("failed to process payment:", err)
    }
    log.Println("Payment processed successfully")
}