package model

type Turma struct {
	ID         int       `json:"id"`
	Nome       string    `json:"nome"`
	Disciplina string    `json:"disciplina"`
	Professor  string    `json:"professor"`
	Alunos     []int     `json:"alunos"`
	Alocacao   *Alocacao `json:"alocacao,omitempty"`
	Ativo      bool      `json:"ativo"`
}
