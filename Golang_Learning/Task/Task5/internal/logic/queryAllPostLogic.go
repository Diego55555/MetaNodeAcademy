// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"net/http"

	"Task5/internal/svc"
	"Task5/internal/types"
	"Task5/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type QueryAllPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询所有文章
func NewQueryAllPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllPostLogic {
	return &QueryAllPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllPostLogic) QueryAllPost() (resp []types.QueryPostResp, err error) {
	user, _ := l.ctx.Value("user").(*model.Users)
	posts, err := l.svcCtx.PostsModel.FindAll(l.ctx, user.Id)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "内部错误")
	}

	for _, post := range posts {
		resp = append(resp, types.QueryPostResp{
			Id:      post.Id,
			Title:   post.Title,
			Content: post.Content,
		})
	}

	return resp, nil
}
