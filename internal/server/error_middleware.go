package server

import (
	"C0lliNN/card-game-service/internal/game"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrorMiddleware struct{}

func NewErrorMiddleware() ErrorMiddleware {
	return ErrorMiddleware{}
}

type ErrorResponse struct {
	Code    int      `json:"-"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func (e *ErrorMiddleware) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) <= 0 {
			return
		}
		err := context.Errors.Last().Err

		log.Println("error processing the request: ", err)

		var response *ErrorResponse

		switch {
		case errors.Is(err, game.ErrDeckNotFound):
			response = newDeckNotFoundResponse(err)
		case errors.Is(err, game.ErrInvalidDrawQuantity):
			response = newInvalidDrawQuantityResponse(err)
		default:
			response = newGenericErrorResponse(err)
		}

		context.JSON(response.Code, response)
	}
}

func newErrorResponse(code int, err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: unwrappedError(err).Error(),
		Details: errorStack(err),
	}
}

func newDeckNotFoundResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}
}

func newInvalidDrawQuantityResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

func newGenericErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Some unexpected error happened",
		Details: errorStack(err),
	}
}

func unwrappedError(err error) error {
	if errors.Unwrap(err) == nil {
		return err
	}

	return unwrappedError(errors.Unwrap(err))
}

func errorStack(err error) []string {
	stack := strings.Split(err.Error(), ":")
	for i := range stack {
		stack[i] = strings.TrimSpace(stack[i])
	}

	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}
