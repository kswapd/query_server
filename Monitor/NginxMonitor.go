package Monitor

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/influxdata/influxdb1-client/v2"
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
		index := indexOf(v.Columns, Value_key_of_last) //index指定value存储位置

		for _, v1 := range v.Values {
			if v1[index] == nil {
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
	var container_uuid string
	indexOfUuid := indexOf(res[0].Series[0].Columns, "container_uuid")
	i := 0
	for res[0].Series[0].Values[i][indexOfUuid] == nil {
		i++
	}
	container_uuid = res[0].Series[0].Values[i][indexOfUuid].(string)
	//	if res[0].Series[0].Values[0][indexOfUuid] != nil {
	//		container_uuid = res[0].Series[0].Values[0][indexOfUuid].(string)
	//	}

	//environment_id
	var environment_id string
	indexOfId := indexOf(res[0].Series[0].Columns, "environment_id")
	environment_id = res[0].Series[0].Values[i][indexOfId].(string)
	//	if res[0].Series[0].Values[0][indexOfId] != nil {
	//		environment_id = res[0].Series[0].Values[0][indexOfId].(string)
	//	}

	//container_name
	var container_name string
	indexOfName := indexOf(res[0].Series[0].Columns, "container_name")
	container_name = res[0].Series[0].Values[i][indexOfName].(string)
	//	if res[0].Series[0].Values[0][indexOfName] != nil {
	//		container_name = res[0].Series[0].Values[0][indexOfName].(string)
	//	}

	//namespace
	var namespace string
	indexOfNamespace := indexOf(res[0].Series[0].Columns, "namespace")
	namespace = res[0].Series[0].Values[i][indexOfNamespace].(string)
	//	if res[0].Series[0].Values[0][indexOfNamespace] != nil {
	//		namespace = res[0].Series[0].Values[0][indexOfNamespace].(string)
	//	}

	//type
	var appType string
	indexOfType := indexOf(res[0].Series[0].Columns, "type")
	appType = res[0].Series[0].Values[i][indexOfType].(string)
	//	if res[0].Series[0].Values[0][indexOfType] != nil {
	//		appType = res[0].Series[0].Values[0][indexOfType].(string)
	//	}
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
