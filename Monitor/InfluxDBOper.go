package Monitor

import (
	"flag"
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

var Cli client.Client
var ArgDbHost = flag.String("influxdb_driver_host", "54.223.73.138:8086", "database host:port")

func CreateInfluxDBClient() {
	var err error
	Cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://" + *ArgDbHost,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

// queryDB convenience function to query the database
func QueryDB(cmd string, db string) (ret []client.Result) {

	q := client.NewQuery(cmd, db, "")
	response, err := Cli.Query(q)
	if err == nil && response.Error() == nil {
		//fmt.Println(response.Results)
	}
	//return res, nil
	return response.Results
}
