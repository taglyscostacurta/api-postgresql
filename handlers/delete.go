package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/taglyscostacurta/api-postgresql/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if rows > 1 {
		log.Printf("error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"error": false,
		"msg":   "Registro removido com sucesso!",
	}

	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
