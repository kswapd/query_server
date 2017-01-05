package Monitor

import (
	"flag"
	"log"
	"github.com/influxdata/influxdb/client/v2"
)

var ArgDbHost = flag.String("influxdb_driver_host", "54.223.73.138:8086", "database host:port")

// queryDB convenience function to query the database
func QueryDB(cmd string, db string) (ret []client.Result) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://" + *ArgDbHost,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	q := client.NewQuery(cmd, db, "")
	response, err := c.Query(q)
	if err == nil && response.Error() == nil {
		//fmt.Println(response.Results)
	}
	//return res, nil
	return response.Results
}
