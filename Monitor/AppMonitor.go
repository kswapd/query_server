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
	var res interface{}
	//确定app type：redis？Nginx？mysql？
	measurementsForConfirmAppType := "connections_total,active_connections,uptime_in_seconds"
	cmdForConfirmAppType := queryCMDFinal(measurementsForConfirmAppType, queryInfon, "*")

	//	fmt.Println("for debug", cmdForConfirmAppType)

	retForConfirmAppType := QueryDB(cmdForConfirmAppType, "appdb")

	//	fmt.Println("debug", retForConfirmAppType)
	if retForConfirmAppType == nil {
		c.JSON(400, res)
		return
	}
	indexOfType := indexOf(retForConfirmAppType[0].Series[0].Columns, "type")
	appType := retForConfirmAppType[0].Series[0].Values[0][indexOfType]
	//	fmt.Println(appType)
	if appType == nil {
		fmt.Println("app type 未知")
	} else {
		appType = appType.(string)
	}

	//确定measurements
	//测试用
	//	appType = "redis"
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
	//	fmt.Println(cmd)

	ret := QueryDB(cmd, "appdb")

	//	fmt.Println(ret)
	//聚合查询结果

	switch appType {
	case "redis":
		{
			res = parseRedisResult(ret)
		}
	case "nginx":
		{
			res = parseNginxResult(ret)
		}
	case "mysql":
		{
			res = parseMySQLResult(ret)
		}
	}

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
