package lists

import (
	"encoding/json"
	"net/http"

	"../common"
)

type lists struct {
	Key    string   `json:",omitempty"`
	Value  []string `json:",omitempty"`
	Append string
}

var Lists map[string][]string

func init() {
	Lists = make(map[string][]string)

}

func GetValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var l lists
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&l)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if value, ok := Lists[l.Key]; ok {
		common.Render.JSON(w, 200, value)
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
	}
}

func PostValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var l lists
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&l)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	Lists[l.Key] = l.Value

	common.Render.JSON(w, 201, l)
}

func DeleteValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var l lists
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&l)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if _, ok := Lists[l.Key]; ok {
		delete(Lists, l.Key)
		common.Render.JSON(w, 204, map[string]string{"deleted": l.Key})
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
	}
}

func PostAppendValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var l lists
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&l)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if value, ok := Lists[l.Key]; ok {
		//reassigning value from key with appended string value from post body
		Lists[l.Key] = append(value, l.Append)
		common.Render.JSON(w, 204, map[string]string{"deleted": l.Key})
		return
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
		return
	}
}

func DeletePopValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var l lists
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&l)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if value, ok := Lists[l.Key]; ok {
		//popping off last element of slice
		value = value[:len(value)-1]

		//reassigning to new value
		Lists[l.Key] = value
		common.Render.JSON(w, 204, map[string]string{"deleted": l.Key})
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
	}
}
