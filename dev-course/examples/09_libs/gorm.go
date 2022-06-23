package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//START_1 OMIT
type Product struct {
	gorm.Model // contains id, createdAt, updatedAt, deletedAt
	Code       string
	Price      uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// ...
	db.AutoMigrate(&Product{}) // Migrate the schema

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, "code = ?", "D42")                        // Find product with code D42
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // Update multiple fields
	db.Delete(&product, 1)                                       // Delete - delete product
}

//END_1 OMIT
