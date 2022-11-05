package logic

import (
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

// AllIngredients
// Risky
// TODO should be limited/paginated
func (l *AllingredientsLogic) AllIngredients() (resp *types.AllIngredientsReply, err error) {
	ingredients, err := l.svcCtx.IngredientModel.FindAll(l.ctx)

	destList := make([]types.IngrediendReply, len(ingredients))

	for i, e := range ingredients {
		destList[i] = types.IngrediendReply(e)
	}

	destIngredients := types.AllIngredientsReply{IngredientList: destList}
	return &destIngredients, err
}
