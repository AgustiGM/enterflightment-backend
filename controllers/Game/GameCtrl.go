package Game

import (
	"awesomeProject/entities"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	//check origin will check the cross region source (note : please not using in production)
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

func SocketHandler(c *gin.Context) {
	//upgrade get request to websocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		//Read Message from client
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		var match entities.Match

		err = json.Unmarshal([]byte(message), &match)
		if err != nil {
			// panic
		}
		//If client message is ping will return pong

		message = []byte("pong")
		//Response message to client
		err = ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
