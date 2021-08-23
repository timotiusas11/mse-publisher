package interfaces

import (
	"encoding/json"
	"mse-publisher/usecases"
	"net/http"
	"strconv"
)

type RecipeInteractor interface {
	GetRecipeById(recipeId int) (res *usecases.Recipe, err error)
	InsertRecipe(usecases.Recipe) (res *usecases.Recipe, err error)
	UpdateRecipe(usecases.Recipe) error
	DeleteRecipe(id int) error
}

type WebServiceHandler struct {
	RecipeInteractor RecipeInteractor
}

func (handler WebServiceHandler) GetRecipeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		recipe, err := handler.RecipeInteractor.GetRecipeById(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result, err := json.Marshal(recipe)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func (handler WebServiceHandler) InsertRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var rec usecases.Recipe

		err := json.NewDecoder(r.Body).Decode(&rec)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, e := handler.RecipeInteractor.InsertRecipe(rec)

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		result, err := json.Marshal(struct {
			Result string
			Data   *usecases.Recipe
		}{"Success!", res})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func (handler WebServiceHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		var rec usecases.Recipe

		err := json.NewDecoder(r.Body).Decode(&rec)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		e := handler.RecipeInteractor.UpdateRecipe(rec)

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		result, err := json.Marshal(struct{ Result string }{"Success!"})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func (handler WebServiceHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		e := handler.RecipeInteractor.DeleteRecipe(id)

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		result, err := json.Marshal(struct{ Result string }{"Success!"})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
