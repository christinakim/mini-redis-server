package strings

import (
	"encoding/json"
	"net/http"

	"../common"
)

type strings struct {
	Key   string `json:",omitempty"`
	Value string `json:",omitempty"`
}

var Strings map[string]string

func init() {
	Strings = make(map[string]string)

}

func GetValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var s strings
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&s)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if value, ok := Strings[s.Key]; ok {
		common.Render.JSON(w, 200, value)
		return
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
		return
	}
}

func PostValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var s strings
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&s)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	Strings[s.Key] = s.Value

	common.Render.JSON(w, 201, s)
	return
}

func DeleteValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var s strings
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&s)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if _, ok := Strings[s.Key]; ok {
		delete(Strings, s.Key)
		common.Render.JSON(w, 204, map[string]string{"deleted": s.Key})
		return
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
		return
	}
}
