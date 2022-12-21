package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/taglyscostacurta/api-postgresql/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("erro ao fazer update: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(todo)

}
