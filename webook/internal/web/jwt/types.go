package jwt

import "github.com/gin-gonic/gin"

type Handler interface {
	CheckSession(ctx *gin.Context, ssid string) error
	ExtractToken(ctx *gin.Context) string
	ClearToken(ctx *gin.Context) error
	SetLoginToken(ctx *gin.Context, uid int64) error
	SetJWTToken(ctx *gin.Context, uid int64, ssid string) error
}
