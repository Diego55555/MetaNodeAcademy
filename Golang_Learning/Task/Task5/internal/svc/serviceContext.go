// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"Task5/internal/config"
	"Task5/internal/middleware"
	"Task5/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config        config.Config
	UsersModel    model.UsersModel
	PostsModel    model.PostsModel
	CommentsModel model.CommentsModel
	Authenticate  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	usersModel := model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource))
	return &ServiceContext{
		Config:        c,
		UsersModel:    usersModel,
		PostsModel:    model.NewPostsModel(sqlx.NewMysql(c.DB.DataSource)),
		CommentsModel: model.NewCommentsModel(sqlx.NewMysql(c.DB.DataSource)),
		Authenticate:  middleware.NewAuthenticateMiddleware(&usersModel).Handle,
	}
}
