package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostsModel = (*customPostsModel)(nil)

type (
	// PostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsModel.
	PostsModel interface {
		postsModel
		withSession(session sqlx.Session) PostsModel
		FindAll(ctx context.Context, userId uint64) ([]*Posts, error)
	}

	customPostsModel struct {
		*defaultPostsModel
	}
)

// NewPostsModel returns a model for the database table.
func NewPostsModel(conn sqlx.SqlConn) PostsModel {
	return &customPostsModel{
		defaultPostsModel: newPostsModel(conn),
	}
}

func (m *customPostsModel) withSession(session sqlx.Session) PostsModel {
	return NewPostsModel(sqlx.NewSqlConnFromSession(session))
}

// 查询所有文章
func (m *customPostsModel) FindAll(ctx context.Context, userId uint64) ([]*Posts, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", postsRows, m.table)
	var resp []*Posts
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
