package LogInfo
type SContainerLogger struct {
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

type SNginxLogger struct {
    Type  string `json:"type"`
    Data struct {
        Container_uuid string `json:"container_uuid"`
        Environment_id string `json:"environment_id"`
        Container_name string `json:"container_name"`
        Namespace string `json:"namespace"`
        App_file string `json:"app_file"`
        Timestamp string `json:"timestamp"`
        Log_info struct {
            Log_time string `json:"log_time"`
            Remote string `json:"remote"`
            Host string `json:"host"`
            Container_uuid string `json:"user"`
            User string `json:"method"`
            Path string `json:"path"`
            Code string `json:"code"`
            Size string `json:"size"`
            Referer string `json:"referer"`
            Agent string `json:"agent"`
        } `json:"log_info"`
    } `json:"data"`
}



type SMysqlLogger struct {
        Type  string `json:"type"`
        Data struct {
          Container_uuid string `json:"container_uuid"`
          Environment_id string `json:"environment_id"`
          Namespace string `json:"namespace"`
          Container_name string `json:"container_name"`
          App_file string `json:"app_file"`
          Timestamp string `json:"timestamp"`
          Log_info struct {
            Log_time string `json:"log_time"`
            Warn_type string `json:"warn_type"`
            Message string `json:"message"`
          } `json:"log_info"`
        }  `json:"data"`
    }






type SRedisLogger struct {
        Type  string `json:"type"`
        Data struct {
          Container_uuid string `json:"container_uuid"`
          Environment_id string `json:"environment_id"`
          Namespace string `json:"namespace"`
          Container_name string `json:"container_name"`
          App_file string `json:"app_file"`
          Timestamp string `json:"timestamp"`
          Log_info struct {
            Log_time string `json:"log_time"`
            Warn_type string `json:"warn_type"`
            Message string `json:"message"`
          } `json:"log_info"`
        }  `json:"data"`
    }

