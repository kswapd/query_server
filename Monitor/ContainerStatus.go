

package Monitor

import (
	"query_server/Common"
	"fmt"
	//"log"
	//"strconv"
	"time"
	//"sort"
	"github.com/gin-gonic/gin"
	//"strings"
)

func QueryContainerStatusHandler(c *gin.Context, queryInfo Common.QueryContainerStatus) {


	const RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	const MyDB = "containerdb"
	var finalQuery string
	
	var containerStatus ContainerStatus
	containerStatus.Return_code = 200

/*
type ContainerStatus struct {
  Return_code int `json:"return_code"`
  Container_num int `json:"container_num"`
  Query_result [] ContainerStatusUnit `json:"query_result"`
}

*/
	var queryValidation = true
	_, err := time.Parse(RFC3339Nano, queryInfo.Start_time)
	if err != nil {
		queryValidation = false
		fmt.Printf("error start time.\n")
	}


	_, err = time.Parse(RFC3339Nano, queryInfo.End_time)
	if err != nil {
		queryValidation = false
		fmt.Printf("error end time.\n")
	}

	finalQuery = "select last(value),environment_id from container_cpu_system_seconds_total"
	if queryValidation {
		finalQuery += fmt.Sprintf(" where time>='%s' and time<='%s' group by container_uuid", queryInfo.Start_time, queryInfo.End_time)
	} else {
    	finalQuery += fmt.Sprintf(" where time<now() - 5m group by container_uuid")
	}
	
	fmt.Println(finalQuery)

	ret := QueryDB(finalQuery, MyDB)
	//fmt.Printf("%#v.\n",ret);
	if len(ret[0].Series) > 0 {
		// monitorResult.
		containerStatus.Container_num = len(ret[0].Series)
		timeInd:= indexOf(ret[0].Series[0].Columns, "time")
		//valInd:= indexOf(ret[0].Series[0].Columns, "first")
		envInd:= indexOf(ret[0].Series[0].Columns, "environment_id")
		for index := 0; index < len(ret[0].Series); index++ {
			se := ret[0].Series[index]

			for valIndex := 0; valIndex < len(se.Values); valIndex++ {

				var cu ContainerStatusUnit
				cu.Container_uuid = fmt.Sprintf("%s", se.Tags["container_uuid"])
				cu.Environment_id = fmt.Sprintf("%s", se.Values[0][envInd])
				//cu.Start_time = fmt.Sprintf("%s", se.Values[0][timeInd])
				cu.End_time = fmt.Sprintf("%s", se.Values[0][timeInd])
				containerStatus.Query_result = append(containerStatus.Query_result, cu)
			}
		}


		
	}else{
		c.JSON(200, gin.H{
            "return_code":  400,
            "err_info":"query not found",
        })
        return 
	}


	//c.JSON(200, monitorResult)
	c.JSON(200, containerStatus)

}
