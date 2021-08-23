package interfaces

import (
	"context"
	"encoding/json"
	"mse-publisher/domain"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nsqio/go-nsq"
)

type recipeRepo struct {
	db  *pgxpool.Pool
	nsq *nsq.Producer
}

func NewRecipeRepository(db *pgxpool.Pool, nsq *nsq.Producer) domain.RecipeRepository {
	return &recipeRepo{db: db, nsq: nsq}
}

func (repo *recipeRepo) FindById(id int) (*domain.Recipe, error) {
	row := repo.db.QueryRow(context.Background(), "SELECT * FROM recipes WHERE id = $1", id)
	var name string
	var ingredients []string
	var isHalal bool
	var isVegetarian bool
	var description string
	var rating float64
	err := row.Scan(&id, &name, &ingredients, &isHalal, &isVegetarian, &description, &rating)
	if err != nil {
		return nil, err
	}
	recipe := domain.Recipe{
		Id:           id,
		Name:         name,
		Ingredients:  ingredients,
		IsHalal:      isHalal,
		IsVegetarian: isVegetarian,
		Description:  description,
		Rating:       rating,
	}

	return &recipe, nil
}

func (repo *recipeRepo) Create(recipe domain.Recipe) (*domain.Recipe, error) {
	var id int

	err := repo.db.QueryRow(context.Background(),
		"INSERT INTO recipes (name, ingredients, isHalal, isVegetarian, description, rating) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		recipe.Name, recipe.Ingredients, recipe.IsHalal, recipe.IsVegetarian, recipe.Description, recipe.Rating,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	recipe.Id = id

	topic := "recipes"

	payload, err := json.Marshal(
		struct {
			Action string
			Data   domain.Recipe
		}{"create", recipe},
	)
	if err != nil {
		return nil, err
	}
	err = repo.nsq.Publish(topic, payload)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (repo *recipeRepo) Update(recipe domain.Recipe) error {
	_, err := repo.db.Exec(context.Background(),
		`UPDATE recipes 
		SET name = $1, 
			ingredients = $2, 
			isHalal = $3,  
			isVegetarian = $4,
			description = $5,
			rating = $6
		WHERE id = $7`,
		recipe.Name, recipe.Ingredients, recipe.IsHalal, recipe.IsVegetarian, recipe.Description, recipe.Rating, recipe.Id,
	)
	if err != nil {
		return err
	}

	topic := "recipes"

	payload, err := json.Marshal(
		struct {
			Action string
			Data   domain.Recipe
		}{"update", recipe},
	)
	if err != nil {
		return err
	}
	err = repo.nsq.Publish(topic, payload)
	if err != nil {
		return err
	}

	return nil
}

func (repo *recipeRepo) Delete(id int) error {
	_, err := repo.db.Exec(context.Background(),
		`DELETE FROM recipes 
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	topic := "recipes"

	payload, err := json.Marshal(
		struct {
			Action string
			Data   domain.Recipe
		}{"delete", domain.Recipe{
			Id: id,
		}},
	)
	if err != nil {
		return err
	}
	err = repo.nsq.Publish(topic, payload)
	if err != nil {
		return err
	}
	return nil
}
