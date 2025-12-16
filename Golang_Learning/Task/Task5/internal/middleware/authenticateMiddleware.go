// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"Task5/internal/config"
	"Task5/model"
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type AuthenticateMiddleware struct {
	usersModel *model.UsersModel
}

func NewAuthenticateMiddleware(usersModel *model.UsersModel) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{usersModel: usersModel}
}

func (m *AuthenticateMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从Header获取token
		tokenString := r.Header.Get("token")
		if tokenString == "" {
			http.Error(w, `{"code": 401, "msg": "缺少认证token"}`, http.StatusUnauthorized)
			return // 关键：拦截请求，不再向后传递
		}

		// 解析和验证token
		var claims jwt.MapClaims
		token, jwtErr := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return config.G_secret, nil
		})
		if jwtErr != nil || !token.Valid {
			http.Error(w, `{"code": 401, "msg": "无效的token"}`, http.StatusUnauthorized)
			return
		}

		user, err := (*m.usersModel).FindOne(r.Context(), uint64(claims["id"].(float64)))
		if err != nil {
			http.Error(w, `{"code": 401, "msg": "用户不存在"}`, http.StatusUnauthorized)
			return
		}

		// 将用户信息存入请求上下文，供后续逻辑使用
		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	}
}
