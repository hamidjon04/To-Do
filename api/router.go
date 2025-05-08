package api

import (
	_ "todo/api/docs"
	"todo/api/handler"
	"todo/api/middleware"
	postgres "todo/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           ToDo API
// @version         1.0
// @description     API for managing ToDo tasks with user authentication.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token.
func Router(storage *postgres.DB) *gin.Engine {
	router := gin.Default()

	// ⚠️ CORS middleware birinchi bo‘lishi kerak
	router.Use(middleware.CORSMiddleware())

	// Swagger (no auth)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewHandler(storage)

	// No auth routes
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	// Auth routes (token kerak bo‘ladi)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware()) // token required
	{
		protected.POST("/todos", h.CreateTodo)
		protected.GET("/todos", h.GetTodos)
		protected.PUT("/todos", h.UpdateTodo)
		protected.DELETE("/todos", h.DeleteTodo)
	}

	return router
}
