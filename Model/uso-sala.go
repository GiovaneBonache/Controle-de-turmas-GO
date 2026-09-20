package model

type UsoSala struct {
	TurmaID       int    `json:"turma_id"`
	TurmaNome     string `json:"turma_nome"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}
