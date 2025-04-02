package main

import (
	"gin/config"
	"gin/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Setup()

	config.Routes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
