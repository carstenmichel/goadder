// Package swaggerui contains a "net/http".Handler that serves the swagger-ui.
package swaggerui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed swagger-ui
	swaggerUIAssets embed.FS
)

// Handler creates a http.Handler that serves the swagger-ui static assets.
func Handler() http.Handler {
	fsys, err := fs.Sub(swaggerUIAssets, "swagger-ui")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(fsys))
}

// Mount mounts a handler serving swagger-ui static assets on e using "/swagger-ui" as
// the prefix.
func Mount(e *echo.Echo) {
	e.GET("/swagger-ui/*", echo.WrapHandler(http.StripPrefix("/swagger-ui/", Handler())))
	e.GET("/swagger-ui", func(ctx echo.Context) error {
		ctx.Redirect(http.StatusTemporaryRedirect, "/swagger-ui/")
		return nil
	})
}
