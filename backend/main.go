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
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Use(corsMiddleware())

	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Static("/frontend", "./frontend/dist/frontend")
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/frontend/index.html")
	})

	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
