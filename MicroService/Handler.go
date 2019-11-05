package MicroService

import (
	"query_server/LogInfo"
	"query_server/Monitor"

	"github.com/gin-gonic/gin"
)

func OnPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func OnQueryLogInfo(c *gin.Context) {
	LogInfo.QueryLogInfo(c)
}

func OnQueryTracingModules(c *gin.Context) {
	LogInfo.QueryTracingModules(c)
}

func OnQueryZipkinInfo(c *gin.Context) {
	LogInfo.QueryZipkinInfo(c)

}

func OnQueryTracingTPS(c *gin.Context) {
	LogInfo.QueryTracingTPS(c)

}

func OnQueryCustomInfo(c *gin.Context) {
	LogInfo.QueryCustomInfo(c)
}

func OnQueryMonitorInfo(c *gin.Context) {
	Monitor.QueryMonitorInfo(c)
}

func OnQueryContainerStatus(c *gin.Context) {
	Monitor.QueryContainerStatus(c)
}
