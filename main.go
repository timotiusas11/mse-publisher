package main

import (
	"fmt"
	"mse-publisher/infrastructure"
	"mse-publisher/interfaces"
	"mse-publisher/usecases"
	"net/http"
)

func main() {
	psqlDB, err := infrastructure.NewPsqlDB()
	if err != nil {
		return
	}
	defer psqlDB.Close()

	prdc, err := infrastructure.NewNsqConn()
	if err != nil {
		return
	}

	repo := interfaces.NewRecipeRepository(psqlDB, prdc)
	recipeInteractor := new(usecases.RecipeInteractor)
	recipeInteractor.RecipeRepository = repo

	webServiceHandler := interfaces.WebServiceHandler{}
	webServiceHandler.RecipeInteractor = recipeInteractor

	http.HandleFunc("/get", webServiceHandler.GetRecipeById)
	http.HandleFunc("/insert", webServiceHandler.InsertRecipe)
	http.HandleFunc("/update", webServiceHandler.UpdateRecipe)
	http.HandleFunc("/delete", webServiceHandler.DeleteRecipe)

	fmt.Println("Starting Web Server at http://localhost:8082/")
	http.ListenAndServe(":8082", nil)
}
