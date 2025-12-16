// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"Task5/internal/logic"
	"Task5/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询所有文章
func QueryAllPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryAllPostLogic(r.Context(), svcCtx)
		resp, err := l.QueryAllPost()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
