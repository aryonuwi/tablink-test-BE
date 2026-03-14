package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type ItemIngredientRepository struct {
	db *pgxpool.Pool
}

func NewItemIngredientRepository(db *pgxpool.Pool) *ItemIngredientRepository {
	return &ItemIngredientRepository{db: db}
}
