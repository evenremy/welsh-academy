package logic

import (
	"api/model/ingredient"
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIngredientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIngredientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIngredientLogic {
	return &AddIngredientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddIngredientLogic) AddIngredient(req *types.AddIngredientReq) (resp *types.IngrediendReply, err error) {
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
