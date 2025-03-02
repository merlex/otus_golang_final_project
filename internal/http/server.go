package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/merlex/otus_golang_final_project/internal/config"
	"github.com/merlex/otus_golang_final_project/internal/logger"
	"github.com/merlex/otus_golang_final_project/internal/service"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.hellowHandler)
	mux.Handle("/fill/", http.StripPrefix("/fill/", h))
	server := &http.Server{
		Addr:              strings.Join([]string{s.ip, s.port}, ":"),
		Handler:           mux,
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
