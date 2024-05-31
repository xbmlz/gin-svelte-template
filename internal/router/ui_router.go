package router

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/ui"
	"io/fs"
	"net/http"
)

const UIIndexFilePath = "build/index.html"
const UIRootFilePath = "build"
const UIAssetsPath = "build/_app"

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

	//r.srv.Engine.Use(static.Serve("/", static.EmbedFolder(ui.UIBuild, UIRootFilePath)))

	// handle the static file by default ui static files
	r.srv.Engine.StaticFS("/_app", http.FS(&_resource{
		fs: ui.UIBuild,
	}))

	r.srv.Engine.NoRoute(func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		filePath := ""

		switch urlPath {
		case "/favicon.png":
			c.Header("content-type", "image/png")
		default:
			filePath = UIIndexFilePath
			c.Header("content-type", "text/html;charset=utf-8")
			c.Header("X-Frame-Options", "DENY")
		}

		file, err := ui.UIBuild.ReadFile(filePath)

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
