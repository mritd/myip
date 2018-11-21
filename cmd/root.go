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

package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/mritd/myip/myip"

	"github.com/spf13/cobra"
)

var listenAddr net.IP
var listenPort int
var db string

var rootCmd = &cobra.Command{
	Use:   "myip",
	Short: "Show my ip address",
	Long: `
Show my ip address.`,
	Run: func(cmd *cobra.Command, args []string) {
		myip.Run(listenAddr, listenPort, db)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IPVarP(&listenAddr, "address", "l", net.ParseIP("0.0.0.0"), "listen address")
	rootCmd.PersistentFlags().IntVarP(&listenPort, "port", "p", 8080, "listen port")
	rootCmd.PersistentFlags().StringVarP(&db, "db", "d", "geoip.mmdb", "geo ip db path")
}
