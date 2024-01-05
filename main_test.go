package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jonathantx/api-go-gin/controllers"
	"github.com/jonathantx/api-go-gin/database"
	"github.com/jonathantx/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutesTests() *gin.Engine {

	routes := gin.Default()

	return routes
}

func CreatAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&aluno)

	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)

}

func TestStatusCodeSaudation(t *testing.T) {
	r := SetupRoutesTests()
	r.GET("/:nome", controllers.Saudation)
	req, _ := http.NewRequest("GET", "/jonathan", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	mockResponse := `{"API diz":"E aí jonathan, tudo beleza?"}`

	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, mockResponse, string(responseBody))
}

func TestListingAllAlunosHandler(t *testing.T) {
	database.ConnectionDB()

	CreatAlunoMock()
	defer DeleteAlunoMock()

	r := SetupRoutesTests()
	r.GET("/alunos", controllers.GetAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	fmt.Println(response.Body)
}
