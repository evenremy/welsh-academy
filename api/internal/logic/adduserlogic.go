package logic

import (
	"api/errorx"
	"context"
	"net/http"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserReply, err error) {
	if req.Username == "" {
		return nil, errorx.NewCodeError(10, "Bad username", http.StatusBadRequest)
	}

	var id int64
	id, err = l.svcCtx.UserModel.InsertReturningId(l.ctx, req.Username, false)
	if err != nil {
		return nil, err
	}

	userReply := types.AddUserReply{UserId: id}
	return &userReply, nil
}
