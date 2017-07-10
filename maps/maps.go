package maps

import (
	"encoding/json"
	"net/http"

	"../common"
)

type maps struct {
	Key      string            `json:",omitempty"`
	MapKey   string            `json:",omitempty"`
	MapValue string            `json:",omitempty"`
	Value    map[string]string `json:",omitempty"`
}

var Maps map[string]map[string]string

func init() {
	Maps = make(map[string]map[string]string)
}

func GetValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var m maps
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&m)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if value, ok := Maps[m.Key]; ok {
		common.Render.JSON(w, 200, value)
		return

	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
		return
	}
}

func PostValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var m maps
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&m)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	Maps[m.Key] = m.Value

	common.Render.JSON(w, 201, m)
	return
}

func DeleteValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var m maps
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&m)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if _, ok := Maps[m.Key]; ok {
		delete(Maps, m.Key)
		common.Render.JSON(w, 204, map[string]string{"deleted": m.Key})
		return
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
		return
	}
}
func GetValueByMapKeyHandler(w http.ResponseWriter, r *http.Request) {
	var m maps
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&m)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if value, ok := Maps[m.Key]; ok {
		if _, ok := value[m.MapKey]; ok {
			Maps[m.Key][m.MapKey] = m.MapValue
			common.Render.JSON(w, 200, value)
			return
		} else {
			common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "map key does not exist"})
			return
		}
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
		return
	}
}

func PostMapValueByMapKeyHandler(w http.ResponseWriter, r *http.Request) {
	var m maps
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&m)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	Maps[m.Key][m.MapKey] = m.MapValue

	common.Render.JSON(w, 201, m)
	return
}

func DeletePopValueByKeyHandler(w http.ResponseWriter, r *http.Request) {
	var m maps
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&m)
	if err != nil {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if value, ok := Maps[m.Key]; ok {
		if _, ok := value[m.MapKey]; ok {
			delete(Maps[m.Key], m.MapKey)
			common.Render.JSON(w, 204, map[string]string{"deleted": m.Key})
			return
		} else {
			common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "map key does not exist"})
			return
		}
	} else {
		common.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "key does not exist"})
		return
	}
}
