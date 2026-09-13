package service

import (
	"errors"
	"strings"

	"api-gin/Model"
)

var salas []model.Sala
var turmas []model.Turma
var proximoSalaID = 1

func ListarSalas() []model.Sala {
	return salas
}

func BuscarSalaPorID(id int) (model.Sala, error) {
	for _, sala := range salas {
		if sala.ID == id {
			return sala, nil
		}
	}

	return model.Sala{}, errors.New("sala não encontrada")
}

func CriarSala(sala model.Sala) (model.Sala, error) {
	if strings.TrimSpace(sala.Nome) == "" {
		return model.Sala{}, errors.New("a sala deve ter um nome")
	}

	if sala.Capacidade <= 0 {
		return model.Sala{}, errors.New("capacidade deve ser maior que zero")
	}

	sala.ID = proximoSalaID
	sala.Ativo = true

	proximoSalaID++

	salas = append(salas, sala)

	return sala, nil
}

func AtualizarSala(id int, salaAtualizada model.Sala) (model.Sala, error) {
	if strings.TrimSpace(salaAtualizada.Nome) == "" {
		return model.Sala{}, errors.New("a sala deve ter um nome")
	}

	if salaAtualizada.Capacidade <= 0 {
		return model.Sala{}, errors.New("capacidade deve ser maior que zero")
	}

	for i, sala := range salas {
		if sala.ID == id {
			salaAtualizada.ID = id
			salaAtualizada.Ativo = sala.Ativo

			salas[i] = salaAtualizada

			return salaAtualizada, nil
		}
	}

	return model.Sala{}, errors.New("sala não encontrada")
}

func ExcluirSala(id int) error {
	for i, sala := range salas {
		if sala.ID == id {
			salas = append(salas[:i], salas[i+1:]...)
			return nil
		}
	}

	return errors.New("sala não encontrada")
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

func VerificarDisponibilidadeSala(
	salaID int,
	diaSemana string,
	horarioInicio string,
	horarioFim string,
) (bool, error) {

	sala, err := BuscarSalaPorID(salaID)
	if err != nil {
		return false, errors.New("sala não encontrada")
	}

	if !sala.Ativo {
		return false, errors.New("sala inativa")
	}

	for _, turma := range turmas {
		if turma.Alocacao == nil {
			continue
		}

		if turma.Alocacao.SalaID != salaID {
			continue
		}

		if turma.Alocacao.DiaSemana != diaSemana {
			continue
		}

		if horarioInicio < turma.Alocacao.HorarioFim &&
			horarioFim > turma.Alocacao.HorarioInicio {
			return false, nil
		}
	}

	return true, nil
}