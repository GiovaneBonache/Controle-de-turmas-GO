package model

type Aluno struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Matricula int    `json:"matricula"`
	Email     string `json:"email"`
	Ativo     bool   `json:"ativo"`
}
