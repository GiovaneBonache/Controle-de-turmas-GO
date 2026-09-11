package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"seuprojeto/model"
)

var salas []model.Sala
var proximoSalaID = 1

func HandleSala(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarSala(w, r)
	case http.MethodPost:
		criarSala(w, r)
	case http.MethodPut:
		atualizarSala(w, r)
	case http.MethodDelete:
		excluirSala(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func listarSala(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salas)
}

func criarSala(w http.ResponseWriter, r *http.Request) {
	var sala model.Sala

	err := json.NewDecoder(r.Body).Decode(&sala)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	sala.ID = proximoSalaID
	proximoSalaID++

	salas = append(salas, sala)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sala)
}

func atualizarSala(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var salaAtualizada model.Sala

	err = json.NewDecoder(r.Body).Decode(&salaAtualizada)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	for i, sala := range salas {
		if sala.ID == id {
			salaAtualizada.ID = id
			salas[i] = salaAtualizada

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(salaAtualizada)
			return
		}
	}

	http.Error(w, "Sala não encontrada", http.StatusNotFound)
}

func excluirSala(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, sala := range salas {
		if sala.ID == id {
			salas = append(salas[:i], salas[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Sala não encontrada", http.StatusNotFound)
}