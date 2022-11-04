package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecipeLogic {
	return &AddRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRecipeLogic) AddRecipe(req *types.AddRecipeReq) (resp *types.AddRecipeReply, err error) {
	// todo: add your logic here and delete this line

	return
}