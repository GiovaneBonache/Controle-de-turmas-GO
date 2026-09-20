package model

type TurmaResumo struct {
	ID               int    `json:"id"`
	Nome             string `json:"nome"`
	Disciplina       string `json:"disciplina"`
	Professor        string `json:"professor"`
	QuantidadeAlunos int    `json:"quantidade_alunos"`
	Alocada          bool   `json:"alocada"`
	Ativo            bool   `json:"ativo"`
}
