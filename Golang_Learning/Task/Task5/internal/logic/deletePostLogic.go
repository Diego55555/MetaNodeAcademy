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

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章
func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (resp *types.CommonResp, err error) {
	user, _ := l.ctx.Value("user").(*model.Users)
	post, err := l.svcCtx.PostsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "未找到文章")
	}
	if user.Id != uint64(post.UserId.Int64) {
		return nil, errors.New(http.StatusBadRequest, "无删除权限")
	}

	err = l.svcCtx.PostsModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "内部错误")
	}

	return &types.CommonResp{Message: "删除成功"}, nil
}
