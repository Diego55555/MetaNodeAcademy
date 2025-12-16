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

type QueryPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询文章
func NewQueryPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostLogic {
	return &QueryPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPostLogic) QueryPost(req *types.QueryPostReq) (resp *types.QueryPostResp, err error) {
	post, err := l.svcCtx.PostsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "未找到文章")
	}

	return &types.QueryPostResp{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}, nil
}
