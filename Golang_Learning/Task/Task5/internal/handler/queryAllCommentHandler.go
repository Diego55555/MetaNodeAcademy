// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"Task5/internal/logic"
	"Task5/internal/svc"
	"Task5/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询所有评论
func QueryAllCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewQueryAllCommentLogic(r.Context(), svcCtx)
		resp, err := l.QueryAllComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
