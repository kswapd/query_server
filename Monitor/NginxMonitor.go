package Monitor

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/influxdata/influxdb/client/v2"
)

func parseNginxResult(res []client.Result) AppNginxJson {
	var appNginxJson AppNginxJson

	if len(res) == 0 {
		return appNginxJson
	}

	nginxResult := make(map[string]map[string]float64) //map["measurement"]map["time"]float64

	//遍历res，取出结果
	for _, v := range res[0].Series {
		nginxResult[v.Name] = make(map[string]float64) //map[time]value
		index := indexOf(v.Columns, "first_value")     //index指定value存储位置

		for _, v1 := range v.Values {
			if v1 == nil {
				continue
			}
			f64, _ := strconv.ParseFloat(string(v1[index].(json.Number)), 64)
			nginxResult[v.Name][v1[0].(string)] = f64
		}
	}

	timeStat := make(map[string]AppNginxStatsJson) //map["time"]AppNginxStatsJson

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

	//container_uuid
	indexOfUuid := indexOf(res[0].Series[0].Columns, "container_uuid")
	container_uuid := res[0].Series[0].Values[0][indexOfUuid].(string)

	//environment_id
	indexOfId := indexOf(res[0].Series[0].Columns, "environment_id")
	environment_id := res[0].Series[0].Values[0][indexOfId].(string)

	//container_name
	indexOfName := indexOf(res[0].Series[0].Columns, "container_name")
	container_name := res[0].Series[0].Values[0][indexOfName].(string)

	//namespace
	indexOfNamespace := indexOf(res[0].Series[0].Columns, "namespace")
	namespace := res[0].Series[0].Values[0][indexOfNamespace].(string)

	//type
	indexOfType := indexOf(res[0].Series[0].Columns, "type")
	appType := res[0].Series[0].Values[0][indexOfType].(string)

	//向目标结构字段添加值
	var anqr []AppNginxQueryResult
	for _, v := range timeStat {
		var qrd AppNginxQueryResultData
		qrd.Stats = v
		qrd.Timestamp = v.Timestamp
		qrd.Container_name = container_name
		qrd.Container_uuid = container_uuid
		qrd.Environment_id = environment_id
		qrd.Namespace = namespace

		var qr AppNginxQueryResult
		qr.Data = qrd
		qr.Type = appType

		anqr = append(anqr, qr)
	}

	appNginxJson.Query_result = anqr
	appNginxJson.Return_code = 200

	return appNginxJson
}
