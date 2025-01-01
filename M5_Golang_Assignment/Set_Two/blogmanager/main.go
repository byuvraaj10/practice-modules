package main

import (
	db "blogmanager/config"
	"blogmanager/controller"
	"blogmanager/middleware"
	"blogmanager/repository"
	"blogmanager/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDatabase()

	// Initialize components
	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	// Set up Gin router
	r := gin.Default()

	// Apply middlewares
	r.Use(middleware.LoggingMiddleware())
	api := r.Group("/api", middleware.AuthMiddleware(db.GetDB()))

	// Blog routes
	api.POST("/blog", blogController.CreateBlog)
	api.GET("/blog/:id", blogController.GetBlog)
	api.GET("/blog", blogController.GetAllBlogs)
	api.PUT("/blog/:id", blogController.UpdateBlog)
	api.DELETE("/blog/:id", blogController.DeleteBlog)

	// Start server
	r.Run(":8080")
}
