package repository

import (
	"context"

	"github.com/aryonuwi/tablink-test-BE/internal/domain"
)

type ItemRepository interface {
	GetByID(ctx context.Context, uuid string) (*domain.Item, error)
}
