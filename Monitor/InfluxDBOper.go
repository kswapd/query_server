package Monitor

import (
	"log"

	"github.com/influxdata/influxdb/client/v2"
)


// queryDB convenience function to query the database
func QueryDB(cmd string, db string) (ret []client.Result) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://54.223.149.26:8086",
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
