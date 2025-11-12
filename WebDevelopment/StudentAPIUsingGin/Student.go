package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

var students []student

func main() {
	r := gin.Default()
	r.POST("/students", func(c *gin.Context) {
		var newStudent student
		if err := c.BindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
			return
		}
		students = append(students, newStudent)
		c.JSON(http.StatusOK, gin.H{"message": "Student Added SuccessFully !!"})

	})
	r.GET("/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, students)

	})
	r.Run(":8080")
}
