package service

import (
	"errors"
	"seuprojeto/model"
)

var alunos []model.Aluno
var proximoAlunoID = 1

func ListarAlunos() []model.Aluno {
	return alunos
}

func CriarAluno(aluno model.Aluno) model.Aluno {
	aluno.ID = proximoAlunoID
	proximoAlunoID++

	alunos = append(alunos, aluno)

	return aluno
}

func AtualizarAluno(id int, alunoAtualizado model.Aluno) (model.Aluno, error) {
	for i, aluno := range alunos {
		if aluno.ID == id {
			alunoAtualizado.ID = id
			alunos[i] = alunoAtualizado

			return alunoAtualizado, nil
		}
	}

	return model.Aluno{}, errors.New("aluno não encontrado")
}

func ExcluirAluno(id int) error {
	for i, aluno := range alunos {
		if aluno.ID == id {
			alunos = append(alunos[:i], alunos[i+1:]...)
			return nil
		}
	}

	return errors.New("aluno não encontrado")
}