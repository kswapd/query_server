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

	/*if queryInfo.Start_time == ""  queryInfo.End_time == "" {
		c.JSON(200, InvalidQuery)
		return
	}*/


	q := elastic.NewBoolQuery()

	
    q = q.Must(elastic.NewTermQuery("_q", "LoadBalanced"))




	
	//q = q.Must(elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))


	// q = q.Must(elastic.NewMatchQuery("data.environment_id", "Network Agent"))

	//dt := elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time)
	// q = q.Must(elastic.NewRangeQuery("data.log_infoq = q.Must(elastic.NewTermQuery("data.container_uuid", queryInfo.Container_uuid)).log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time))

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

