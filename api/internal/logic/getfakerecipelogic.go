package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"

	"api/internal/svc"
	"api/internal/types"
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
	Id             int64                          `json:"id" faker:"oneof: 1, 100, 55"`
	Title          string                         `json:"title" faker:"sentence"`
	Description    string                         `json:"description" faker:"paragraph"`
	IngredientList []types.IngredientWithQuantity `json:"ingredientList" faker:"-"`
	StageList      []types.Stage                  `json:"stageList" faker:"-"`
}

type fakeIngredientWithQuantity struct {
	IngredientId int64   `json:"ingredientId" faker:"oneof: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10"`
	Quantity     float32 `json:"quantity" faker:"oneof: 1.0, 5.0, 15, 100, 250, 700"`
	Unit         string  `json:"unit" faker:"oneof: slices, ml, cl, cup, g, L"`
}

type fakeStage struct {
	Order       int32  `json:"order" faker:"oneof: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10"`
	Description string `json:"description" faker:"paragraph"`
}

func (l *GetFakeRecipeLogic) GetFakeRecipe(req *types.Recipe) error {
	frecipe := fakeRecipe{}

	return nil
}
