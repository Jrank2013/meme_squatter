package web

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type server struct {
	server         http.Server
	shutdownSignal chan struct{}
}

func NewServer(address string) (*server, error) {
	h := gin.New()
	h.Use(gin.Recovery())

	h.StaticFS("/public", PublicAssets)

	t, err := loadTemplate()

	if err != nil {
		return nil, err
	}

	h.SetHTMLTemplate(t)

	h.GET("/", index)

	return &server{
		server: http.Server{
			Addr:              address,
			Handler:           h,
			ReadHeaderTimeout: time.Second * 30,
		},
	}, nil
}

// loadTemplate loads templates embedded by go-assets-builder
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		splitName := strings.Split(name, "/")
		t, err = t.New(splitName[len(splitName)-1]).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (s *server) Start() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		sig := <-c
		fmt.Printf("Shutdown with signal %s\n", sig)
		s.shutdown()
	}()

	err := s.server.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	start := time.Now()

	s.waitForShutdown()

	end := time.Now()

	fmt.Printf("Shutdown took %f\n", end.Sub(start).Seconds())

	return nil

}

func (s *server) shutdown() {
	s.shutdownSignal = make(chan struct{})
	s.server.Shutdown(context.Background())
	close(s.shutdownSignal)
}

func (s *server) waitForShutdown() {
	<-s.shutdownSignal
}
