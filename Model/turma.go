package model

type turma struct{
	ID  int   `json:"id"`
	QTDALUNOS  int   `json:"qtdAlunos"`
	ALUNOS  int   `json:"alunos"`
	SALAS  int   `json:"salas"`
}