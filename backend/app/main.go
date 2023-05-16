package main

import (
	"LilaGames/resource"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	gameHandler := resource.NewGameHandler()

	// games API group
	r.GET("/games", gameHandler.GetPopularGameMode)
	r.POST("/games", gameHandler.WriteGameMode)

	port := os.Getenv("PORT")

	err := r.Run(":" + port) // listen and serve on 0.0.0.0:port
	if err != nil {
		fmt.Println(err.Error())
	}
}
