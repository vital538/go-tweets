package main

import (
	"go-tweets/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

}
