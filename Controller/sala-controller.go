package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api-gin/Model"
	"api-gin/Service"
)

type CriarSalaRequest struct {
	Nome       string   `json:"nome"`
	Capacidade int      `json:"capacidade"`
	Recursos   []string `json:"recursos"`
}

func ListarSalas(c *gin.Context) {
	salas := service.ListarSalas()
	c.JSON(http.StatusOK, salas)
}

func BuscarSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	sala, err := service.BuscarSalaPorID(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrSalaNaoEncontrada):
			c.JSON(http.StatusNotFound, gin.H{
				"erro": err.Error(),
			})

		case errors.Is(err, service.ErrDiaSemanaInvalido),
			errors.Is(err, service.ErrHorarioInvalido),
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

	c.JSON(http.StatusOK, sala)
}

func CriarSala(c *gin.Context) {
	var request CriarSalaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "dados da sala inválidos",
		})
		return
	}

	sala := model.Sala{
		Nome:       request.Nome,
		Capacidade: request.Capacidade,
		Recursos:   request.Recursos,
	}

	salaCriada, err := service.CriarSala(sala)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "sala cadastrada com sucesso",
		"sala":     salaCriada,
	})
}

func AtualizarSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var sala model.Sala

	if err := c.ShouldBindJSON(&sala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	salaAtualizada, err := service.AtualizarSala(id, sala)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salaAtualizada)
}

func ExcluirSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := service.ExcluirSala(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func ConsultarGradeSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido",
		})
		return
	}

	grade, err := service.ConsultarGradeSala(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, grade)
}

func VerificarDisponibilidadeSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido",
		})
		return
	}

	dia := c.Query("dia")
	inicio := c.Query("inicio")
	fim := c.Query("fim")

	if dia == "" || inicio == "" || fim == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "dia, inicio e fim são obrigatórios",
		})
		return
	}

	disponivel, err := service.VerificarDisponibilidadeSala(
		id,
		dia,
		inicio,
		fim,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sala_id":    id,
		"dia":        dia,
		"inicio":     inicio,
		"fim":        fim,
		"disponivel": disponivel,
	})
}

func ListarSalasDisponiveis(c *gin.Context) {
	dia := c.Query("dia")
	inicio := c.Query("inicio")
	fim := c.Query("fim")

	if dia == "" || inicio == "" || fim == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "informe dia, inicio e fim",
		})
		return
	}

	salas, err := service.ListarSalasDisponiveis(
		dia,
		inicio,
		fim,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quantidade": len(salas),
		"salas":      salas,
	})
}
