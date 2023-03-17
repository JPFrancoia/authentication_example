package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	data_registry "local/auth_example/api/data_registy"
	"local/auth_example/api/handlers/auth"
)

func health(c *gin.Context) {

	// Make sure the database is reachable
	err := data_registry.PingDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "unhealthy")
		return
	}

	c.JSON(http.StatusOK, "healthy")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func main() {

	// Crash early if connection to DB fails
	if err := data_registry.InitDB(os.Getenv("DATABASE_URL")); err != nil {
		log.Fatal("Failed to open a database connection: ", err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/health", health)

	router.GET("/", index)
	router.GET("/login", login)

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/:provider", auth.Login)
		authGroup.GET("/callback", auth.AuthCallback)
	}

	router.Run("localhost:8080")
}
