package main

import (
	"awesomeProject/controllers/Game"
	"awesomeProject/controllers/GameWS"
	"awesomeProject/controllers/Movies"
	"awesomeProject/controllers/Songs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	hub := GameWS.NewGameRoomHub()
	go hub.RunLobby()
	router.GET("/movies", Movies.GetMovies)
	router.GET("/movies/:id", Movies.GetMovie)
	router.POST("/movies", Movies.CreateMovies)

	router.POST("/games", Game.CreateMatch)
	router.POST("/games/:id", Game.JoinMatch)
	router.GET("/games", Game.GetAllMatches)
	router.GET("/games/:id",
		func(c *gin.Context) {
			GameWS.ServeWs(&hub, c.Writer, c.Request, c)
		})

	router.GET("/songs", Songs.GetSongs)
	router.GET("/songs/playlist", Songs.GetPlaylist)
	router.GET("/songs/:id", Songs.GetSong)
	router.GET("/songs/upvotes", Songs.GetUpvotes)

	router.PUT("/songs/:id/upvotes", Songs.AddUpvote) //put upvote
	router.PUT("/songs/:id", Songs.AddSong)           //put to playlist

	router.Use(cors.Default())

	router.Run("localhost:8080")

}
