package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecipeLogic {
	return &GetRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecipeLogic) GetRecipe() (resp *types.FullRecipeReply, err error) {
	// todo: add your logic here and delete this line

	return
}
