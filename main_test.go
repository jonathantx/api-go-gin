package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jonathantx/api-go-gin/controllers"
	"github.com/jonathantx/api-go-gin/database"
	"github.com/jonathantx/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutesTests() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	return routes
}

func CreatAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12345678911", RG: "123456789"}
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

func TestGetAlunoCPFHandler(t *testing.T) {
	database.ConnectionDB()

	CreatAlunoMock()
	defer DeleteAlunoMock()

	r := SetupRoutesTests()

	r.GET("/alunos/cpf/:cpf", controllers.SearchAlunoCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678911", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetAlunoID(t *testing.T) {
	database.ConnectionDB()

	CreatAlunoMock()
	defer DeleteAlunoMock()

	r := SetupRoutesTests()

	r.GET("/alunos/:id", controllers.GetAlunoId)

	path := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var alunoMock models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	fmt.Println(alunoMock.Nome)

	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678911", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteAluno(t *testing.T) {
	database.ConnectionDB()

	CreatAlunoMock()

	r := SetupRoutesTests()
	r.DELETE("/alunos/:id", controllers.DeleteAluno)

	path := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
