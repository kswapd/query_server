//package structjson
package Monitor

import (
	"strings"
)

//type ContainerJson struct {
//	Type string `json:"type"`
//	Data struct {
//		Container_uuid string `json:"container_uuid"`
//		Environment_id string `json:"environment_id"`
//		Namespace      string `json:"namespace"`
//		Container_name string `json:"container_name"`
//		Timestamp      string `json:"timestamp"`
//		Log_info       struct {
//			Log_time string `json:"log_time"`
//			Source   string `json:"source"`
//			Message  string `json:"message"`
//		} `json:"log_info"`
//	} `json:"data"`
//}

type QueryJson struct {
	Query_type      string `json:"query_type"`
	Container_uuid  string `json:"container_uuid"`
	Environment_id  string `json:"environment_id"`
	Start_time      string `json:"start_time"`
	End_time        string `json:"end_time"`
	Query_content   string `json:"query_content"`
	Length_per_page int    `json:"length_per_page"`
	Page_index      int    `json:"page_index"`
}

/*
type ContainerJson struct {
	Timestamp      string `json:"timestamp"`
	Container_uuid string `json:"container_uuid"`
	Environment_id string `json:"environment_id"`
	Container_name string `json:"container_name"`
	Namespace      string `json:"namespace"`
	Stats          []struct {
		Timestamp                                        string `json:"timestamp"`
		Container_cpu_usage_seconds_total                int    `json:"container_cpu_usage_seconds_total"`
		Container_cpu_user_seconds_total                 int    `json:"container_cpu_user_seconds_total"`
		Container_cpu_system_seconds_total               int    `json:"container_cpu_system_seconds_total"`
		Container_memory_usage_bytes                     int    `json:"container_memory_usage_bytes"`
		Container_memory_limit_bytes                     int    `json:"container_memory_limit_bytes"`
		Container_memory_cache                           int    `json:"container_memory_cache"`
		Container_memory_rss                             int    `json:"container_memory_rss"`
		Container_memory_swap                            int    `json:"container_memory_swap"`
		Container_network_receive_bytes_total            int    `json:"container_network_receive_bytes_total"`
		Container_network_receive_packets_total          int    `json:"container_network_receive_packets_total"`
		Container_network_receive_packets_dropped_total  int    `json:"container_network_receive_packets_dropped_total"`
		Container_network_receive_errors_total           int    `json:"container_network_receive_errors_total"`
		Container_network_transmit_bytes_total           int    `json:"container_network_transmit_bytes_total"`
		Container_network_transmit_packets_total         int    `json:"container_network_transmit_packets_total"`
		Container_network_transmit_packets_dropped_total int    `json:"container_network_transmit_packets_dropped_total"`
		Container_network_transmit_errors_total          int    `json:"container_network_transmit_errors_total"`
		Container_filesystem                             []struct {
			Container_filesystem_name     string `json:"container_filesystem_name"`
			Container_filesystem_type     string `json:"container_filesystem_type"`
			Container_filesystem_capacity int    `json:"container_filesystem_capacity"`
			Container_filesystem_usage    int    `json:"container_filesystem_usage"`
		} `json:"container_filesystem"`

		Container_diskio_service_bytes_async     int `json:"container_diskio_service_bytes_async"`
		Container_diskio_service_bytes_read      int `json:"container_diskio_service_bytes_read"`
		Container_diskio_service_bytes_sync      int `json:"container_diskio_service_bytes_sync"`
		Container_diskio_service_bytes_total     int `json:"container_diskio_service_bytes_total"`
		Container_diskio_service_bytes_write     int `json:"container_diskio_service_bytes_write"`
		Container_tasks_state_nr_sleeping        int `json:"container_tasks_state_nr_sleeping"`
		Container_tasks_state_nr_running         int `json:"container_tasks_state_nr_running"`
		Container_tasks_state_nr_stopped         int `json:"container_tasks_state_nr_stopped"`
		Container_tasks_state_nr_uninterruptible int `json:"container_tasks_state_nr_uninterruptible"`
		Container_tasks_state_nr_io_wait         int `json:"container_tasks_state_nr_io_wait"`
	} `json:"stats"`
}
*/
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
type AppMySQLStatsInfo struct {
	Timestamp                                     string `json:"timestamp"`
	Connections_total                             int    `json:"connection_total"`
	Command_query_total                           int    `json:"command_query_total"`
	Command_insert_total                          int    `json:"command_insert_total"`
	Command_update_total                          int    `json:"command_update_total"`
	Command_delete_total                          int    `json:"command_delete_total"`
	Commands_total                                int    `json:"commands_total"`
	Handlers_total                                int    `json:"handlers_total"`
	Connection_errors_total                       int    `json:"connection_errors_total"`
	Buffer_pool_pages                             int    `json:"buffer_pool_pages"`
	Thread_connected                              int    `json:"thread_connected"`
	Max_connections                               int    `json:"max_connections"`
	Query_response_time_seconds                   int    `json:"query_reponse_time_seconds"`
	Read_query_response_time_seconds              int    `json:"read_query_response_time"`
	Write_query_response_time_seconds             int    `json:"write_query_response_time_seconds"`
	Queries_inside_innodb                         int    `json:"queries_inside_innodb"`
	Queries_in_queue                              int    `json:"queries_in_queue"`
	Read_views_open_inside_innodb                 int    `json:"read_views_open_inside_innodb"`
	Table_statistics_rows_read_total              int    `json:"table_statistics_rows_read_total"`
	Table_statistics_rows_changed_total           int    `json:"table_statistics_rows_changed_total"`
	Table_statistics_rows_changed_x_indexes_total int    `json:"table_statistics_rows_changed_x_indexes_total"`
	Qql_lock_waits_total                          int    `json:"sql_lock_waits_total"`
	External_lock_waits_total                     int    `json:"external_lock_waits_total"`
	Table_io_waits_total                          int    `json:"table_io_waits_total"`
	Table_io_waits_seconds_total                  int    `json:"table_io_waits_seconds_total"`
}

