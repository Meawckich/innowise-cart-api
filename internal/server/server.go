package server

import (
	"context"
	"log"

	"cart-api/internal/config"

	//"cart-api/internal/pkg/endpoint"
	//"cart-api/internal/repository"

	"net/http"

	_ "cart-api/docs"
	//httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	cfg *config.Config
	mux *http.ServeMux
}

func NewServer(config *config.Config, mux *http.ServeMux) *Server {
	return &Server{
		cfg: config,
		mux: mux,
	}
}

func (s *Server) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    s.cfg.Port,
		Handler: s.mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v", err)
		}
	}()

	log.Printf("listening on %s", s.cfg.Port)
	<-ctx.Done()

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
	}

	log.Println("final")
	return nil
}

// func (s *Server) Run() error {

// 	wrappedMux := middleware.NewLoggingMiddleware(router)

// 	muxServer := &http.Server{
// 		Addr:    s.cfg.Port,
// 		Handler: wrappedMux,
// 	}

// 	log.Printf("Server listen on port:  %s", s.cfg.Port)

// 	if err := muxServer.ListenAndServe(); err != nil {
// 		return err
// 	}

// 	return nil
// }
