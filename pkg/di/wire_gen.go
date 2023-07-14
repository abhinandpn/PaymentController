// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/abhinandpn/StockX/pkg/api"
	"github.com/abhinandpn/StockX/pkg/api/handler"
	"github.com/abhinandpn/StockX/pkg/config"
	"github.com/abhinandpn/StockX/pkg/db"
	"github.com/abhinandpn/StockX/pkg/repository"
	"github.com/abhinandpn/StockX/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
