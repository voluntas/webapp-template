package webapp

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	zlog "github.com/rs/zerolog/log"
	sqlc "github.com/voluntas/webapp/gen/sqlc"
)

// TODO(v): 定数どうにかする
var SQLITE_PATH = "webapp.db"

type Server struct {
	revision string
	config   Config
	query    *sqlc.Queries

	echo         *echo.Echo
	echoExporter *echo.Echo
}

// ここでデータベース接続したままにする
func NewConn(_ Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", SQLITE_PATH)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewServer(r string, c Config, conn *sql.DB) (*Server, error) {
	server := &Server{
		revision: r,
		config:   c,
		query:    sqlc.New(conn),
	}

	server.setupEcho()
	return server, nil
}

func (s *Server) Start(address string, port int) error {
	if s.config.Https {
		return s.echo.StartTLS(fmt.Sprintf("%s:%d", address, port), s.config.CetificatePath, s.config.PrivateKeyPath)
	}
	return s.echo.Start(fmt.Sprintf("%s:%d", address, port))
}

func (s *Server) StartExporter(address string, port int) error {
	return s.echoExporter.Start(fmt.Sprintf("%s:%d", address, port))
}

func (s *Server) setupEcho() {
	e := echo.New()

	e.Validator = &Validator{validator: validator.New()}

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(s.LoggingMiddleware())

	// LB からのヘルスチェック専用 API
	e.GET("/.ok", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"revision": s.revision,
		})
	})

	e.POST("/create-account", s.apiCreateAccount)

	adminGroup := e.Group("/admin", s.BasicAuthMiddlware())
	adminRoutes(s, adminGroup)

	apiGroup := e.Group("/api", s.BasicAuthMiddlware())
	apiRoutes(s, apiGroup)

	echoExporter := echo.New()
	echoExporter.HideBanner = true
	prom := prometheus.NewPrometheus("echo", nil)

	e.Use(prom.HandlerFunc)
	prom.SetMetricsPath(echoExporter)

	s.echo = e
	s.echoExporter = echoExporter
}

func adminRoutes(s *Server, g *echo.Group) {
	g.POST("/list-accounts", s.adminListAccounts)
}

func apiRoutes(s *Server, g *echo.Group) {
	g.POST("/create-account", s.apiCreateAccount)
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func (s *Server) LoggingMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogLatency:       true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogReferer:       true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogResponseSize:  true,
		Skipper: func(c echo.Context) bool {
			// /.ok の時はログを吐き出さない
			return strings.HasPrefix(c.Request().URL.Path, "/.ok")
		},
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			stop := time.Now()
			var requestSize int
			if len(v.ContentLength) > 0 {
				requestSize, _ = strconv.Atoi(v.ContentLength)
			}
			zlog.Info().
				Err(v.Error).
				Time("start", v.StartTime).
				Int64("letency", int64(v.Latency)).
				Str("human_letency", stop.Sub(v.StartTime).String()).
				Str("protocol", v.Host).
				Str("remote_ip", v.RemoteIP).
				Str("host", v.Host).
				Str("method", v.Method).
				Str("url", v.URI).
				Str("referer", v.Referer).
				Str("user_agent", v.UserAgent).
				Int("status", v.Status).
				Int("request_size", requestSize).
				Int64("response_size", v.ResponseSize).
				Msg("request")
			return nil
		},
	})
}
