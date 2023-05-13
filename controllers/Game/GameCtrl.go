package Game

import (
	"awesomeProject/data"
	"awesomeProject/entities"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
)

var lock = &sync.Mutex{}

var upgrader = websocket.Upgrader{
	//check origin will check the cross region source (note : please not using in production)
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

var cnt = 0
var uri string = "mongodb://127.0.0.1:27017/"

func SocketHandler(c *gin.Context) {
	//upgrade get request to websocket protocol
	id, _ := strconv.Atoi(c.Param("id"))
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		var match entities.Match
		err = json.Unmarshal(message, &match)
		if err != nil {
			panic("Formatting error in JSON")
		}

		var cm, _ = Repo.GetMatchById(id)
		cm.Board = match.Board
		if cm.Turn == cm.User1 {
			cm.Turn = cm.User2
		} else {
			cm.Turn = cm.User1
		}
		Repo.Save(cm)

		aux, err := json.Marshal(cm)
		err = ws.WriteMessage(mt, aux)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func CreateMatch(c *gin.Context) {
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	var newMatch entities.Match

	if err := c.BindJSON(&newMatch); err != nil {
		return
	}
	newMatch = Repo.AddMatch(newMatch)
	c.IndentedJSON(http.StatusCreated, newMatch)
}

func JoinMatch(c *gin.Context) {
	var joinMatch entities.Match
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.BindJSON(&joinMatch); err != nil {
		return
	}
	var currentMatch entities.Match
	currentMatch, _ = Repo.GetMatchById(id)
	if currentMatch.User2 == "" && currentMatch.Password == joinMatch.Password {
		currentMatch.User2 = joinMatch.User2
		currentMatch.Board = "---------"
		currentMatch.Turn = currentMatch.User1
		Repo.Save(currentMatch)
	} else {
		panic("Todo")
	}

	// Add the new album to the slice.

	c.IndentedJSON(http.StatusCreated, currentMatch)
}

func GetAllMatches(c *gin.Context) {
	var Repo, _ = data.NewMongoRepo(c, uri, "enterflight")
	var list []entities.Match = Repo.GetAllMatches()
	c.IndentedJSON(http.StatusOK, list)
}
