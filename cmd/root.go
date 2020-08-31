/*
Copyright © 2020 Luke Reed

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	port    string
	version string
)

type DateTimeReturn struct {
	DateTime   time.Time `json:"dateTime"`
	AppVersion string    `json:"appVersion"`
}

func init() {
	rootCmd.PersistentFlags().StringVar(&port, "port", "8080", "port default is 8080")
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "dtapi",
	Short: "A simple json API that returns the datetime and version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the version of datetime",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(VERSION string) {
	version = VERSION
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func getUnixTime() time.Time {
	return time.Now()
}

func getVersion() string {
	return version
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	ret := DateTimeReturn{
		DateTime:   getUnixTime(),
		AppVersion: getVersion(),
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
