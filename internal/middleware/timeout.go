package middleware

import (
	gtimeout "github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/internal/config"
	"server/pkg/constant"
	"server/pkg/response"
	"time"
)

type Timeout struct {
	log    *zap.Logger
	config *config.Possess
}

func NewTimeoutMiddleware(log *zap.Logger, c *config.Possess) *Timeout {
	return &Timeout{
		log:    log,
		config: c,
	}
}

func (t *Timeout) Handle() gin.HandlerFunc {
	tot := t.config.System.Timeout
	// 默认超时时间为5秒
	if tot <= 2 {
		tot = 5
	}
	return gtimeout.New(
		gtimeout.WithTimeout(time.Duration(tot)*time.Second),
		gtimeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		gtimeout.WithResponse(func(c *gin.Context) {
			response.FailWithDetail(constant.RequestTimeoutCode, "请求超时，请刷新再使用。。。", c)
		}),
	)
}
