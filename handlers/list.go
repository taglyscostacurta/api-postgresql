package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/taglyscostacurta/api-postgresql/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("erro ao buscar todos os registros: %v", err)
	}

	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
