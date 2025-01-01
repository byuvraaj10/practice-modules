package main

import (
	"ecommerce-inventory/config"
	"ecommerce-inventory/controller"
	"ecommerce-inventory/middleware"
	"ecommerce-inventory/repository"
	"ecommerce-inventory/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := config.InitializeDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Initialize components
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Set up router
	router := gin.Default()
	router.Use(middleware.LoggingMiddleware())

	// User routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Product routes with authentication
	authorized := router.Group("/", middleware.AuthMiddleware())
	{
		authorized.POST("/product", middleware.ValidationMiddleware(), productController.AddProduct)
		authorized.GET("/product/:id", productController.GetProduct)
		authorized.PUT("/product/:id", productController.UpdateProduct)
		authorized.DELETE("/product/:id", productController.DeleteProduct)
		authorized.GET("/products", productController.GetAllProducts)
	}

	// Start server
	router.Run(":8080")
}
