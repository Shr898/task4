package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	// gin router with default middleware
	r := gin.Default()

	// GET endpoint
	r.GET("/task4" , func(c *gin.Context){
		//Return a JSON response
		c.JSON(http.StatusOK , gin.H{
			"message" : "This is a task4 app",
		})
	})

	// New task4 feature
	r.GET("/task4-new-feature" , func(c *gin.Context){
		c.JSON(http.StatusOK , gin.H{
			"message" : "This is a new feature for task4",
		})
	})
	
	fmt.Println("Started the task4 go application on port 8080")
	r.Run()
}