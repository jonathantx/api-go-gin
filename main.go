package main

import (
	"github.com/jonathantx/api-go-gin/database"
	"github.com/jonathantx/api-go-gin/models"
	"github.com/jonathantx/api-go-gin/routes"
)

func main() {

	database.ConnectionDB()

	models.Alunos = []models.Aluno{
		{Nome: "Jonathan Teixeira", CPF: "99999879-5", RG: "35654789"},
		{Nome: "Ana", CPF: "123654789-5", RG: "63587981"},
	}

	routes.HandleRequests()
}
