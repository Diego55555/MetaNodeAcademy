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

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新文章
func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostReq) (resp *types.CommonResp, err error) {
	user, _ := l.ctx.Value("user").(*model.Users)
	post, err := l.svcCtx.PostsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "未找到文章")
	}
	if user.Id != uint64(post.UserId.Int64) {
		return nil, errors.New(http.StatusBadRequest, "无修改权限")
	}

	post.Title = req.Title
	post.Content = req.Content
	err = l.svcCtx.PostsModel.Update(l.ctx, post)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "内部错误")
	}

	return &types.CommonResp{Message: "更新成功"}, nil
}
