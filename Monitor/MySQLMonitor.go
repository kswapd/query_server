package Monitor

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/influxdata/influxdb/client/v2"
)

//原理参照nginx
func parseMySQLResult(res []client.Result) AppMySQLJson {

	var appMySQLJson AppMySQLJson

	if len(res) == 0 {
		return appMySQLJson
	}
	mysqlResult := make(map[string]map[string]float64)

	//遍历res，取出结果
	for _, v := range res[0].Series {
		mysqlResult[v.Name] = make(map[string]float64)
		index := indexOf(v.Columns, Value_key_of_last)

		for _, v1 := range v.Values {

			if v1[index] == nil {
				continue
			}

			f64, _ := strconv.ParseFloat(string(v1[index].(json.Number)), 64)
			mysqlResult[v.Name][v1[0].(string)] = f64
		}
	}

	timeStat := make(map[string]AppMySQLStatsJson)

	for k, v := range mysqlResult {
		for k1, val := range v {
			info := timeStat[k1]
			info.Timestamp = k1
			switch k {

			case "connections_total":
				{
					info.Connections_total = val
				}
			case "command_query_total":
				{
					info.Command_query_total = val
				}
			case "command_insert_total":
				{
					info.Command_insert_total = val
				}
			case "command_update_total":
				{
					info.Command_update_total = val
				}
			case "command_delete_total":
				{
					info.Command_delete_total = val
				}
			case "commands_total":
				{
					info.Commands_total = val
				}
			case "handlers_total":
				{
					info.Handlers_total = val
				}
			case "connection_errors_total":
				{
					info.Connection_errors_total = val
				}
			case "buffer_pool_pages":
				{
					info.Buffer_pool_pages = val
				}
			case "thread_connected":
				{
					info.Thread_connected = val
				}
			case "max_connections":
				{
					info.Max_connections = val
				}
			case "query_reponse_time_seconds":
				{
					info.Query_response_time_seconds = val
				}
			case "read_query_response_time":
				{
					info.Read_query_response_time_seconds = val
				}
			case "write_query_response_time_seconds":
				{
					info.Write_query_response_time_seconds = val
				}
			case "queries_inside_innodb":
				{
					info.Queries_inside_innodb = val
				}
			case "queries_in_queue":
				{
					info.Queries_in_queue = val
				}
			case "read_views_open_inside_innodb":
				{
					info.Read_views_open_inside_innodb = val
				}
			case "table_statistics_rows_read_total":
				{
					info.Table_statistics_rows_read_total = val
				}
			case "table_statistics_rows_changed_total":
				{
					info.Table_statistics_rows_changed_total = val
				}
			case "table_statistics_rows_changed_x_indexes_total":
				{
					info.Table_statistics_rows_changed_x_indexes_total = val
				}
			case "sql_lock_waits_total":
				{
					info.Sql_lock_waits_total = val
				}
			case "external_lock_waits_total":
				{
					info.External_lock_waits_total = val
				}
			case "table_io_waits_total":
				{
					info.Table_io_waits_total = val
				}
			case "table_io_waits_seconds_total":
				{
					info.Table_io_waits_seconds_total = val
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
	var amqr []AppMySQLQueryResult
	for _, v := range timeStat {
		var qrd AppMySQLQueryResultData
		qrd.Stats = v
		qrd.Timestamp = v.Timestamp
		qrd.Container_name = container_name
		qrd.Container_uuid = container_uuid
		qrd.Environment_id = environment_id
		qrd.Nmespace = namespace

		var mqr AppMySQLQueryResult
		mqr.Data = qrd
		mqr.Type = appType

		amqr = append(amqr, mqr)
	}

	appMySQLJson.Query_result = amqr
	appMySQLJson.Return_code = 200

	return appMySQLJson
}
