//go:build wireinject

package main

import (
	"github.com/Wenkun2001/We-Red-Book/webook/internal/repository"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/repository/cache"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/repository/dao"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/service"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/web"
	ijwt "github.com/Wenkun2001/We-Red-Book/webook/internal/web/jwt"
	"github.com/Wenkun2001/We-Red-Book/webook/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		ioc.InitDB,
		ioc.InitRedis,
		ioc.InitLogger,
		// DAO  部分
		dao.NewUserDAO,

		// cache 部分
		cache.NewCodeCache, cache.NewUserCache,

		// repository 部分
		repository.NewCacheUserRepository,
		repository.NewCodeRepository,

		// service 部分
		ioc.InitSMSService,
		ioc.InitWechatService,
		service.NewUserService,
		service.NewCodeService,

		// handler 部分
		web.NewUserHandler,
		ijwt.NewRedisJWTHandler,
		web.NewOAuth2WechatHandler,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
