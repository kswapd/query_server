package LogInfo
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "query_server/Common"
    "log"
    "encoding/json"
    elastic "gopkg.in/olivere/elastic.v5"
    "golang.org/x/net/context"
)
func QueryLogInfo(c *gin.Context) {
        var queryInfo Common.QueryLogJson
	    c.BindJSON(&queryInfo)
	    //c.JSON(200, gin.H{"type": queryInfo.Query_type})

      client, err := elastic.NewClient(elastic.SetURL("http://223.202.32.60:8056"))
      if err != nil {
          log.Fatalln("Error: ", err)
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

     // termQuery := elastic.NewTermQuery("type", "*")
      searchResult, err2 := client.Search().
          Index("fluentd_from_container_to_es.log-2016.12.05").   // search in index "twitter"          Query(termQuery).   // specify the query          Sort("type", true).
          From(0).Size(10).
          Pretty(true).       
          Do(context.TODO())

      if err2 != nil {
          // Handle error
          panic(err2)
      }
      fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

     // fmt.Println(*searchResult)
      fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

      var t LogContainerJson
      if searchResult.Hits.TotalHits > 0 {
        //fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            
            err := json.Unmarshal(*hit.Source, &t)

            if err != nil {
                // Deserialization failed
            }

            // Work with tweet
           // fmt.Print(t)
            fmt.Printf("%+v\n", t)
            fmt.Println(t.Data.Log_info.Message)
        }
      } else {
          // No hits
          fmt.Print("Found no tweets\n")
      }
	    c.JSON(200, t)
    }

