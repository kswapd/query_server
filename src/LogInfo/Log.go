package LogInfo
type LogContainerJson struct {
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