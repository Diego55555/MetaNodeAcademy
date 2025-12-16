// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"database/sql"
	"net/http"

	"Task5/internal/svc"
	"Task5/internal/types"
	"Task5/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文章
func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.CommonResp, err error) {
	user, _ := l.ctx.Value("user").(*model.Users)
	post := model.Posts{Title: req.Title,
		Content: req.Content,
		UserId: sql.NullInt64{
			Int64: int64(user.Id),
			Valid: true,
		}}
	_, err = l.svcCtx.PostsModel.Insert(l.ctx, &post)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "创建文章失败")
	}

	return &types.CommonResp{Message: "创建成功"}, nil
}
