package web

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type server struct {
	server http.Server
}

func NewServer(address string) *server {
	return &server{
		server: http.Server{
			Addr: address,
		},
	}
}

func Start() {
	cwd, _ := os.Getwd()
	r := gin.Default()

	r.Static("/public", path.Join(cwd, "pkg", "public"))

	r.LoadHTMLGlob(path.Join(cwd, "pkg", "templates", "*"))

	r.GET("/", index)

	r.Run(":8080")

}
