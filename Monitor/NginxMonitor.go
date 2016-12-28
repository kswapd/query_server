package Monitor

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/influxdata/influxdb/client/v2"
)

func parseNginxResult(res []client.Result) AppNginxJson {
	var appNginxJson AppNginxJson
	nginxResult := make(map[string]map[string]float64)

	//遍历res，取出结果
	for _, v := range res[0].Series {
		nginxResult[v.Name] = make(map[string]float64) //map[time]value
		index := indexOf(v.Columns, "value")           //哪个位置存储value

		for _, v1 := range v.Values {
			f64, _ := strconv.ParseFloat(string(v1[index].(json.Number)), 64)
			nginxResult[v.Name][v1[0].(string)] = f64
		}
	}

	timeStat := make(map[string]AppNginxStatsJson)

	for k, v := range nginxResult {
		for k1, val := range v {
			info := timeStat[k1]
			info.Timestamp = k1
			switch k {

			case "active_connections":
				{
					info.Active_connections = val
				}
			case "accepts":
				{
					info.Accepts = val
				}
			case "handled":
				{
					info.Handled = val
				}
			case "requests":
				{
					info.Requests = val
				}
			case "reading":
				{
					info.Reading = val
				}
			case "writing":
				{
					info.Writing = val
				}
			case "waiting":
				{
					info.Waiting = val
				}
			default:
				{
					fmt.Println("err measurements")
				}
			}
			timeStat[k1] = info
		}
	}

	var tmp []AppNginxStatsJson

	for _, v := range timeStat {
		tmp = append(tmp, v)
	}
	appNginxJson.Data.Stats = tmp

	//container_uuid
	indexOfUuid := indexOf(res[0].Series[0].Columns, "container_uuid")
	//	fmt.Println(indexOfUuid)
	appNginxJson.Data.Container_uuid = res[0].Series[0].Values[0][indexOfUuid].(string)
	//	fmt.Println(appRedisJson.Data.Container_uuid)

	//environment_id
	indexOfId := indexOf(res[0].Series[0].Columns, "environment_id")
	appNginxJson.Data.Environment_id = res[0].Series[0].Values[0][indexOfId].(string)

	//container_name
	indexOfName := indexOf(res[0].Series[0].Columns, "container_name")
	//	fmt.Println(indexOfName)
	appNginxJson.Data.Container_name = res[0].Series[0].Values[0][indexOfName].(string)

	//namespace
	indexOfNamespace := indexOf(res[0].Series[0].Columns, "namespace")
	appNginxJson.Data.Namespace = res[0].Series[0].Values[0][indexOfNamespace].(string)

	//type
	//	indexOfType := indexOf(res[0].Series[0].Columns, "type")
	//	appRedisJson.Data.Container_uuid = res[0].Series[0].Values[0][indexOfType].(string)

	return appNginxJson
}
