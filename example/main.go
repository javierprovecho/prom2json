package main

import (
	"fmt"

	"github.com/javierprovecho/prom2json"
)

func main() {
	metrics, err := prom2json.FetchMetricFamilies("http://146.185.151.73:9100/metrics")

	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"%f",
		metrics["node_network_transmit_bytes"].
			GetMetric()[0].
			GetGauge().
			GetValue(),
	)
}
