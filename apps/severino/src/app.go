package src

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"

	"backend/apps/severino/src/handlers"
	"backend/apps/severino/src/middlewares"
	"backend/apps/severino/src/services"
	validator "backend/libs/httpvalidator"
	log "backend/libs/logger"
)

type Application struct {
	authHandler       *handlers.AuthHandler
	trackHandler      *handlers.TrackHandler
	shelfHandler      *handlers.ShelfHandler
	albumHandler      *handlers.AlbumHandler
	userHandler       *handlers.UserHandler
	accountHandler    *handlers.AccountHandler
	productHandler    *handlers.ProductHandler
	homeHandler       *handlers.HomeHandler
	preferenceHandler *handlers.PreferenceHandler
	subjectHandler    *handlers.SubjectHandler
}

var connections []*grpc.ClientConn
var createdAt string

func init() {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	createdAt = time.Now().In(loc).Format("2006-01-02 15:04:05.000000")
}

func New(serviceContainer services.Container) *Application {
	return &Application{
		authHandler:       handlers.NewAuthHandler(serviceContainer),
		trackHandler:      handlers.NewTrackHandler(serviceContainer),
		albumHandler:      handlers.NewAlbumHandler(serviceContainer),
		shelfHandler:      handlers.NewShelfHandler(serviceContainer),
		userHandler:       handlers.NewUserHandler(serviceContainer),
		accountHandler:    handlers.NewAccountHandler(serviceContainer),
		productHandler:    handlers.NewProductHandler(serviceContainer),
		homeHandler:       handlers.NewHomeHandler(serviceContainer),
		preferenceHandler: handlers.NewPreferenceHandler(serviceContainer),
		subjectHandler:    handlers.NewSubjectHandler(serviceContainer),
	}
}

func (this *Application) GetServer() *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewHTTPValidator()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"X-Requested-With", "X-Request-ID", "Content-Type", "Accept"},
	}))
	e.Use(middleware.BodyLimit("5M"))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	e.HideBanner = true
	e.Use(middleware.RequestID())
	e.Debug = true

	env := os.Getenv("APP_ENV")
	if env == "production" {
		e.Use(middleware.Recover())
		e.Debug = false
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	e.GET("/status", func(c echo.Context) error {
		host, _ := os.Hostname()
		var status string
		for _, conn := range connections {
			status = fmt.Sprintf("%s\n%s: %s", status, conn.Target(), conn.GetState().String())
		}

		result := fmt.Sprintf("Running on %s since %s\n%s", host, createdAt, status)
		return c.String(http.StatusOK, result)
	})

	auth := e.Group("/auth")
	codes := auth.Group("/code")
	codes.POST("/generate", this.accountHandler.CreateAccount)
	codes.POST("/validate", this.accountHandler.ValidateAccount)

	login := auth.Group("/login")
	login.POST("/generate", this.accountHandler.CreateAccount)
	login.POST("/app/validate", this.accountHandler.LoginApp)
	login.POST("/web/validate", this.accountHandler.LoginWeb)

	user := auth.Group("/users")
	user.POST("", this.userHandler.CreateUser)

	valid := e.Group("/validate")
	valid.Use(this.authHandler.Authenticate)
	valid.GET("/app", this.authHandler.ReturnValidate)

	api := e.Group("/api")
	api.Use(middlewares.LogRequestID())
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// req := c.Request()
			// contentType := req.Header.Get("Content-Type")
			// if contentType != "application/json" {
			//	return c.JSON(http.StatusBadRequest, "Please set Content-Type: application/json in the HTTP header.")
			// }
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			log.SetRequestID(requestID)
			log.Info(c.Path(), c.QueryString())
			return next(c)
		}
	})
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().After(func() {
				log.Info(c.Path())
			})
			return next(c)
		}
	})

	api.GET("/me", this.authHandler.GetAuthenticatedUser)
	api.GET("/logout", this.authHandler.Logout)

	albums := api.Group("/albums")
	albums.POST("", this.albumHandler.Create)
	albums.PUT("/:id", this.albumHandler.Update)
	albums.GET("", this.albumHandler.ReadAll)
	albums.GET("/:id", this.albumHandler.ReadOne)
	albums.DELETE("/:id", this.albumHandler.Delete)
	albums.PUT("/publish/:id", this.albumHandler.Publish)
	albums.PUT("/unpublish/:id", this.albumHandler.Unpublish)
	albums.GET("/published", this.albumHandler.ReadAllPublishedOnly)

	tracks := api.Group("/tracks")
	tracks.GET("", this.trackHandler.ReadAll)
	tracks.POST("", this.trackHandler.Create)
	tracks.GET("/:id", this.trackHandler.ReadOne)
	tracks.PUT("/:id", this.trackHandler.Update)
	tracks.DELETE("/:id", this.trackHandler.Delete)

	shelves := api.Group("/shelves")
	shelves.GET("", this.shelfHandler.ReadAll)
	shelves.GET("/latest", this.shelfHandler.ReadAll)
	shelves.GET("/:id", this.shelfHandler.ReadOne)
	shelves.GET("/published/albums", this.shelfHandler.FetchShelvesOnlyWithPublishedAlbums)
	shelves.POST("", this.shelfHandler.Create)
	shelves.PUT("/:id", this.shelfHandler.Update)
	shelves.DELETE("/:id", this.shelfHandler.Delete)

	users := api.Group("/users")
	users.GET("", this.userHandler.ReadAllUsers)

	products := api.Group("/products")
	products.GET("", this.productHandler.ReadAll)
	products.GET("/:id", this.productHandler.ReadOne)
	products.POST("", this.productHandler.Create)
	products.PUT("/:id", this.productHandler.Update)

	home := api.Group("/home")
	home.GET("", this.homeHandler.ReadAll)

	preference := api.Group("/preferences/:type")
	preference.PUT("", this.preferenceHandler.Update)

	subjects := api.Group("/subjects")
	subjects.POST("", this.subjectHandler.Create)
	subjects.PUT("/:id", this.subjectHandler.Update)
	subjects.GET("/:id", this.subjectHandler.ReadOne)
	subjects.GET("", this.subjectHandler.ReadAll)
	subjects.DELETE("/:id", this.subjectHandler.Delete)

	return e
}

func SetRpcConnections(c []*grpc.ClientConn) {
	connections = c
}
