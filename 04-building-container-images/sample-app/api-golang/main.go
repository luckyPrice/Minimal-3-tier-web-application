package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"api-golang/database"
)

func init() {
	errDB := database.InitDB(os.Getenv("DATABASE_URL"))
	if errDB != nil {
		log.Fatalf("⛔ Unable to connect to database: %v\n", errDB)
	} else {
		log.Println("DATABASE CONNECTED 🥇")
	}

}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		database.GetTime(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
