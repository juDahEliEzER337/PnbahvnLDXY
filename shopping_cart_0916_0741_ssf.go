// 代码生成时间: 2025-09-16 07:41:56
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// 定义购物车商品结构体
type CartItem struct {
    gorm.Model
    ProductID  uint
    Quantity  int
}

// 定义购物车结构体
type ShoppingCart struct {
    gorm.Model
    Items       []CartItem // 购物车中的商品列表
}

// 初始化数据库
func initDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("shopping_cart.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database"))
    }

    // 迁移模式
    db.AutoMigrate(&CartItem{}, &ShoppingCart{})
    return db
}

// 添加商品到购物车
func addToCart(db *gorm.DB, cartID, productID uint, quantity int) error {
    var cart ShoppingCart
    // 根据购物车ID查询购物车
    if err := db.First(&cart, cartID).Error; err != nil {
        return err
    }

    // 检查商品是否已存在购物车中
    for _, item := range cart.Items {
        if item.ProductID == productID {
            return fmt.Errorf("product already exists in cart")
        }
    }

    // 创建新的购物车商品项
    newItem := CartItem{ProductID: productID, Quantity: quantity}
    // 将新商品添加到购物车
    db.Model(&cart).Association("Items").Append(newItem)
    return nil
}

// 主函数
func main() {
    db := initDB()
    defer db.Close()

    // 示例：向购物车添加商品
    cartID := uint(1) // 假设购物车ID为1
    productID := uint(101) // 假设产品ID为101
    quantity := 2 // 购买数量为2

    if err := addToCart(db, cartID, productID, quantity); err != nil {
        fmt.Println("Error adding to cart: ", err)
    } else {
        fmt.Println("Product added to cart successfully")
    }
}
