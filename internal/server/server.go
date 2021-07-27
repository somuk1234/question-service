package server

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Engine *gin.Engine

func Start(ctx context.Context) {
	app := fx.New(
		fx.Provide(NewHttpRouter),
		fx.Invoke(registerHooks),
		fx.Invoke(registerRoutes),
	)
	app.Run()
}

func NewHttpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func registerHooks(lifecycle fx.Lifecycle, router *gin.Engine) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				log.Println("Starting application in :8080")
				go router.Run(":8080")
				return nil
			},
			OnStop: func(context.Context) error {
				log.Println("Stopping application")
				return nil
			},
		},
	)
}
