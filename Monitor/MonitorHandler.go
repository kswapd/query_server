package Monitor

import (
	"fmt"
	"log"
	"query_server/Common"

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
	queryInfo.Container_uuid = c.Query("container_uuid")
	queryInfo.End_time = c.Query("end_time")
	queryInfo.Environment_id = c.Query("environment_id")
	queryInfo.Query_type = c.Query("query_type")
	queryInfo.Start_time = c.Query("start_time")
	queryInfo.Time_step = c.Query("time_step")

	fmt.Println(queryInfo)
	//	c.BindJSON(&queryInfo)

	//fmt.Println(queryInfo)

	switch queryInfo.Query_type {
	case "container":
		QueryContainerMonitorInfo(c, queryInfo)
		break
	case "app":
		//QueryAppMonitorInfo(c, queryInfo)
		//		fmt.Println(queryInfo)
		queryPerformanceHandler(c, queryInfo)
		break
	default:
		log.Fatalln("Error, invalid query type.")
		break

	}

}
