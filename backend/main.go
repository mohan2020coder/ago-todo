package main

import (
	"log"
	"net/http"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Serve Angular static files
	r.StaticFS("/", http.Dir("./frontend/dist/frontend"))

	// Connect to SQLite database
	var err error
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Auto Migrate database schema
	db.AutoMigrate(&models.Todo{})

	// Define API routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", GetTodos)
		v1.POST("/todos", CreateTodo)
		v1.GET("/todos/:id", GetTodo)
		v1.PUT("/todos/:id", UpdateTodo)
		v1.DELETE("/todos/:id", DeleteTodo)
	}

	// Start server
	r.Run(":8080")
}

// Handlers
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	todo.Completed = false
	todo.CreatedAt = time.Now()
	db.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if db.First(&todo, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if db.First(&todo, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.BindJSON(&todo)
	db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if db.First(&todo, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
