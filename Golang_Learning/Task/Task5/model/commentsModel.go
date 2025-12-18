package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentsModel = (*customCommentsModel)(nil)

type (
	// CommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentsModel.
	CommentsModel interface {
		commentsModel
		withSession(session sqlx.Session) CommentsModel
		FindAll(ctx context.Context, postId uint64) ([]*Comments, error)
	}

	customCommentsModel struct {
		*defaultCommentsModel
	}
)

// NewCommentsModel returns a model for the database table.
func NewCommentsModel(conn sqlx.SqlConn) CommentsModel {
	return &customCommentsModel{
		defaultCommentsModel: newCommentsModel(conn),
	}
}

func (m *customCommentsModel) withSession(session sqlx.Session) CommentsModel {
	return NewCommentsModel(sqlx.NewSqlConnFromSession(session))
}

// 查询所有文章
func (m *customCommentsModel) FindAll(ctx context.Context, postId uint64) ([]*Comments, error) {
	query := fmt.Sprintf("select %s from %s where `post_id` = ?", commentsRows, m.table)
	var resp []*Comments
	err := m.conn.QueryRowsCtx(ctx, &resp, query, postId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
