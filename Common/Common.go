package Common
type QueryLogJson struct {
    Query_type string `json:"query_type"` 
    Container_uuid string `json:"container_uuid"`
    Environment_id string `json:"environment_id"`
    Start_time string `json:"start_time"`
    End_time string `json:"end_time"`
    Query_content string `json:"query_content"`
    Length_per_page int `json:"length_per_page"`
    Page_index int `json:"page_index"`
}


type QueryMonitorJson struct {
    Query_type string `json:"query_type"` 
    Container_uuid string `json:"container_uuid"`
    Environment_id string `json:"environment_id"`
    Start_time string `json:"start_time"`
    End_time string `json:"end_time"`
    Time_step string `json:"time_step"`
}




