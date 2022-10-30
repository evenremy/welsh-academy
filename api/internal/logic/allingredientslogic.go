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

func (l *AllingredientsLogic) Allingredients() (resp *types.AllIngredentsReply, err error) {
	// todo: add your logic here and delete this line

	return
}
