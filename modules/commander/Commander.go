/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package commander

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joanbono/color"
	"github.com/tidwall/gjson"
)

var yellow = color.New(color.Bold, color.FgYellow).SprintfFunc()
var red = color.New(color.Bold, color.FgRed).SprintfFunc()
var cyan = color.New(color.Bold, color.FgCyan).SprintfFunc()
var green = color.New(color.Bold, color.FgGreen).SprintfFunc()
var redBG = color.New(color.Bold, color.FgWhite, color.BgHiRed).SprintfFunc()
var cyanBG = color.New(color.Bold, color.FgBlack, color.BgHiCyan).SprintfFunc()
var yellowBG = color.New(color.Bold, color.FgBlack, color.BgHiYellow).SprintfFunc()
var greenBG = color.New(color.Bold, color.FgBlack, color.BgHiGreen).SprintfFunc()

// Skipping SSL verification
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//Proxy:           http.ProxyURL(proxyUrl),
}

var client = &http.Client{Timeout: time.Second * 5, Transport: tr}

func Printer(test string) {
	fmt.Fprintf(color.Output, "%v %v TEST\n", red(" [-] ERROR"), test)
	fmt.Fprintf(color.Output, "%v %v TEST\n", yellow(" [!] ALERT"), test)
	fmt.Fprintf(color.Output, "%v %v TEST\n", cyan(" [i] INFO"), test)
	fmt.Fprintf(color.Output, "%v %v TEST\n", green(" [+] SUCCESS"), test)
}

func GetMetrics(target, port, Location, apikey string) {
	var endpoint string
	if apikey != "" {
		endpoint = "http://" + target + ":" + port + "/" + apikey + "/v0.1/scan/" + Location
	} else {
		endpoint = "http://" + target + ":" + port + "/v0.1/scan/" + Location
	}

	resp, err := client.Get(endpoint)

	if err != nil {
		fmt.Fprintf(color.Output, "%v Can't perform request to %v.\n", red(" [-] ERROR"), endpoint)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(color.Output, "%v Scan ID %v not found.\n", red(" [-] ERROR"), Location)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		//println(string(body))
		fmt.Fprintf(color.Output, "%v Retrieving Metrics from task %v \n", yellow(" [!] ALERT"), Location)

		//Retrieving info from response
		scan_status := gjson.Get(string(body), "scan_status")
		crawl_requests_made := gjson.Get(string(body), "scan_metrics.crawl_requests_made")
		crawl_requests_queued := gjson.Get(string(body), "scan_metrics.crawl_requests_queued")
		audit_queue_items_completed := gjson.Get(string(body), "scan_metrics.audit_queue_items_completed")
		audit_queue_items_waiting := gjson.Get(string(body), "scan_metrics.audit_queue_items_waiting")
		audit_requests_made := gjson.Get(string(body), "scan_metrics.audit_requests_made")
		audit_network_errors := gjson.Get(string(body), "scan_metrics.audit_network_errors")
		issue_events := gjson.Get(string(body), "scan_metrics.issue_events")

		// Printing the info
		fmt.Fprintf(color.Output, "\t %v Scan status %v\n", cyan(" [i] INFO:"), scan_status)
		fmt.Fprintf(color.Output, "\t %v %v Requests made\n", cyan(" [i] INFO:"), crawl_requests_made)
		fmt.Fprintf(color.Output, "\t %v %v Requests queued\n", cyan(" [i] INFO:"), crawl_requests_queued)
		fmt.Fprintf(color.Output, "\t %v %v Audit items completed\n", cyan(" [i] INFO:"), audit_queue_items_completed)
		fmt.Fprintf(color.Output, "\t %v %v Audit items waiting\n", cyan(" [i] INFO:"), audit_queue_items_waiting)
		fmt.Fprintf(color.Output, "\t %v %v Audit requests made\n", cyan(" [i] INFO:"), audit_requests_made)
		fmt.Fprintf(color.Output, "\t %v %v Audit network errors\n", cyan(" [i] INFO:"), audit_network_errors)
		fmt.Fprintf(color.Output, "\t %v %v Issue events\n", cyan(" [i] INFO:"), issue_events)
	}
}

func GetScan(target, port, Location, exportFolder, apikey string) {
	var issue_uniq string = ""
	var endpoint string
	if apikey != "" {
		endpoint = "http://" + target + ":" + port + "/" + apikey + "/v0.1/scan/" + Location
	} else {
		endpoint = "http://" + target + ":" + port + "/v0.1/scan/" + Location
	}

	resp, err := client.Get(endpoint)

	if err != nil {
		fmt.Fprintf(color.Output, "%v Can't perform request to %v.\n", red(" [-] ERROR:"), endpoint)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(color.Output, "%v Scan ID %v not found.\n", red(" [-] ERROR:"), Location)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(color.Output, "%v Retrieving Issues from task %v \n", yellow(" [!] ALERT:"), Location)

		value := gjson.Get(string(body), "issue_events")
		raw_issues := value.String()[1 : len(value.String())-1]

		// at raw_issues we have every field in json format
		// println(raw_issues)
		issue_names := gjson.Get(raw_issues, "..#.issue.name")
		severity := gjson.Get(raw_issues, "..#.issue.severity")
		//println(issue_names.String())

		var index []string
		for _, criticity := range severity.Array() {
			index = append(index, criticity.String())
		}

		if len(index) == 0 {
			fmt.Fprintf(color.Output, "%v No issues found.\n", cyan(" [i] INFO:"))
		}

		// Printing issue names removing duplicates
		for k, name := range issue_names.Array() {
			if issue_uniq != name.String() {
				if index[k] == "low" {
					fmt.Fprintf(color.Output, "\t %v %v \n", greenBG("[*] LOW:"), name.String())
				} else if index[k] == "high" {
					fmt.Fprintf(color.Output, "\t %v %v \n", redBG("[*] HIGH:"), name.String())
				} else if index[k] == "medium" {
					fmt.Fprintf(color.Output, "\t %v %v \n", yellowBG("[*] MEDIUM:"), name.String())
				} else if index[k] == "info" {
					fmt.Fprintf(color.Output, "\t %v %v \n", cyanBG("[i] INFO:"), name.String())
				}
			}
			issue_uniq = name.String()
		}

		if exportFolder != "" {
			fmt.Fprintf(color.Output, "%v Exporting raw json to %v \n", yellow(" [!] ALERT:"), exportFolder)

			if _, err := os.Stat(exportFolder); !os.IsNotExist(err) {
				// Write raw_issues to file
				f, err := os.Create(exportFolder + "/Burp_export.json")
				if err != nil {
					panic(err)
				}
				defer f.Close()
				f.WriteString(value.String())
				f.Sync()

			} else {
				fmt.Fprintf(color.Output, "%v Folder %v don't exists.\n", red(" [-] ERROR:"), exportFolder)
			}
		}
	}
}
