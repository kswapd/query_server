package LogInfo

import (
	"fmt"
	"query_server/Common"

	"github.com/gin-gonic/gin"
	// "log"
	"encoding/json"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

const (
	ESUrl string = "http://223.202.32.59:8056"
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

	/* m3 := map[string]string{
	   "a": "aa",
	   "b": "bb",
	 }*/

)

func QueryContainerLog(c *gin.Context, queryInfo Common.QueryLogJson) {

	client, err := elastic.NewClient(elastic.SetURL(ESUrl))
	pageIndex := 0
	lengthPerPage := 50
	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return

	}

	if queryInfo.Container_uuid == "" || queryInfo.Start_time == "" || queryInfo.End_time == "" || queryInfo.Page_index == "" || queryInfo.Length_per_page == "" {
		c.JSON(200, InvalidQuery)
		return
	}

	if pageIndex, err = strconv.Atoi(queryInfo.Page_index); err != nil {
		c.JSON(200, InvalidQuery)
		return
	}

	if lengthPerPage, err = strconv.Atoi(queryInfo.Length_per_page); err != nil {
		c.JSON(200, InvalidQuery)
		return
	}

	/* count, err := client.Count("fluentd_from_container_to_es.log-2016.12.01").Do(context.TODO())
	   if err != nil {
	     fmt.Printf("error:%#v.\n",err)
	   }

	   fmt.Printf("No condition got %d.\n",  count)


	   fmt.Println(queryInfo.Container_uuid)*/

	q := elastic.NewBoolQuery()

	//      q = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid))

	//	q = q.Must(elastic.NewMatchQuery("type", queryInfo.Query_type))
	//q = q.Must(elastic.NewMatchQuery("data.log_info.source", "stdout"))
	//q = q.Should(elastic.NewTermQuery("type", "log_container"))
	q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
	q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	// q = q.Must(elastic.NewMatchQuery("data.environment_id", "Network Agent"))

	//dt := elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time)
	// q = q.Must(elastic.NewRangeQuery("data.log_infoq = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid)).log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	src, err := q.Source()
	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return

	}
	data, err := json.Marshal(src)
	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("fluentd_from_container_to_es.log-*") //.Type("film")
	search = search.Query(q)                                              //.Filter(andFilter)

	search = search.From(pageIndex - 1).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	var logResult SQueryContainerLogResult
	var t SContainerLogger
	logResult.Return_code = 200
	logResult.Current_query_result_length = 10
	logResult.All_query_result_length = 100

	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_query_result_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Current_query_result_length = int64(len(searchResult.Hits.Hits))
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}
			logResult.Query_result = append(logResult.Query_result, t)
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}

		c.JSON(200, logResult)

	} else {
		// No hits
		c.JSON(200, QueryNoResult)
	}

}

func QueryAppLog(c *gin.Context, queryInfo Common.QueryLogJson) {

	client, err := elastic.NewClient(elastic.SetURL(ESUrl))
	pageIndex := 0
	lengthPerPage := 50
	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return

	}

	if queryInfo.Container_uuid == "" || queryInfo.Start_time == "" || queryInfo.End_time == "" || queryInfo.Page_index == "" || queryInfo.Length_per_page == "" {
		c.JSON(200, InvalidQuery)
		return
	}

	if pageIndex, err = strconv.Atoi(queryInfo.Page_index); err != nil {
		c.JSON(200, InvalidQuery)
		return
	}

	if lengthPerPage, err = strconv.Atoi(queryInfo.Length_per_page); err != nil {
		c.JSON(200, InvalidQuery)
		return
	}

	/* count, err := client.Count("fluentd_from_container_to_es.log-2016.12.01").Do(context.TODO())
	   if err != nil {
	     fmt.Printf("error:%#v.\n",err)
	   }

	   fmt.Printf("No condition got %d.\n",  count)


	   fmt.Println(queryInfo.Container_uuid)*/

	q := elastic.NewBoolQuery()

	//      q = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid))

	//q = q.Must(elastic.NewMatchQuery("type", "log_file_container"))
	//q = q.Must(elastic.NewMatchQuery("data.log_info.source", "stdout"))
	//q = q.Should(elastic.NewTermQuery("type", "log_container"))
	q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
	q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	// q = q.Must(elastic.NewMatchQuery("data.environment_id", "Network Agent"))

	//dt := elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time)
	// q = q.Must(elastic.NewRangeQuery("data.log_infoq = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid)).log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	src, err := q.Source()
	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return

	}
	data, err := json.Marshal(src)
	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	ss := string(data)
	fmt.Println(ss)

	search := client.Search().Index("fluentd_from_*_to_es.log-*") //.Type("film")
	search = search.Query(q)                                      //.Filter(andFilter)

	search = search.From(pageIndex - 1).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	var appType string
	//	var s SContainerLogger

	var logNginxResult SQueryNginxLogResult
	var logRedisResult SQueryRedisLogResult
	var logMysqlResult SQueryMysqlLogResult

	//var t SNginxLogger
	//var t interface{}
	var tNginx SNginxLogger
	var tMysql SMysqlLogger
	var tRedis SRedisLogger

	/*logResult.Return_code = 200
	  logResult.Current_query_result_length = 10
	  logResult.All_query_result_length = 100*/

	if len(searchResult.Hits.Hits) > 0 {

		//		err := json.Unmarshal(*searchResult.Hits.Hits[0].Source, &s)

		//		if err != nil {
		//			// Deserialization failed
		//		}

		//		appType = s.Type

		var d interface{}
		err := json.Unmarshal(*searchResult.Hits.Hits[0].Source, &d)
		if err != nil {
			// Deserialization failed
		}
		appType = (d.(map[string]interface{}))["type"].(string)
		fmt.Printf("App type:%s.\n", appType)

		if strings.Contains(appType, "nginx") {
			t := tNginx
			logResult := logNginxResult

			logResult.Return_code = 200
			logResult.All_query_result_length = searchResult.Hits.TotalHits
			logResult.Current_query_result_length = int64(len(searchResult.Hits.Hits))
			for _, hit := range searchResult.Hits.Hits {
				err := json.Unmarshal(*hit.Source, &t)
				if err != nil {
				}
				logResult.Query_result = append(logResult.Query_result, t)
			}
			c.JSON(200, logResult)
		} else if strings.Contains(appType, "mysql") {
			t := tMysql
			logResult := logMysqlResult
			logResult.Return_code = 200
			logResult.All_query_result_length = searchResult.Hits.TotalHits
			logResult.Current_query_result_length = int64(len(searchResult.Hits.Hits))
			for _, hit := range searchResult.Hits.Hits {
				err := json.Unmarshal(*hit.Source, &t)
				if err != nil {
				}
				logResult.Query_result = append(logResult.Query_result, t)
			}
			c.JSON(200, logResult)

		} else if strings.Contains(appType, "redis") {
			t := tRedis
			logResult := logRedisResult
			logResult.Return_code = 200
			logResult.All_query_result_length = searchResult.Hits.TotalHits
			logResult.Current_query_result_length = int64(len(searchResult.Hits.Hits))
			for _, hit := range searchResult.Hits.Hits {
				err := json.Unmarshal(*hit.Source, &t)
				if err != nil {
				}
				logResult.Query_result = append(logResult.Query_result, t)
			}
			c.JSON(200, logResult)

		} else {
			c.JSON(200, QueryNoResult)
			return
		}

	} else {
		c.JSON(200, QueryNoResult)
	}

}

