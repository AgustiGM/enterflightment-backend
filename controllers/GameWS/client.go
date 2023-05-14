// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package GameWS

import (
	"awesomeProject/data"
	"awesomeProject/entities"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2048
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	user string
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

var uri string = "mongodb://localhost:27017/"

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var Repo, _ = data.NewMongoRepo(context.TODO(), uri, "enterflight")
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var match entities.Match
		err = json.Unmarshal(message, &match)
		currentState, _ := Repo.GetMatchById(match.ID)
		// check if it's first user
		if currentState.User2 == "" && currentState.User1 != c.user {
			currentState.User2 = c.user
			currentState.Turn = currentState.User1
			Repo.Save(currentState)
		} else if currentState.Turn == c.user {
			if currentState.User1 == c.user {
				currentState.Turn = currentState.User2

			} else {
				currentState.Turn = currentState.User1
			}
			currentState.Board = match.Board
			Repo.Save(currentState)
		} else {

		}

		if err != nil {
			panic("Formatting error in JSON")
		}

		message, _ = json.Marshal(currentState)
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(gameRoomHub *GameRoomHub, w http.ResponseWriter, r *http.Request, ctx *gin.Context) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	username := ctx.Query("username")
	username1 := ctx.Query("t")
	if username == "" {
		username = username1
	}
	val, ok := gameRoomHub.hubList[id]

	var client *Client

	if !ok {
		hub := NewHub()
		go hub.Run()
		client = &Client{hub: &hub, user: username, conn: conn, send: make(chan []byte, 512)}
		gameRoomHub.hubList[id] = hub
		client.hub.register <- client
	} else {
		client = &Client{hub: &val, user: username, conn: conn, send: make(chan []byte, 512)}
		gameRoomHub.hubList[id].register <- client
	}

	go client.writePump()
	go client.readPump()
}
