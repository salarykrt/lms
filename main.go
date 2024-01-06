package main

import (
	configuration "salarykart/config"
	"salarykart/ctypes"
	"salarykart/handlers"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func GetRouter(logHandler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(logHandler)
	return router
}

func InitRouter(r *gin.Engine, app *handlers.App) (*gin.Engine, error) {

	// set the validators
	if err := ctypes.RegisterValidators(); err != nil {
		slog.Error("Failed to register the validators", slog.String("err", err.Error()))
		return nil, err
	}

	handlers.RegisterRouter(r, app)
	return r, nil
}

func main() {
	//setup logs
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)

	// Get the environment from the ENV variable or default to "development"
	err := configuration.InitEnv()
	if err != nil {
		slog.Error("Error loading environment file", err)
		return
	}

	conn, err := handlers.InitDB()
	if err != nil {
		return
	}

	router, err := InitRouter(GetRouter(sloggin.New(logger)), handlers.CreateApp(conn))
	if err != nil {
		return
	}

	err = router.Run()
	if err != nil {
		slog.Error("Failed to run the router", err)
		return
	}
}
