package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	//"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type event struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type allEvents []event

var events = allEvents{
	{
		Id:    "1",
		Title: "Title",
		Desc:  "Desc",
	},
}

/*type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}*/

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)

	fmt.Println(events)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	for _, event := range events {
		if eventId == event.Id {
			json.NewEncoder(w).Encode(event)
		}
	}
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	var updatedEvent event

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(requestBody, &updatedEvent)

	for i, event := range events {
		if event.Id == eventId {
			event.Desc = updatedEvent.Desc
			event.Title = updatedEvent.Title
			events = append(events[:i], event)
			json.NewEncoder(w).Encode(event)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.Id == eventID {
			events = append(events[:i], events[i+1:]...)
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
