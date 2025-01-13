package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oGalhardo/EstudandoGo/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
