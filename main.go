package main

import (
	"awesomeProject/controllers/Game"
	"awesomeProject/controllers/GameWS"
	"awesomeProject/controllers/Movies"
	"awesomeProject/controllers/Songs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	router := gin.Default()
	hub := GameWS.NewGameRoomHub()
	go hub.RunLobby()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	router.GET("/video/:id", ServeVideo)

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
	router.GET("/songs/:id/file", ServeSong)

	router.PUT("/songs/:id/upvotes", Songs.AddUpvote) //put upvote
	router.PUT("/songs/:id", Songs.AddSong)           //put to playlist

	router.Run("0.0.0.0:8080")

}

func ServeVideo(c *gin.Context) {
	videoId := c.Param("id")
	videoPath := "resources"
	if videoId == "1" {
		videoPath += "/video1.mp4"
	} else {
		videoPath += "/Video.mp4"
	}

	videoData, err := ioutil.ReadFile(videoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read video file"})
		return
	}

	c.Data(http.StatusOK, "video/mp4", videoData)
}

func ServeSong(c *gin.Context) {
	songId := c.Param("id")
	var songPath string
	switch {
	case songId == "1":
		songPath = "resources/song1.mp3"
		break
	case songId == "2":
		songPath = "resources/song2.mp3"
		break
	case songId == "3":
		songPath = "resources/song3.mp3"
		break
	case songId == "4":
		songPath = "resources/song4.mp3"
		break
	default:
		songPath = "resources/song1.mp3"
	}

	songData, err := ioutil.ReadFile(songPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read video file"})
		return
	}

	c.Data(http.StatusOK, "audio/mpeg", songData)
}
