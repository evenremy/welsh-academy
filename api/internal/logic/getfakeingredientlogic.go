package logic

import (
	"context"
	"github.com/go-faker/faker/v4"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFakeIngredientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFakeIngredientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFakeIngredientLogic {
	return &GetFakeIngredientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type fakeAddIngredientReq struct {
	Name string `json:"name" faker:"unique,oneof:beer,eggs,cheddar,flour,jam,mustard,tabasco,worcestershire sauce,bread"`
}

func (l *GetFakeIngredientLogic) GetFakeIngredient() (resp *types.AddIngredientReq, err error) {
	fakeIngredient := fakeAddIngredientReq{}
	err = faker.FakeData(&fakeIngredient)
	if err != nil {
		faker.ResetUnique()
		return nil, err
	}
	ingredient := types.AddIngredientReq(fakeIngredient)
	return &ingredient, nil
}
