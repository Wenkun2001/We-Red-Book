package startup

import (
	"github.com/Wenkun2001/We-Red-Book/webook/internal/service/oauth2/wechat"
	"github.com/Wenkun2001/We-Red-Book/webook/pkg/logger"
)

func InitWechatService(l logger.LoggerV1) wechat.Service {
	return wechat.NewService("", "", l)
}
