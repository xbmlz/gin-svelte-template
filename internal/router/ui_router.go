package router

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/ui"
)

const UIIndexFilePath = "dist/index.html"
const UIRootFilePath = "dist"
const UIAssetsPath = "dist/assets"

type UIRouter struct {
	log core.Logger
	srv core.HTTPServer
}

// _resource is an interface that provides static file, it's a private interface
type _resource struct {
	fs embed.FS
}

// Open to implement the interface by http.FS required
func (r *_resource) Open(name string) (fs.File, error) {
	name = fmt.Sprintf(UIAssetsPath+"/%s", name)
	return r.fs.Open(name)
}

func NewUIRouter(log core.Logger, srv core.HTTPServer) UIRouter {
	return UIRouter{log, srv}
}

func (r UIRouter) Setup() {
	r.log.Debug("UI router is setup")

	// handle the static file by default ui static files
	r.srv.Engine.StaticFS("/assets", http.FS(&_resource{
		fs: ui.UIDist,
	}))

	r.srv.Engine.NoRoute(func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		filePath := ""

		switch urlPath {
		case "/favicon.svg":
			c.Header("content-type", "image/svg+xml")
		default:
			filePath = UIIndexFilePath
			c.Header("content-type", "text/html;charset=utf-8")
			c.Header("X-Frame-Options", "DENY")
		}

		file, err := ui.UIDist.ReadFile(filePath)

		if err != nil {
			r.log.Error(err)
			c.Status(http.StatusNotFound)
			return
		}
		c.Header("content-type", "text/html;charset=utf-8")
		c.Header("X-Frame-Options", "DENY")
		c.String(http.StatusOK, string(file))
	})
}
