package Monitor

import (
	"query_server/Common"
	//"encoding/json"
	"fmt"
	//"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/influxdata/influxdb/client/v2"
)

//最终命令
func queryCMDFinal(measurements string, qp Common.QueryMonitorJson, functions string) string {
	cmd := "SELECT " + functions + " FROM " + measurements

	cmd += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", qp.Container_uuid)
	cmd += fmt.Sprintf("\"environment_id\"='%s' AND ", qp.Environment_id)
	//cmd += fmt.Sprintf("time>='%s' AND time<='%s' GROUP BY time(%s)", qp.Start_time, qp.End_time, qp.Time_step)
	cmd += fmt.Sprintf("time>='%s' AND time<='%s'", qp.Start_time, qp.End_time)

	//cmd += fmt.Sprintf("limit %d", limit)
	return cmd
}

func queryPerformanceHandler(c *gin.Context, queryInfon Common.QueryMonitorJson) {
	//确定app type：redis？Nginx？mysql？
	measurementsForConfirmAppType := "connections_total,active_connections,uptime_in_seconds"
	cmdForConfirmAppType := queryCMDFinal(measurementsForConfirmAppType, queryInfon, "*")
	retForConfirmAppType := QueryDB(cmdForConfirmAppType)

	indexOfType := indexOf(retForConfirmAppType[0].Series[0].Columns, "type")
	appType := retForConfirmAppType[0].Series[0].Values[0][indexOfType]
	if appType == nil {
		fmt.Println("app type 未知")
	} else {
		appType = appType.(string)
	}

	//确定measurements
	//测试用
	appType = "redis"
	var measurements string
	switch appType {
	case "redis":
		{
			measurements = commandMeasurementsRedis()
		}
	case "nginx":
		{
			measurements = commandMeasurementsNginx()
		}
	case "mysql":
		{
			measurements = commandMeasurementsMySQL()
		}
	}

	cmd := queryCMDFinal(measurements, queryInfon, "*")

	//cmd = "select mean(*) from used_memory_rss,used_memory_peak limit 2"
	fmt.Println(cmd)

	ret := QueryDB(cmd)

	//	fmt.Println(ret)
	//聚合查询结果
	var res interface{}
	switch appType {
	case "redis":
		{
			res = parseRedisResult(ret)
		}
	case "nginx":
		{
			//			res = parseNginxResult(ret)
		}
	case "mysql":
		{
			//			res = parseMySQLResult(ret)
		}
	}

	//注：这里应该从查询结果中提取相应字段值更合适
	//	res.Type = queryInfon.Query_type
	//	res.Data.Container_uuid = queryInfon.Container_uuid
	//	res.Data.Environment_id = queryInfon.Environment_id

	c.JSON(200, res)
}

func indexOf(strs []string, dst string) int {
	for k, v := range strs {
		if v == dst {
			return k
		}
	}
	return -1 //未找到dst，返回-1
}
