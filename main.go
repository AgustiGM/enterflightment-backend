package main

import (
	"awesomeProject/controllers/Game"
	"awesomeProject/controllers/Movies"
	"awesomeProject/controllers/Songs"
	"github.com/gin-contrib/cors"
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

	router.GET("/songs", Songs.GetSongs)
	router.GET("/songs/:id", Songs.GetSong)
	router.PUT("/songs/:id", Songs.AddSong)
	router.PUT("/songs/:id/upvote", Songs.UpvoteSong)

	router.Use(cors.Default())

	router.Run("localhost:8080")
	//go Movies.HttpServer("localhost:8082")

}
