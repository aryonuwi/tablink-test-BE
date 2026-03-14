package fiberhandler

import (
	"errors"
	"strconv"

	"github.com/aryonuwi/tablink-test-BE/internal/dto"
	"github.com/aryonuwi/tablink-test-BE/internal/pkg/pagination"
	"github.com/aryonuwi/tablink-test-BE/internal/pkg/response"
	appvalidator "github.com/aryonuwi/tablink-test-BE/internal/pkg/validator"
	"github.com/aryonuwi/tablink-test-BE/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type IngredientHandler struct {
	usecase   *usecase.IngredientUsecase
	validator *appvalidator.Validator
}

func NewIngredientHandler(
	usecase *usecase.IngredientUsecase,
	validator *appvalidator.Validator,
) *IngredientHandler {
	return &IngredientHandler{
		usecase:   usecase,
		validator: validator,
	}
}

func (h *IngredientHandler) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	pg := pagination.Normalize(page, limit)

	data, total, err := h.usecase.List(c.Context(), pg.Page, pg.Limit)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	meta := dto.PaginationMeta{
		Page:       pg.Page,
		Limit:      pg.Limit,
		Total:      total,
		TotalPages: pagination.TotalPages(total, pg.Limit),
	}

	return response.SuccessWithMeta(c, "success", data, meta)
}

func (h *IngredientHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateIngredientRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	if validationErrors := h.validator.ValidateStruct(req); validationErrors != nil {
		return response.ValidationError(c, validationErrors)
	}

	result, err := h.usecase.Create(c.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrIngredientNameExists) {
			return response.Error(c, fiber.StatusConflict, err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Created(c, "ingredient created", result)
}

func (h *IngredientHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	var req dto.UpdateIngredientRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	if validationErrors := h.validator.ValidateStruct(req); validationErrors != nil {
		return response.ValidationError(c, validationErrors)
	}

	result, err := h.usecase.Update(c.Context(), uuid, req)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrIngredientNotFound):
			return response.Error(c, fiber.StatusNotFound, err.Error())
		case errors.Is(err, usecase.ErrIngredientNameExists):
			return response.Error(c, fiber.StatusConflict, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return response.Success(c, "ingredient updated", result)
}

func (h *IngredientHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	err := h.usecase.Delete(c.Context(), uuid)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}

	return response.Success(c, "ingredient deleted", nil)
}
