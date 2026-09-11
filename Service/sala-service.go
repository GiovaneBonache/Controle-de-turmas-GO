package service

import (
	"errors"
	"seuprojeto/model"
)

var salas []model.Sala
var proximoSalaID = 1

func ListarSalas() []model.Sala {
	return salas
}

func CriarSala(sala model.Sala) model.Sala {
	sala.ID = proximoSalaID
	proximoSalaID++

	salas = append(salas, sala)

	return sala
}

func AtualizarSala(id int, salaAtualizada model.Sala) (model.Sala, error) {
	for i, sala := range salas {
		if sala.ID == id {
			salaAtualizada.ID = id
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