func QueryCustomLogFile(c *gin.Context, queryInfo Common.QueryLogJson) {
	fmt.Println("custom...")
	client, err := elastic.NewClient(elastic.SetURL(ESUrl))
	pageIndex := 0
	lengthPerPage := 50

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return
	}

	if queryInfo.Container_uuid == "" || queryInfo.Start_time == "" || queryInfo.End_time == "" || queryInfo.Page_index == "" || queryInfo.Length_per_page == "" {
		c.JSON(200, InvalidQuery)
		return
	}

	if pageIndex, err = strconv.Atoi(queryInfo.Page_index); err != nil {
		c.JSON(200, InvalidQuery)
		return
	}

	if lengthPerPage, err = strconv.Atoi(queryInfo.Length_per_page); err != nil {
		c.JSON(200, InvalidQuery)
		return
	}

	//	elastic.NewTermsAggregation()

	q := elastic.NewBoolQuery()

	//      q = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid))

	q = q.Must(elastic.NewMatchQuery("type", queryInfo.Query_type))
	//q = q.Must(elastic.NewMatchQuery("data.log_info.source", "stdout"))
	//q = q.Should(elastic.NewTermQuery("type", "log_container"))
	q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
	q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	fmt.Println(queryInfo.File_name)
	//	q = q.Must(elastic.NewMatchQuery("data.app_file", queryInfo.File_name))

	// q = q.Must(elastic.NewMatchQuery("data.environment_id", "Network Agent"))

	//dt := elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time)
	// q = q.Must(elastic.NewRangeQuery("data.log_infoq = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid)).log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	src, err := q.Source()
	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return

	}

	data, err := json.Marshal(src)
	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	ss := string(data)
	fmt.Println("ss:", ss)

	search := client.Search().Index("fluentd_from_*_to_es.log-*") //.Type("film")
	search = search.Query(q)                                      //.Filter(andFilter)

	search = search.From(pageIndex - 1).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	//	fmt.Println()

	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}

	//	fmt.Println(searchResult.Hits.Hits[0].Source)

	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	var customLogResult SQueryCustomLogResult
	var customLogger SCustomLogger

	if len(searchResult.Hits.Hits) > 0 {
		customLogResult.Return_code = 200

		for _, hit := range searchResult.Hits.Hits {
			err := json.Unmarshal(*hit.Source, &customLogger)
			if err != nil {
				//解码异常处理
			}
			customLogResult.Query_result = append(customLogResult.Query_result, customLogger)
		}
		c.JSON(200, customLogResult)
	} else {
		c.JSON(200, QueryNoResult)
	}
}

func QueryLogInfo(c *gin.Context) {
	var queryInfo Common.QueryLogJson

	//  c.BindJSON(&queryInfo)

	queryInfo.Query_type = c.Query("query_type")
	queryInfo.Container_uuid = c.Query("container_uuid")
	queryInfo.Environment_id = c.Query("environment_id")
	queryInfo.Start_time = c.Query("start_time")
	queryInfo.End_time = c.Query("end_time")
	queryInfo.Query_content = c.Query("query_content")
	queryInfo.Length_per_page = c.Query("length_per_page")
	queryInfo.Page_index = c.Query("page_index")
	queryInfo.File_name = c.Query("file_name")

	//c.BindJSON(&queryInfo)
	//c.JSON(200, gin.H{"type": queryInfo.Query_type})

	fmt.Printf("%#v.\n", queryInfo)

	switch queryInfo.Query_type {
	case "container":
		QueryContainerLog(c, queryInfo)
	case "app":
		QueryAppLog(c, queryInfo)
	case "custom_log":
		QueryCustomLogFile(c, queryInfo)
	default:
		c.JSON(200, InvalidQuery)
		return
	}

}
