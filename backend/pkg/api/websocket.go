package api

import (
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	connections = make(map[*websocket.Conn]map[string]interface{})
	mu          sync.Mutex
)
var broadcast = make(chan interface{})

func (api *APIServer) Websocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in websocket")

	// upgrade to websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error upgrading to websocket", err)
		return
	}

	// add defer to close the websocket after use
	defer func() {
		fmt.Println("closing websocket")
		api.CloseSocket(conn)
	}()

	// change status to online

	fmt.Println("changing status to online")

	utils.GoOnline(api.db, api.id)

	// send logged in message to every connected user
	mu.Lock()
	for msgConn := range connections {
		onlineResponse := make(map[string]interface{})
		onlineResponse["statusChange"] = true
		onlineResponse["id"] = api.id
		onlineResponse["username"] = api.username
		onlineResponse["online"] = 1

		if err := msgConn.WriteJSON(onlineResponse); err != nil {
			fmt.Println("error writing onlineresponse message", err)
		}
	}

	// adds new connection to the connections map
	connections[conn] = map[string]interface{}{
		"id":       api.id,
		"username": api.username,
	}
	mu.Unlock()

	// after websocket has been set up start listening for messages
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("err reading websocket message", err)
			return
		}
		fmt.Println("got a readJSON")
		fmt.Println("message:", msg.Message)
		fmt.Println("receiverId:", msg.ReceiverId)
		broadcast <- msg
	}
}

func (api *APIServer) HandleWebsocketConnections() {
	for msg := range broadcast {
		fmt.Println("got a broadcast")
		sendPrivateMessage(msg)
	}
}

func sendPrivateMessage(msg interface{}) {
	mu.Lock()
	defer mu.Unlock()
	// add to db

	// send message to receiver
	// sends to every connection right now
	for msgConn, value := range connections {
		responseData := make(map[string]interface{})
		responseData["response"] = "hello world"
		responseData["id"] = value["id"]
		responseData["username"] = value["username"]

		if err := msgConn.WriteJSON(responseData); err != nil {
			fmt.Println("error writing onlineresponse message", err)
		}
	}
}

func (api *APIServer) CloseSocket(conn *websocket.Conn) {
	// delete connection from connections
	mu.Lock()
	defer mu.Unlock()

	delete(connections, conn)
	conn.Close()

	// change online status
	utils.GoOffline(api.db, api.id)

	// send every connection a status change
	for msgConn := range connections {
		onlineResponse := make(map[string]interface{})
		onlineResponse["statusChange"] = true
		onlineResponse["id"] = api.id
		onlineResponse["username"] = api.username
		onlineResponse["online"] = -1

		if err := msgConn.WriteJSON(onlineResponse); err != nil {
			fmt.Println("error writing closesocket message", err)
		}
	}
}