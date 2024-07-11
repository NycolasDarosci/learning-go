package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	allExhibitions := data.GetAll()
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	if id != "" {
		finalId, err := strconv.Atoi(id)
		if err == nil && finalId < len(allExhibitions) {
			json.NewEncoder(w).Encode(allExhibitions[finalId])
		} else {
			http.Error(w, "Out of index", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(allExhibitions)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	var exhibition data.Exhibition
	err := json.NewDecoder(r.Body).Decode(&exhibition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	exhibition.Add()
	w.Write([]byte("Post succesfully!"))
}
