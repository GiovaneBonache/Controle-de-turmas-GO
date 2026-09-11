package model

type Alocacao struct {
	SalaID        int    `json:"sala_id"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}