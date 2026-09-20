package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api-gin/Model"
	"api-gin/Service"
)

type AlocarSalaRequest struct {
	SalaID        int    `json:"sala_id"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}

type MatriculaAlunoDTO struct {
	AlunoID int `json:"aluno_id"`
}

func ListarTurmas(c *gin.Context) {
	turmas := service.ListarTurmas()
	c.JSON(http.StatusOK, turmas)
}

func BuscarTurma(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	turma, err := service.BuscarTurmaPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, turma)
}

func CriarTurma(c *gin.Context) {
	var turma model.Turma

	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	turmaCriada, err := service.CriarTurma(turma)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, turmaCriada)
}

func AtualizarTurma(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var turma model.Turma

	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	turmaAtualizada, err := service.AtualizarTurma(id, turma)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, turmaAtualizada)
}

func ExcluirTurma(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := service.ExcluirTurma(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func MatricularAluno(c *gin.Context) {
	turmaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID da turma inválido"})
		return
	}

	var request MatriculaAlunoDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}
	if request.AlunoID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "aluno_id inválido"})
		return
	}

	err = service.MatricularAluno(turmaID, request.AlunoID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTurmaNaoEncontrada),
			errors.Is(err, service.ErrAlunoNaoEncontrado),
			errors.Is(err, service.ErrSalaAlocadaNaoEncontrada):

			c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		case errors.Is(err, service.ErrAlunoJaMatriculado),
			errors.Is(err, service.ErrConflitoHorarioAluno):

			c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		case errors.Is(err, service.ErrCapacidadeInsuficiente):

			c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "aluno matriculado com sucesso",
	})
}
func ListarAlunosDaTurma(c *gin.Context) {
	turmaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID da turma inválido"})
		return
	}

	alunos, err := service.ListarAlunosDaTurma(turmaID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alunos)
}

func AlocarSala(c *gin.Context) {
	turmaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID da turma inválido",
		})
		return
	}

	var request AlocarSalaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	if request.SalaID <= 0 ||
		request.DiaSemana == "" ||
		request.HorarioInicio == "" ||
		request.HorarioFim == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "sala_id, dia_semana, horario_inicio e horario_fim são obrigatórios",
		})
		return
	}

	err = service.AlocarSala(
		turmaID,
		request.SalaID,
		request.DiaSemana,
		request.HorarioInicio,
		request.HorarioFim,
	)

	if err != nil {
		switch {

		case errors.Is(err, service.ErrTurmaNaoEncontrada),
			errors.Is(err, service.ErrSalaNaoEncontrada):

			c.JSON(http.StatusNotFound, gin.H{
				"erro": err.Error(),
			})

		case errors.Is(err, service.ErrCapacidadeInsuficiente):

			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"erro": err.Error(),
			})

		case errors.Is(err, service.ErrConflitoSala),
			errors.Is(err, service.ErrConflitoHorarioAluno):

			c.JSON(http.StatusConflict, gin.H{
				"erro": err.Error(),
			})

		case errors.Is(err, service.ErrHorarioInvalido),
			errors.Is(err, service.ErrDiaSemanaInvalido),
			errors.Is(err, service.ErrTurmaInativa),
			errors.Is(err, service.ErrSalaInativa):

			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "sala alocada com sucesso",
	})
}
