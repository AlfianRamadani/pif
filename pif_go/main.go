package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	router := gin.Default()

	router.GET("/test-performance", func(c *gin.Context) {
		const n = 10
		
		result := fibonacci(n)
		
		c.JSON(200, gin.H{
			"framework": "Gin",
			"language":  "Go",
			"task":      "Calculate Fibonacci(" + strconv.Itoa(n) + ")",
			"result":    result,
		})
	})

	router.Run(":8080")
}