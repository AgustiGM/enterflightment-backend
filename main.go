package main

import (
	"awesomeProject/controllers/Game"
	"awesomeProject/controllers/Movies"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/movies", Movies.GetMovies)
	router.GET("/movies/:id", Movies.GetMovie)
	router.POST("/movies", Movies.CreateMovies)

	router.GET("/", Game.SocketHandler)
	router.Run("localhost:8080")
}
