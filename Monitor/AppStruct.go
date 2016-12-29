//package structjson
package Monitor

import (
	"strings"
)

//type framwork struct {
//	Return_code int `json:"return_code"`
//	Query_result
//}

//type Data struct {
//	Environment_id string        `json:"environment_id"`
//	Container_name string        `json:"container_name"`
//	Container_uuid string        `json:"container_uuid"`
//	Namespace      string        `json:"namespace"`
//	TimeStamp      string        `json:"timestamp"`
//	Stats          []interface{} `json:"stats"`
//}

//type QueryJson struct {
//	Query_type      string `json:"query_type"`
//	Container_uuid  string `json:"container_uuid"`
//	Environment_id  string `json:"environment_id"`
//	Start_time      string `json:"start_time"`
//	End_time        string `json:"end_time"`
//	Query_content   string `json:"query_content"`
//	Length_per_page int    `json:"length_per_page"`
//	Page_index      int    `json:"page_index"`
//}

type QueryPerformanceJson struct {
	Query_type     string `json:"query_type"`
	Container_uuid string `json:"container_uuid"`
	Environment_id string `json:"environment_id"`
	Start_time     string `json:"start_time"`
	End_time       string `json:"end_time"`
	Time_step      string `json:"time_step"`
}

/*
app
*/
type AppMySQLStatsJson struct {
	Timestamp                                     string  `json:"timestamp"`
	Connections_total                             float64 `json:"connection_total"`
	Command_query_total                           float64 `json:"command_query_total"`
	Command_insert_total                          float64 `json:"command_insert_total"`
	Command_update_total                          float64 `json:"command_update_total"`
	Command_delete_total                          float64 `json:"command_delete_total"`
	Commands_total                                float64 `json:"commands_total"`
	Handlers_total                                float64 `json:"handlers_total"`
	Connection_errors_total                       float64 `json:"connection_errors_total"`
	Buffer_pool_pages                             float64 `json:"buffer_pool_pages"`
	Thread_connected                              float64 `json:"thread_connected"`
	Max_connections                               float64 `json:"max_connections"`
	Query_response_time_seconds                   float64 `json:"query_reponse_time_seconds"`
	Read_query_response_time_seconds              float64 `json:"read_query_response_time"`
	Write_query_response_time_seconds             float64 `json:"write_query_response_time_seconds"`
	Queries_inside_innodb                         float64 `json:"queries_inside_innodb"`
	Queries_in_queue                              float64 `json:"queries_in_queue"`
	Read_views_open_inside_innodb                 float64 `json:"read_views_open_inside_innodb"`
	Table_statistics_rows_read_total              float64 `json:"table_statistics_rows_read_total"`
	Table_statistics_rows_changed_total           float64 `json:"table_statistics_rows_changed_total"`
	Table_statistics_rows_changed_x_indexes_total float64 `json:"table_statistics_rows_changed_x_indexes_total"`
	Sql_lock_waits_total                          float64 `json:"sql_lock_waits_total"`
	External_lock_waits_total                     float64 `json:"external_lock_waits_total"`
	Table_io_waits_total                          float64 `json:"table_io_waits_total"`
	Table_io_waits_seconds_total                  float64 `json:"table_io_waits_seconds_total"`
}

//type AppMySQLJson struct {
//	Type string `json:"type"`
//	Data struct {
//		Container_uuid string              `json:"container_uuid"`
//		Environment_id string              `json:"environment_id"`
//		Container_name string              `json:"container_name"`
//		Namespace      string              `json:"namespace"`
//		Stats          []AppMySQLStatsInfo `json:"stats"`
//	} `json:"data"`
//}

type AppMySQLJson struct {
	Return_code  int                   `json:"return_code"`
	Query_result []AppMySQLQueryResult `json:"query_result"`
}

type AppMySQLQueryResult struct {
	Type string                  `json:"type"`
	Data AppMySQLQueryResultData `json:"data"`
}

