package LogInfo

import (
	"encoding/json"
	_ "flag"
	"fmt"
	"query_server/Common"
	_ "reflect"
	"strconv"
	_ "strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

func ZipkinStatsTPS(c *gin.Context, queryInfo Common.QueryZipkinSpan) SQueryZipkinStatsTPSResult {

	var logResult SQueryZipkinStatsTPSResult
	var spanStat SZipkinStatsTPS

	var maxLen, defaultLookback int64
	maxLen = 50
	EsHostArr = strings.Split(*ArgEsHost, ",")
	fmt.Printf("ES host info %s,%q.\n", *ArgEsHost, EsHostArr)
	client, err := elastic.NewClient(elastic.SetURL(EsHostArr...),
		elastic.SetSniff(false))

	if err != nil {
		c.JSON(200, ConnElasticsearchErr)
		return logResult

	}

	q := elastic.NewBoolQuery()

	//agg := elastic.TermsAggregation()
	//agg.
	//	q = q.Ag
	if queryInfo.Query_type == "" || queryInfo.Query_type == "all" {
		//return
	}

	ti := time.Now()
	timestamp := ti.Unix()
	fmt.Println("当前本时区时间：", ti)
	fmt.Println("当前本时区时间时间戳：", timestamp)

	/*matchDruidQuery := elastic.NewQueryStringQuery("hystrix")
	matchDruidQuery.DefaultField("name")

	qSub := elastic.NewBoolQuery()
	qSub = qSub.Should(matchDruidQuery)
	q = q.Must(qSub)*/

	defaultLookback = 60 * 60
	if queryInfo.Lookback > 0 {
		defaultLookback = queryInfo.Lookback
		end_micro_timestamp := timestamp * 1000
		start_micro_timestamp := end_micro_timestamp - queryInfo.Lookback*1000
		q = q.Must(elastic.NewRangeQuery("timestamp_millis").Gt(start_micro_timestamp).Lt(end_micro_timestamp).Format("epoch_millis"))
	} else if queryInfo.Start_time != "" && queryInfo.End_time != "" {
		start_time, err := time.ParseInLocation("2006-01-02 15:04:05", queryInfo.Start_time, time.Local)
		if err != nil {
			fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return logResult
		}

		end_time, err := time.ParseInLocation("2006-01-02 15:04:05", queryInfo.End_time, time.Local)
		if err != nil {
			fmt.Printf("error: %q.\n", err)
			c.JSON(200, InvalidQuery)
			return logResult
		}

		start_micro_timestamp := start_time.Unix() * 1000
		end_micro_timestamp := end_time.Unix() * 1000
		//q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))

		q = q.Must(elastic.NewRangeQuery("timestamp_millis").Gt(start_micro_timestamp).Lt(end_micro_timestamp).Format("epoch_millis"))

	} else {
		end_micro_timestamp := timestamp * 1000
		start_micro_timestamp := end_micro_timestamp - defaultLookback*1000
		//q = q.Must(elastic.NewRangeQuery("timestamp").Gt(start_micro_timestamp).Lt(end_micro_timestamp))
		q = q.Must(elastic.NewRangeQuery("timestamp_millis").Gt(start_micro_timestamp).Lt(end_micro_timestamp).Format("epoch_millis"))

	}

	src, err := q.Source()
	if err != nil {
		fmt.Printf("error: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return logResult

	}
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Printf("error2: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return logResult
	}
	s := string(data)
	fmt.Println("query:", s)

	search := client.Search().Index("zipkin:span-*") //.Type("film")
	search = search.Query(q)
	//search.Size(0)
	//if()
	//service := service.Aggregation("trace_name", agg)
	//subAgg := elastic.NewTermsAggregation().Field("genre")
	//elastic.NewDateHistogramAggregation

	var interSec int64 = 600
	if queryInfo.Interval > 0 {
		interSec = queryInfo.Interval

	}
	inter := fmt.Sprintf("%d%s", interSec, "S")
	dateAgg := elastic.NewDateHistogramAggregation().Field("timestamp_millis").Interval(inter).Format("epoch_millis")
	//.SubAggregation("genres_by_year", subAgg)
	agg := elastic.NewTermsAggregation().Field("localEndpoint.serviceName")
	agg = agg.SubAggregation("int_time", dateAgg)

	src, err = agg.Source()
	if err != nil {
		fmt.Printf("error55: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return logResult

	}
	data, err = json.Marshal(src)
	if err != nil {
		fmt.Printf("error55: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return logResult
	}
	s = string(data)
	fmt.Println("aggs:", s)

	search = search.Aggregation("inn", agg)
	search = search.Size(0)

	search = search.Sort("timestamp", false)
	if queryInfo.Max_len > 0 {
		maxLen = queryInfo.Max_len
	} //.Filter(andFilter)
	//search = search.From(0).Size(int(maxLen))
	_ = maxLen
	//search = search.Sort("data.log_info.log_time", false)
	//search = search.From((pageIndex - 1) * lengthPerPage).Size(lengthPerPage)

	searchResult, err := search.Do(context.Background())

	if err != nil {
		fmt.Printf("error3: %q.\n", err)
		c.JSON(200, ErrElasticsearch)
		return logResult
	}
	fmt.Printf("Found a total of %d ,%d result, took %d milliseconds.\n", searchResult.TotalHits(), searchResult.Hits.TotalHits, searchResult.TookInMillis)

	//var minDur=0, maxDur=0, avgDur=0, sumDur=0, count=0
	logResult.Ret_code = 200
	logResult.Ret_length = 10
	logResult.Type = "stats"

	//第二种方式

	/*a := searchResult.Aggregations
	b, _ := json.Marshal(searchResult.Aggregations)
	//fmt.Println(string(b))
	json.Unmarshal(b, searchResult.Aggregations.)
	for _, v := range a.inn.Buckets {
		fmt.Println(v)
	}*/
	//spanStat.Min = 1000000
	fmt.Printf("%+v\n", searchResult.Hits.Hits)
	fmt.Printf("%+v\n", searchResult.Aggregations)
	//b, _ := json.Marshal(searchResult.Aggregations)
	//#fmt.Println(string(b))
	logResult.All_length = 0
	logResult.Ret_length = 0
	if aggr, found := searchResult.Aggregations.Terms("inn"); found {
		//spanStat.

		fmt.Printf("%+v\n", aggr)
		for _, bucket := range aggr.Buckets {
			// JSON doesn't have integer types: All numeric values are float64
			strValue, ok := bucket.Key.(string)
			if !ok {
				panic("expected a float64")
			}
			spanStat.Doc_Count = bucket.DocCount
			spanStat.Service_Name = strValue
			// Iterate over the sub-aggregation
			if subAgg, found := bucket.Terms("int_time"); found {
				var unitStats SZipkinStatsPerDuration
				for _, subBucket := range subAgg.Buckets {
					unitStats.Doc_Count = subBucket.DocCount
					unitStats.Start_Time = int64(subBucket.Key.(float64))
					if unitStats.Doc_Count > 0 {
						unitStats.TPS, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", unitStats.Doc_Count/interSec), 64)
					}
					spanStat.Query_Result_Duration = append(spanStat.Query_Result_Duration, unitStats)
				}
			}
			logResult.Ret = append(logResult.Ret, spanStat)
			logResult.All_length = logResult.All_length + 1
			logResult.Ret_length = logResult.Ret_length + 1
			logResult.Ret_code = 200
			logResult.Type = "tps"

		}
	}
	c.JSON(200, logResult)
	return logResult

}
