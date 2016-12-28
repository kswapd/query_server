package Monitor

import (
	//	"common"
	"encoding/json"
	"fmt"
	"strconv"

	//"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/client/v2"
)

////最终命令
//func queryCMDFinal(measurements string, qp QueryPerformanceJson, functions string) string {
//	cmd := "SELECT " + functions + " FROM " + measurements

//	cmd += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", qp.Container_uuid)
//	cmd += fmt.Sprintf("\"environment_id\"='%s' AND ", qp.Environment_id)
//	//cmd += fmt.Sprintf("time>='%s' AND time<='%s' GROUP BY time(%s)", qp.Start_time, qp.End_time, qp.Time_step)
//	cmd += fmt.Sprintf("time>='%s' AND time<='%s'", qp.Start_time, qp.End_time)

//	//cmd += fmt.Sprintf("limit %d", limit)
//	return cmd
//}

//func queryPerformanceHandler(c *gin.Context) {
//	//存储json格式的查询请求
//	var queryInfon QueryPerformanceJson
//	c.BindJSON(&queryInfon)

//	fmt.Println(queryInfon)

//	//json->command
//	//这里存在冗余的查询
//	//	measurements := commandMeasurementsMySQL()
//	//	measurements += commandMeasurementsNginx()
//	measurements := commandMeasurementsRedis()

//	cmd := queryCMDFinal(measurements, queryInfon, "*")

//	//cmd = "select mean(*) from used_memory_rss,used_memory_peak limit 2"
//	fmt.Println(cmd)

//	ret := QueryDB(cmd)

//	fmt.Println(ret)
//	//聚合查询结果
//	res := parseRedisResult(ret)

//	//注：这里应该从查询结果中提取相应字段值更合适
//	res.Type = queryInfon.Query_type
//	res.Data.Container_uuid = queryInfon.Container_uuid
//	res.Data.Environment_id = queryInfon.Environment_id

//	c.JSON(200, res)
//}

func parseNginxResult(res []client.Result) AppRedisJson {
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

	return appRedisJson
}

//func indexOf(strs []string, dst string) int {
//	for k, v := range strs {
//		if v == dst {
//			return k
//		}
//	}
//	return -1 //未找到dst，返回-1
//}
