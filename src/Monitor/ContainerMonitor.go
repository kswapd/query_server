package Monitor

import (
	"Common"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func QueryContainerMonitorInfo(c *gin.Context, queryInfo Common.QueryMonitorJson) {

	var monitorResult QueryMonitorResultJson
	//ret := Monitor.QueryDB("select * from /.*/ limit 10")
	var finalQuery string
	timeStr := ""
	const TimeFormat = "2006-01-02 15:04:05"
	const RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	const InfluxTimeFormat = "2006-01-02T15:04:05.999Z"
	//var err error
	//MetricsName-->TimeStamp-->value
	timeNameStatResult := make(map[string]map[string]int)

	var queryValidation = true

	startTime, err := time.Parse(RFC3339Nano, queryInfo.Start_time)
	if err != nil {
		queryValidation = false
		log.Fatalln("Error: ", err)
	}

	endTime, err := time.Parse(RFC3339Nano, queryInfo.End_time)
	if err != nil {
		queryValidation = false
		log.Fatalln("Error: ", err)
	}

	st := startTime.Format(RFC3339Nano)
	et := endTime.Format(RFC3339Nano)

	finalQuery = fmt.Sprintf("select * from /.*/ where time > '%s' and time < '%s' order by time desc limit 10", st, et)
	fmt.Println(startTime, endTime)
	fmt.Println(finalQuery)
	_ = queryValidation

	ret := QueryDB(finalQuery)

	if len(ret[0].Series) > 0 {
		// monitorResult.
		monitorResult.Timestamp = fmt.Sprintf("%s", ret[0].Series[0].Values[0][0])
		monitorResult.Container_uuid = fmt.Sprintf("%s", ret[0].Series[0].Values[0][14])
		monitorResult.Environment_id = fmt.Sprintf("%s", ret[0].Series[0].Values[0][24])
		monitorResult.Container_name = fmt.Sprintf("%s", ret[0].Series[0].Values[0][1])
		monitorResult.Namespace = fmt.Sprintf("%s", ret[0].Series[0].Values[0][24])
	}

	for index := 0; index < len(ret[0].Series); index++ {
		se := ret[0].Series[index]
		timeNameStatResult[se.Name] = make(map[string]int)

		for valIndex := 0; valIndex < len(se.Values); valIndex++ {
			timeStr = fmt.Sprintf("%s", se.Values[valIndex][0])
			valStr := fmt.Sprintf("%s", se.Values[valIndex][28])
			val, err := strconv.Atoi(valStr)
			_ = err
			//fmt.Printf("%d :%s,%s,%s\n", index, se.Name, se.Values[valIndex][28], se.Values[valIndex][0])
			//fmt.Println(reflect.TypeOf(se.Name))
			timeNameStatResult[se.Name][timeStr] = val
		}
	}
	timeStat := make(map[string]*StatsInfo)

	for k, v := range timeNameStatResult {
		_ = k
		t := v
		for k1, val := range t {
			if _, ok := timeStat[k1]; !ok {
				timeStat[k1] = new(StatsInfo)
				//intTime,error := strconv.Atoi(k1)
				intNanoTime, error := strconv.ParseInt(k1, 10, 64)
				if error != nil {
					log.Fatalln("Error: ", err)
				}

				// timeStat[k1].Timestamp = time.Unix(intTime/1000000000, 0).Format(RFC3339Nano,)
				timeStat[k1].Timestamp = time.Unix(0, intNanoTime).Format(RFC3339Nano)
				//t := time.SecondsToLocalTime(1305861602)
			}
			info := timeStat[k1]
			switch k {
			case "cpu_usage_per_cpu":
				//StatsInfo.Container_cpu_usage_seconds_total =
				break
			case "cpu_usage_system":
				info.Container_cpu_system_seconds_total = val
				break
			case "cpu_usage_total":
				info.Container_cpu_usage_seconds_total = val
				break
			case "cpu_usage_user":
				info.Container_cpu_user_seconds_total = val
				break
			case "fs_limit":

				break
			case "fs_usage":
				break
			case "load_average":
				break
			case "memory_usage":
				info.Container_memory_usage_bytes = val
				break
			case "memory_working_set":
				break
			case "rx_bytes":
				info.Container_network_receive_bytes_total = val
				break
			case "rx_errors":
				info.Container_network_receive_errors_total = val
				break
			case "tx_bytes":
				info.Container_network_transmit_bytes_total = val
				break
			case "tx_errors":
				info.Container_network_transmit_errors_total = val
				break
			default:
				fmt.Println("Error metric name.")
			}

		}
	}

	monitorResult.Stats = make([]StatsInfo, len(timeStat))
	index := 0
	for k, _ := range timeStat {
		//fmt.Printf("%#v.\n",timeStat[k]);
		monitorResult.Stats[index] = *timeStat[k]
		index++
	}

	_ = ret
	_ = monitorResult

	c.JSON(200, monitorResult)

}
