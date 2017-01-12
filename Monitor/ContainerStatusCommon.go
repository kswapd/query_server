package Monitor
type ContainerStatusUnit struct {
          Container_uuid string `json:"container_uuid"`
          Environment_id string `json:"environment_id"`
          Start_time     string `json:"start_time"`
          End_time       string `json:"end_time"`
    }

type ContainerStatus struct {
  Return_code int `json:"return_code"`
  Container_num int `json:"container_num"`
  Query_result [] ContainerStatusUnit `json:"query_result"`
}

