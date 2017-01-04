package Monitor

import (
	"query_server/Common"
	"fmt"
	"log"
	"strconv"
	"time"
	"sort"
	"github.com/gin-gonic/gin"
	"strings"
)

func QueryContainerMonitorInfo(c *gin.Context, queryInfo Common.QueryMonitorJson) {


	const RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	const MyDB = "containerdb"
	var containerMonitorTag ContainerMonitorTag
	var finalTagQuery string
	var finalMetricQuery string
	
	var containerMonitor QueryContainerMonitor
	containerMonitor.Return_code = 200

	var containerMonitorKeys []string
	//var err error
	//MetricsName-->TimeStamp-->value
	timeNameStatResult := make(map[string]map[string]int)

	var queryValidation = true
	timeStr := ""
	_, err := time.Parse(RFC3339Nano, queryInfo.Start_time)
	if err != nil {
		queryValidation = false
		log.Fatalln("Error: ", err)
	}

	_, err = time.Parse(RFC3339Nano, queryInfo.End_time)
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

	
	fmt.Println(finalTagQuery)
	_ = queryValidation

	ret := QueryDB(finalTagQuery, MyDB)
	//fmt.Printf("%#v.\n",ret);
	if len(ret[0].Series) > 0 {
		// monitorResult.
		timeInd:= indexOf(ret[0].Series[0].Columns, "time")
		uuidInd := indexOf(ret[0].Series[0].Columns, "container_uuid")
		envIdInd := indexOf(ret[0].Series[0].Columns, "environment_id")
		nameInd := indexOf(ret[0].Series[0].Columns, "container_name")
		namespaceInd := indexOf(ret[0].Series[0].Columns, "namespace")
		typeInd := indexOf(ret[0].Series[0].Columns, "type")

		containerMonitorTag.Timestamp = fmt.Sprintf("%s", ret[0].Series[0].Values[0][timeInd])
		containerMonitorTag.Container_uuid = fmt.Sprintf("%s", ret[0].Series[0].Values[0][uuidInd])
		containerMonitorTag.Environment_id = fmt.Sprintf("%s", ret[0].Series[0].Values[0][envIdInd])
		containerMonitorTag.Container_name = fmt.Sprintf("%s", ret[0].Series[0].Values[0][nameInd])
		containerMonitorTag.Namespace = fmt.Sprintf("%s", ret[0].Series[0].Values[0][namespaceInd])
		containerMonitorTag.Type = fmt.Sprintf("%s", ret[0].Series[0].Values[0][typeInd])
	}else{
		c.JSON(200, gin.H{
            "return_code":  400,
            "err_info":"query not found",
        })
        return 
	}

	//sample(*, 1) 
	finalMetricQuery = "select first(*) from /.*/"
	finalMetricQuery += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", queryInfo.Container_uuid)
    finalMetricQuery += fmt.Sprintf("\"environment_id\"='%s' AND ", queryInfo.Environment_id)
    finalMetricQuery += fmt.Sprintf("time>='%s' AND time<='%s' group by time(%ss)", queryInfo.Start_time, queryInfo.End_time, queryInfo.Time_step)




    finalMetricQuery += "; select time,value from /.*/"
	finalMetricQuery += fmt.Sprintf(" WHERE \"container_uuid\"='%s' AND ", queryInfo.Container_uuid)
    finalMetricQuery += fmt.Sprintf("\"environment_id\"='%s' AND ", queryInfo.Environment_id)
    finalMetricQuery += fmt.Sprintf("time>='%s' AND time<='%s' order by time desc limit 1", queryInfo.Start_time, queryInfo.End_time)





    fmt.Println(finalMetricQuery)

    ret = QueryDB(finalMetricQuery, MyDB)

    //fmt.Printf("%#v.\n",ret[1]);
    //fmt.Printf("%#v.\n",ret);

    //timeInd := indexOf(ret[0].Series[0].Columns, "time")
	//valInd := indexOf(ret[0].Series[0].Columns, "")


	for index := 0; index < len(ret[0].Series); index++ {
		se := ret[0].Series[index]
		timeNameStatResult[se.Name] = make(map[string]int)

		for valIndex := 0; valIndex < len(se.Values); valIndex++ {

			if(se.Values[valIndex][2] == nil){
				continue
			}
			timeStr = fmt.Sprintf("%s", se.Values[valIndex][0])
			valStr := fmt.Sprintf("%s", se.Values[valIndex][2])
			val, err := strconv.Atoi(valStr)
			_ = err
			//fmt.Printf("%d :%s,%s,%s\n", index, se.Name, se.Values[valIndex][28], se.Values[valIndex][0])
			//fmt.Println(reflect.TypeOf(se.Name))
			timeNameStatResult[se.Name][timeStr] = val
		}
	}



	for index := 0; index < len(ret[1].Series); index++ {
		se := ret[1].Series[index]

		if timeNameStatResult[se.Name] == nil{
			timeNameStatResult[se.Name] = make(map[string]int)
		}

		for valIndex := 0; valIndex < len(se.Values); valIndex++ {

			if(se.Values[valIndex][1] == nil){
				continue
			}
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
				timeStat[k1].Data.Stats.Timestamp = k1

				containerMonitorKeys = append(containerMonitorKeys, k1)
				//time.Unix(0, intNanoTime).Format(RFC3339Nano)
				//t := time.SecondsToLocalTime(1305861602)
			}

			info := timeStat[k1]


			var fs ContainerFileSystem
			fsIndex := -1
			
			//fmt.Println(k)
			if strings.Contains(k, "container_filesystem_capacity_"){
				fs.Container_filesystem_name  = strings.TrimPrefix(k, "container_filesystem_capacity_")
            	fs.Container_filesystem_type  = "default"
              	fs.Container_filesystem_capacity = val

				for i, _ := range info.Data.Stats.Container_filesystem {
					if info.Data.Stats.Container_filesystem[i].Container_filesystem_name == fs.Container_filesystem_name {
						fsIndex = i
	              	 	info.Data.Stats.Container_filesystem[i].Container_filesystem_type = fs.Container_filesystem_type 
	              	 	info.Data.Stats.Container_filesystem[i].Container_filesystem_capacity = fs.Container_filesystem_capacity
						break
					}
				}

				if fsIndex == -1 {
					info.Data.Stats.Container_filesystem = append(info.Data.Stats.Container_filesystem, fs)
				}

				continue
				 
			}else if strings.Contains(k, "container_filesystem_usage_"){

				 fs.Container_filesystem_name  = strings.TrimPrefix(k, "container_filesystem_usage_")
              	 fs.Container_filesystem_type  = "default"
				 fs.Container_filesystem_usage = val

				 for i, _ := range info.Data.Stats.Container_filesystem {
					if info.Data.Stats.Container_filesystem[i].Container_filesystem_name == fs.Container_filesystem_name {
						fsIndex = i
	              	 	info.Data.Stats.Container_filesystem[i].Container_filesystem_type = fs.Container_filesystem_type 
	              	 	info.Data.Stats.Container_filesystem[i].Container_filesystem_usage = fs.Container_filesystem_usage
						break
					}
				}

				if fsIndex == -1{
					info.Data.Stats.Container_filesystem = append(info.Data.Stats.Container_filesystem, fs)
				}

				continue

			}

			

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
