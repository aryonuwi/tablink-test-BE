package repository

import (
	"context"

	"github.com/aryonuwi/tablink-test-BE/internal/domain"
)

type ItemIngredientRepository interface {
	ReplaceItemIngredients(ctx context.Context, itemUUID string, ingredientUUIDs []string) error
	DeleteRelation(ctx context.Context, itemUUID, ingredientUUID string) error
	ListIngredientsByItemID(ctx context.Context, itemUUID string) ([]domain.Ingredient, error)
}
