package Monitor
type ContainerJson struct {
        Type  string `json:"type"`
        Data struct {
          Container_uuid string `json:"container_uuid"`
          Environment_id string `json:"environment_id"`
          Namespace string `json:"namespace"`
          Container_name string `json:"container_name"`
          Timestamp string `json:"timestamp"`
          Log_info struct {
            Log_time string `json:"log_time"`
            Source string `json:"source"`
            Message string `json:"message"`
          } `json:"log_info"`
        }  `json:"data"`
    }

type ContainerFileSystem struct {
              Container_filesystem_name string `json:"container_filesystem_name"`
              Container_filesystem_type string `json:"container_filesystem_type"`
              Container_filesystem_capacity int `json:"container_filesystem_capacity"`
              Container_filesystem_usage int `json:"container_filesystem_usage"`
} 

type StatsInfo struct {
          //Timestamp string `json:"timestamp"`
          Container_cpu_usage_seconds_total int `json:"container_cpu_usage_seconds_total"`
          Container_cpu_user_seconds_total int `json:"container_cpu_user_seconds_total"`
          Container_cpu_system_seconds_total int `json:"container_cpu_system_seconds_total"`
          Container_memory_usage_bytes int `json:"container_memory_usage_bytes"`
          Container_memory_limit_bytes int `json:"container_memory_limit_bytes"`
          Container_memory_cache int `json:"container_memory_cache"`
          Container_memory_rss int `json:"container_memory_rss"`
          Container_memory_swap int `json:"container_memory_swap"`
          Container_network_receive_bytes_total int `json:"container_network_receive_bytes_total"`
          Container_network_receive_packets_total int `json:"container_network_receive_packets_total"`
          Container_network_receive_packets_dropped_total int `json:"container_network_receive_packets_dropped_total"`
          Container_network_receive_errors_total int `json:"container_network_receive_errors_total"`
          Container_network_transmit_bytes_total int `json:"container_network_transmit_bytes_total"`
          Container_network_transmit_packets_total int `json:"container_network_transmit_packets_total"`
          Container_network_transmit_packets_dropped_total int `json:"container_network_transmit_packets_dropped_total"`
          Container_network_transmit_errors_total int `json:"container_network_transmit_errors_total"`
          Container_filesystem []  ContainerFileSystem `json:"container_filesystem"`
          
          Container_diskio_service_bytes_async int `json:"container_diskio_service_bytes_async"`
          Container_diskio_service_bytes_read int `json:"container_diskio_service_bytes_read"`
          Container_diskio_service_bytes_sync int `json:"container_diskio_service_bytes_sync"`
          Container_diskio_service_bytes_total int `json:"container_diskio_service_bytes_total"`
          Container_diskio_service_bytes_write int `json:"container_diskio_service_bytes_write"`
          Container_tasks_state_nr_sleeping int `json:"container_tasks_state_nr_sleeping"`
          Container_tasks_state_nr_running int `json:"container_tasks_state_nr_running"`
          Container_tasks_state_nr_stopped int `json:"container_tasks_state_nr_stopped"`
          Container_tasks_state_nr_uninterruptible int `json:"container_tasks_state_nr_uninterruptible"`
          Container_tasks_state_nr_io_wait int `json:"container_tasks_state_nr_io_wait"`
        }



type ContainerMonitorTag struct {
      Timestamp string `json:"timestamp"`
      Container_uuid string `json:"container_uuid"`
      Environment_id string `json:"environment_id"`
      Container_name string `json:"container_name"`
      Namespace string `json:"namespace"`
      Type string `json:"type"`
    }



type QueryMonitorUnit struct {
      Type  string `json:"type"`
      Data struct {
        Timestamp string `json:"timestamp"`
        Container_uuid string `json:"container_uuid"`
        Environment_id string `json:"environment_id"`
        Container_name string `json:"container_name"`
        Namespace string `json:"namespace"`
        Stats StatsInfo  `json:"stats"`
      } `json:"data"`
    }

type QueryContainerMonitor struct {
  Return_code int `json:"return_code"`
  Query_result [] QueryMonitorUnit `json:"query_result"`
}


type LoginCommand struct {
    Username string `json:"username"`
    Password string `json:"password"`
}