type AppMySQLQueryResultData struct {
	Environment_id string            `json:"environment_id"`
	Container_name string            `json:"container_name"`
	Container_uuid string            `json:"container_uuid"`
	Nmespace       string            `json:"namespace"`
	Timestamp      string            `json:"timestamp"`
	Stats          AppMySQLStatsJson `json:"stats"`
}

type AppNginxStatsJson struct {
	Timestamp          string  `json:"timestamp"`
	Active_connections float64 `json:"active_connections"`
	Accepts            float64 `json:"accepts"`
	Handled            float64 `json:"handled"`
	Requests           float64 `json:"requests"`
	Reading            float64 `json:"reading"`
	Writing            float64 `json:"writing"`
	Waiting            float64 `json:"waiting"`
}

//type AppNginxJson struct {
//	Type string `json:"type"`
//	Data struct {
//		Container_uuid string              `json:"container_uuid"`
//		Environment_id string              `json:"environment_id"`
//		Container_name string              `json:"container_name"`
//		Namespace      string              `json:"namespace"`
//		Stats          []AppNginxStatsJson `json:"stats"`
//	} `json:"data"`
//}

type AppNginxJson struct {
	Return_code  int                   `json:"return_code"`
	Query_result []AppNginxQueryResult `json:"query_result"`
}

type AppNginxQueryResult struct {
	Type string                  `json:"type"`
	Data AppNginxQueryResultData `json:"data"`
}

type AppNginxQueryResultData struct {
	Environment_id string            `json:"environment_id"`
	Container_name string            `json:"container_name"`
	Container_uuid string            `json:"container_uuid"`
	Nmespace       string            `json:"namespace"`
	Timestamp      string            `json:"timestamp"`
	Stats          AppNginxStatsJson `json:"stats"`
}

type AppRedisStatsJson struct {
	Timestamp                    string  `json:"timestamp"`
	Uptime_in_seconds            float64 `json:"uptime_in_seconds"`
	Connected_clients            float64 `json:"connected_clients"`
	Blocked_clients              float64 `json:"blocked_clients"`
	Used_memory                  float64 `json:"used_memory"`
	Used_memory_rss              float64 `json:"used_memory_rss"`
	Used_memory_peak             float64 `json:"used_memory_peak"`
	Used_memory_lua              float64 `json:"used_memory_lua"`
	Max_memory                   float64 `json:"max_memory"`
	Mem_fragmentation_ratio      float64 `json:"mem_fragmentation_ratio"`
	Rdb_changes_since_last_save  float64 `json:"rdb_changes_since_last_save"`
	Rdb_last_bgsave_time_sec     float64 `json:"rdb_last_bgsave_time_sec"`
	Rdb_current_bgsave_time_sec  float64 `json:"rdb_current_bgsave_time_sec"`
	Aof_enabled                  float64 `json:"aof_enabled"`
	Aof_rewrite_in_progress      float64 `json:"aof_rewrite_in_progress"`
	Aof_rewrite_scheduled        float64 `json:"aof_rewrite_scheduled"`
	Aof_last_rewrite_time_sec    float64 `json:"aof_last_rewrite_time_sec"`
	Aof_current_rewrite_time_sec float64 `json:"aof_current_rewrite_time_sec"`
	Total_connections_received   float64 `json:"total_connections_received"`
	Total_commands_processed     float64 `json:"total_commands_processed"`
	Total_net_input_bytes        float64 `json:"total_net_input_bytes"`
	Total_net_output_bytes       float64 `json:"total_net_output_bytes"`
	Rejected_connections         float64 `json:"rejected_connections"`
	Expired_keys                 float64 `json:"expired_keys"`
	Evicted_keys                 float64 `json:"evicted_keys"`
	Keyspace_hits                float64 `json:"keyspace_hits"`
	Keyspace_misses              float64 `json:"keyspace_misses"`
	Pubsub_channels              float64 `json:"pubsub_channels"`
	Pubsub_patterns              float64 `json:"pubsub_patterns"`
	Loading                      float64 `json:"loading"`
	Connected_slaves             float64 `json:"connected_slaves"`
	Repl_backlog_size            float64 `json:"repl_backlog_size"`
	Used_cpu_sys                 float64 `json:"used_cpu_sys"`
	Used_cpu_user                float64 `json:"used_cpu_user"`
	Used_cpu_sys_children        float64 `json:"used_cpu_sys_children"`
	Used_cpu_user_children       float64 `json:"used_cpu_user_children"`
}

