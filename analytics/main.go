package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	r.Run(":8000")

}
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade error: ", err)
		return
	}
	defer conn.Close()

	for {
		num := rand.Intn(100)
		err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", num)))
		if err != nil {
			fmt.Println("write error: ", err)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
