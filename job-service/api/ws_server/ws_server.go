package wsserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type BroadcastMessage struct {
	MessageType string
	Data        interface{}
}

type WebSocketServer struct {
	upgrader  *websocket.Upgrader
	clients   map[*websocket.Conn]bool
	broadcast chan BroadcastMessage
}

func NewWebSocketServer(upgrader *websocket.Upgrader, clients map[*websocket.Conn]bool, broadcast chan BroadcastMessage) *WebSocketServer {
	return &WebSocketServer{upgrader: upgrader, clients: clients, broadcast: broadcast}
}

func (wss *WebSocketServer) HandleWebSocket(c *gin.Context) {
	ws, err := wss.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	wss.clients[ws] = true

	for {
		var msg BroadcastMessage

		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error: %v", err)
			delete(wss.clients, ws)
			break
		}

		wss.broadcast <- msg
	}
}

func (wss *WebSocketServer) BroadcastMessage() {
	for {
		msg := <-wss.broadcast

		for client := range wss.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error: %v", err)
				client.Close()
				delete(wss.clients, client)
			}
		}
	}
}
