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

type Config struct {
	Domains map[string]string `yaml:"domains"`
	Server  struct {
		Port int16 `yaml:"port"`
	} `yaml:"server"`
}

func configMiddleware(c *Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		host := c.Domains[strings.Split(ctx.Request.Host, ":")[0]]
		ctx.Set("host", host)
		ctx.Next()
	}

}

func NewServer(address string, config *Config) (*server, error) {
	h := gin.New()

	t, err := loadTemplate()

	if err != nil {
		return nil, err
	}

	h.SetHTMLTemplate(t)

	h.Use(gin.Recovery())
	h.Use(configMiddleware(config))
	h.StaticFS("/public", PublicAssets)

	h.GET("/", index)

	return &server{
		server: http.Server{
			Addr:              fmt.Sprintf(":%d", config.Server.Port),
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
