package Monitor

import (
	"fmt"
	"query_server/Common"

	"github.com/gin-gonic/gin"
)

const (
	Functions_init = "*"
)

//用于查询的最终命令
func queryCMDFinal(measurements string, qp Common.QueryMonitorJson, functions string) string {
	cmd := "SELECT " + functions + " FROM " + measurements
	cmd += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", qp.Container_uuid)
	cmd += fmt.Sprintf("\"environment_id\"='%s' AND ", qp.Environment_id)
	cmd += fmt.Sprintf("time>='%s' AND time<='%s'", qp.Start_time, qp.End_time)

	return cmd
}

func queryPerformanceHandler(c *gin.Context, queryInfon Common.QueryMonitorJson) {
	var res interface{}

	var appType string //确定app type：redis？Nginx？mysql？

	/*
		从redis，nginx，mysql中各选取一个measurement进行查询，获取查询结果，以确定app type
	*/
	measurementsForConfirmAppType := "connections_total,active_connections,uptime_in_seconds"
	cmdForConfirmAppType := queryCMDFinal(measurementsForConfirmAppType, queryInfon, Functions_init)
	cmdForConfirmAppType += " limit 1"

	retForConfirmAppType := QueryDB(cmdForConfirmAppType, "appdb")

	if len(retForConfirmAppType[0].Series) <= 0 {
		c.JSON(200, gin.H{
			"return_code": 400,
			"err_info":    "query not found",
		})
		return
	}
	indexOfType := indexOf(retForConfirmAppType[0].Series[0].Columns, "type")
	appType = retForConfirmAppType[0].Series[0].Values[0][indexOfType].(string)

	//确定measurements
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

<<<<<<< HEAD
	ret := QueryDB(cmd)
=======
	//cmd = "select mean(*) from used_memory_rss,used_memory_peak limit 2"
	//	fmt.Println(cmd)

	ret := QueryDB(cmd, "appdb")
>>>>>>> 0725cfd5117e542f85532784131e10b17dd5f95d

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
