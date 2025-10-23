package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//This is the main function

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello Magic Stream Movies Server")
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
