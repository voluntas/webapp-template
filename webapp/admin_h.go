package webapp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

func (s *Server) adminListAccounts(c echo.Context) error {
	accounts, err := s.query.ListAccounts(c.Request().Context())
	if err != nil {
		zlog.Error().Err(err).Send()
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"accounts": accounts,
	})
}
