// 代码生成时间: 2025-09-02 17:14:10
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Product represents a product in the shopping cart
type Product struct {
    gorm.Model
    Code  string  "json:"code"" // Unique product code
    Price float64 "json:"price"" // Price of the product
}

// Cart represents a shopping cart
type Cart struct {
    gorm.Model
    Items []CartItem "json:"items" gorm:""" // Items in the cart
}

// CartItem represents an item in the shopping cart
type CartItem struct {
    gorm.Model
    ProductID uint   "json:"product_id"" // ID of the product
    Quantity  int    "json:"quantity""    // Quantity of the product
}

// DB is a global variable for the database connection
var DB *gorm.DB

// SetupDatabase initializes the database connection
func SetupDatabase() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("shopping_cart.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    return DB.AutoMigrate(&Product{}, &Cart{}, &CartItem{})
}

// AddProductToCart adds a product to the cart
func AddProductToCart(cartID uint, productID uint, quantity int) error {
    var cart Cart
    if err := DB.First(&cart, cartID).Error; err != nil {
        return err
    }

    // Check if the product is already in the cart
    var item CartItem
    DB.Where(&CartItem{ProductID: productID}).First(&item)
    if item.ID != 0 {
        // Update the quantity if the product is already in the cart
        return DB.Model(&cart).Where(&CartItem{ProductID: productID}).Update("quantity", gorm.Expr("quantity + ?