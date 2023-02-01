package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wUpgrader = websocket.Upgrader{} // use default options

var serverAddr = flag.String("addr", "localhost:8080", "http service address")

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := wUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("server recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}
