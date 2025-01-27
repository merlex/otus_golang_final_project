package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/merlex/otus-image-previewer/internal/config"
	"github.com/merlex/otus-image-previewer/internal/logger"
	"github.com/merlex/otus-image-previewer/internal/service"
)

type Server struct {
	ctx     context.Context
	ip      string
	port    string
	log     *logger.Logger
	srv     *http.Server
	service service.ImageService
}

func NewServer(ctx context.Context, logger *logger.Logger, conf *config.HTTPServerConfig,
	service service.ImageService,
) *Server {
	return &Server{ctx: ctx, log: logger, ip: conf.IP, port: conf.Port, service: service}
}

func (s *Server) Start(ctx context.Context) {
	h := NewProxyHandler(ctx, s.log, s.service)

	r := mux.NewRouter().UseEncodedPath()
	r.Path("/").HandlerFunc(h.hellowHandler)
	r.PathPrefix("/fill/").Handler(http.StripPrefix("/fill/", h))
	server := &http.Server{
		Addr:              strings.Join([]string{s.ip, s.port}, ":"),
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}
	s.srv = server
	s.ctx = ctx

	s.log.Info(fmt.Sprintf("http server start on port %s", s.port))
	go func() {
		_ = server.ListenAndServe()
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	s.Stop(ctx)
}

func (s *Server) Stop(ctx context.Context) {
	err := s.srv.Shutdown(ctx)
	if err != nil {
		s.log.Errorf("could not shutdown http server %v", err)
		return
	}
	s.log.Info("http server stopped")
}
