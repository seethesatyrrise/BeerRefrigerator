package server

import (
	"BeerRefrigerator/pkg/config"
	"BeerRefrigerator/pkg/database"
	"BeerRefrigerator/pkg/refrigerator/repo"
	"BeerRefrigerator/pkg/refrigerator/rest"
	"BeerRefrigerator/pkg/refrigerator/service"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) (*server, error) {
	db, err := database.NewDatabaseConnection(cfg)
	if err != nil {
		return nil, err
	}

	refrigeratorRepo := repo.NewRepo(db)
	refrigeratorService := service.NewService(refrigeratorRepo)
	refrigeratorRest := rest.NewRest(refrigeratorService)

	router := gin.New()
	api := router.Group("/api/v1")

	refrigeratorRest.Register(api)

	return &server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.ServerPort,
			Handler: router,
		},
	}, nil
}

func (s *server) Start() error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	fmt.Println("server started")
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	if s.httpServer != nil {
		return s.httpServer.Shutdown(ctx)
	}

	fmt.Println("server stopped")
	return nil
}
