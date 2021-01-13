package controllers

import (
	"context"
	"errors"
	"net/http"
	"proof-concept/logger-meli-cbt-items-api/pruebalog/controllers/middlewares"
	"proof-concept/logger-meli-cbt-items-api/pruebalog/logger"
	"proof-concept/logger-meli-cbt-items-api/pruebalog/logger/tracing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

var router *gin.Engine

func init() {
	router = mlhandlers.DefaultMeliRouter()

	router.Use(middlewares.Tracing)
}

// Controller s
func Controller() {
	router.GET("/prueba-log", func(ctxGin *gin.Context) {
		ctx := tracing.WithFlow(ctxGin.Request.Context(), "flujo de prueba")

		if err := handler(ctx); err != nil {
			logger.Error(ctx, "error en endpoint", err)
			apiErr := apierrors.NewInternalServerApiError("error en endpoint de prueba", err)
			//c.JSON(apiErr.Status(), apiErr)
			ctxGin.AbortWithStatusJSON(apiErr.Status(), apiErr)
			//c.AbortWithError(apiErr.Status(), apiErr)
			return
		}

		ctxGin.String(http.StatusOK, "llamada completada exitosamente")
	})

	router.Run(":8080")
}

func handler(ctx context.Context) error {
	logger.Infof(ctx, "ingresó al handler %s", "prueba logger")
	return errors.New("se ha generado un error de prueba")
	//return nil
}
