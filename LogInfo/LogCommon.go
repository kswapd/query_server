package LogInfo
import (
	"strings"
	"flag"
	"github.com/gin-gonic/gin"

)


const (
	ESUrl string = "http://192.168.100.224:8056"
  //ESUrlList = []string{"str1", "str2", "str3", "str4"}
)

var (
	QueryNoResult = gin.H{
		"return_code": 400,
		"err_info":    "query not found",
	}
	ConnElasticsearchErr = gin.H{
		"return_code": 401,
		"err_info":    "elastic search connection error",
	}
	ErrElasticsearch = gin.H{
		"return_code": 402,
		"err_info":    "elastic search error",
	}
	InvalidQuery = gin.H{
		"return_code": 403,
		"err_info":    "invalid query",
	}
	ArgEsHost = flag.String("elasticsearch_cluster_host", "http://192.168.1.238:9200", "host1:port1, host2:port2")
			
	EsHostArr = strings.Split(*ArgEsHost,",");

)


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


type TracingModules struct {
    Name          string `json:"name"`
    Type   		string `json:"type"`
}

type SZipkinSpan struct {

		TraceId string `json:"traceId"`
		Duration int64 `json:"duration"`
		Name      string `json:"name"`
		Id string `json:"id"`
		Timestamp_millis      int64 `json:"timestamp_millis"`

		Tags   map[string]interface{}  `json:"tags"`
}

 type SQueryZipkinResult struct {
	Return_code                 int64              `json:"return_code"`
	Current_query_result_length int64              `json:"current_query_result_length"`
	All_query_result_length     int64              `json:"all_query_result_length"`
	Type                        string             `json:"type"`
	Query_result                []SZipkinSpan `json:"query_result"`
}


type SZipkinStats struct {


		Type string `json:"type"`
		Annotation string `json:"annotation"`
		Name      string `json:"name"`
		Counts      int64 `json:"counts"`
		All_Hits      int64 `json:"all_hits"`
		Max 		int64 `json:"max"`
		Min 		int64 `json:"min"`
		Avg 		int64 `json:"avg"`
		Sum      int64 `json:"sum"`
}

 type SQueryZipkinStatsResult struct {
	Ret_code                 int64              `json:"ret_code"`
	Ret_length int64              `json:"ret_length"`
	All_length int64              `json:"all_length"`
	Type                        string             `json:"type"`
	Ret                []SZipkinStats `json:"ret"`
}

type SQueryContainerLogResult struct {
	Return_code                 int64              `json:"return_code"`
	Current_query_result_length int64              `json:"current_query_result_length"`
	All_query_result_length     int64              `json:"all_query_result_length"`
	Type                        string             `json:"type"`
	Query_result                []SContainerLogger `json:"query_result"`
}

type SQueryNginxLogResult struct {
	Return_code                 int64          `json:"return_code"`
	Current_query_result_length int64          `json:"current_query_result_length"`
	All_query_result_length     int64          `json:"all_query_result_length"`
	Type                        string         `json:"type"`
	Query_result                []SNginxLogger `json:"query_result"`
}

type SQueryMysqlLogResult struct {
	Return_code                 int64          `json:"return_code"`
	Current_query_result_length int64          `json:"current_query_result_length"`
	All_query_result_length     int64          `json:"all_query_result_length"`
	Type                        string         `json:"type"`
	Query_result                []SMysqlLogger `json:"query_result"`
}

type SQueryRedisLogResult struct {
	Return_code                 int64          `json:"return_code"`
	Current_query_result_length int64          `json:"current_query_result_length"`
	All_query_result_length     int64          `json:"all_query_result_length"`
	Type                        string         `json:"type"`
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
	Type                        string          `json:"type"`
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



type SFileLogger struct {
  File_name                 string           `json:"file_name"`
  Log_start_time string          `json:"log_start_time"`
  Log_end_time     string          `json:"log_end_time"`
}


type SQueryCustomFileResult struct {
  Return_code                 int64           `json:"return_code"`
  Type                        string          `json:"type"`
  Query_result                []SFileLogger `json:"query_result"`
}



