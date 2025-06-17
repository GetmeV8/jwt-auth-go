package main

import (
	"go-jwt-auth/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init (){
	err:=godotenv.Load()
     if err != nil{
		log.Fatal("Error Loading DotEnv")
	 }
}

func main() {
	// Set Gin to debug mode
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world using gin")
	})
	routes.AuthRouter(router)

	router.Run(":8080")
}
