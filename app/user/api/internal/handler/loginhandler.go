package handler

import (
	"net/http"
	"ws_chat/app/user/api/internal/logic"
	"ws_chat/app/user/api/internal/svc"
	"ws_chat/app/user/api/internal/types"
	"ws_chat/common/xresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Response(r, w, nil, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		xresp.Response(r, w, resp, err)
	}
}
