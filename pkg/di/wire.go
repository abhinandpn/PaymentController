//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/abhinandpn/PaymentController/pkg/api"
	handler "github.com/abhinandpn/PaymentController/pkg/api/handler"
	config "github.com/abhinandpn/PaymentController/pkg/config"
	db "github.com/abhinandpn/PaymentController/pkg/db"
	repository "github.com/abhinandpn/PaymentController/pkg/repository"
	usecase "github.com/abhinandpn/PaymentController/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
