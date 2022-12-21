package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/taglyscostacurta/api-postgresql/models"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("erro ao fazer update: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if rows > 1 {
		log.Printf("error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"error": false,
		"msg":   fmt.Sprintf("dados atualizados con sucesso!"),
	}

	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
