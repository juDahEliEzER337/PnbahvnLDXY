// 代码生成时间: 2025-08-07 14:38:30
package main

import (
	"fmt"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product represents a product in the inventory.
type Product struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);not null"`
	Price float64 `gorm:"type:decimal(10,2);not null"`
	Stock int     `gorm:"type:int;not null"`
}

// InventoryService handles the business logic of the inventory.
type InventoryService struct {
	db *gorm.DB
}

// NewInventoryService creates a new instance of InventoryService.
func NewInventoryService() *InventoryService {
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	return &InventoryService{db: db}
}

// AddProduct adds a new product to the inventory.
func (s *InventoryService) AddProduct(name string, price float64, stock int) error {
	product := Product{Name: name, Price: price, Stock: stock}
	result := s.db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateProduct updates an existing product in the inventory.
func (s *InventoryService) UpdateProduct(id uint, name string, price float64, stock int) error {
	product := Product{ID: id}
	result := s.db.Model(&product).Updates(Product{Name: name, Price: price, Stock: stock})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteProduct deletes a product from the inventory by its ID.
func (s *InventoryService) DeleteProduct(id uint) error {
	product := Product{ID: id}
	result := s.db.Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetProduct retrieves a product from the inventory by its ID.
func (s *InventoryService) GetProduct(id uint) (Product, error) {
	var product Product
	result := s.db.First(&product, id)
	if result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

// ListProducts retrieves all products in the inventory.
func (s *InventoryService) ListProducts() ([]Product, error) {
	var products []Product
	result := s.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func main() {
	service := NewInventoryService()
	
	// Create a new product
	if err := service.AddProduct("Laptop", 1200.99, 10); err != nil {
		log.Println("Error adding product: ", err)
	}
	
	// Update an existing product
	if err := service.UpdateProduct(1, "Laptop Pro", 1500.99, 5); err != nil {
		log.Println("Error updating product: ", err)
	}
	
	// Get a product by ID
	product, err := service.GetProduct(1)
	if err != nil {
		log.Println("Error getting product: ", err)
	} else {
		fmt.Printf("Product: %+v
", product)
	}
	
	// List all products
	products, err := service.ListProducts()
	if err != nil {
		log.Println("Error listing products: ", err)
	} else {
		fmt.Println("Products: ")
		for _, p := range products {
			fmt.Printf("%+v
", p)
		}
	}
	
	// Delete a product by ID
	if err := service.DeleteProduct(1); err != nil {
		log.Println("Error deleting product: ", err)
	}
}
