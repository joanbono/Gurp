package commander

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

var yellow = color.New(color.Bold, color.FgYellow).SprintfFunc()
var red = color.New(color.Bold, color.FgRed).SprintfFunc()
var cyan = color.New(color.Bold, color.FgCyan).SprintfFunc()
var green = color.New(color.Bold, color.FgGreen).SprintfFunc()

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

func GetMetrics(target, port, Location string) {
	var endpoint string = "http://" + target + ":" + port + "/v0.1/scan/" + Location

	resp, err := client.Get(endpoint)

	if err != nil {
		fmt.Fprintf(color.Output, "%v Can't perform request to %v.\n", red(" [-] ERROR"), endpoint)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(color.Output, "%v Scan ID %v not found.\n", red(" [-] ERROR"), Location)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		//var data = json.NewDecoder(resp.Body)
		//fmt.Fprintf(color.Output, "%v Response body:\n %v \n", cyan(" [i] INFO"), string(body))
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

func GetScan(target, port, Location string) {
	var issue_uniq string = ""
	var endpoint string = "http://" + target + ":" + port + "/v0.1/scan/" + Location

	resp, err := client.Get(endpoint)

	if err != nil {
		fmt.Fprintf(color.Output, "%v Can't perform request to %v.\n", red(" [-] ERROR"), endpoint)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(color.Output, "%v Scan ID %v not found.\n", red(" [-] ERROR"), Location)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		//var data = json.NewDecoder(resp.Body)
		//fmt.Fprintf(color.Output, "%v Response body:\n %v \n", cyan(" [i] INFO"), string(body))
		fmt.Fprintf(color.Output, "%v Retrieving alerts from task %v \n", yellow(" [!] ALERT"), Location)

		value := gjson.Get(string(body), "issue_events")
		//println(value.String())
		//println(value.String())
		//test := gjson.Get(value.String(), "..#.id")
		t1 := value.String()[1 : len(value.String())-1]
		//println(t1)
		issue_names := gjson.Get(t1, "..#.issue.name")
		severity := gjson.Get(t1, "..#.issue.severity")
		//println(issue_names.String())

		var index []string
		for _, criticity := range severity.Array() {
			index = append(index, criticity.String())
		}

		//var test []string(severity.Array())
		// Printing issue names removing duplicates
		for k, name := range issue_names.Array() {
			if issue_uniq != name.String() {
				//println(index[k])
				fmt.Fprintf(color.Output, "\t %v --%v-- %v \n", cyan(" [!] INFO:"), index[k], name.String())
				//println(name.String())
			}
			issue_uniq = name.String()
		}
	}
}
