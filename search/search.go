package search

import (
	"net/http"
	"regexp"

	"../common"
	"../lists"
	"../maps"
	"../strings"

	"github.com/gorilla/mux"
)

func GetSearchHandler(w http.ResponseWriter, r *http.Request) {
	// getting search term
	vars := mux.Vars(r)
	query := vars["query"]

	//creating regex that will return a match if query term exists in the string
	regex, err := regexp.Compile(".*" + query + " .*")

	if err != nil {
		common.Render.JSON(w, 500, map[string]string{"error": err.Error()})
	}

	var results []string
	for k, _ := range strings.Strings {
		if regex.MatchString(k) {
			results = append(results, k)
		}
	}
	for k, _ := range lists.Lists {
		if regex.MatchString(k) {
			results = append(results, k)
		}
	}

	for k, _ := range maps.Maps {
		if regex.MatchString(k) {
			results = append(results, k)
		}
	}

	common.Render.JSON(w, 200, results)

}
