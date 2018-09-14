package main

import (
	"encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Ride struct {
	ID int `json:"id,omitempty"`
	Departfrom string `json:departfrom,omitempty`
	Arriveat string `json:arriveat,omitempty`
	Date string `json:date,omitempty`
	Time string `jaom:time,omitempty`
}

var rides []Ride

func GetAllRides(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rides)
}

func main() {
	router := mux.NewRouter()
	rides = append(rides, Ride{ID: 1, Departfrom: "Yaba", Arriveat: "Ikoyi", Date: "27/07/2018", Time: "4:30"})
	rides = append(rides, Ride{ID: 2, Departfrom: "Ikoyi", Arriveat: "Yaba", Date: "27/07/2018", Time: "4:30"})
	rides = append(rides, Ride{ID: 3, Departfrom: "Lekki 1", Arriveat: "Ikeja", Date: "27/07/2018", Time: "4:30"})

	router.HandleFunc("/rides", GetAllRides).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
