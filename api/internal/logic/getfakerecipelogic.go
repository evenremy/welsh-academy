package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetFakeRecipeLogic) GetFakeRecipe(req *types.Recipe) error {
	// todo: add your logic here and delete this line

	return nil
}
