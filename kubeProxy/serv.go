package main

import (
	"kuberclient"
	"net/http"
	"time"
)

var kuber kuberclient.Cluster

func main() {
	// sTime := time.Now()
	// println(etime.Seconds())
	go PollingCluster()

	http.Handle("/", new(staticHandler))
	http.ListenAndServe(":5000", nil)
}

func PollingCluster() {
	cnt := 1

	kuber.SetCluster()

	for {
		if cnt%2 == 0 {
			kuber.SetNodeUsage()
		}

		if cnt%5 == 0 {
			kuber.SetPodInfo()
		}

		if cnt%60 == 0 {
			kuber.SetCluster()
			cnt = 1
		}

		cnt++
		time.Sleep(time.Second * 1)
	}

}
