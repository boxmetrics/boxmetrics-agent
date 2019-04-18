package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws/v1", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/websocket.html")
	})

	var httpErr error

	if _, err := os.Stat("./certificates/server.crt"); err == nil {
		fmt.Println("Server running at https://localhost:9090")
		httpErr = http.ListenAndServeTLS(":9090", "certificates/server.crt", "certificates/server.key", nil)

		if httpErr != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	} else {
		fmt.Println("Server running at http://localhost:8080")
		httpErr = http.ListenAndServe(":8080", nil)

		if httpErr != nil {
			log.Fatal("ListenAndServe: ", err)
		}

	}

}
