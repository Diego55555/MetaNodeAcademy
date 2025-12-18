// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"net/http"

	"Task5/internal/svc"
	"Task5/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
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

func (l *QueryAllCommentLogic) QueryAllComment(req *types.QueryCommentReq) (resp []types.QueryCommentResp, err error) {
	comments, err := l.svcCtx.CommentsModel.FindAll(l.ctx, req.PostId)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "内部错误")
	}

	for _, comment := range comments {
		resp = append(resp, types.QueryCommentResp{
			Id:      comment.Id,
			Content: comment.Content,
		})
	}

	return resp, nil
}
