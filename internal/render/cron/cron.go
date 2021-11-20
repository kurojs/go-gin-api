package cron

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
	db     db.Repo
}

func New(logger *zap.Logger, db db.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (h *handler) Add() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("cron_task_add", nil)
	}
}

func (h *handler) Edit() core.HandlerFunc {
	type editRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type editResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	return func(ctx core.Context) {
		req := new(editRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		obj := new(editResponse)
		obj.HashID = req.Id
		ctx.HTML("cron_task_edit", obj)
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("cron_task_list", nil)
	}
}
