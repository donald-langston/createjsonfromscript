package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getHostName() (string, error) {
	return os.Hostname()
}

func getIpAddress() (string, bool) {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipnet := addr.(*net.IPNet).IP
			// To4() will return nil if not a ipv4 address
			if !ipnet.IsLoopback() && ipnet.To4() != nil {
				return ipnet.To4().String(), true
			}
		}
	}
	return "", false
}

func getTimeStamp() string {
	now := time.Now()
	timestamp := now.Format("3:04pm")
	return timestamp
}

func main() {

	if len(os.Args) > 1 {
		mapB := make(map[string]string)

		for _, arg := range os.Args[1:] {
			if arg == "hostname" {
				hostname, _ := getHostName()
				mapB[arg] = hostname
			} else if arg == "ipaddress" {
				ipaddress, success := getIpAddress()
				if success {
					mapB[arg] = ipaddress
				}
			} else if arg == "timestamp" {
				mapB[arg] = getTimeStamp()
			} else {
				fmt.Println("Unknown argument was passed")
			}
		}

		data, _ := json.Marshal(mapB)
		f, err := os.Create("test.json")
		check(err)
		defer f.Close()
		f.WriteString(string(data))
		f.Sync()
	}
}
