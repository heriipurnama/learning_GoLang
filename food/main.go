package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// roll is model for food API
type Roll struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredient  string `json:"ingredient"`
}

// init rolls var as a slice
var rolls []Roll

// show all data
func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)

}

// show single data
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // function buat setting prameter http buat di parsing
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

// add single data
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll                          // data dari http di simpan disini
	json.NewDecoder(r.Body).Decode(&newRoll)  // req. dr http req.
	newRoll.ID = strconv.Itoa(len(rolls) + 1) // *helper : create auto increment ID
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

// update single data
func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...) // diguakn buat hapus dan menambahkan data baru
			/*
			   append(rolls[:i] ==> untuk mencari objek yang akan di update
			   rolls[i+1:]...) ==> untuk mencari objek yang akan di append
			*/
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}

}

// delete single sushi
func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...) // delete data
			break
		}
	}
}

func main() {
	// Generate Mock Data
	rolls = append(rolls,
		Roll{
			ID:          "1",
			Name:        "Salmon",
			Description: "crab stick, tamago sushi with salmon and cheese",
			Ingredient:  "salmon, nuri, soy sauce, rice",
		},
		Roll{
			ID:          "2",
			Name:        "california roll",
			Description: "crab stick, tamago sushi with salmon and cheese",
			Ingredient:  "salmon, nuri, soy sauce, rice",
		},
	)

	// init router
	router := mux.NewRouter()

	//Handle Endpoints/Routing
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("PUT")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	// connection
	log.Fatal(http.ListenAndServe(":8000", router))
}
