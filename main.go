package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fmt.Println(params)
	for _, item := range people {
		fmt.Println("ITEM is ", item)
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeoplesEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In GET people")
	json.NewEncoder(w).Encode(people)
}

func CreatePersonDetails(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	// json.NewDecoder(w).En

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {

	fmt.Println("hello REST")
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Mayur", Lastname: "Patange", Address: &Address{City: "Solapur", State: "Maharashtra"}})
	people = append(people, Person{ID: "2", Firstname: "Nikhil", Lastname: "Kondaa", Address: &Address{City: "Solapur", State: "Maharashtra"}})
	router.HandleFunc("/people", GetPeoplesEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonDetails).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", router))
}
