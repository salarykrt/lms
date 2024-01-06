package handlers

import (
	"log/slog"
	configuration "salarykart/config"
	"salarykart/db"
	"salarykart/docs"
	"salarykart/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Conn *db.Conn
}

func InitDB() (*db.Conn, error) {
	conn, err := db.NewConn()
	if err != nil {
		slog.Error("Failed to init the DB", slog.String("err", err.Error()))
		return nil, err
	}
	slog.Info("DB initialized successfully")

	if env := configuration.GetEnvironment(); env == "development" {
		if err = models.MigrateAll(conn); err != nil {
			slog.Error("Failed to migrate the DB", slog.String("err", err.Error()))
			return nil, err
		}
		slog.Info("DB migrated successfully")
	}

	return conn, nil
}

func CreateApp(conn *db.Conn) *App {
	return &App{conn}
}

func setupSwaggerRoutes(router *gin.Engine, basePath string) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = basePath
}

func RegisterRouter(router *gin.Engine, app *App) {
	basePath := "/api/v1"
	setupSwaggerRoutes(router, basePath)
	// v1 := router.Group(basePath)
	// {
	// 	v1.POST("/lead", app.InsertLead)             // Insert lead Initial data
	// 	v1.POST("/docs-upload", app.UploadDocuments) // Upload documents
	// 	v1.POST("/download", app.DownloadDocs)       // Download documents
	// }
}
