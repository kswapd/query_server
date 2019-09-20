package LogInfo

import (
	"encoding/json"
	"fmt"
	"query_server/Common"
	_ "strconv"
	"strings"
	_ "flag"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
    _ "reflect"
    "time"
)

var IsAll = false
func DoQueryZipkinInfo(c *gin.Context, queryInfo Common.QueryZipkinSpan) {
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return

	}

	q := elastic.NewBoolQuery()
    q = q.Must(elastic.NewTermQuery("_q", "LoadBalanced"))
	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-2019-09-19") //.Type("film")
	search = search.Query(q)                                 //.Filter(andFilter)

	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	var maxLen int64
	maxLen = 50
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	var logResult SQueryZipkinResult
	var t SZipkinSpan
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






func DoZipkinStats(c *gin.Context, queryInfo Common.QueryZipkinSpan) {
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}
    q = q.Must(elastic.NewTermQuery("_q", "LoadBalanced"))
	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-2019-09-19") //.Type("film")
	search = search.Query(q)                                 //.Filter(andFilter)

	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)
	var maxLen int64
	maxLen = 50
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"
	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		c.JSON(200, logResult)

	} else {
		// No hits
		
	}

}



func ZipkinStatsLoadBalanced(c *gin.Context, queryInfo Common.QueryZipkinSpan)  SZipkinStats{


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type

	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}




    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    q = q.Must(elastic.NewTermQuery("_q", "LoadBalanced"))
    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }
	
	

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)                                 //.Filter(andFilter)

	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)
	search = search.Sort("timestamp", false)
	var maxLen int64
	maxLen = 50
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = "LoadBalanced"
		spanStat.Type = "lb"
		spanStat.Name = "负载均衡组件"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}


func ZipkinStatsGravity(c *gin.Context, queryInfo Common.QueryZipkinSpan)  SZipkinStats{


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	var anno = "Gravity"
	spanStat.Type  = queryInfo.Query_type
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}


    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    q = q.Must(elastic.NewTermQuery("_q", "GravityNode"))
    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }
	
	

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)                                 //.Filter(andFilter)

	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)
	search = search.Sort("timestamp", false)
	var maxLen int64
	maxLen = 50
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "gravity"
		spanStat.Name = "工作流引擎"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}











func ZipkinStatsFeign(c *gin.Context, queryInfo Common.QueryZipkinSpan)  SZipkinStats{


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "FeignClient"
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}


    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)


    q = q.Must(elastic.NewTermQuery("_q", anno))
    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }
	
	

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)                                 //.Filter(andFilter)

	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)
	search = search.Sort("timestamp", false)
	var maxLen int64
	maxLen = 50
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "feign"
		spanStat.Name = "Feign组件"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}








func ZipkinStatsCache(c *gin.Context, queryInfo Common.QueryZipkinSpan)  SZipkinStats{


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "Cacheable|CachePut|CacheEvict"

	
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}

    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    	matchCacheablePhraseQuery := elastic.NewQueryStringQuery("cacheable*")
    	matchCacheablePhraseQuery.DefaultField("name")

    	matchCachePutPhraseQuery := elastic.NewQueryStringQuery("cacheput")
    	matchCachePutPhraseQuery.DefaultField("name")

    	matchCacheEvictPhraseQuery := elastic.NewQueryStringQuery("cacheevict")
    	matchCacheEvictPhraseQuery.DefaultField("name")

		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(matchCacheablePhraseQuery)
		qSub = qSub.Should(matchCachePutPhraseQuery)
		qSub = qSub.Should(matchCacheEvictPhraseQuery)
		q = q.Must(qSub)

    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)    
	search = search.Sort("timestamp", false)
	var maxLen int64
	maxLen = 50
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "cache"
		spanStat.Name = "分布式缓存"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}










func ZipkinStatsDruid(c *gin.Context, queryInfo Common.QueryZipkinSpan)  SZipkinStats{


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "AlibabaDruid"

	var maxLen int64
	maxLen = 50
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}
	
    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    	matchDruidQuery := elastic.NewQueryStringQuery("druid*")
    	matchDruidQuery.DefaultField("name")


		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(matchDruidQuery)
		q = q.Must(qSub)

    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)    

	search = search.Sort("timestamp", false)
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "druid"
		spanStat.Name = "数据库连接池"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}










func ZipkinStatsMysql(c *gin.Context, queryInfo Common.QueryZipkinSpan)  SZipkinStats{


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "Mysql"

	var maxLen int64
	maxLen = 50
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}
	
    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    	matchDruidQuery := elastic.NewQueryStringQuery("mysqldbservice")
    	matchDruidQuery.DefaultField("remoteEndpoint.serviceName")


		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(matchDruidQuery)
		q = q.Must(qSub)

    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)    

	search = search.Sort("timestamp", false)
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "mysql"
		spanStat.Name = "Mysql数据库"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}






func ZipkinStatsHystrix(c *gin.Context, queryInfo Common.QueryZipkinSpan)  SZipkinStats{


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "Hystrix"

	var maxLen int64
	maxLen = 50
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}
	
    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    	matchDruidQuery := elastic.NewQueryStringQuery("hystrix")
    	matchDruidQuery.DefaultField("name")

		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(matchDruidQuery)
		q = q.Must(qSub)

    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)    

	search = search.Sort("timestamp", false)
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "hystrix"
		spanStat.Name = "熔断降级组件"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}








