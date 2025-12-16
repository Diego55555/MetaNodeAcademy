// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"net/http"
	"time"

	"Task5/internal/config"
	"Task5/internal/svc"
	"Task5/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	storedUser, err := l.svcCtx.UsersModel.FindOneByUserName(l.ctx, req.UserName)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password)); err != nil {
		return nil, errors.New(http.StatusForbidden, "密码不正确")
	}

	// 生成 JWT
	expirationTime := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        storedUser.Id,
		"user_name": storedUser.UserName,
		"exp":       expirationTime,
	})

	tokenString, err := token.SignedString(config.G_secret)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "生成Token失败")
	}

	return &types.LoginResp{
		ExpiresAt: expirationTime,
		Token:     tokenString,
		User: types.UserInfo{
			ID:       storedUser.Id,
			UserName: storedUser.UserName,
		},
	}, nil
}
