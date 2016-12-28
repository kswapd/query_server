package Monitor

import (
	"Common"
	//	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

/*
func QueryAppMonitorInfo(c *gin.Context, queryInfo Common.QueryMonitorJson) {
	c.JSON(200, gin.H{
		"message": "Todo app monitor system.",
	})
}
*/
func QueryMonitorInfo(c *gin.Context) {

	var queryInfo Common.QueryMonitorJson

	c.BindJSON(&queryInfo)

	//fmt.Println(queryInfo)

	switch queryInfo.Query_type {
	case "container":
		QueryContainerMonitorInfo(c, queryInfo)
		break
	case "app":
		//QueryAppMonitorInfo(c, queryInfo)
		queryPerformanceHandler(c, queryInfo)
		break
	default:
		log.Fatalln("Error, invalid query type.")
		break

	}

}
