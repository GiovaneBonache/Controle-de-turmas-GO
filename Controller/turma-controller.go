package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"seuprojeto/model"
	"seuprojeto/service"
)

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