func ZipkinStatsJupiter(c *gin.Context, queryInfo Common.QueryZipkinSpan) SZipkinStats {


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "Jupiter"

	var maxLen int64
	maxLen = 50
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}
	
    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    	matchDruidQuery := elastic.NewQueryStringQuery("dev-jupiter")
    	matchDruidQuery.DefaultField("localEndpoint.serviceName")


		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(matchDruidQuery)
		q = q.Must(qSub)

    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)    

	search = search.Sort("timestamp", false)
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "gateway"
		spanStat.Name = "网关组件"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}




func ZipkinStatsHttp(c *gin.Context, queryInfo Common.QueryZipkinSpan) SZipkinStats {


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "Http"

	var maxLen int64
	maxLen = 50
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}
	
    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    	matchDruidQuery := elastic.NewQueryStringQuery("dev-jupiter")
    	matchDruidQuery.DefaultField("localEndpoint.serviceName")


		//qSub := elastic.NewBoolQuery()
		//qSub = qSub.Should(matchDruidQuery)
		//q = q.Must(qSub)

		q = q.Must(elastic.NewTermQuery("_q", "http.method"))

    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)    

	search = search.Sort("timestamp", false)
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "http"
		spanStat.Name = "Http"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}



func ZipkinStatsMQ(c *gin.Context, queryInfo Common.QueryZipkinSpan) SZipkinStats {


	var logResult SQueryZipkinStatsResult
	var t SZipkinSpan
	var spanStat SZipkinStats
	spanStat.Type  = queryInfo.Query_type
	var anno = "RabbitMQ"

	var maxLen int64
	maxLen = 50
	EsHostArr = strings.Split(*ArgEsHost,",");
	fmt.Printf("ES host info %s,%q.\n",  ArgEsHost, EsHostArr);
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return spanStat

	}

	q := elastic.NewBoolQuery()

	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}
	
    ti := time.Now()
    timestamp := ti.Unix()
    fmt.Println("当前本时区时间：", ti)
    fmt.Println("当前本时区时间时间戳：", timestamp)

    	matchMQQuery := elastic.NewQueryStringQuery("rabbitmq")
    	matchMQQuery.DefaultField("remoteEndpoint.serviceName")

    	matchNameQuery := elastic.NewQueryStringQuery("on-message")
    	matchNameQuery.DefaultField("name")
		qSub := elastic.NewBoolQuery()
		qSub = qSub.Should(matchMQQuery)
		qSub = qSub.Should(matchNameQuery)
		q = q.Must(qSub)

		//q = q.Must(elastic.NewTermQuery("_q", "http.method"))

    if queryInfo.Lookback > 0{
    	ti := time.Now()
    	timestamp := ti.Unix()
    	end_micro_timestamp := timestamp * 1000 * 1000
    	start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
    }

    if queryInfo.Start_time != "" && queryInfo.End_time != "" {
    	start_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.Start_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    end_time, err := time.ParseInLocation("2006-01-02 15:04:05",queryInfo.End_time , time.Local)
	    if err != nil {
	       	fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return spanStat
	    }

	    start_micro_timestamp:= start_time.Unix() * 1000 * 1000
    	end_micro_timestamp := end_time.Unix() * 1000 * 1000
    	q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

    }

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	s := string(data)
	fmt.Println(s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)    

	search = search.Sort("timestamp", false)

	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
    }                             //.Filter(andFilter)
    search = search.From(0).Size(int(maxLen))
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.TODO())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return spanStat
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	
	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0 
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Ret_length = int64(len(searchResult.Hits.Hits))

		spanStat.Counts = int64(len(searchResult.Hits.Hits))
		spanStat.All_Hits = searchResult.Hits.TotalHits
		spanStat.Annotation = anno
		spanStat.Type = "mq"
		spanStat.Name = "消息中间件"
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				// Deserialization failed
			}

			spanStat.Sum += t.Duration
			if spanStat.Min > t.Duration {
				spanStat.Min = t.Duration
			}
			if spanStat.Max < t.Duration {
				spanStat.Max = t.Duration
			}
			// Work with tweet
			// fmt.Print(t)
			//  fmt.Printf("%v\n", t)
			//fmt.Println(t.Data.Log_info.Message)
		}
		if spanStat.Counts > 0 {
			spanStat.Avg = spanStat.Sum / spanStat.Counts
		}
		logResult.Ret = append(logResult.Ret, spanStat)
		if !IsAll {
			c.JSON(200, spanStat)
		}
		return spanStat

	} else {
		// No hits
		if !IsAll {
			c.JSON(200, QueryNoResult)
		}
		return spanStat
	}

}



func ZipkinStatsAll(c *gin.Context, queryInfo Common.QueryZipkinSpan) {
		


		var logResult SQueryZipkinStatsResult
		var spanStat SZipkinStats
		logResult.Ret_length = 0
		queryInfo.Query_type = "hystrix"
		spanStat = ZipkinStatsHystrix(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "lb"
		spanStat = ZipkinStatsLoadBalanced(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "gateway"
		spanStat = ZipkinStatsJupiter(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "druid"
		spanStat = ZipkinStatsDruid(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "feign"
		spanStat = ZipkinStatsFeign(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "cache"
		spanStat = ZipkinStatsCache(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "mysql"
		spanStat = ZipkinStatsMysql(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "gravity"
		spanStat = ZipkinStatsGravity(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "http"
		spanStat = ZipkinStatsHttp(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		queryInfo.Query_type = "mq"
		spanStat = ZipkinStatsMQ(c, queryInfo)
		logResult.Ret = append(logResult.Ret, spanStat)
		logResult.Ret_length ++ 

		logResult.All_length = logResult.Ret_length
		logResult.Ret_code = 200

		c.JSON(200, logResult)




}

