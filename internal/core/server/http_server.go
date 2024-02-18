package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/api"
	"server/internal/config"
	"server/internal/middleware"
	"time"
)

type HttpServer struct {
	Server *http.Server
}

func RunServer(apiGroup *api.Group, log *zap.Logger, c *config.Possess, cors *middleware.Cors, timeout *middleware.Timeout) *HttpServer {

	// 路由
	router := gin.Default()
	router.Use(cors.Handle(), timeout.Handle())
	router.GET("timeout", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.JSON(http.StatusOK, "timeout")
	})
	// 注册路由
	apiGroup.InitApi(router)

	time.Sleep(500 * time.Millisecond)

	log.Info(fmt.Sprintf("[项目运行于]：http://127.0.0.1:%d", c.System.Port))

	return &HttpServer{
		Server: &http.Server{
			Addr:           fmt.Sprintf(":%d", c.System.Port),
			Handler:        router,
			ReadTimeout:    20 * time.Second,
			WriteTimeout:   20 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
