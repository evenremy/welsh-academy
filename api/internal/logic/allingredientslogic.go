package logic

import (
	"api/model/ingredient"
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllingredientsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllingredientsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllingredientsLogic {
	return &AllingredientsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type destIngredientList []types.IngrediendReply
type sourceIngredientsList []ingredient.Ingredients

func (l *AllingredientsLogic) Allingredients() (resp *types.AllIngredientsReply, err error) {
	var ingredients sourceIngredientsList
	ingredients, err = l.svcCtx.IngredientModel.FindAll(l.ctx)
	destIngredients := types.AllIngredientsReply{
		IngredientList: destIngredientList(ingredients)} // FIXME
	return &destIngredients, err
}
