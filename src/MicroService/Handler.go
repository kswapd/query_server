package MicroService
import (
    "github.com/gin-gonic/gin"
    "Monitor"
    "LogInfo"
)
func OnPing(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
}


func OnQueryLogInfo(c *gin.Context){
  LogInfo.QueryLogInfo(c)
}


func OnQueryMonitorInfo(c *gin.Context) {
  Monitor.QueryMonitorInfo(c)
}