package pruebalib

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
)

// Controller s
func Controller() {
	router := mlhandlers.DefaultMeliRouter()

	router.GET("/prueba-libreria", func(c *gin.Context) {
		if err := handler(); err != nil {
			logger.Error("error en endpoint", err, "prueba:tags")
			apiErr := apierrors.NewInternalServerApiError("error en wrapeado", err)
			c.AbortWithStatusJSON(apiErr.Status(), apiErr)
			return
		}

		c.FullPath()
		c.String(http.StatusOK, "llamada completada exitosamente")
	})

	router.Run(":8081")
}

func handler() error {
	return errors.New("error de negocio generado")
	//return nil
}
