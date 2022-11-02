package svc

import (
	"api/internal/config"
	"api/model/ingredient"
	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config          config.Config
	IngredientModel ingredient.IngredientsModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := postgres.New(c.Postgre.Datasource)
	// TODO : init db schema

	return &ServiceContext{
		Config:          c,
		IngredientModel: ingredient.NewIngredientsModel(conn),
	}
}
