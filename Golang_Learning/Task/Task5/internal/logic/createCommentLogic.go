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

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评论
func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CommonResp, err error) {
	user, _ := l.ctx.Value("user").(*model.Users)
	comment := model.Comments{
		Content: req.Content,
		UserId: sql.NullInt64{
			Int64: int64(user.Id),
			Valid: true,
		},
		PostId: sql.NullInt64{
			Int64: int64(req.PostId),
			Valid: true,
		},
	}
	_, err = l.svcCtx.CommentsModel.Insert(l.ctx, &comment)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "创建评论失败")
	}

	return &types.CommonResp{Message: "创建成功"}, nil
}
