package MicroService
import (
    "github.com/gin-gonic/gin"
)


func Start(){
    r := gin.Default()
    r.GET("/ping", OnPing)
    r.POST("/queryLogInfo", OnQueryLogInfo)
    r.POST("/queryMonitorInfo",OnQueryMonitorInfo)
    r.Run() // listen and serve on 0.0.0.0:8080
}