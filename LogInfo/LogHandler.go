package LogInfo
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "query_server/Common"
   // "log"
    "encoding/json"
    elastic "gopkg.in/olivere/elastic.v5"
    "golang.org/x/net/context"
)

const(
  ESUrl string = "http://223.202.32.60:8056"
)

var (

    QueryNoResult = gin.H{
            "return_code":  400,
            "err_info":"query not found",
      }
      ErrElasticsearch = gin.H{
            "return_code":  401,
            "err_info":"elastic search connection error",
      }
      InvalidQuery = gin.H{
            "return_code":  402,
            "err_info":"invalid query",
      }

     /* m3 := map[string]string{
        "a": "aa",
        "b": "bb",
      }*/

)
func QueryContainerLog(c *gin.Context, queryInfo Common.QueryLogJson) {

   client, err := elastic.NewClient(elastic.SetURL(ESUrl))
      if err != nil {
          c.JSON(200, ErrElasticsearch)
          return

      }  



//.BodyString(mapping)
/*_, err = client.CreateIndex("twitter").Do(context.TODO())
if err != nil {
    // Handle error
    panic(err)
}*/

// Add a document to the index
/*tweet := Tweet{User: "olivere", Message: "Take Five"}
_, err = client.Index().
    Index("twitter").
    Type("tweet").
    Id("1").
    BodyJson(tweet).
    Refresh(true).
    Do()
if err != nil {
    // Handle error
    panic(err)
}*/

      //termQuery := elastic.NewTermQuery("type", "log_file_container")
      //_=termQuery



      q := elastic.NewBoolQuery()
     // q = q.Must(elastic.NewTermQuery("type", "container"))
      q = q.Should(elastic.NewTermQuery("type", "log_file_container"))
      q = q.Should(elastic.NewTermQuery("type", "log_container"))


       dt := elastic.NewRangeQuery("data.log_info.log_time").Gt(queryInfo.Start_time).Lt(queryInfo.End_time)
      q = q.Must(dt)

     /* src, err := q.Source()
      if err != nil {
        panic(err)
      }
      data, err := json.Marshal(src)
      if err != nil {
        panic(err)
      }
      s := string(data)
      fmt.Println(s)*/


      //q = q.Should(elastic.NewTermQuery("type", "log_file_container"))
      //q = q.Or(elastic.NewTermQuery("type", "log_container"))
      //q = q.Should(elastic.NewTermQuery("type", "log_container"))

      //termFilter_Timestamps := elastic.NewTermsFilter("utc_unix_timestamp", utc_unix_timestamps)
      //termFilter_ExperienceIds := elastic.NewTermsFilter("experience_id", experience_ids)
     // andFilter := elastic.NewAndFilter()
    //  andFilter.Add(termFilter_Timestamps)
     // andFilter.Add(termFilter_ExperienceIds)

    //q = q.Filter(elastic.NewTermQuery("account", 1))
      //timeRangeFilter := elastic.NewRangeFilter("@timestamp").Gte(1442100154219).Lte(1442704954219)

            count, err := client.Count("fluentd_from_container_to_es.log-*").Do(context.TODO())
            if err != nil {
              fmt.Printf("error:%#v.\n",err)
            }
           
              fmt.Printf("expected Count = %d; got %d", 3, count)
        




      search := client.Search().Index("fluentd_from_container_to_es.log-*")//.Type("film")
      

      //search = search.Query(elastic.NewMatchAllQuery())
      search = search.Query(q)//.Filter(andFilter)







      search = search.From(10).Size(2)




      searchResult, err2 := search.Do(context.TODO())

 /*     termFilter := elastic.NewBoolQuery()
filter := elastic.NewBoolQuery().Must(elastic.NewTermQuery("active", 1))
filter = filter.Should(elastic.NewTermQuery("types.raw", "App"))
filter = filter.Should(elastic.NewTermQuery("types.raw", "Ho"))
filter = filter.Should(elastic.NewTermQuery("types.raw", "Pe"))
filter = filter.Should(elastic.NewTermQuery("types.raw", "St"))
termFilter = termFilter.Filter(filter)*/

      /*searchResult, err2 := client.Search().
          Index("fluentd_from_container_to_es.log-2016.12.01").   // search in index "twitter"          Query(termQuery).   // specify the query          Sort("type", true).
          Query(termQuery).
          From(0).Size(10).
          Pretty(true).       
          Do(context.TODO())*/



      if err2 != nil {
          // Handle error
          panic(err2)
      }
      fmt.Printf("Found a total of %d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.TookInMillis)

      var logResult SQueryContainerLogResult
      var t SContainerLogger
      logResult.Return_code = 200
      logResult.Current_query_result_length = 10
      logResult.All_query_result_length = 100

      if searchResult.Hits.TotalHits > 0 {
        logResult.Current_query_result_length = searchResult.Hits.TotalHits
        //fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

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
	    //c.BindJSON(&queryInfo)
	    //c.JSON(200, gin.H{"type": queryInfo.Query_type})






  fmt.Printf("%#v.\n", queryInfo)


  switch queryInfo.Query_type {
    case "container":
      QueryContainerLog(c, queryInfo)
    case "app":
      QueryAppLog(c, queryInfo)
    default:
      c.JSON(200, InvalidQuery)
      return
  }



     
	    
    }

