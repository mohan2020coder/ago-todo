package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var DB *gorm.DB

func main() {
	// Initialize the database connection
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the Todo model
	DB.AutoMigrate(&Todo{})

	// Initialize the Gin router
	r := gin.Default()
	r.Use(corsMiddleware())

	// Serve static files from the Angular dist directory
	r.StaticFS("/static", http.Dir("../todo/dist/todo"))

	// API endpoints for CRUD operations
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	// Handle requests to non-existing API routes by serving index.html
	r.NoRoute(func(c *gin.Context) {
		// Serve index.html from the Angular build directory
		http.ServeFile(c.Writer, c.Request, "../todo/dist/todo/index.html")
	})

	// Start the Gin server on port 8080
	r.Run(":8080")
}

// Middleware to handle CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// Handlers for CRUD operations
func getTodos(c *gin.Context) {
	var todos []Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&Todo{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
