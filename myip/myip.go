/*
 * Copyright 2018 mritd <mritd1234@gmail.com>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package myip

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/oschwald/geoip2-golang"
)

var dbPath string

func remoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := remoteIp(r)
	var addr string
	db, err := geoip2.Open(dbPath)
	if err != nil {
		logrus.Error(err)
		return
	} else {
		city, err := db.City(net.ParseIP(ip))
		if err != nil {
			logrus.Error(err)
			addr = "Unknown"
		} else {
			addr = fmt.Sprintf("%s %s %s", city.Continent.Names["zh-CN"], city.Country.Names["zh-CN"], city.City.Names["zh-CN"])
		}
	}

	_, err = w.Write([]byte(fmt.Sprintf("IP: %s\nAddress: %s\n", ip, addr)))
	if err != nil {
		logrus.Error(err)
	}
}

func getIP2Json(w http.ResponseWriter, r *http.Request) {
	ip := remoteIp(r)
	var addr string
	db, err := geoip2.Open(dbPath)
	if err != nil {
		logrus.Error(err)
		return
	} else {
		city, err := db.City(net.ParseIP(ip))
		if err != nil {
			logrus.Error(err)
			addr = "Unknown"
		} else {
			addr = fmt.Sprintf("%s %s %s", city.Continent.Names["zh-CN"], city.Country.Names["zh-CN"], city.City.Names["zh-CN"])
		}
	}

	_, err = w.Write([]byte(fmt.Sprintf(`{"IP":"%s","Address":"%s"}`, ip, addr)))
	if err != nil {
		logrus.Error(err)
	}
}

func Run(host net.IP, port int, db string) {
	dbPath = db
	fmt.Printf("Server listening at: %s\n", fmt.Sprintf("%s:%d", host, port))
	http.HandleFunc("/", getIP)
	http.HandleFunc("/json", getIP2Json)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
