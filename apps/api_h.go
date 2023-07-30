package webapp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

func (s *Server) apiCreateAccount(c echo.Context) error {
	name, err := generateRandomString(16)
	if err != nil {
		zlog.Error().Err(err).Send()
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err := s.query.CreateAccount(c.Request().Context(), name); err != nil {
		zlog.Error().Err(err).Send()
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{})
}
