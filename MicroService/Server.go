package MicroService

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/ping", OnPing)
	r.POST("/queryLogInfo", OnQueryLogInfo)
	r.GET("/queryLogInfo", OnQueryLogInfo)


	r.POST("/tracing-stats", OnQueryZipkinInfo)
	r.GET("/tracing-stats", OnQueryZipkinInfo)
	r.Run(":8100")
}
