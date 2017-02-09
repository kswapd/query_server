package LogInfo

type SContainerLogger struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string `json:"container_uuid"`
		Environment_id string `json:"environment_id"`
		Namespace      string `json:"namespace"`
		Container_name string `json:"container_name"`
		Timestamp      string `json:"timestamp"`
		Log_info       struct {
			Log_time string `json:"log_time"`
			Source   string `json:"source"`
			Message  string `json:"message"`
		} `json:"log_info"`
	} `json:"data"`
}

/* {
   "data": {
     "environment_id": "29619",
     "container_name": "testlog-mynginx07-1",
     "container_uuid": "be4b5b3f-01b1-4962-89fa-bd1548996def",
     "namespace": "testlog",
     "log_info": {
       "source": "stdout",
       "message": "2017/01/13 12:55:53 ~~~ End of Kafka Send.\r",
       "log_time": "2017-01-13T12:55:53.000+00:00"
     },
     "timestamp": "2017-01-13T12:55:53.352+00:00"
   },
   "type": "container"
 },*/

type SQueryContainerLogResult struct {
	Return_code                 int64              `json:"return_code"`
	Current_query_result_length int64              `json:"current_query_result_length"`
	All_query_result_length     int64              `json:"all_query_result_length"`
	Query_result                []SContainerLogger `json:"query_result"`
}

type SQueryNginxLogResult struct {
	Return_code                 int64          `json:"return_code"`
	Current_query_result_length int64          `json:"current_query_result_length"`
	All_query_result_length     int64          `json:"all_query_result_length"`
	Query_result                []SNginxLogger `json:"query_result"`
}

type SQueryMysqlLogResult struct {
	Return_code                 int64          `json:"return_code"`
	Current_query_result_length int64          `json:"current_query_result_length"`
	All_query_result_length     int64          `json:"all_query_result_length"`
	Query_result                []SMysqlLogger `json:"query_result"`
}

type SQueryRedisLogResult struct {
	Return_code                 int64          `json:"return_code"`
	Current_query_result_length int64          `json:"current_query_result_length"`
	All_query_result_length     int64          `json:"all_query_result_length"`
	Query_result                []SRedisLogger `json:"query_result"`
}

type SNginxLogger struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string `json:"container_uuid"`
		Environment_id string `json:"environment_id"`
		Container_name string `json:"container_name"`
		Namespace      string `json:"namespace"`
		App_file       string `json:"app_file"`
		Timestamp      string `json:"timestamp"`
		Log_info       struct {
			Log_time       string `json:"log_time"`
			Remote         string `json:"remote"`
			Host           string `json:"host"`
			Container_uuid string `json:"user"`
			User           string `json:"method"`
			Path           string `json:"path"`
			Code           string `json:"code"`
			Size           string `json:"size"`
			Referer        string `json:"referer"`
			Agent          string `json:"agent"`
		} `json:"log_info"`
	} `json:"data"`
}

type SMysqlLogger struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string `json:"container_uuid"`
		Environment_id string `json:"environment_id"`
		Namespace      string `json:"namespace"`
		Container_name string `json:"container_name"`
		App_file       string `json:"app_file"`
		Timestamp      string `json:"timestamp"`
		Log_info       struct {
			Log_time  string `json:"log_time"`
			Warn_type string `json:"warn_type"`
			Message   string `json:"message"`
		} `json:"log_info"`
	} `json:"data"`
}

type SRedisLogger struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string `json:"container_uuid"`
		Environment_id string `json:"environment_id"`
		Namespace      string `json:"namespace"`
		Container_name string `json:"container_name"`
		App_file       string `json:"app_file"`
		Timestamp      string `json:"timestamp"`
		Log_info       struct {
			Log_time  string `json:"log_time"`
			Warn_type string `json:"warn_type"`
			Message   string `json:"message"`
		} `json:"log_info"`
	} `json:"data"`
}

type SQueryCustomLogResult struct {
	Return_code                 int64           `json:"return_code"`
	Current_query_result_length int64           `json:"current_query_result_length"`
	All_query_result_length     int64           `json:"all_query_result_length"`
	Query_result                []SCustomLogger `json:"query_result"`
}

type SCustomLogger struct {
	Type string `json:"type"`
	Data struct {
		Container_uuid string `json:"container_uuid"`
		Environment_id string `json:"environment_id"`
		App_file       string `json:"app_file"`
		Timestamp      string `json:"timestamp"`
		Log_info       struct {
			Log_time  string `json:"log_time"`
			Warn_type string `json:"warn_type"`
			Message   string `json:"message"`
		} `json:"log_info"`
	} `json:"data"`
}
