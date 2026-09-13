package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api-gin/Model"
	"api-gin/Service"
)

func ListarAlunos(c *gin.Context) {
	alunos := service.ListarAlunos()
	c.JSON(http.StatusOK, alunos)
}

func BuscarAluno(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	aluno, err := service.BuscarAlunoPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func CriarAluno(c *gin.Context) {
	var aluno model.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	alunoCriado, err := service.CriarAluno(aluno)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, alunoCriado)
}

func AtualizarAluno(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var aluno model.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	alunoAtualizado, err := service.AtualizarAluno(id, aluno)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alunoAtualizado)
}

func ExcluirAluno(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := service.ExcluirAluno(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}