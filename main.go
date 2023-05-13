package main

import (
	"awesomeProject/controllers/Game"
	"awesomeProject/controllers/Movies"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	router.GET("/movies", Movies.GetMovies)
	router.GET("/movies/:id", Movies.GetMovie)
	router.POST("/movies", Movies.CreateMovies)
	router.POST("/games", Game.CreateMatch)
	router.POST("/games/:id", Game.JoinMatch)
	router.GET("/games", Game.GetAllMatches)
	router.GET("/", Game.SocketHandler)
	router.Run("localhost:8080")
}
