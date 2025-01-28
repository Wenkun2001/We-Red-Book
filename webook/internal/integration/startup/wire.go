//go:build wireinject

package startup

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

var thirdPartySet = wire.NewSet( // 第三方依赖
	InitDB,
	InitRedis,
	InitLogger)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		thirdPartySet,

		// DAO 部分
		dao.NewUserDAO,
		dao.NewArticleGormDAO,

		// cache 部分
		cache.NewCodeCache, cache.NewUserCache,

		// repository 部分
		repository.NewCachedUserRepository,
		repository.NewCodeRepository,
		repository.NewCachedArticleRepository,

		// Service 部分
		ioc.InitSMSService,
		ioc.InitWechatService,
		service.NewUserService,
		service.NewCodeService,
		service.NewArticleService,
		InitWechatService,

		// handler 部分
		web.NewUserHandler,
		ijwt.NewRedisJWTHandler,
		web.NewOAuth2WechatHandler,
		web.NewArticleHandler,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}

func InitArticleHandler() *web.ArticleHandler {
	wire.Build(
		thirdPartySet,
		dao.NewArticleGormDAO,
		service.NewArticleService,
		web.NewArticleHandler,
		repository.NewCachedArticleRepository)
	return &web.ArticleHandler{}
}
