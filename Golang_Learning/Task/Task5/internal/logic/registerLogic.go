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
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "内部错误")
	}

	user := model.Users{UserName: req.UserName, Password: string(hashedPassword)}
	_, err = l.svcCtx.UsersModel.Insert(l.ctx, &user)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "创建用户失败")
	}

	return &types.RegisterResp{Message: "User registered successfully"}, nil
}
