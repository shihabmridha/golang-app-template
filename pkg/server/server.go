package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/shihabmridha/golang-app-template/pkg/logging"
)

type Server struct {
	ip   string
	port string
}

func New(ip, port string) *Server {
	return &Server{
		ip:   ip,
		port: port,
	}
}

func (s *Server) ServeHttp(ctx context.Context, handler http.Handler) error {
	logger := logging.FromContext(ctx)

	logger.Infof("listening on :%s", s.port)

	addr := fmt.Sprintf(s.ip + ":" + s.port)

	if err := http.ListenAndServe(addr, handler); err != nil {
		return fmt.Errorf("server - server.ServeHttp: %w", err)
	}

	return nil
}

func (s *Server) Addr() string {
	return net.JoinHostPort(s.ip, s.port)
}

func (s *Server) Ip() string {
	return s.ip
}

func (s *Server) Port() string {
	return s.port
}
