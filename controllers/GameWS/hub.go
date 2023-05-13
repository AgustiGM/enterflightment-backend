// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package GameWS

import (
	"awesomeProject/entities"
	"encoding/json"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type GameRoomHub struct {
	// Inbound messages from the clients.
	broadcast chan *GameRoomMove

	// Register requests from the clients.
	register chan *GameRoomRegister

	// Unregister requests from clients.
	unregister chan *GameRoomRegister

	hubList map[int]Hub
}

func NewHub() Hub {
	return Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

type GameRoomRegister struct {
	c  Client
	id int
}

type GameRoomMove struct {
	c         Client
	id        int
	newStatus entities.Match
}

func NewGameRoomHub() GameRoomHub {
	return GameRoomHub{
		broadcast:  make(chan *GameRoomMove),
		register:   make(chan *GameRoomRegister),
		unregister: make(chan *GameRoomRegister),
		hubList:    make(map[int]Hub),
	}
}

func (h *GameRoomHub) RunLobby() {
	for {
		select {
		case client := <-h.register:
			h.hubList[client.id].register <- &client.c
		case client := <-h.unregister:
			if _, ok := h.hubList[client.id]; ok {
				delete(h.hubList, client.id)
				close(client.c.send)
			}
		case message := <-h.broadcast:
			aux, _ := json.Marshal(message.newStatus)
			select {
			case message.c.send <- aux:
			default:
				close(message.c.send)
				delete(h.hubList, message.id)
			}

		}
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
