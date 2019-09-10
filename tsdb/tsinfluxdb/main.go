package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {
	client := newClient()
	//randMetric(client)
	//meanCpuUsage(client, "private")
	queryValues(client)
	_ = client.Close()
}

func newClient() *influxdb.Client {
	client, err := influxdb.New(nil, influxdb.WithToken("HRnUjPMp-33wAOenMTDWwwpelfjweQt18EHIj2f54MG_XFoCc1KMJbo5ME0TxitBTdyWJOfI3Lv3okD5GaYBmg=="))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// 随机写入
func randMetric(client *influxdb.Client) {
	clusters := []string{"public", "private"}
	eventTime := time.Now().Add(time.Second * -20)

	metrics := make([]influxdb.Metric, 0, 20*100)
	// create metrics
	for t := 0; t < 20; t++ {
		for i := 0; i < 100; i++ {
			clusterIndex := rand.Intn(len(clusters))
			tags := map[string]string{
				"cluster": clusters[clusterIndex],
				"host":    fmt.Sprintf("192.168.%d.%d", clusterIndex, rand.Intn(100)),
			}

			fields := map[string]interface{}{
				"cpu_usage":  rand.Float64() * 100.0,
				"disk_usage": rand.Float64() * 100.0,
			}
			row := influxdb.NewRowMetric(fields, "node_status", tags, eventTime.Add(time.Second*10))
			metrics = append(metrics, row)
		}
	}
	// write
	if err := client.Write(context.Background(), "my-bucket", "my-org", metrics...); err != nil {
		log.Fatal(err)
	}

	log.Printf("write ok\n")
}

func queryValues(client *influxdb.Client) {
	query := `
		from(bucket: "my-bucket")
		  |> range(start: -12h)
		  |> keep(columns: ["host"])
		  |> distinct(column: "host")
		`
	reader, err := client.QueryCSV(context.Background(), query, "my-org")
	if err != nil {
		log.Fatalf("query err,%+v", err)
	}

	result, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", result)
}

// 使用flux语言读取
func meanCpuUsage(client *influxdb.Client, cluster string) {
	// 喜欢需要的时间戳
	//startTime := "2019-09-10T00:23:37.076439Z"
	startTime := "-3h"
	format := `
	from(bucket: "my-bucket")
	|> range(start: %s)
	|> filter(fn: (r) => 
		r._measurement == "node_status" and
		r._field == "cpu_usage" and
		r.cluster == "%s"
	)
	|> window(every:10m)
	|> mean()
	`

	query := fmt.Sprintf(format, startTime, cluster)
	reader, err := client.QueryCSV(context.Background(), query, "my-org")
	if err != nil {
		log.Fatalf("query err,%+v", err)
	}

	result, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", result)
}
