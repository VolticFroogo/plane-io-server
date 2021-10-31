package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: actually check origin in production.
		return true
	},
	EnableCompression: true,
}

func handleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	player := &Player{
		Conn:          conn,
		Buffer:        NewBuffer(State{}),
		KnownEntities: make(map[uint32]bool),
	}

	defer player.Despawn()

	player.ID = server.NextEntityID
	server.NextEntityID++

	if server.EntityMap[0] == nil {
		server.EntityMap[0] = player
	} else {
		next := server.EntityMap[0]
		for next.GetNext() != nil {
			next = next.GetNext()
		}
		next.SetNext(player)
		player.SetPrev(next)
	}

	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
				return
			}

			log.Println(err)
			return
		}

		HandleMessage(player, bytes)
	}
}
