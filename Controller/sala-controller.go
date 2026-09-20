package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api-gin/Model"
	"api-gin/Service"
)

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
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sala)
}

func CriarSala(c *gin.Context) {
	var sala model.Sala

	if err := c.ShouldBindJSON(&sala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	salaCriada, err := service.CriarSala(sala)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, salaCriada)
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
