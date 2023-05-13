package main

import (
	"awesomeProject/controllers/Movies"
	"awesomeProject/controllers/Songs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Global string
var GlobalList []string
var Global3 = "myvalue3"

func init() {
	Global, GlobalList = InitApp()
}

func InitApp() (string, []string) {
	return "myvalue", []string{"one", "two", "three"}
}

func main() {

	router := gin.Default()
	router.GET("/movies", Movies.GetMovies)
	router.GET("/movies/:id", Movies.GetMovie)
	router.POST("/movies", Movies.CreateMovies)

	router.GET("/songs", Songs.GetSongs)
	router.GET("/songs/:id", Songs.GetSong)
	router.PUT("/songs/:id", Songs.AddSong)
	router.GET("/songs/:id/upvote", Songs.UpvoteSong)

	router.Use(cors.Default())

	router.Run("localhost:8080")
	//go Movies.HttpServer("localhost:8082")

}

//func main() {
//	http.Handle("/", new(viewHandler))
//	http.ListenAndServe(":8080", nil)
//}
//
//func (vh *viewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	path := r.URL.Path[1:]
//	log.Println(path)
//	data, err := ioutil.ReadFile(string(path))
//
//	if err == nil {
//
//		var contentType string
//
//		if strings.HasSuffix(path, ".html") {
//			contentType = "text/html"
//		} else if strings.HasSuffix(path, ".css") {
//			contentType = "text/css"
//		} else if strings.HasSuffix(path, ".js") {
//			contentType = "application/javascript"
//		} else if strings.HasSuffix(path, ".png") {
//			contentType = "image/png"
//		} else if strings.HasSuffix(path, ".jpg") {
//			contentType = "image/jpeg"
//		}
//
//		w.Header().Add("Content-Type", contentType)
//		w.Write(data)
//	} else {
//		log.Println("ERROR!")
//		w.WriteHeader(404)
//		w.Write([]byte("404 - " + http.StatusText(404)))
//	}
//}
