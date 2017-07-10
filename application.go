package main

import (
	"log"
	"net/http"

	"./lists"
	"./maps"
	"./search"
	"./strings"

	"../../../github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello anson!\n"))
}

func main() {
	r := mux.NewRouter()

	// routes consist of a path and a handler function.
	r.HandleFunc("/", IndexHandler)

	//search handlers
	r.HandleFunc("/api/v1/search/{query}", search.GetSearchHandler).Methods("GET")

	//string handlers
	r.HandleFunc("/api/v1/strings/get", strings.GetValueByKeyHandler).Methods("GET")
	r.HandleFunc("/api/v1/strings/set", strings.PostValueByKeyHandler).Methods("POST")
	r.HandleFunc("/api/v1/strings/delete", lists.DeleteValueByKeyHandler).Methods("DELETE")

	//list handlers
	r.HandleFunc("/api/v1/lists/get", lists.GetValueByKeyHandler).Methods("GET")
	r.HandleFunc("/api/v1/lists/set", lists.PostValueByKeyHandler).Methods("POST")
	r.HandleFunc("/api/v1/lists/delete", lists.DeleteValueByKeyHandler).Methods("DELETE")
	r.HandleFunc("/api/v1/lists/append", lists.PostAppendValueByKeyHandler).Methods("POST")
	r.HandleFunc("/api/v1/lists/pop", lists.DeletePopValueByKeyHandler).Methods("DELETE")

	// map handlers
	r.HandleFunc("/api/v1/maps/get", maps.GetValueByKeyHandler).Methods("GET")
	r.HandleFunc("/api/v1/maps/set", maps.PostValueByKeyHandler).Methods("POST")
	r.HandleFunc("/api/v1/maps/delete", maps.DeleteValueByKeyHandler).Methods("DELETE")
	r.HandleFunc("/api/v1/maps/mapget", maps.GetValueByMapKeyHandler).Methods("GET")
	r.HandleFunc("/api/v1/maps/mapset", maps.PostMapValueByMapKeyHandler).Methods("POST")
	r.HandleFunc("/api/v1/maps/mapdelete", maps.DeletePopValueByKeyHandler).Methods("DELETE")

	// Bind to a port and pass our router in
	log.Println("init finished, listening.....")
	log.Fatal(http.ListenAndServe(":8000", r))
}
