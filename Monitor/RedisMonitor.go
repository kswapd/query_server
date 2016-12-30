package Monitor

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/influxdata/influxdb/client/v2"
)

//原理参照nginx
func parseRedisResult(res []client.Result) AppRedisJson {

	var appRedisJson AppRedisJson

	if len(res) == 0 {
		return appRedisJson
	}

	redisResult := make(map[string]map[string]float64)

	//遍历res，取出结果
	for _, v := range res[0].Series {
		redisResult[v.Name] = make(map[string]float64)
		index := indexOf(v.Columns, "first_value")

		for _, v1 := range v.Values {
			if v1[index].(string) == "" {
				continue
			}
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

	//container_uuid
	indexOfUuid := indexOf(res[0].Series[0].Columns, "container_uuid")
	container_uuid := res[0].Series[0].Values[0][indexOfUuid].(string)

	//environment_id
	indexOfId := indexOf(res[0].Series[0].Columns, "environment_id")
	environment_id := res[0].Series[0].Values[0][indexOfId].(string)

	//container_name
	indexOfName := indexOf(res[0].Series[0].Columns, "container_name")
	container_name := res[0].Series[0].Values[0][indexOfName].(string)

	//namespace
	indexOfNamespace := indexOf(res[0].Series[0].Columns, "namespace")
	namespace := res[0].Series[0].Values[0][indexOfNamespace].(string)

	//type
	indexOfType := indexOf(res[0].Series[0].Columns, "type")
	appType := res[0].Series[0].Values[0][indexOfType].(string)

	var arqr []AppRedisQueryResult
	for _, v := range timeStat {
		var qrd AppRedisQueryResultData
		qrd.Stats = v
		qrd.Timestamp = v.Timestamp
		qrd.Container_name = container_name
		qrd.Container_uuid = container_uuid
		qrd.Environment_id = environment_id
		qrd.Nmespace = namespace

		var mqr AppRedisQueryResult
		mqr.Data = qrd
		mqr.Type = appType

		arqr = append(arqr, mqr)
	}

	appRedisJson.Query_result = arqr
	appRedisJson.Return_code = 200

	return appRedisJson
}
