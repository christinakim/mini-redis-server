package search

import (
	"net/http"
	"regexp"

	"../common"
	"../lists"
	"../maps"
	"../strings"

	"../../../../github.com/gorilla/mux"
)

type searchResults struct {
	String []string
	List   []string
	Map    []string
}

func GetSearchHandler(w http.ResponseWriter, r *http.Request) {
	// getting search term
	vars := mux.Vars(r)
	query := vars["query"]

	//creating regex that will return a match if query term exists in the string
	regex, err := regexp.Compile(`.*` + query + `.*`)

	if err != nil {
		common.Render.JSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	var results searchResults
	for k, _ := range strings.Strings {
		if regex.MatchString(k) {
			results.String = append(results.String, k)
		}
	}
	for k, _ := range lists.Lists {
		if regex.MatchString(k) {
			results.List = append(results.List, k)
		}
	}

	for k, _ := range maps.Maps {
		if regex.MatchString(k) {
			results.Map = append(results.Map, k)
		}
	}

	common.Render.JSON(w, 200, results)
	return
}
