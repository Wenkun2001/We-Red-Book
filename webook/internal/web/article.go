package web

import (
	"github.com/Wenkun2001/We-Red-Book/webook/internal/domain"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/service"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/web/jwt"
	"github.com/Wenkun2001/We-Red-Book/webook/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleHandler struct {
	svc service.ArticleService
	l   logger.LoggerV1
}

func NewArticleHandler(l logger.LoggerV1,
	svc service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		l:   l,
		svc: svc,
	}
}

func (h *ArticleHandler) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/articles")

	//g.PUT("/", h.Edit)
	g.POST("/edit", h.Edit)
	g.POST("/publish", h.Publish)
}

// Edit 接收 Article 输入，返回一个 ID，文章的 ID
func (h *ArticleHandler) Edit(ctx *gin.Context) {
	type Req struct {
		Id      int64
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	uc := ctx.MustGet("user").(jwt.UserClaims)
	id, err := h.svc.Save(ctx, domain.Article{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
		Author: domain.Author{
			Id: uc.Uid,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Msg: "系统错误",
		})
		h.l.Error("保存文章数据失败",
			logger.Int64("uid", uc.Uid),
			logger.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Data: id,
	})
}

func (h *ArticleHandler) Publish(ctx *gin.Context) {
	type Req struct {
		Id      int64
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	//val, ok := ctx.Get("user")
	//if !ok {
	//	ctx.JSON(http.StatusOK, Result{
	//		Code: 4,
	//		Msg:  "未登录",
	//	})
	//	return
	//}
	uc := ctx.MustGet("user").(jwt.UserClaims)
	id, err := h.svc.Publish(ctx, domain.Article{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
		Author: domain.Author{
			Id: uc.Uid,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Msg:  "系统错误",
			Code: 5,
		})
		h.l.Error("发表文章失败",
			logger.Int64("uid", uc.Uid),
			logger.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Data: id,
	})
}
