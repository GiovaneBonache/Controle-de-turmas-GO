package service

import (
	"errors"
	"strings"

	"api-gin/Model"
)

var turmas []model.Turma
var proximoTurmaID = 1

func ListarTurmas() []model.Turma {
	return turmas
}

func BuscarTurmaPorID(id int) (model.Turma, error) {
	for _, turma := range turmas {
		if turma.ID == id {
			return turma, nil
		}
	}

	return model.Turma{}, errors.New("turma não encontrada")
}

func CriarTurma(turma model.Turma) (model.Turma, error) {
	if strings.TrimSpace(turma.Nome) == "" {
		return model.Turma{}, errors.New("a turma deve ter um nome")
	}

	if strings.TrimSpace(turma.Disciplina) == "" {
		return model.Turma{}, errors.New("a turma deve ter uma disciplina")
	}

	if strings.TrimSpace(turma.Professor) == "" {
		return model.Turma{}, errors.New("a turma deve ter um professor")
	}

	turma.ID = proximoTurmaID
	turma.Ativo = true

	proximoTurmaID++

	turmas = append(turmas, turma)

	return turma, nil
}

func AtualizarTurma(id int, turmaAtualizada model.Turma) (model.Turma, error) {
	if strings.TrimSpace(turmaAtualizada.Nome) == "" {
		return model.Turma{}, errors.New("a turma deve ter um nome")
	}

	if strings.TrimSpace(turmaAtualizada.Disciplina) == "" {
		return model.Turma{}, errors.New("a turma deve ter uma disciplina")
	}

	if strings.TrimSpace(turmaAtualizada.Professor) == "" {
		return model.Turma{}, errors.New("a turma deve ter um professor")
	}

	for i, turma := range turmas {
		if turma.ID == id {
			turmaAtualizada.ID = id
			turmaAtualizada.Ativo = turma.Ativo

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

func ConsultarGradeSala(salaID int) ([]model.UsoSala, error) {
	sala, err := BuscarSalaPorID(salaID)
	if err != nil {
		return nil, errors.New("sala não encontrada")
	}

	if !sala.Ativo {
		return nil, errors.New("sala inativa")
	}

	var grade []model.UsoSala

	for _, turma := range turmas {
		if turma.Alocacao != nil && turma.Alocacao.SalaID == salaID {
			grade = append(grade, model.UsoSala{
				TurmaID:       turma.ID,
				TurmaNome:     turma.Nome,
				DiaSemana:     turma.Alocacao.DiaSemana,
				HorarioInicio: turma.Alocacao.HorarioInicio,
				HorarioFim:    turma.Alocacao.HorarioFim,
			})
		}
	}

	return grade, nil
}