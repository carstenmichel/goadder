//go:generate oapi-codegen --config ./oapi-codegen-config.yaml spec.yaml

package backend

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

//go:embed spec.yaml
var spec []byte

var _ ServerInterface = &Warehouse{}

type Warehouse struct {
}

func RegisterWarehouse(e *echo.Echo) ServerInterface {
	handler := &Warehouse{}

	RegisterHandlersWithBaseURL(e, handler, "api")

	e.GET("/q/swagger-ui", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/swagger-ui/index.html?spec=/q/spec.yaml")
	})
	e.GET("/q/spec.yaml", func(context echo.Context) error {
		return context.Blob(http.StatusOK, "text/x-yaml", spec)
	})
	return handler
}

func (w *Warehouse) AddNumbers(ctx echo.Context) error {
	log.Printf("AddNumbers has started")
	var newNumbers Numbers

	err := ctx.Bind(&newNumbers)
	log.Printf("working with %v\n", newNumbers)
	if err != nil {
		log.Printf("Error %v\n", err)
		ctx.JSON(http.StatusBadRequest, "Invalid format for Numbers")
		return nil
	}
	floatR := newNumbers.Numberone + newNumbers.Numbertwo
	rr := Result{
		Resultingvalue: &floatR,
	}
	return ctx.JSON(http.StatusOK, rr)
}
