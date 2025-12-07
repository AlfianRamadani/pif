package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()

	router.GET("/test-performance", func(c *gin.Context) {

		const minN = 5
		const maxN = 30 

		n := rand.Intn(maxN-minN+1) + minN

		result := fibonacci(n)

		c.JSON(200, gin.H{
			"framework": "Gin",
			"language":  "Go",
			"task":      "Calculate Fibonacci(" + strconv.Itoa(n) + ")", // Menggunakan n yang acak
			"result":    result,
		})
	})

	router.Run(":8080")
}