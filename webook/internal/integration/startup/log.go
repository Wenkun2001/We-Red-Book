package startup

import "github.com/Wenkun2001/We-Red-Book/webook/pkg/logger"

func InitLogger() logger.LoggerV1 {
	return logger.NewNopLogger()
}
