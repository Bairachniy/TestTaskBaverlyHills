package httphandlers

import (
	"baverly/service/app"
	"baverly/service/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SignUpHandler struct {
	e             *echo.Echo
	signUpService app.SignUpIService
}

func NewSignUpHandler(e *echo.Echo, sUpService app.SignUpIService) *SignUpHandler {
	return &SignUpHandler{
		e:             e,
		signUpService: sUpService,
	}
}

func (h *SignUpHandler) SignUp(ctx echo.Context) error {
	var req SignUpInfo
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.signUpService.SignUp(ctx.Request().Context(), domain.SignUpInfo{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "OK")
}
