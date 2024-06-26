package service

import (
	"baverly/service/app"
	"baverly/service/domain"
	"baverly/service/infrastructure/httphandlers"
	"baverly/service/infrastructure/repo"
	"context"
	"crypto/subtle"
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func InitService(dbConn *sql.DB, e *echo.Echo) {
	signUpRepo := repo.NewSignUpRepo(dbConn)
	authRepo := repo.NewAuthRepo(dbConn)

	signUpService := app.NewSignUpService(signUpRepo)
	authService := app.NewAuth(authRepo)

	signUpHandlers := httphandlers.NewSignUpHandler(e, signUpService)

	e.POST("/signup", signUpHandlers.SignUp)

	registered := e.Group("/v1", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {

		userInfo, err := authService.Auth(c.Request().Context(), domain.SignUpInfo{Username: username})
		if err != nil {
			return false, err
		}

		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(userInfo.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(userInfo.Password)) == 1 {
			ctx := context.WithValue(c.Request().Context(), "userInfo", userInfo)

			c.SetRequest(c.Request().Clone(ctx))

			return true, nil
		}
		return false, nil
	}))

	registered.GET("/get", func(c echo.Context) error {
		ctx := c.Request().Context()
		userInfo, ok := ctx.Value("userInfo").(domain.SignUpInfo)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, userInfo)
	})
}
