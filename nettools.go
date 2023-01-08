package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type TrafficCache struct {
	Uploads   []int64 `json:"uploads"`
	Downloads []int64 `json:"downloads"`
}

var networkCard string
var ct TrafficCache
var nowUpload, nowDownload int64

func init() {
	cmd := exec.Command("bash", "-c", "route  | grep default  | awk '{print $8}'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	networkCard = strings.ReplaceAll(string(out), "\n", "")
	nowUpload, nowDownload = getTraffic()
	ct.Downloads = []int64{}
	ct.Uploads = []int64{}
	cacheTraffic()
}

func cacheTraffic() {
	u, d := getTraffic()
	if d-nowDownload != 0 && u-nowUpload != 0 {
		ct.Downloads = append(ct.Downloads, d-nowDownload)
		ct.Uploads = append(ct.Uploads, u-nowUpload)
	}
	if len(ct.Downloads) > 5 {
		ct.Downloads = ct.Downloads[len(ct.Downloads)-5:]
	}
	if len(ct.Uploads) > 5 {
		ct.Uploads = ct.Uploads[len(ct.Uploads)-5:]
	}
	time.AfterFunc(time.Minute, cacheTraffic)
}

func removeEmpty(l []string) []string {
	var new []string = []string{}
	for i := range l {
		if l[i] != "" {
			new = append(new, l[i])
		}
	}
	return new
}

func getTraffic() (int64, int64) {
	file, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		fmt.Println("OpenErr:", err)
	}

	t := strings.Split(string(file), "\n")
	var n []string

	for i := 2; i < len(t); i++ {
		info := removeEmpty(strings.Split(t[i], " "))
		if len(info) != 0 && info[0] == networkCard+":" {
			n = info[1:]
		}
	}

	receive, _ := strconv.Atoi(n[0])
	transmit, _ := strconv.Atoi(n[8])

	// Receive Transmit
	return int64(transmit), int64(receive)
}
