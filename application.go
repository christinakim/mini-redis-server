package main

import (
	"log"
	"net/http"

	"./lists"
	"./maps"
	"./strings"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello anson!\n"))
}

func GetValueFromKeyHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", IndexHandler)
	//search handlers
	/*
		SEARCH-KEYS query: Returns a List<String> containing all keys matching the query.
					You can decide the structure of query and how it is evaluated to search for keys.
					The intent of SEARCH-KEYS is to give clients the ability to interact with groups of values by embedding semantic information in their keys.
					For example, clients might use keys like login-attempt:<username>:<attempt-number> to store failed login attempts, and would use SEARCH-KEYS to determine how many failed attempts had been made against a specific username.

			r.HandleFunc("/api/v1/search/{search}", search.GetSearchHandler).Methods("GET")
	*/

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
	/*
		maps
			GET key: Return the Map value identified by key
			SET key value: Instantiate or overwrite a Map identified by key with value value
			DELETE key: Delete the Map identified by key
			MAPGET key mapkey: Return the String identified by mapkey from within the Map identified by key.
			MAPSET key mapkey mapvalue: Add the mapping mapkey -> mapvalue to the Map identified by key.
			MAPDELETE key mapkey: Delete the value identified by mapkey from the Map identified by key.
	*/

	r.HandleFunc("/api/v1/maps/get", maps.GetValueByKeyHandler).Methods("GET")
	r.HandleFunc("/api/v1/maps/set", maps.PostValueByKeyHandler).Methods("POST")
	r.HandleFunc("/api/v1/maps/delete", maps.DeleteValueByKeyHandler).Methods("DELETE")
	r.HandleFunc("/api/v1/maps/mapget", maps.GetValueByMapKeyHandler).Methods("GET")
	r.HandleFunc("/api/v1/maps/mapset", maps.PostMapValueByMapKeyHandler).Methods("POST")
	r.HandleFunc("/api/v1/maps/mapdelete", maps.DeletePopValueByKeyHandler).Methods("DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