type AppMySQLJson struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string              `json:"container_uuid"`
		Environment_id string              `json:"environment_id"`
		Container_name string              `json:"container_name"`
		Namespace      string              `json:"namespace"`
		Stats          []AppMySQLStatsInfo `json:"stats"`
	} `json:"data"`
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

type AppNginxJson struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string              `json:"container_uuid"`
		Environment_id string              `json:"environment_id"`
		Container_name string              `json:"container_name"`
		Namespace      string              `json:"namespace"`
		Stats          []AppNginxStatsJson `json:"stats"`
	} `json:"data"`
}

type AppRedisStatsJson struct {
	Timestamp               string  `json:"timestamp"`
	Uptime_in_seconds       float64 `json:"uptime_in_seconds"`
	Connected_clients       float64 `json:"connected_clients"`
	Blocked_clients         float64 `json:"blocked_clients"`
	Used_memory             float64 `json:"used_memory"`
	Used_memory_rss         float64 `json:"used_memory_rss"`
	Used_memory_peak        float64 `json:"used_memory_peak"`
	Used_memory_lua         float64 `json:"used_memory_lua"`
	Max_memory              float64 `json:"max_memory"`
	Mem_fragmentation_ratio float64 `json:"mem_fragmentation_ratio"`

	// # Persistence
	Rdb_changes_since_last_save  float64 `json:"rdb_changes_since_last_save"`
	Rdb_last_bgsave_time_sec     float64 `json:"rdb_last_bgsave_time_sec"`
	Rdb_current_bgsave_time_sec  float64 `json:"rdb_current_bgsave_time_sec"`
	Aof_enabled                  float64 `json:"aof_enabled"`
	Aof_rewrite_in_progress      float64 `json:"aof_rewrite_in_progress"`
	Aof_rewrite_scheduled        float64 `json:"aof_rewrite_scheduled"`
	Aof_last_rewrite_time_sec    float64 `json:"aof_last_rewrite_time_sec"`
	Aof_current_rewrite_time_sec float64 `json:"aof_current_rewrite_time_sec"`

	// # Stats
	Total_connections_received float64 `json:"total_connections_received"`
	Total_commands_processed   float64 `json:"total_commands_processed"`
	Total_net_input_bytes      float64 `json:"total_net_input_bytes"`
	Total_net_output_bytes     float64 `json:"total_net_output_bytes"`
	Rejected_connections       float64 `json:"rejected_connections"`
	Expired_keys               float64 `json:"expired_keys"`
	Evicted_keys               float64 `json:"evicted_keys"`
	Keyspace_hits              float64 `json:"keyspace_hits"`
	Keyspace_misses            float64 `json:"keyspace_misses"`
	Pubsub_channels            float64 `json:"pubsub_channels"`
	Pubsub_patterns            float64 `json:"pubsub_patterns"`

	// # Replication
	Loading           float64 `json:"loading"`
	Connected_slaves  float64 `json:"connected_slaves"`
	Repl_backlog_size float64 `json:"repl_backlog_size"`

	// # CPU
	Used_cpu_sys           float64 `json:"used_cpu_sys"`
	Used_cpu_user          float64 `json:"used_cpu_user"`
	Used_cpu_sys_children  float64 `json:"used_cpu_sys_children"`
	Used_cpu_user_children float64 `json:"used_cpu_user_children"`
}

type AppRedisJson struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string              `json:"container_uuid"`
		Environment_id string              `json:"environment_id"`
		Container_name string              `json:"container_name"`
		Namespace      string              `json:"namespace"`
		Stats          []AppRedisStatsJson `json:"stats"`
	} `json:"data"`
}

func commandMeasurementsMySQL() string {
	measurements := []string{

		//mysql
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

		//Nginx
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

		//Redis
		"uptime_in_seconds",

		// # Clients
		"connected_clients",
		"blocked_clients",

		// # Memory
		"used_memory",
		"used_memory_rss",
		"used_memory_peak",
		"used_memory_lua",
		"max_memory",
		"mem_fragmentation_ratio",

		// # Persistence
		"rdb_changes_since_last_save",
		"rdb_last_bgsave_time_sec",
		"rdb_current_bgsave_time_sec",
		"aof_enabled",
		"aof_rewrite_in_progress",
		"aof_rewrite_scheduled",
		"aof_last_rewrite_time_sec",
		"aof_current_rewrite_time_sec",

		// # Stats
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

		// # Replication
		"loading",
		"connected_slaves",
		"repl_backlog_size",

		// # CPU
		"used_cpu_sys",
		"used_cpu_user",
		"used_cpu_sys_children",
		"used_cpu_user_children",
	}

	str := strings.Join(measurements, ",")
	return str
}
