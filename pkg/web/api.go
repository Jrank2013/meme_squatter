package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type server struct {
	server         http.Server
	shutdownSignal chan struct{}
}

func NewServer(address string) (*server, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	h := gin.New()
	h.Use(gin.Recovery())

	h.Static("/public", path.Join(cwd, "pkg", "public"))
	h.LoadHTMLGlob(path.Join(cwd, "pkg", "templates", "*.html"))
	h.GET("/", index)

	return &server{
		server: http.Server{
			Addr:              address,
			Handler:           h,
			ReadHeaderTimeout: time.Second * 30,
		},
	}, nil
}

func (s *server) Start() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		sig := <-c
		fmt.Printf("Shutdown with signal %s\n", sig)
		s.shutdown()
	}()

	s.server.ListenAndServe()
	start := time.Now()

	s.waitForShutdown()

	end := time.Now()

	fmt.Printf("Shutdown took %f\n", end.Sub(start).Seconds())

}

func (s *server) shutdown() {
	s.shutdownSignal = make(chan struct{})
	s.server.Shutdown(context.Background())
	close(s.shutdownSignal)
}

func (s *server) waitForShutdown() {
	<-s.shutdownSignal
}
