package handlers

import (
	"backend/models"
	"backend/shared"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(ws)
    defer ws.Close()
    shared.Clients[ws] = true

    for {
        var msg models.Job
        if err := ws.ReadJSON(&msg); err != nil {
            log.Printf("error: %v", err)
            delete(shared.Clients, ws)
            break
        }
    }
}

func HandleMessages() {
    for {
        job := <-shared.Broadcast
        for client := range shared.Clients {
            err := client.WriteJSON(job)
            if err != nil {
                log.Printf("error: %v", err)
                client.Close()
                delete(shared.Clients, client)
            }
        }
    }
}
