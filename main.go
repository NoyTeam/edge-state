package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chenjiandongx/pinger"
)

func pingApi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var hostList []string = []string{}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if r.Form["host"] != nil {
		json.Unmarshal([]byte(r.Form["host"][0]), &hostList)
		stats, err := pinger.ICMPPing(nil, hostList...)
		if err != nil {
			fmt.Println("Err:", err)
		}
		jsonStr, _ := json.Marshal(stats)
		fmt.Fprintln(w, `{"status":"ok","data":`+string(jsonStr)+`}`)
	} else {
		fmt.Fprintln(w, `{"status":"err"}`)
	}
}

func trafficApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	jsonStr, _ := json.Marshal(ct)
	fmt.Fprintln(w, string(jsonStr))
}

func main() {
	http.HandleFunc("/noyedge/ping", pingApi)
	http.HandleFunc("/noyedge/traffic", trafficApi)

	fmt.Println("Start NoyPL...")
	http.ListenAndServe(":18000", nil)
}
