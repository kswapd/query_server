package Common

type QueryLogJson struct {
	Query_type      string `json:"query_type"`
	File_name      string `json:"file_name"`
	Container_uuid  string `json:"container_uuid"`
	Environment_id  string `json:"environment_id"`
	Start_time      string `json:"start_time"`
	End_time        string `json:"end_time"`
	Query_content   string `json:"query_content"`
	Length_per_page string    `json:"length_per_page"`
	Page_index      string    `json:"page_index"`
}


type QueryCustomJson struct {
	Container_uuid  string `json:"container_uuid"`
	Environment_id  string `json:"environment_id"`
	Start_time      string `json:"start_time"`
	End_time        string `json:"end_time"`
}


type QueryMonitorJson struct {
	Query_type     string `json:"query_type"`
	Container_uuid string `json:"container_uuid"`
	Environment_id string `json:"environment_id"`
	Start_time     string `json:"start_time"`
	End_time       string `json:"end_time"`
	Time_step      string `json:"time_step"`
}

type QueryContainerStatus struct {
	Query_type     string `json:"query_type"`
	Container_uuid string `json:"container_uuid"`
	Start_time     string `json:"start_time"`
	End_time       string `json:"end_time"`
}
