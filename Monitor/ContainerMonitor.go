package Monitor

import (
	"query_server/Common"
	"fmt"
	"log"
	"strconv"
	"time"
	"sort"
	"github.com/gin-gonic/gin"
)

func QueryContainerMonitorInfo(c *gin.Context, queryInfo Common.QueryMonitorJson) {

	var containerMonitorTag ContainerMonitorTag

	//monitorResult.Return_code = "200"
	//ret := Monitor.QueryDB("select * from /.*/ limit 10")
	var finalTagQuery string
	var finalMetricQuery string
	timeStr := ""
	const TimeFormat = "2006-01-02 15:04:05"
	const RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	const InfluxTimeFormat = "2006-01-02T15:04:05.999Z"

	var containerMonitor QueryContainerMonitor
	containerMonitor.Return_code = 200

	var containerMonitorKeys []string
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

	//st := startTime.Format(RFC3339Nano)
	//et := endTime.Format(RFC3339Nano)

	//finalQuery = "select * from /.*/"
	//finalQuery += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", queryInfo.Container_uuid)
    //finalQuery += fmt.Sprintf("\"environment_id\"='%s' AND ", queryInfo.Environment_id)
    //finalQuery += fmt.Sprintf("time>='%s' AND time<='%s' order by time limit 10", queryInfo.Start_time, queryInfo.End_time)

	finalTagQuery = "select * from /.*/"
	finalTagQuery += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", queryInfo.Container_uuid)
    finalTagQuery += fmt.Sprintf("\"environment_id\"='%s' AND ", queryInfo.Environment_id)
    finalTagQuery += fmt.Sprintf("time>='%s' AND time<='%s' limit 1", queryInfo.Start_time, queryInfo.End_time)

	


	fmt.Println(startTime, endTime)
	fmt.Println(finalTagQuery)
	_ = queryValidation

	ret := QueryDB(finalTagQuery)

	if len(ret[0].Series) > 0 {
		// monitorResult.
		containerMonitorTag.Timestamp = fmt.Sprintf("%s", ret[0].Series[0].Values[0][0])
		containerMonitorTag.Container_uuid = fmt.Sprintf("%s", ret[0].Series[0].Values[0][2])
		containerMonitorTag.Environment_id = fmt.Sprintf("%s", ret[0].Series[0].Values[0][3])
		containerMonitorTag.Container_name = fmt.Sprintf("%s", ret[0].Series[0].Values[0][1])
		containerMonitorTag.Namespace = fmt.Sprintf("%s", ret[0].Series[0].Values[0][4])
		containerMonitorTag.Type = fmt.Sprintf("%s", ret[0].Series[0].Values[0][5])
	}else{
		c.JSON(200, gin.H{
            "return_code":  400,
            "err_info":"query not found",
        })
        return 
	}


	finalMetricQuery = "select first(*) from /.*/"
	finalMetricQuery += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", queryInfo.Container_uuid)
    finalMetricQuery += fmt.Sprintf("\"environment_id\"='%s' AND ", queryInfo.Environment_id)
    finalMetricQuery += fmt.Sprintf("time>='%s' AND time<='%s' group by time(%ss)", queryInfo.Start_time, queryInfo.End_time, queryInfo.Time_step)

    fmt.Println(finalMetricQuery)




    ret = QueryDB(finalMetricQuery)

    //fmt.Printf("%#v.\n",ret);
	for index := 0; index < len(ret[0].Series); index++ {
		se := ret[0].Series[index]
		timeNameStatResult[se.Name] = make(map[string]int)

		for valIndex := 0; valIndex < len(se.Values); valIndex++ {
			timeStr = fmt.Sprintf("%s", se.Values[valIndex][0])
			valStr := fmt.Sprintf("%s", se.Values[valIndex][1])
			val, err := strconv.Atoi(valStr)
			_ = err
			//fmt.Printf("%d :%s,%s,%s\n", index, se.Name, se.Values[valIndex][28], se.Values[valIndex][0])
			//fmt.Println(reflect.TypeOf(se.Name))
			timeNameStatResult[se.Name][timeStr] = val
		}
	}
	timeStat := make(map[string]*QueryMonitorUnit)

	for k, v := range timeNameStatResult {
		_ = k
		t := v
		for k1, val := range t {
			if _, ok := timeStat[k1]; !ok {
				timeStat[k1] = new(QueryMonitorUnit)
				//intTime,error := strconv.Atoi(k1)
				/*fmt.Println(k1)
				intNanoTime, error := strconv.ParseInt(k1, 10, 64)
				if error != nil {
					log.Fatalln("Error: ", err)
				}*/

				// timeStat[k1].Timestamp = time.Unix(intTime/1000000000, 0).Format(RFC3339Nano,)
				timeStat[k1].Data.Timestamp = k1

				//timeStat[k1].Timestamp = fmt.Sprintf("%s", ret[0].Series[0].Values[0][0])
				timeStat[k1].Data.Container_uuid = containerMonitorTag.Container_uuid
				timeStat[k1].Data.Environment_id = containerMonitorTag.Environment_id
				timeStat[k1].Data.Container_name = containerMonitorTag.Container_name 
				timeStat[k1].Data.Namespace = containerMonitorTag.Namespace
				timeStat[k1].Type = containerMonitorTag.Type

				containerMonitorKeys = append(containerMonitorKeys, k1)
				//time.Unix(0, intNanoTime).Format(RFC3339Nano)
				//t := time.SecondsToLocalTime(1305861602)
			}
			info := timeStat[k1]
			switch k {
			case "cpu_usage_per_cpu":
				//StatsInfo.Container_cpu_usage_seconds_total =
				break
			case "fs_limit":

				break
			case "fs_usage":
				break
			case "load_average":
				break
			case "memory_usage":
				info.Data.Stats.Container_memory_usage_bytes = val
				break
			case "memory_working_set":
				break
			case "container_network_receive_packets_total":
				info.Data.Stats.Container_network_receive_bytes_total = val
				break
			case "container_network_receive_errors_total":
				info.Data.Stats.Container_network_receive_errors_total = val
				break
			case "container_network_transmit_bytes_total":
				info.Data.Stats.Container_network_transmit_bytes_total = val
				break
			case "container_network_transmit_errors_total":
				info.Data.Stats.Container_network_transmit_errors_total = val
				break
			case "container_tasks_state_nr_sleeping":
				info.Data.Stats.Container_tasks_state_nr_sleeping = val
				break
			case "container_tasks_state_nr_io_wait":
				info.Data.Stats.Container_tasks_state_nr_io_wait = val
				break
			case "container_network_transmit_packets_total":
				info.Data.Stats.Container_network_transmit_packets_total = val
				break
			case "container_memory_usage_bytes":
				info.Data.Stats.Container_memory_usage_bytes = val
				break
			case "container_memory_swap":
				info.Data.Stats.Container_memory_swap = val
				break
			case "container_memory_cache":
				info.Data.Stats.Container_memory_cache = val
				break
			case "container_cpu_usage_seconds_total":
				info.Data.Stats.Container_cpu_usage_seconds_total = val
				break
			case "container_memory_limit_bytes":
				info.Data.Stats.Container_memory_limit_bytes = val
				break
			case "container_diskio_service_bytes_read":
				info.Data.Stats.Container_diskio_service_bytes_read = val
				break
			case "container_cpu_system_seconds_total":
				info.Data.Stats.Container_cpu_system_seconds_total = val
				break
			case "container_tasks_state_nr_uninterruptible":
				info.Data.Stats.Container_tasks_state_nr_uninterruptible = val
				break
			case "container_tasks_state_nr_running":
				info.Data.Stats.Container_tasks_state_nr_running = val
				break
			case "container_network_transmit_packets_dropped_total":
				info.Data.Stats.Container_network_transmit_packets_dropped_total = val
				break
			case "container_network_receive_packets_dropped_total":
				info.Data.Stats.Container_network_receive_packets_dropped_total = val
				break
			case "container_network_receive_bytes_total":
				info.Data.Stats.Container_network_receive_bytes_total = val
				break
			case "container_memory_rss":
				info.Data.Stats.Container_memory_rss = val
				break
			case "container_cpu_user_seconds_total":
				info.Data.Stats.Container_cpu_user_seconds_total = val
				break
			case "container_tasks_state_nr_stopped":
				info.Data.Stats.Container_tasks_state_nr_stopped = val
				break
			case "container_diskio_service_bytes_write":
				info.Data.Stats.Container_diskio_service_bytes_write = val
				break
			case "container_diskio_service_bytes_total":
				info.Data.Stats.Container_diskio_service_bytes_total = val
				break
			case "container_diskio_service_bytes_sync":
				info.Data.Stats.Container_diskio_service_bytes_sync = val
				break
			case "container_diskio_service_bytes_async":
				info.Data.Stats.Container_diskio_service_bytes_async = val
				break
			default:
				fmt.Printf("Error metric name:%s.\n", k)
			}

		}
	}

	//monitorResult.Stats = make([]StatsInfo, len(timeStat))
	containerMonitor.Query_result = make([]QueryMonitorUnit, len(timeStat))
	index := 0
	/*for k, _ := range timeStat {
		//fmt.Printf("%#v.\n",timeStat[k]);
		//monitorResult.Stats[index] = *timeStat[k]
		containerMonitor.Query_result[index]=  *timeStat[k]
		
		index++
	}*/
	sort.Strings(containerMonitorKeys) 
	for _, k := range containerMonitorKeys {
       // fmt.Println("Key:", k, "Value:", m[k])
		containerMonitor.Query_result[index]=  *timeStat[k]
		//fmt.Println(k)
		index ++
    }

	//c.JSON(200, monitorResult)
	c.JSON(200, containerMonitor)

}
