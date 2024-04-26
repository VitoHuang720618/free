package app

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (cm *ConnectionManager) HandleLogin(w http.ResponseWriter, r *http.Request) {
	login(cm.upgrader, w, r)
}

// here ... logic
func login(upgrader websocket.Upgrader, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
