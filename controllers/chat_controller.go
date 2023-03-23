package controllers

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
type message struct {
	Room string `json:"room"`
	Name string `json:"name"`
	Body string `json:"body"`
}
type room struct {
	sync.Mutex
	connections []*websocket.Conn
}
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
var rooms = make(map[string]*room)
func HandleChat() gin.HandlerFunc { 
	return func(c *gin.Context) {
		w:= c. Writer
		r:= c. Request
		conn,err := upgrader.Upgrade(w,r,nil)
		if err != nil{
			fmt.Println(err)
			return
		}
		msg  := message{}
		err  = conn.ReadJSON(&msg)
		if err != nil{
			fmt.Println(err)
			return
		}
		room := getRoom(msg.Room)
		room.Lock()
		room.connections = append(room.connections,conn)
		room.Unlock()
		broadcast(msg)
	} 
}

func getRoom(name string)*room {
	r,ok := rooms[name]
	if !ok {
		r = &room{}
		rooms[name] = r
	}
	return r
}

func broadcast(msg message) {
    r := getRoom(msg.Room)
    r.Lock()
    defer r.Unlock()
    for _, conn := range r.connections {
        err := conn.WriteJSON(msg)
        if err != nil {
            fmt.Println(err)
        }
    }
}
