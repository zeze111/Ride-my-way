package main

import (
	"encoding/json"
	"fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAllRides(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(rides)
}

func CreateRide(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var ride Ride
	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	ride.ID = bson.NewObjectId()
	if err := model.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, ride)
}

func main() {
	router := mux.NewRouter()
	// rides = append(rides, Ride{ID: 1, Departfrom: "Yaba", Arriveat: "Ikoyi", Date: "27/07/2018", Time: "4:30"})
	// rides = append(rides, Ride{ID: 2, Departfrom: "Ikoyi", Arriveat: "Yaba", Date: "27/07/2018", Time: "4:30"})
	// rides = append(rides, Ride{ID: 3, Departfrom: "Lekki 1", Arriveat: "Ikeja", Date: "27/07/2018", Time: "4:30"})

	router.HandleFunc("/rides", GetAllRides).Methods("GET")
	router.HandleFunc("/ride", CreateRide).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
