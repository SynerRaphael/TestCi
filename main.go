package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Route simple pour Hello World
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":3000") // écoute sur le port 3000
}
