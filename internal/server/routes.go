package server

import (
	"io/fs"
	"net/http"

	"tess-personal/internal/web"

	"github.com/a-h/templ"
)

func addRoutes(
	mux *http.ServeMux,
) {
	assetsDir, _ := fs.Sub(web.Assets, "assets")
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assetsDir))))
	mux.Handle("/", templ.Handler(web.Home()))
	mux.Handle("/contact", handleContact())
}
