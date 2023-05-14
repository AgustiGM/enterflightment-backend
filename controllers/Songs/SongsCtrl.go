package Songs

import (
	"awesomeProject/data"
	"awesomeProject/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cuan s'activa el upvote --> ini de SongList
// cuan un user fa upvote -->
//
//	|--> si no està a la llista, l'afegeixo // si está no el deixo votar
//	|--> poso un +1 a la posició corresponent de la canço en UpvoteList
type UpvoteList struct {
	Songlist   []entities.Song
	Upvotelist []int
	Userslist  []string
}

// si está buida, poso una random
var Playlist []entities.Song

var Upvotes UpvoteList

//func init() {
//	Playlist = make([]entities.Song, 0)
//	Playlist = append(Playlist, songs[0])
//	Playlist = append(Playlist, songs[1])
//	Playlist = append(Playlist, songs[2])
//	Upvotes = UpvoteList{
//		Songlist:   make([]entities.Song, len(Playlist)),
//		Upvotelist: make([]int, len(Playlist)),
//		Userslist:  make([]string, 0),
//	}
//}

var uri string = "mongodb://127.0.0.1:27017/"

func GetSongs(c *gin.Context) {
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	result, err := Repo.GetAllSongs()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "songs not found"})
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func GetSong(c *gin.Context) {
	id := c.Param("id")
	result, err := data.GetSongById(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "song not found"})
	}
}

func GetPlaylist(c *gin.Context) {
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	result, err := Repo.GetPlaylist()
	if err == nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "song not found"})
	}
}

func GetUpvotes(c *gin.Context) {
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	result, err := Repo.GetUpvotes()
	if err == nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Upvotes found"})
	}
}

func AddUpvote(c *gin.Context) {
	id := c.Param("id")
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")

	err := Repo.AddUpvote(id)
	if err == nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Upvote added"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "NOT UPVOTED"})
	}
}

func AddSong(c *gin.Context) {
	id := c.Param("id")
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")

	result, err := Repo.AddSongToPlaylist(id)
	fmt.Printf("\n******** result: %s\n", result)
	if err == nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "song not found"})
	}
}

func EmptyPlaylist(c *gin.Context) {
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	err := Repo.EmptyPlaylist()
	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Playlist emptied"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Playlist not emptied"})
	}
}

func PrepareNextSong(c *gin.Context) {
	var Repo, _ = data.NewMongoRepo(nil, uri, "enterflight")
	err1, err2 := Repo.PrepareNextSong()
	if err1 == nil && err2 == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "next song ready"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "something failed"})
	}
}
