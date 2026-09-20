package service

import (
	"errors"
	"strings"
	"time"

	"api-gin/Model"
)

var turmas []model.Turma
var proximoTurmaID = 1

var (
	ErrSalaNaoEncontrada = errors.New("sala não encontrada")
	ErrTurmaInativa      = errors.New("turma inativa")
	ErrSalaInativa       = errors.New("sala inativa")
	ErrHorarioInvalido   = errors.New("horário inválido")
	ErrConflitoSala      = errors.New("sala ocupada nesse horário")
	ErrDiaSemanaInvalido = errors.New("dia da semana inválido")
)

func ListarTurmas() []model.TurmaResumo {
	var resultado []model.TurmaResumo

	for _, turma := range turmas {
		resultado = append(resultado, model.TurmaResumo{
			ID:               turma.ID,
			Nome:             turma.Nome,
			Disciplina:       turma.Disciplina,
			Professor:        turma.Professor,
			QuantidadeAlunos: len(turma.Alunos),
			Alocada:          turma.Alocacao != nil,
			Ativo:            turma.Ativo,
		})
	}

	return resultado
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

	turma.Alunos = []int{}
	turma.Alocacao = nil

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

			turmas[i].Nome = turmaAtualizada.Nome
			turmas[i].Disciplina = turmaAtualizada.Disciplina
			turmas[i].Professor = turmaAtualizada.Professor

			return turmas[i], nil
		}
	}

	return model.Turma{}, ErrTurmaNaoEncontrada
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

func alunoEstaNaTurma(turma model.Turma, alunoID int) bool {
	for _, id := range turma.Alunos {
		if id == alunoID {
			return true
		}
	}

	return false
}

func horariosSobrepostos(
	inicioNovo string,
	fimNovo string,
	inicioExistente string,
	fimExistente string,
) bool {

	return inicioNovo < fimExistente &&
		fimNovo > inicioExistente
}

func alunoTemConflitoHorario(
	alunoID int,
	turmaID int,
	diaSemana string,
	horarioInicio string,
	horarioFim string,
) bool {

	for _, turma := range turmas {

		if turma.ID == turmaID {
			continue
		}

		if turma.Alocacao == nil {
			continue
		}

		if !alunoEstaNaTurma(turma, alunoID) {
			continue
		}

		if !strings.EqualFold(
			turma.Alocacao.DiaSemana,
			diaSemana,
		) {
			continue
		}

		if horariosSobrepostos(
			horarioInicio,
			horarioFim,
			turma.Alocacao.HorarioInicio,
			turma.Alocacao.HorarioFim,
		) {
			return true
		}
	}

	return false
}

func validarHorario(inicio string, fim string) error {
	horarioInicio, err := time.Parse("15:04", inicio)
	if err != nil {
		return ErrHorarioInvalido
	}

	horarioFim, err := time.Parse("15:04", fim)
	if err != nil {
		return ErrHorarioInvalido
	}

	if !horarioInicio.Before(horarioFim) {
		return ErrHorarioInvalido
	}

	return nil
}
func AlocarSala(
	turmaID int,
	salaID int,
	diaSemana string,
	horarioInicio string,
	horarioFim string,
) error {

	turmaIndex := -1

	for i, turma := range turmas {
		if turma.ID == turmaID {
			turmaIndex = i
			break
		}
	}

	if turmaIndex == -1 {
		return ErrTurmaNaoEncontrada
	}

	turma := turmas[turmaIndex]

	if !turma.Ativo {
		return ErrTurmaInativa
	}

	sala, err := BuscarSalaPorID(salaID)
	if err != nil {
		return ErrSalaNaoEncontrada
	}

	if !sala.Ativo {
		return ErrSalaInativa
	}

	if !validarDiaSemana(diaSemana) {
		return ErrDiaSemanaInvalido
	}

	if err := validarHorario(horarioInicio, horarioFim); err != nil {
		return err
	}

	if len(turma.Alunos) > sala.Capacidade {
		return ErrCapacidadeInsuficiente
	}

	for _, outraTurma := range turmas {

		if outraTurma.ID == turmaID {
			continue
		}

		if outraTurma.Alocacao == nil {
			continue
		}

		if outraTurma.Alocacao.SalaID != salaID {
			continue
		}

		if !strings.EqualFold(
			outraTurma.Alocacao.DiaSemana,
			diaSemana,
		) {
			continue
		}

		if horariosSobrepostos(
			horarioInicio,
			horarioFim,
			outraTurma.Alocacao.HorarioInicio,
			outraTurma.Alocacao.HorarioFim,
		) {
			return ErrConflitoSala
		}
	}

	for _, alunoID := range turma.Alunos {
		if alunoTemConflitoHorario(
			alunoID,
			turmaID,
			diaSemana,
			horarioInicio,
			horarioFim,
		) {
			return ErrConflitoHorarioAluno
		}
	}

	turmas[turmaIndex].Alocacao = &model.Alocacao{
		SalaID:        salaID,
		DiaSemana:     diaSemana,
		HorarioInicio: horarioInicio,
		HorarioFim:    horarioFim,
	}

	return nil
}

func validarDiaSemana(dia string) bool {
	diasValidos := []string{
		"segunda",
		"terca",
		"quarta",
		"quinta",
		"sexta",
		"sabado",
		"domingo",
	}

	for _, diaValido := range diasValidos {
		if strings.EqualFold(dia, diaValido) {
			return true
		}
	}

	return false
}
