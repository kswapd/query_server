package LogInfo

import (
	"encoding/json"
	"fmt"
	"query_server/Common"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)



func QueryContainerLog(c *gin.Context, queryInfo Common.QueryLogJson) {
	

        

	//client, err := elastic.NewClient(elastic.SetURL("http://192.168.100.224:8056", "http://192.168.100.225:8056", "http://192.168.100.226:8056"))
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

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

	// q = q.Must(elastic.NewMatchQuery("type", "log_file_container"))
	//q = q.Must(elastic.NewMatchQuery("data.log_info.source", "stdout"))
	//q = q.Should(elastic.NewTermQuery("type", "log_container"))

 if queryInfo.Container_eid != ""{
    qSubc := elastic.NewBoolQuery()
   //q.Must(elastic.NewTermQuery("data.app_file.keyword", queryInfo.File_name))
   // q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_eid))
    q = q.Must(qSubc)
 }else{
    q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    //q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
 }



	
	q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	if queryInfo.Query_content != "" {
		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.source", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.message", queryInfo.Query_content))
		q = q.Must(qSub)
	}

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

	search := client.Search().Index("container_to_es.log*") //.Type("film")
	search = search.Query(q)                                 //.Filter(andFilter)

	search = search.Sort("data.log_info.log_time", false)
	search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

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
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))
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

	q := elastic.NewBoolQuery()

	//      q = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid))

	//q = q.Must(elastic.NewMatchQuery("type", "log_file_container"))
	//q = q.Must(elastic.NewMatchQuery("data.log_info.source", "stdout"))
	//q = q.Should(elastic.NewTermQuery("type", "log_container"))
	//q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))


if queryInfo.Container_eid != ""{
    qSubc := elastic.NewBoolQuery()
   //q.Must(elastic.NewTermQuery("data.app_file.keyword", queryInfo.File_name))
   // q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_eid))
    q = q.Must(qSubc)
 }else{
    q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    //q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
 }


	q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	if queryInfo.Query_content != "" {
		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.warn_type", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.message", queryInfo.Query_content))

		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.remote", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.host", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.user", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.method", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.path", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.code", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.size", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.referer", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.agent", queryInfo.Query_content))

		q = q.Must(qSub)
	}

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

	search := client.Search().Index("app_*_to_es.log*") //.Type("film")
	search = search.Query(q)                             //.Filter(andFilter)

	search = search.Sort("data.log_info.log_time", false)
	search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	var appType string
	var s SContainerLogger

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

		err := json.Unmarshal(*searchResult.Hits.Hits[0].Source, &s)

		if err != nil {
			// Deserialization failed
		}

		appType = s.Type
		fmt.Printf("App type:%s.\n", appType)
		if strings.Contains(appType, "nginx") {
			t := tNginx
			logResult := logNginxResult
			logResult.Type = "nginx"
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
			logResult.Type = "mysql"
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
			logResult.Type = "redis"
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

func QueryCustomLog(c *gin.Context, queryInfo Common.QueryLogJson) {
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))
	pageIndex := 0
	lengthPerPage := 50

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return
	}

	if queryInfo.Container_uuid == "" || queryInfo.Start_time == "" || queryInfo.End_time == "" || queryInfo.Page_index == "" || queryInfo.Length_per_page == ""  ||queryInfo.File_name == ""{
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

	q := elastic.NewBoolQuery()
	//q = q.Must(elastic.NewMatchQuery("type", "custom_log"))
	//	q = q.Must(elastic.NewMatchQuery("data.log_info.source", "stdout"))
	//	q = q.Should(elastic.NewTermQuery("type", "log_container"))
	//q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))

  if queryInfo.Container_eid != ""{
    qSubc := elastic.NewBoolQuery()
   //q.Must(elastic.NewTermQuery("data.app_file.keyword", queryInfo.File_name))
   // q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_eid))
    q = q.Must(qSubc)
 }else{
    q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    //q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
 }



  q = q.Must(elastic.NewTermQuery("data.app_file.keyword", queryInfo.File_name))
	q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))
	

	if queryInfo.Query_content != "" {
		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.warn_type", queryInfo.Query_content))
		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.message", queryInfo.Query_content))
		q = q.Must(qSub)
	}
	//	if queryInfo.Query_content != "" {
	//		qSub := elastic.NewBoolQuery()
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.warn_type", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.message", queryInfo.Query_content))

	//		//	q = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid))
	//		q = q.Must(elastic.NewMatchQuery("type", "custom_log"))
	//		//	q = q.Must(elastic.NewMatchQuery("data.log_info.source", "stdout"))
	//		//	q = q.Should(elastic.NewTermQuery("type", "log_container"))
	//		q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
	//		q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.remote", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.host", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.user", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.method", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.path", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.code", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.size", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.referer", queryInfo.Query_content))
	//		qSub = qSub.Should(elastic.NewMatchQuery("data.log_info.agent", queryInfo.Query_content))

	//		q = q.Must(qSub)
	//	}

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

	search := client.Search().Index("custom_log_to_es.log*") //.Type("film")
	search = search.Query(q)                                  //.Filter(andFilter)

	search = search.Sort("data.log_info.log_time", false)
	search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	var customLogger SCustomLogger
	var customResult SQueryCustomLogResult
	customResult.Return_code = 200
	customResult.Current_query_result_length = int64(len(searchResult.Hits.Hits))
	customResult.All_query_result_length = searchResult.Hits.TotalHits
	customResult.Type = "custom_log"

	if len(searchResult.Hits.Hits) > 0 {

		for _, hit := range searchResult.Hits.Hits {
			err := json.Unmarshal(*hit.Source, &customLogger)
			if err != nil {

			}
			customResult.Query_result = append(customResult.Query_result, customLogger)
		}

		c.JSON(200, customResult)

	} else {
		c.JSON(200, QueryNoResult)
	}

}

