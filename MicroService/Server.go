package MicroService
import (
    "github.com/gin-gonic/gin"
)


func Start(){
    r := gin.Default()
    r.GET("/ping", OnPing)
    r.POST("/queryLogInfo", OnQueryLogInfo)
    r.POST("/queryMonitorInfo",OnQueryMonitorInfo)
    r.GET("/queryMonitorInfo",OnQueryMonitorInfo)


    r.GET("/logyun/selectPerformanceInfo.do", OnQueryMonitorInfo)
    r.POST("/logyun/selectPerformanceInfo.do", OnQueryMonitorInfo)
    r.GET("/logyun/selectLogInfo.do",OnQueryLogInfo)
    

    r.GET("/queryContainerStatus",OnQueryContainerStatus)
    r.POST("/queryContainerStatus",OnQueryContainerStatus)
    /*r.GET("/queryMonitorInfo", func(c *gin.Context) {

        id := c.Query("id")
        page := c.DefaultQuery("page", "0")
        name := c.PostForm("name")
        message := c.PostForm("message")

        fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
    })*/
    //r.Run() // listen and serve on 0.0.0.0:8080
    r.Run(":8100")
}