package Monitor

import (
	"query_server/Common"
	"encoding/json"
	"fmt"
	"strconv"

		"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/client/v2"
)




func queryPerformanceHandler(c *gin.Context, queryInfon Common.QueryMonitorJson) {
//	//确定app type：redis？Nginx？mysql？
//	measurementsForConfirmAppType := "connections_total,active_connections,uptime_in_seconds"
//	cmdForConfirmAppType := queryCMDFinal(measurementsForConfirmAppType, queryInfon, "*")
//	retForConfirmAppType := QueryDB(cmdForConfirmAppType)

//	indexOfType := indexOf(retForConfirmAppType[0].Series[0].Columns, "type")
//	appType := retForConfirmAppType[0].Series[0].Values[0][indexOfType]

//	//确定measurements
//	var measurements string
//	switch appType {
//	case "redis":
//		{
//			measurements = commandMeasurementsRedis()
//		}
//	case "nginx":
//		{
//			measurements = commandMeasurementsNginx()
//		}
//	case "mysql":
//		{
//			measurements = commandMeasurementsMySQL()
//		}
//	}

//	cmd := queryCMDFinal(measurements, queryInfon, "*")

//	//cmd = "select mean(*) from used_memory_rss,used_memory_peak limit 2"
//	//	fmt.Println(cmd)

//	ret := QueryDB(cmd)

//	//	fmt.Println(ret)
//	//聚合查询结果
//	switch appType {
//	case "redis":
//		{
//			res := parseRedisResult(ret)
//		}
//	case "nginx":
//		{
//			res := parseNginxResult(ret)
//		}
//	case "mysql":
//		{
//			res := parseMySQLResult(ret)
//		}
//	}

//	//注：这里应该从查询结果中提取相应字段值更合适
//	res.Type = queryInfon.Query_type
//	res.Data.Container_uuid = queryInfon.Container_uuid
//	res.Data.Environment_id = queryInfon.Environment_id

//	c.JSON(200, res)
}

func parseRedisResult(res []client.Result) AppRedisJson {
	var appRedisJson AppRedisJson
	redisResult := make(map[string]map[string]float64)

	//遍历res，取出结果
	for _, v := range res[0].Series {
		redisResult[v.Name] = make(map[string]float64) //map[time]value
		index := indexOf(v.Columns, "value")           //哪个位置存储value

		for _, v1 := range v.Values {
			f64, _ := strconv.ParseFloat(string(v1[index].(json.Number)), 64)
			redisResult[v.Name][v1[0].(string)] = f64
		}
	}

	timeStat := make(map[string]AppRedisStatsJson)

	for k, v := range redisResult {
		for k1, val := range v {
			info := timeStat[k1]
			info.Timestamp = k1
			switch k {
			case "uptime_in_seconds":
				{
					info.Uptime_in_seconds = val
				}
			case "connected_clients":
				{
					info.Connected_clients = val
				}
			case "blocked_clients":
				{
					info.Blocked_clients = val
				}
			case "used_memory":
				{
					info.Used_memory = val
				}
			case "used_memory_rss":
				{
					info.Used_memory_rss = val
				}
			case "used_memory_peak":
				{
					info.Used_memory_peak = val
				}
			case "used_memory_lua":
				{
					info.Used_memory_lua = val
				}
			case "max_memory":
				{
					info.Max_memory = val
				}
			case "mem_fragmentation_ratio":
				{
					info.Mem_fragmentation_ratio = val
				}
			case "rdb_changes_since_last_save":
				{
					info.Rdb_changes_since_last_save = val
				}
			case "rdb_last_bgsave_time_sec":
				{
					info.Rdb_last_bgsave_time_sec = val
				}
			case "rdb_current_bgsave_time_sec":
				{
					info.Rdb_current_bgsave_time_sec = val
				}
			case "aof_enabled":
				{
					info.Aof_enabled = val
				}
			case "aof_rewrite_in_progress":
				{
					info.Aof_rewrite_in_progress = val
				}
			case "aof_rewrite_scheduled":
				{
					info.Aof_rewrite_scheduled = val
				}
			case "aof_last_rewrite_time_sec":
				{
					info.Aof_last_rewrite_time_sec = val
				}
			case "aof_current_rewrite_time_sec":
				{
					info.Aof_current_rewrite_time_sec = val
				}
			case "total_connections_received":
				{
					info.Total_connections_received = val
				}
			case "total_commands_processed":
				{
					info.Total_commands_processed = val
				}
			case "total_net_input_bytes":
				{
					info.Total_net_input_bytes = val
				}
			case "total_net_output_bytes":
				{
					info.Total_net_output_bytes = val
				}
			case "rejected_connections":
				{
					info.Rejected_connections = val
				}
			case "expired_keys":
				{
					info.Expired_keys = val
				}
			case "evicted_keys":
				{
					info.Evicted_keys = val
				}
			case "keyspace_hits":
				{
					info.Keyspace_hits = val
				}
			case "keyspace_misses":
				{
					info.Keyspace_misses = val
				}
			case "pubsub_channels":
				{
					info.Pubsub_channels = val
				}
			case "pubsub_patterns":
				{
					info.Pubsub_patterns = val
				}
			case "loading":
				{
					info.Loading = val
				}
			case "connected_slaves":
				{
					info.Connected_slaves = val
				}
			case "repl_backlog_size":
				{
					info.Repl_backlog_size = val
				}
			case "used_cpu_sys":
				{
					info.Used_cpu_sys = val
				}
			case "used_cpu_user":
				{
					info.Used_cpu_user = val
				}
			case "used_cpu_sys_children":
				{
					info.Used_cpu_sys_children = val
				}
			case "used_cpu_user_children":
				{
					info.Used_cpu_user_children = val
				}

			default:
				{
					fmt.Println("err measurements")
				}
			}
			timeStat[k1] = info
		}
	}

	var tmp []AppRedisStatsJson

	for _, v := range timeStat {
		tmp = append(tmp, v)
	}
	appRedisJson.Data.Stats = tmp

	//container_uuid
	indexOfUuid := indexOf(res[0].Series[0].Columns, "container_uuid")
	//	fmt.Println(indexOfUuid)
	appRedisJson.Data.Container_uuid = res[0].Series[0].Values[0][indexOfUuid].(string)
	//	fmt.Println(appRedisJson.Data.Container_uuid)

	//environment_id
	indexOfId := indexOf(res[0].Series[0].Columns, "environment_id")
	appRedisJson.Data.Environment_id = res[0].Series[0].Values[0][indexOfId].(string)

	//container_name
	indexOfName := indexOf(res[0].Series[0].Columns, "container_name")
	//	fmt.Println(indexOfName)
	appRedisJson.Data.Container_name = res[0].Series[0].Values[0][indexOfName].(string)

	//namespace
	indexOfNamespace := indexOf(res[0].Series[0].Columns, "namespace")
	appRedisJson.Data.Namespace = res[0].Series[0].Values[0][indexOfNamespace].(string)

	//type
	//	indexOfType := indexOf(res[0].Series[0].Columns, "type")
	//	appRedisJson.Data.Container_uuid = res[0].Series[0].Values[0][indexOfType].(string)

	//	fmt.Println(appRedisJson.Data.Container_name)
	//	fmt.Println(appRedisJson.Data.Container_uuid)
	//	fmt.Println(appRedisJson.Data.Environment_id)
	//	fmt.Println(appRedisJson.Data.Namespace)

	return appRedisJson
}

func indexOf(strs []string, dst string) int {
	for k, v := range strs {
		if v == dst {
			return k
		}
	}
	return -1 //未找到dst，返回-1
}