//type AppRedisJson struct {
//	Type string `json:"type"`
//	Data struct {
//		Container_uuid string              `json:"container_uuid"`
//		Environment_id string              `json:"environment_id"`
//		Container_name string              `json:"container_name"`
//		Namespace      string              `json:"namespace"`
//		Stats          []AppRedisStatsJson `json:"stats"`
//	} `json:"data"`
//}

type AppRedisJson struct {
	Return_code  int                   `json:"return_code"`
	Query_result []AppRedisQueryResult `json:"query_result"`
}

type AppRedisQueryResult struct {
	Type string                  `json:"type"`
	Data AppRedisQueryResultData `json:"data"`
}

type AppRedisQueryResultData struct {
	Environment_id string            `json:"environment_id"`
	Container_name string            `json:"container_name"`
	Container_uuid string            `json:"container_uuid"`
	Nmespace       string            `json:"namespace"`
	Timestamp      string            `json:"timestamp"`
	Stats          AppRedisStatsJson `json:"stats"`
}

func commandMeasurementsMySQL() string {
	measurements := []string{
		"connections_total",
		"command_query_total",
		"command_insert_total",
		"command_update_total",
		"command_delete_total",
		"commands_total",
		"handlers_total",
		"connection_errors_total",
		"buffer_pool_pages",
		"thread_connected",
		"max_connections",
		"query_response_time_seconds",
		"read_query_response_time_seconds",
		"write_query_response_time_seconds",
		"queries_inside_innodb",
		"queries_in_queue",
		"read_views_open_inside_innodb",
		"table_statistics_rows_read_total",
		"table_statistics_rows_changed_total",
		"table_statistics_rows_changed_x_indexes_total",
		"sql_lock_waits_total",
		"external_lock_waits_total",
		"sql_lock_waits_seconds_total",
		"external_lock_waits_seconds_total",
		"table_io_waits_total",
		"table_io_waits_seconds_total",
	}
	str := strings.Join(measurements, ",")
	return str
}
func commandMeasurementsNginx() string {
	measurements := []string{
		"active_connections",
		"accepts",
		"handled",
		"requests",
		"reading",
		"writing",
		"waiting",
	}

	str := strings.Join(measurements, ",")
	return str
}

func commandMeasurementsRedis() string {
	measurements := []string{
		"uptime_in_seconds",
		"connected_clients",
		"blocked_clients",
		"used_memory",
		"used_memory_rss",
		"used_memory_peak",
		"used_memory_lua",
		"max_memory",
		"mem_fragmentation_ratio",
		"rdb_changes_since_last_save",
		"rdb_last_bgsave_time_sec",
		"rdb_current_bgsave_time_sec",
		"aof_enabled",
		"aof_rewrite_in_progress",
		"aof_rewrite_scheduled",
		"aof_last_rewrite_time_sec",
		"aof_current_rewrite_time_sec",
		"total_connections_received",
		"total_commands_processed",
		"total_net_input_bytes",
		"total_net_output_bytes",
		"rejected_connections",
		"expired_keys",
		"evicted_keys",
		"keyspace_hits",
		"keyspace_misses",
		"pubsub_channels",
		"pubsub_patterns",
		"loading",
		"connected_slaves",
		"repl_backlog_size",
		"used_cpu_sys",
		"used_cpu_user",
		"used_cpu_sys_children",
		"used_cpu_user_children",
	}

	str := strings.Join(measurements, ",")
	return str
}
