package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathantx/api-go-gin/database"
	"github.com/jonathantx/api-go-gin/models"
)

func GetAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
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
