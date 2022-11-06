package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"

	"api/internal/svc"
	"api/internal/types"
	"github.com/go-faker/faker/v4"
)

type GetFakeRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFakeRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFakeRecipeLogic {
	return &GetFakeRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type fakeRecipe struct {
	Title          string                             `json:"title" faker:"sentence"`
	Description    string                             `json:"description" faker:"paragraph"`
	IngredientList []types.LinkIngredientWithQuantity `json:"ingredientList" faker:"-"`
	StageList      []types.Stage                      `json:"stageList" faker:"-"`
}

type fakeIngredientWithQuantity struct {
	IngredientId int64   `json:"ingredientId" faker:"unique,oneof: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10"`
	Quantity     float32 `json:"quantity" faker:"oneof: 1.0, 5.0, 15, 100, 250, 700"`
	Unit         string  `json:"unit" faker:"oneof: slices, ml, cl, cup, g, L"`
}

type fakeStage struct {
	Order       int32  `json:"order" faker:"unique,oneof: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10"`
	Description string `json:"description" faker:"paragraph"`
}

func (l *GetFakeRecipeLogic) GetFakeRecipe() (*types.AddRecipeReq, error) {
	faker.ResetUnique()

	// Generate fake ingredients
	ingredients := make([]types.LinkIngredientWithQuantity, 10)
	for i := range ingredients {
		fakeIngr := fakeIngredientWithQuantity{}
		err := faker.FakeData(&fakeIngr)
		if err != nil {
			return nil, err
		}
		ingredients[i] = types.LinkIngredientWithQuantity(fakeIngr)
	}

	// Generate fake stages
	stages := make([]types.Stage, 10)
	for i := range stages {
		fakeStage := fakeStage{}
		err := faker.FakeData(&fakeStage)
		if err != nil {
			return nil, err
		}
		stages[i] = types.Stage(fakeStage)
	}

	frecipe := fakeRecipe{}
	err := faker.FakeData(&frecipe)
	if err != nil {
		return nil, err
	}

	frecipe.IngredientList = ingredients
	frecipe.StageList = stages

	recipe := types.AddRecipeReq(frecipe)
	return &recipe, nil
}
