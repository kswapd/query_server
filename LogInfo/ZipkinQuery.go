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
)
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
	logResult.Return_code = 200
	logResult.Current_query_result_length = 10
	logResult.All_query_result_length = 100
	spanStat.Min = 1000000
	if len(searchResult.Hits.Hits) > 0 {
		logResult.All_query_result_length = searchResult.Hits.TotalHits
		//fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		logResult.Current_query_result_length = int64(len(searchResult.Hits.Hits))

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
		logResult.Query_result = append(logResult.Query_result, spanStat)
		c.JSON(200, logResult)

	} else {
		// No hits
		c.JSON(200, QueryNoResult)
	}

}

