// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"Task5/internal/svc"
	"Task5/internal/types"
	"Task5/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户名
func NewGetUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserNameLogic {
	return &GetUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserNameLogic) GetUserName(req *types.GetUserNameReq) (resp *types.GetUserNameResp, err error) {
	value := l.ctx.Value("user")
	user, _ := value.(*model.Users)

	return &types.GetUserNameResp{UserName: user.UserName}, nil
}
