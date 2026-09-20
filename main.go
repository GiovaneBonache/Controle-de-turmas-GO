package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"api-gin/Controller"
)

func main() {

	r := gin.New()

	// Uso dos Middlewares globais nativos e personalizados
	r.Use(gin.Recovery())

	// 4. Mapeamento de Rotas sob Grupo Versionado
	v1 := r.Group("/api/v1")
	{
		// Monitoramento da API
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":    "healthy",
				"timestamp": time.Now(),
				"version":   "1.0.0",
			})
		})

		// Domínio de Turmas (Classes)
		v1.POST("/alunos", controller.CriarAluno)
		v1.GET("/alunos", controller.ListarAlunos)
		v1.GET("/alunos/:id", controller.BuscarAluno)
		v1.PUT("/alunos/:id", controller.AtualizarAluno)
		v1.DELETE("/alunos/:id", controller.ExcluirAluno)

		v1.POST("/salas", controller.CriarSala)
		v1.GET("/salas", controller.ListarSalas)
		v1.GET("/salas/:id/grade", controller.ConsultarGradeSala)
		v1.GET("/salas/:id/disponibilidade", controller.VerificarDisponibilidadeSala)
		v1.GET("/salas/:id", controller.BuscarSala)
		v1.PUT("/salas/:id", controller.AtualizarSala)
		v1.DELETE("/salas/:id", controller.ExcluirSala)

		v1.POST("/turmas", controller.CriarTurma)
		v1.POST("/turmas/:id/alunos", controller.MatricularAluno)
		v1.POST("/turmas/:id/alocar", controller.AlocarSala)
		v1.GET("/turmas", controller.ListarTurmas)
		v1.GET("/turmas/:id/alunos", controller.ListarAlunosDaTurma)
		v1.GET("/turmas/:id", controller.BuscarTurma)
		v1.PUT("/turmas/:id", controller.AtualizarTurma)
		v1.DELETE("/turmas/:id", controller.ExcluirTurma)
	}

	r.Run(":8080")
}
