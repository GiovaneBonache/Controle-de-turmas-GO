package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"seuprojeto/model"
)

var alunos []model.Aluno
var proximoAlunoID = 1

func HandleAluno(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarAluno(w, r)
	case http.MethodPost:
		criarAluno(w, r)
	case http.MethodPut:
		atualizarAluno(w, r)
	case http.MethodDelete:
		excluirAluno(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func listarAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

func criarAluno(w http.ResponseWriter, r *http.Request) {
	var aluno model.Aluno

	err := json.NewDecoder(r.Body).Decode(&aluno)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	aluno.ID = proximoAlunoID
	proximoAlunoID++

	alunos = append(alunos, aluno)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aluno)
}

func atualizarAluno(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var alunoAtualizado model.Aluno

	err = json.NewDecoder(r.Body).Decode(&alunoAtualizado)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	for i, aluno := range alunos {
		if aluno.ID == id {
			alunoAtualizado.ID = id
			alunos[i] = alunoAtualizado

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(alunoAtualizado)
			return
		}
	}

	http.Error(w, "Aluno não encontrado", http.StatusNotFound)
}

func excluirAluno(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, aluno := range alunos {
		if aluno.ID == id {
			alunos = append(alunos[:i], alunos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Aluno não encontrado", http.StatusNotFound)
}