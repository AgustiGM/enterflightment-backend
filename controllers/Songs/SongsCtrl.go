package Songs

import (
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

var songs = []entities.Song{
	{
		ID:     "1",
		Title:  "Call me",
		Artist: "Blondie",
		Length: 10.00,
		Album:  "Call me",
		Year:   "1980",
	},
	{
		ID:     "2",
		Title:  "Bette Davis Eyes",
		Artist: "Kim Carnes",
		Length: 10.00,
		Album:  "Mistaken Identity",
		Year:   "1981",
	},
	{
		ID:     "3",
		Title:  "Physical",
		Artist: "Olive Newton-John",
		Length: 10.00,
		Album:  "Physical",
		Year:   "1982",
	},
	{
		ID:     "4",
		Title:  "Every Breath You Take",
		Artist: "The Police",
		Length: 10.00,
		Album:  "Synchronicity",
		Year:   "1983",
	},
	{
		ID:     "5",
		Title:  "When Doves Cry",
		Artist: "Prince",
		Length: 10.00,
		Album:  "Purple Rain",
		Year:   "1984",
	},
	{
		ID:     "6",
		Title:  "Careless Whisper",
		Artist: "George Michael",
		Length: 10.00,
		Album:  "Make It Big",
		Year:   "1985",
	},
	{
		ID:     "7",
		Title:  "That's What Friends Are For",
		Artist: "Dionne and Friends",
		Length: 10.00,
		Album:  "Friends",
		Year:   "1986",
	},
	{
		ID:     "8",
		Title:  "Walk Like an Egyptian",
		Artist: "The Bangles",
		Length: 10.00,
		Album:  "Different Light",
		Year:   "1987",
	},
	{
		ID:     "9",
		Title:  "Faith",
		Artist: "George Michael",
		Length: 10.00,
		Album:  "Faith",
		Year:   "1988",
	},
	{
		ID:     "10",
		Title:  "Roll With It",
		Artist: "Steve Winwood",
		Length: 10.00,
		Album:  "Roll With It",
		Year:   "1989",
	},
	{
		ID:     "11",
		Title:  "Another Day in Paradise",
		Artist: "Phil Collins",
		Length: 10.00,
		Album:  "But Seriously",
		Year:   "1990",
	},
	{
		ID:     "12",
		Title:  "Unbelievable",
		Artist: "EMF",
		Length: 10.00,
		Album:  "Schubert Dip",
		Year:   "1991",
	},
	{
		ID:     "13",
		Title:  "End of the Road",
		Artist: "Boyz II Men",
		Length: 10.00,
		Album:  "Boomerang",
		Year:   "1992",
	},
	{
		ID:     "14",
		Title:  "I Will Always Love You",
		Artist: "Whitney Houston",
		Length: 10.00,
		Album:  "The Bodyguard",
		Year:   "1993",
	},
	{
		ID:     "15",
		Title:  "The Sign",
		Artist: "Ace of Base",
		Length: 10.00,
		Album:  "Happy Nation",
		Year:   "1994",
	},
	{
		ID:     "16",
		Title:  "Gangsta's Paradise",
		Artist: "Coolio",
		Length: 10.00,
		Album:  "Dangerous Minds",
		Year:   "1995",
	},
	{
		ID:     "17",
		Title:  "Macarena",
		Artist: "Los del Rio",
		Length: 10.00,
		Album:  "A mi me gusta",
		Year:   "1996",
	},
	{
		ID:     "18",
		Title:  "Something About the Way You Look Tonight",
		Artist: "Elton John",
		Length: 10.00,
		Album:  "The Big Picture",
		Year:   "1997",
	},
	{
		ID:     "19",
		Title:  "Too Close",
		Artist: "Next",
		Length: 10.00,
		Album:  "Rated Next",
		Year:   "1998",
	},
	{
		ID:     "20",
		Title:  "Believe",
		Artist: "Cher",
		Length: 10.00,
		Album:  "Believe",
		Year:   "1999",
	},
	{
		ID:     "21",
		Title:  "Breathe",
		Artist: "Faith Hill",
		Length: 10.00,
		Album:  "Breathe",
		Year:   "2000",
	},
	{
		ID:     "22",
		Title:  "Hanging by a Moment",
		Artist: "Lifehouse",
		Length: 10.00,
		Album:  "No Name Face",
		Year:   "2001",
	},
}

// si está buida, poso una random
var Playlist []entities.Song

var Upvotes UpvoteList

func init() {
	Playlist = make([]entities.Song, 0)
	Playlist = append(Playlist, songs[0])
	Playlist = append(Playlist, songs[1])
	Playlist = append(Playlist, songs[2])
	Upvotes = UpvoteList{
		Songlist:   make([]entities.Song, len(Playlist)),
		Upvotelist: make([]int, len(Playlist)),
		Userslist:  make([]string, 0),
	}
}

func GetSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, songs)
}

func GetSong(c *gin.Context) {
	id := c.Param("id")
	for _, s := range songs {
		if s.ID == id {
			c.IndentedJSON(http.StatusOK, s)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "song not found"})
}

func AddSong(c *gin.Context) {
	id := c.Param("id")
	var newSong entities.Song
	for _, s := range songs {
		if s.ID == id {
			newSong = s
			return
		}
	}
	//l'afegim a la llista de sonant actualment
	newSong.ID = id
}

func UpvoteSong(c *gin.Context) {
	id := c.Param("id")
	user := c.Query("username")
	fmt.Println("\n******** user: %s\nid: %s\n", user, id)
	fmt.Println("\n******** Playlist: %s\n", Playlist)
	fmt.Println("\n******** UpvoteList: %s\n", Upvotes)

	upvoteIndex := -1
	for i, s := range Playlist {
		if s.ID == id {
			upvoteIndex = i
			break
		}
	}
	Upvotes.Userslist = append(Upvotes.Userslist, user)
	Upvotes.Upvotelist[upvoteIndex]++

	//for i, s := range songs {
	//	if s.ID == id {
	//		songs[i].Upvotes++
	//		c.IndentedJSON(http.StatusOK, songs[i])
	//		return
	//	}
	//}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
}
