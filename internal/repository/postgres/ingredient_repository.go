package postgres

import (
	"context"
	"errors"
	"strings"

	"github.com/aryonuwi/tablink-test-BE/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IngredientRepository struct {
	db *pgxpool.Pool
}

func NewIngredientRepository(db *pgxpool.Pool) *IngredientRepository {
	return &IngredientRepository{db: db}
}

func (r *IngredientRepository) List(ctx context.Context, limit, offset int) ([]domain.Ingredient, int64, error) {
	query := `
		SELECT uuid, name, cause_alergy, type, status, created_at, updated_at, deleted_at
		FROM tm_ingredient
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var data []domain.Ingredient
	for rows.Next() {
		var row domain.Ingredient
		err := rows.Scan(
			&row.UUID,
			&row.Name,
			&row.CauseAlergy,
			&row.Type,
			&row.Status,
			&row.CreatedAt,
			&row.UpdatedAt,
			&row.DeletedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		data = append(data, row)
	}

	countQuery := `
		SELECT COUNT(*)
		FROM tm_ingredient
		WHERE deleted_at IS NULL
	`

	var total int64
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (r *IngredientRepository) GetByID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	query := `
		SELECT uuid, name, cause_alergy, type, status, created_at, updated_at, deleted_at
		FROM tm_ingredient
		WHERE uuid = $1 AND deleted_at IS NULL
	`

	var row domain.Ingredient
	err := r.db.QueryRow(ctx, query, uuid).Scan(
		&row.UUID,
		&row.Name,
		&row.CauseAlergy,
		&row.Type,
		&row.Status,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (r *IngredientRepository) Create(ctx context.Context, data domain.Ingredient) (*domain.Ingredient, error) {
	query := `
		INSERT INTO tm_ingredient (name, cause_alergy, type, status)
		VALUES ($1, $2, $3, $4)
		RETURNING uuid, name, cause_alergy, type, status, created_at, updated_at, deleted_at
	`

	var row domain.Ingredient
	err := r.db.QueryRow(ctx, query,
		strings.TrimSpace(data.Name),
		data.CauseAlergy,
		data.Type,
		data.Status,
	).Scan(
		&row.UUID,
		&row.Name,
		&row.CauseAlergy,
		&row.Type,
		&row.Status,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (r *IngredientRepository) Update(ctx context.Context, uuid string, data domain.Ingredient) (*domain.Ingredient, error) {
	query := `
		UPDATE tm_ingredient
		SET name = $2,
		    cause_alergy = $3,
		    type = $4,
		    status = $5,
		    updated_at = NOW()
		WHERE uuid = $1
		  AND deleted_at IS NULL
		RETURNING uuid, name, cause_alergy, type, status, created_at, updated_at, deleted_at
	`

	var row domain.Ingredient
	err := r.db.QueryRow(ctx, query,
		uuid,
		strings.TrimSpace(data.Name),
		data.CauseAlergy,
		data.Type,
		data.Status,
	).Scan(
		&row.UUID,
		&row.Name,
		&row.CauseAlergy,
		&row.Type,
		&row.Status,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (r *IngredientRepository) SoftDelete(ctx context.Context, uuid string) error {
	query := `
		UPDATE tm_ingredient
		SET deleted_at = NOW(),
		    updated_at = NOW()
		WHERE uuid = $1
		  AND deleted_at IS NULL
	`

	cmdTag, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("ingredient not found")
	}

	return nil
}

func (r *IngredientRepository) ExistsByName(ctx context.Context, name string, excludeUUID *string) (bool, error) {
	name = strings.TrimSpace(name)

	if excludeUUID == nil {
		query := `
			SELECT EXISTS(
				SELECT 1
				FROM tm_ingredient
				WHERE LOWER(name) = LOWER($1)
				  AND deleted_at IS NULL
			)
		`
		var exists bool
		err := r.db.QueryRow(ctx, query, name).Scan(&exists)
		return exists, err
	}

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM tm_ingredient
			WHERE LOWER(name) = LOWER($1)
			  AND deleted_at IS NULL
			  AND uuid <> $2
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, name, *excludeUUID).Scan(&exists)
	return exists, err
}
