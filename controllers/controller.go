package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathantx/api-go-gin/database"
	"github.com/jonathantx/api-go-gin/models"
)

func GetAlunos(c *gin.Context) {

	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.JSON(200, alunos)
}

func GetAlunoId(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func Saudation(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E aí " + nome + ", tudo beleza?",
	})
}

func CreatNewAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
