package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

type Service struct {
	Project string
	Application string
	Lane string
	CIID int
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	services := make([]Service,2,10)

	simpleservice := Service{
		Project: "Project1",
		Application: "App1",
		Lane: "stable",
		CIID: 54321}

	anotherservice := Service{
		Project: "Project2",
		Application: "App2",
		Lane: "stable",
		CIID: 54322}

	services[0] = simpleservice
	services[1] = anotherservice

	b, err := json.Marshal(services)
	if err != nil {
		fmt.Println("error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
