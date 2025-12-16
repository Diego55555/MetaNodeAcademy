// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"Task5/internal/svc"
	"Task5/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询所有评论
func NewQueryAllCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllCommentLogic {
	return &QueryAllCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllCommentLogic) QueryAllComment() (resp []types.QueryCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
