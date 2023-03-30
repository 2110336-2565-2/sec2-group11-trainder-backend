package ws

import (
	"fmt"
	"net/http"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}
type GetRoomsResponse struct {
	Status int       `json:"status"`
	Rooms  []RoomRes `json:"rooms"`
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID      string `json:"id"`
	Trainer string `json:"trainer"`
	Trainee string `json:"trainee"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error")
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Trainer: req.Trainer,
		Trainee: req.Trainee,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		RoomID:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomRes struct {
	ID      string `json:"id"`
	Trainer string `json:"trainer"`
	Trainee string `json:"trainee"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)
	username, err := tokens.ExtractTokenUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, rooms)
	}

	for _, r := range h.hub.Rooms {
		if r.Trainee == username || r.Trainer == username {
			rooms = append(rooms, RoomRes{
				ID:      r.ID,
				Trainer: r.Trainer,
				Trainee: r.Trainee,
			})
		}
	}

	c.JSON(http.StatusOK, GetRoomsResponse{
		Status: http.StatusOK,
		Rooms:  rooms,
	})
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	// fmt.Println("roomId", roomId)
	// _, f := h.hub.Rooms[roomId]
	// fmt.Println("f", f)

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
