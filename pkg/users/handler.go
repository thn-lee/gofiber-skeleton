package users

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/gofiber-skelton/pkg/domain"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

// NewPostHandler will initialize the post resource endpoint
func NewUserHandler(router fiber.Router, userUseCase domain.UserUseCase) {

	handler := &UserHandler{
		UserUseCase: userUseCase,
	}

	router.Get("/:id", handler.GetUser())
	router.Post("/", handler.CreateUser())
	router.Patch("/:id", handler.EditUser())
	router.Delete("/:id", handler.DeleteUser())
}

func (h *UserHandler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		userID := c.Params("id")

		user, err := h.UserUseCase.GetUser(userID)
		if err != nil {
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		responseForm.Result = map[string]interface{}{
			"user": user,
		}

		if len(responseForm.Errors) == 0 {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}

func (h *UserHandler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		if len(responseForm.Errors) == 0 {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}

func (h *UserHandler) EditUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		if len(responseForm.Errors) == 0 {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}

func (h *UserHandler) DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		if len(responseForm.Errors) == 0 {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
