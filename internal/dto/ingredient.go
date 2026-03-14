package dto

type CreateIngredientRequest struct {
	Name        string `json:"name" validate:"required"`
	CauseAlergy bool   `json:"cause_alergy"`
	Type        int    `json:"type" validate:"oneof=0 1 2"`
	Status      int    `json:"status" validate:"oneof=0 1"`
}

type UpdateIngredientRequest struct {
	Name        string `json:"name" validate:"required"`
	CauseAlergy bool   `json:"cause_alergy"`
	Type        int    `json:"type" validate:"oneof=0 1 2"`
	Status      int    `json:"status" validate:"oneof=0 1"`
}
