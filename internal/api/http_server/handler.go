package http_server

import (
	"chatsearch/internal/domain"
	"chatsearch/internal/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	core *domain.SearchEngine
}

func NewHandler(core *domain.SearchEngine) *Handler {
	return &Handler{core: core}
}

func (h *Handler) Query(ctx *gin.Context) {
	var prompt string
	err := ctx.ShouldBind(&prompt)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid query parameters:%s\n", err)
		return
	}

	res, err := h.core.Query(prompt)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "%s", err)
		return
	}

	data, err := json.Marshal(res)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "%s", err)
		return
	}

	ctx.String(http.StatusOK, "%s", data)
}

func (h *Handler) Insert(ctx *gin.Context) {
	var conversationList []*model.Conversation
	err := ctx.ShouldBind(&conversationList)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid post data")
		return
	}

	err = h.core.Insert(conversationList)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "%s", err)
		return
	}

	ctx.String(http.StatusOK, "successfully insert data")
}

func (h *Handler) Update(ctx *gin.Context) {
	var conversationList []*model.Conversation
	err := ctx.ShouldBind(&conversationList)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid post data")
		return
	}

	err = h.core.Update(conversationList)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "%s", err)
		return
	}

	ctx.String(http.StatusOK, "successfully update data")
}

func (h *Handler) Delete(ctx *gin.Context) {
	var ids []int64
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid parameters")
		return
	}

	err = h.core.Delete(ids)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "%s", err)
		return
	}

	ctx.String(http.StatusOK, "successfully delete data")
}
