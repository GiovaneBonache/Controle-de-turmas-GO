package service

import (
	"errors"
	"strings"

	"seuprojeto/model"
)

var alunos []model.Aluno
var proximoAlunoID = 1

func ListarAlunos() []model.Aluno {
	return alunos
}

func BuscarAlunoPorID(id int) (model.Aluno, error) {
	for _, aluno := range alunos {
		if aluno.ID == id {
			return aluno, nil
		}
	}

	return model.Aluno{}, errors.New("aluno não encontrado")
}

func CriarAluno(aluno model.Aluno) (model.Aluno, error) {
	if strings.TrimSpace(aluno.Nome) == "" {
		return model.Aluno{}, errors.New("o aluno deve ter um nome")
	}

	if aluno.Matricula <= 0 {
		return model.Aluno{}, errors.New("o aluno deve ter uma matrícula válida")
	}

	if strings.TrimSpace(aluno.Email) == "" {
		return model.Aluno{}, errors.New("o aluno deve ter um e-mail")
	}

	for _, alunoExistente := range alunos {
		if alunoExistente.Matricula == aluno.Matricula {
			return model.Aluno{}, errors.New("matrícula já cadastrada")
		}
	}

	aluno.ID = proximoAlunoID
	aluno.Ativo = true

	proximoAlunoID++

	alunos = append(alunos, aluno)

	return aluno, nil
}

func AtualizarAluno(id int, alunoAtualizado model.Aluno) (model.Aluno, error) {
	if strings.TrimSpace(alunoAtualizado.Nome) == "" {
		return model.Aluno{}, errors.New("o aluno deve ter um nome")
	}

	if alunoAtualizado.Matricula <= 0 {
		return model.Aluno{}, errors.New("o aluno deve ter uma matrícula válida")
	}

	if strings.TrimSpace(alunoAtualizado.Email) == "" {
		return model.Aluno{}, errors.New("o aluno deve ter um e-mail")
	}

	for _, aluno := range alunos {
		if aluno.Matricula == alunoAtualizado.Matricula && aluno.ID != id {
			return model.Aluno{}, errors.New("matrícula já cadastrada")
		}
	}

	for i, aluno := range alunos {
		if aluno.ID == id {
			alunoAtualizado.ID = id
			alunoAtualizado.Ativo = aluno.Ativo

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