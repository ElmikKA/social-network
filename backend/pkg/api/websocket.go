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
	connections = make(map[*websocket.Conn]models.Connection)
	mu          sync.Mutex
)

var broadcast = make(chan models.Message)

func (h *Handler) Websocket(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
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
		h.CloseSocket(conn)
	}()

	// change status to online

	fmt.Println("changing status to online")

	h.store.GoOnline(h.id)

	// send logged in message to every connected user
	mu.Lock()
	for msgConn := range connections {
		onlineResponse := make(map[string]interface{})
		onlineResponse["statusChange"] = true
		onlineResponse["id"] = h.id
		onlineResponse["username"] = h.username
		onlineResponse["online"] = 1

		if err := msgConn.WriteJSON(onlineResponse); err != nil {
			fmt.Println("error writing onlineresponse message", err)
		}
	}

	// adds new connection to the connections map
	connections[conn] = models.Connection{
		Id:       h.id,
		Username: h.username,
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
		msg.UserId = connections[conn].Id

		fmt.Println("got a readJSON")
		fmt.Println("message:", msg.Message)
		fmt.Println("receiverId:", msg.ReceiverId)
		fmt.Println(msg)
		broadcast <- msg
	}
}

func (h *Handler) HandleWebsocketConnections() {
	for msg := range broadcast {
		fmt.Println("got a broadcast")
		fmt.Println(msg.ReceiverId)
		if msg.ReceiverId != 0 {
			h.sendPrivateMessage(msg)
		} else {
			h.sendGroupMessage(msg)
		}
	}
}
func (h *Handler) sendGroupMessage(msg models.Message) {
	fmt.Println(msg)
	fmt.Println("group message")

	mu.Lock()
	defer mu.Unlock()

	// add to db
	err := h.store.AddMessage(msg)
	if err != nil {
		fmt.Println("err addin group message", err)
		return
	}

	// send message to all group members
	responseData := make(map[string]interface{})
	onlineMembers, err := h.store.GetOnlineGroupMembers(msg.UserId)
	fmt.Println(onlineMembers)
	if err != nil {
		fmt.Println("err getting online group members", err)
		return
	}
	for msgConn, value := range connections {
		if utils.ContainsInt(onlineMembers, value.Id) {
			// send message
			responseData["response"] = msg.Message
			responseData["username"] = value.Username
			if err := msgConn.WriteJSON(responseData); err != nil {
				fmt.Println("error writing onlineresponse message", err)
			}
		}
	}
}

func (h *Handler) sendPrivateMessage(msg models.Message) {
	fmt.Println("private message")
	fmt.Println(msg)
	mu.Lock()
	defer mu.Unlock()
	// add to db

	err := h.store.AddMessage(msg)
	if err != nil {
		fmt.Println("error adding message", err)
		return
	}

	// send message to receiver

	responseData := make(map[string]interface{})
	for msgConn, value := range connections {
		if value.Id == msg.ReceiverId {
			// send message
			responseData["response"] = msg.Message
			responseData["id"] = value.Id
			responseData["username"] = value.Username
			if err := msgConn.WriteJSON(responseData); err != nil {
				fmt.Println("error writing onlineresponse message", err)
			}
			return
		}
	}
	fmt.Println("receiver not online")
}

func (h *Handler) CloseSocket(conn *websocket.Conn) {
	// delete connection from connections
	mu.Lock()
	defer mu.Unlock()

	delete(connections, conn)
	conn.Close()
	user := connections[conn]

	// change online status
	h.store.GoOffline(user.Id)

	// send every connection a status change
	for msgConn := range connections {
		onlineResponse := make(map[string]interface{})
		onlineResponse["statusChange"] = true
		onlineResponse["id"] = user.Id
		onlineResponse["username"] = user.Username
		onlineResponse["online"] = -1

		if err := msgConn.WriteJSON(onlineResponse); err != nil {
			fmt.Println("error writing closesocket message", err)
		}
	}
}
