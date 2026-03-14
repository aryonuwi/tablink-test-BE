package repository

import (
	"context"

	"github.com/aryonuwi/tablink-test-BE/internal/domain"
)

type IngredientRepository interface {
	List(ctx context.Context, limit, offset int) ([]domain.Ingredient, int64, error)
	GetByID(ctx context.Context, uuid string) (*domain.Ingredient, error)
	Create(ctx context.Context, data domain.Ingredient) (*domain.Ingredient, error)
	Update(ctx context.Context, uuid string, data domain.Ingredient) (*domain.Ingredient, error)
	SoftDelete(ctx context.Context, uuid string) error
	ExistsByName(ctx context.Context, name string, excludeUUID *string) (bool, error)
}
