package Monitor

import (
	//	"fmt"
	"log"
	"query_server/Common"

	//	"fmt"

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

	//	c.BindJSON(&queryInfo)
	/*
	   type QueryMonitorJson struct {
	       Query_type string `json:"query_type"`
	       Container_uuid string `json:"container_uuid"`
	       Environment_id string `json:"environment_id"`
	       Start_time string `json:"start_time"`
	       End_time string `json:"end_time"`
	       Time_step string `json:"time_step"`
	   }
	*/
	queryInfo.Query_type = c.Query("query_type")
	queryInfo.Container_uuid = c.Query("container_uuid")
	queryInfo.Environment_id = c.Query("environment_id")
	queryInfo.Start_time = c.Query("start_time")
	queryInfo.End_time = c.Query("end_time")
	queryInfo.Time_step = c.Query("time_step")
	//c.BindJSON(&queryInfo)

	//	fmt.Println(queryInfo)

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




func QueryContainerStatus(c *gin.Context) {
	var queryInfo Common.QueryContainerStatus

	queryInfo.Query_type = c.Query("query_type")
	queryInfo.Container_uuid = c.Query("container_uuid")
	queryInfo.Start_time = c.Query("start_time")
	queryInfo.End_time = c.Query("end_time")
	//c.BindJSON(&queryInfo)

	//	fmt.Println(queryInfo)

	switch queryInfo.Query_type {
	case "container":
		QueryContainerStatusHandler(c, queryInfo)
	default:
		QueryContainerStatusHandler(c, queryInfo)

	}

}
