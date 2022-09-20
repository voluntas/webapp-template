package webapp

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) BasicAuthMiddlware() echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(
		middleware.BasicAuthConfig{
			Skipper: func(c echo.Context) bool {
				return s.config.SkipBasicAuth
			},
			Validator: func(username, password string, _ echo.Context) (bool, error) {
				if subtle.ConstantTimeCompare([]byte(username), []byte(s.config.BasicAuthUsername)) == 1 &&
					subtle.ConstantTimeCompare([]byte(password), []byte(s.config.BasicAuthPassword)) == 1 {
					return true, nil
				}
				return false, nil
			},
		})
}
