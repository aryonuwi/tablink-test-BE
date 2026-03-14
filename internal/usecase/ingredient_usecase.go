package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/aryonuwi/tablink-test-BE/internal/domain"
	"github.com/aryonuwi/tablink-test-BE/internal/dto"
	"github.com/aryonuwi/tablink-test-BE/internal/repository"
)

var (
	ErrIngredientNameExists = errors.New("ingredient name already exists")
	ErrIngredientNotFound   = errors.New("ingredient not found")
)

type IngredientUsecase struct {
	repo repository.IngredientRepository
}

func NewIngredientUsecase(repo repository.IngredientRepository) *IngredientUsecase {
	return &IngredientUsecase{repo: repo}
}

func (u *IngredientUsecase) List(ctx context.Context, page, limit int) ([]domain.Ingredient, int64, error) {
	offset := (page - 1) * limit
	return u.repo.List(ctx, limit, offset)
}

func (u *IngredientUsecase) Create(ctx context.Context, req dto.CreateIngredientRequest) (*domain.Ingredient, error) {
	req.Name = strings.TrimSpace(req.Name)

	exists, err := u.repo.ExistsByName(ctx, req.Name, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrIngredientNameExists
	}

	payload := domain.Ingredient{
		Name:        req.Name,
		CauseAlergy: req.CauseAlergy,
		Type:        req.Type,
		Status:      req.Status,
	}

	return u.repo.Create(ctx, payload)
}

func (u *IngredientUsecase) Update(ctx context.Context, uuid string, req dto.UpdateIngredientRequest) (*domain.Ingredient, error) {
	req.Name = strings.TrimSpace(req.Name)

	_, err := u.repo.GetByID(ctx, uuid)
	if err != nil {
		return nil, ErrIngredientNotFound
	}

	exists, err := u.repo.ExistsByName(ctx, req.Name, &uuid)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrIngredientNameExists
	}

	payload := domain.Ingredient{
		Name:        req.Name,
		CauseAlergy: req.CauseAlergy,
		Type:        req.Type,
		Status:      req.Status,
	}

	return u.repo.Update(ctx, uuid, payload)
}

func (u *IngredientUsecase) Delete(ctx context.Context, uuid string) error {
	return u.repo.SoftDelete(ctx, uuid)
}
