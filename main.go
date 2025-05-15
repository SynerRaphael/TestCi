package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Route simple pour Hello World
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Holàààààà")
	})

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Erreur lors du paramétrage des proxies : %v", err)
	}

	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur : %v", err)
	}
}
