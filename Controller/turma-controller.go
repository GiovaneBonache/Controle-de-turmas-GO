package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"seuprojeto/model"
)

var turmas []model.Turma
var proximoTurmaID = 1

func HandleTurma(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarTurma(w, r)
	case http.MethodPost:
		criarTurma(w, r)
	case http.MethodPut:
		atualizarTurma(w, r)
	case http.MethodDelete:
		excluirTurma(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func listarTurma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(turmas)
}

func criarTurma(w http.ResponseWriter, r *http.Request) {
	var turma model.Turma

	err := json.NewDecoder(r.Body).Decode(&turma)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	turma.ID = proximoTurmaID
	proximoTurmaID++

	turmas = append(turmas, turma)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(turma)
}

func atualizarTurma(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var turmaAtualizada model.Turma

	err = json.NewDecoder(r.Body).Decode(&turmaAtualizada)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	for i, turma := range turmas {
		if turma.ID == id {
			turmaAtualizada.ID = id
			turmas[i] = turmaAtualizada

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(turmaAtualizada)
			return
		}
	}

	http.Error(w, "Turma não encontrada", http.StatusNotFound)
}

func excluirTurma(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, turma := range turmas {
		if turma.ID == id {
			turmas = append(turmas[:i], turmas[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Turma não encontrada", http.StatusNotFound)
}