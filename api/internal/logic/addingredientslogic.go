package logic

import (
	"api/model/ingredient"
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIngredientsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIngredientsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIngredientsLogic {
	return &AddIngredientsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddIngredientsLogic) AddIngredients(req *types.AddIngredientReq) (resp *types.IngrediendReply, err error) {
	newIngredient := &ingredient.Ingredients{Name: req.Name}
	resultId, err := l.svcCtx.IngredientModel.InsertReturningId(l.ctx, newIngredient)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	reply := &types.IngrediendReply{
		Name: newIngredient.Name,
		Id:   resultId}
	return reply, nil
}
