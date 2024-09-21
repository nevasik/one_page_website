package server

import (
	"Odnostranishka/config"
	lg "gitlab.com/nevasik7/logger"
	"net/http"
)

type Server struct {
	cfg        *config.Config
	httpServer *http.Server
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg:        cfg,
		httpServer: &http.Server{Addr: "0.0.0.0:" + cfg.Server.Port},
	}
}

func (s *Server) Run() {
	var err error
	if err = s.init(); err != nil {
		lg.Panic(err.Error())
	}

	go func() {
		lg.Infof("Starting server on %v", s.httpServer.Addr)
		if err = s.httpServer.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			lg.Panicf("ListenAndServe: {%s}", err)
		}
	}()
	s.listen()
}
