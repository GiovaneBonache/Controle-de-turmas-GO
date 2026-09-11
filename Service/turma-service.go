package service

import (
	"errors"
	"seuprojeto/model"
)

var turmas []model.Turma
var proximoTurmaID = 1

func ListarTurmas() []model.Turma {
	return turmas
}

func CriarTurma(turma model.Turma) model.Turma {
	turma.ID = proximoTurmaID
	proximoTurmaID++

	turmas = append(turmas, turma)

	return turma
}

func AtualizarTurma(id int, turmaAtualizada model.Turma) (model.Turma, error) {
	for i, turma := range turmas {
		if turma.ID == id {
			turmaAtualizada.ID = id
			turmas[i] = turmaAtualizada

			return turmaAtualizada, nil
		}
	}

	return model.Turma{}, errors.New("turma não encontrada")
}

func ExcluirTurma(id int) error {
	for i, turma := range turmas {
		if turma.ID == id {
			turmas = append(turmas[:i], turmas[i+1:]...)
			return nil
		}
	}

	return errors.New("turma não encontrada")
}