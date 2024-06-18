package main

import (
    "log"
    "net/http"
	"fmt"
    "backend/handlers"
    "backend/scheduler"
    "github.com/gorilla/mux"
	"backend/middleware"
    "backend/shared"
)

func main() {
    shared.InitJobQueue()

    r := mux.NewRouter()

    r.HandleFunc("/",handlers.Home)
    r.HandleFunc("/jobs", handlers.SubmitJob).Methods("POST")
    r.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
    r.HandleFunc("/ws", handlers.HandleConnections)

    go handlers.HandleMessages()
    go scheduler.ScheduleJobs()

	fmt.Println("Starting server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", middleware.GetCorsConfig().Handler(r)))
}