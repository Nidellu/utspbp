package main

import (
	"fmt"
	"log"
	"net/http"
	"uts/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rooms", controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/detailrooms", controller.GetDetailRooms).Methods("GET")
	router.HandleFunc("/participants", controller.InsertParticipants).Methods("POST")
	router.HandleFunc("/delparticipants", controller.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