func QueryLogInfo(c *gin.Context) {
	var queryInfo Common.QueryLogJson

	//  c.BindJSON(&queryInfo)

	queryInfo.Query_type = c.Query("query_type")
	queryInfo.Container_uuid = c.Query("container_uuid")
  queryInfo.Container_eid = c.Query("container_eid")
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
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  *ArgEsHost, EsHostArr);
	switch queryInfo.Query_type {
	case "container":
		QueryContainerLog(c, queryInfo)
	case "app":
		QueryAppLog(c, queryInfo)
	case "custom_log":
		QueryCustomLog(c, queryInfo)
	default:
		c.JSON(200, InvalidQuery)
		return
	}

}

func QueryZipkinInfo(c *gin.Context) {
	var queryInfo Common.QueryZipkinSpan

	//  c.BindJSON(&queryInfo)

	queryInfo.Query_type = c.Query("query_type")
	queryInfo.Start_time = c.Query("start_time")
	queryInfo.End_time = c.Query("end_time")
	queryInfo.Lookback,_ = strconv.ParseInt(c.Query("lookback"), 10, 64)
	queryInfo.Max_len, _ = strconv.ParseInt(c.Query("max_len"), 10, 64)
	fmt.Printf("%#v.\n", queryInfo)


	fmt.Printf("ES host info %s,%q.\n",  *ArgEsHost, EsHostArr);
	IsAll = false
	switch queryInfo.Query_type {

	case "all":
		IsAll = true
		ZipkinStatsAll(c, queryInfo)
	case "hystrix":
		ZipkinStatsHystrix(c, queryInfo)
	case "lb":
		ZipkinStatsLoadBalanced(c, queryInfo)
	case "gateway":
		ZipkinStatsJupiter(c, queryInfo)
	case "druid":
		ZipkinStatsDruid(c, queryInfo)
	case "feign":
		ZipkinStatsFeign(c, queryInfo)
	case "cache":
		ZipkinStatsCache(c, queryInfo)
	case "mysql":
		ZipkinStatsMysql(c, queryInfo)
	case "gravity":
		ZipkinStatsGravity(c, queryInfo)
	case "http":
		ZipkinStatsHttp(c, queryInfo)
	default:
		IsAll = true
		ZipkinStatsAll(c, queryInfo)
		return
	}

}

func QueryCustomInfo(c *gin.Context) {
	var queryInfo Common.QueryCustomJson


	queryInfo.Container_uuid = c.Query("container_uuid")
  queryInfo.Container_eid = c.Query("container_eid")
	queryInfo.Environment_id = c.Query("environment_id")
	queryInfo.Start_time = c.Query("start_time")
	queryInfo.End_time = c.Query("end_time")
	//c.BindJSON(&queryInfo)
	//c.JSON(200, gin.H{"type": queryInfo.Query_type})

	fmt.Printf("%#v.\n", queryInfo)
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))
	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return

	}

	if queryInfo.Container_uuid == "" {
		c.JSON(200, InvalidQuery)
		return
	}
	 q := elastic.NewBoolQuery()
	//q := elastic.NewMatchAllQuery()

   if queryInfo.Container_eid != ""{
    qSubc := elastic.NewBoolQuery()
   //q.Must(elastic.NewTermQuery("data.app_file.keyword", queryInfo.File_name))
   // q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    qSubc = qSubc.Should(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_eid))
    q = q.Must(qSubc)
 }else{
    q = q.Must(elastic.NewTermQuery("data.container_uuid.keyword", queryInfo.Container_uuid))
    //q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
 }
	  //q = q.Must(elastic.NewMatchQuery("data.container_uuid", queryInfo.Container_uuid))
	  //q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))*/

	//agg := elastic.NewTermsAggregation().Field("data.app_file")

	/* src, err := q.Source()
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
	   fmt.Println(s)*/
	//all := NewMatchAllQuery()
	search := client.Search().Index("custom_log_to_es.log*").Query(q) //.Pretty(true)
	// search = search.Query(q)//.Filter(andFilter)
	agg := elastic.NewTermsAggregation().Field("data.app_file.keyword").Size(10).OrderByCountDesc()
	search = search.Aggregation("files", agg)
	//search = search.From(pageIndex-1).Size(lengthPerPage)
	// search = search.Sort("data.log_info.log_time", false)
	searchResult, err := search.Do(context.TODO())

	if err != nil {
		c.JSON(200, ErrElasticsearch)
		return
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	//fmt.Printf("%#v---\n", searchResult.Hits.Hits)
	if aggret, found := searchResult.Aggregations.Terms("files"); found {

  		fmt.Printf("Found aggs :%d.\n", len(aggret.Buckets))
  		

      var fileLogger SFileLogger
    var fileResult SQueryCustomFileResult
    fileResult.Return_code = 200
    fileResult.Type = "custom_log"
    fileLogger.Log_start_time = "2017-01-20T06:11:01.820+00:00"
    fileLogger.Log_end_time = "2017-01-20T06:11:01.820+00:00"

  	for _, bucket := range aggret.Buckets {
  			//Genres[bucket.Key.(string)] = bucket.DocCount
  			fmt.Printf("%s, %d...\n", bucket.Key.(string), bucket.DocCount)
        fileLogger.File_name = bucket.Key.(string)
        fileResult.Query_result = append(fileResult.Query_result, fileLogger)
  		//fmt.Printf("%#v...\n", *Genres)
  	}
    c.JSON(200, fileResult)

  }else{
    c.JSON(200, QueryNoResult)
  }

}
