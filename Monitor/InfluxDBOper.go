package Monitor

import (
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB = "containerdb"
	//username = "bubba"
	//password = "bumblebeetuna"
)

// queryDB convenience function to query the database
func QueryDB(cmd string) (ret []client.Result) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://54.223.73.138:8086",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	q := client.NewQuery(cmd, MyDB, "")
	response, err := c.Query(q)
	if err == nil && response.Error() == nil {
		//fmt.Println(response.Results)
	}
	//return res, nil
	return response.Results
}
