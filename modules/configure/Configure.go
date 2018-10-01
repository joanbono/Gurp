package configure

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
	strip "github.com/grokify/html-strip-tags-go"
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

// [+] Check Burp Connection alive
// [ ] Check if API is valid
// [ ] Initialize Burp

func Configurer(test string) {
	fmt.Fprintf(color.Output, "%v %v TEST\n", red(" [-] ERROR"), test)
	fmt.Fprintf(color.Output, "%v %v TEST\n", yellow(" [!] ALERT"), test)
	fmt.Fprintf(color.Output, "%v %v TEST\n", cyan(" [i] INFO"), test)
	fmt.Fprintf(color.Output, "%v %v TEST\n", green(" [+] SUCCESS"), test)
}

func CheckBurp(target, port string) (response bool) {
	var endpoint string = "http://" + target + ":" + port + "/v0.1/"

	resp, err := client.Get(endpoint)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 {
		return true
	} else {
		return false
	}

}

func ScanConfig(target, port, urls, username, password string) (ScanLocation string) {
	var endpoint string = "http://" + target + ":" + port + "/v0.1/scan"
	var url_string string
	// At the moment, this only allows 1 url to be scanned
	if username == "" && password == "" {
		fmt.Fprintf(color.Output, " %v Setting up scanner...\n", cyan("[i] INFO"))
		url_string = `{"urls":["` + urls + `"]}`
	} else {
		fmt.Fprintf(color.Output, " %v Setting up scanner using credentials %v:%v\n", cyan("[i] INFO"), username, password)
		url_string = `{"application_logins":[{"password":"` + password + `","username":"` + username + `"}],"urls":["` + urls + `"]}`
	}
	var body = []byte(url_string)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error")
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error")
	}
	Location := resp.Header.Get("Location")
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		fmt.Println("Error")
	}

	return Location
}

func GetDescription(target, port, issueName string) {
	//println(issueName)

	var endpoint string = "http://" + target + ":" + port + "/v0.1/knowledge_base/issue_definitions"

	resp, err := client.Get(endpoint)

	if err != nil {
		fmt.Fprintf(color.Output, "%v Can't perform request to %v.\n", red(" [-] ERROR:"), endpoint)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(color.Output, "%v Resource not found.\n", red(" [-] ERROR:"))
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Fprintf(color.Output, "%v Retrieving Issues from task %v \n", yellow(" [!] ALERT:"), Location)
		fmt.Fprintf(color.Output, "%v Fetching '%v' information...\n", cyan(" [i] INFO:"), issueName)
		//println(string(body))
		raw_issues := string(body)[1 : len(string(body))-1]
		var descriptionSelected string = `..#[name="` + issueName + `"]`
		//value := gjson.Get(raw_issues, "..#.name")
		value := gjson.Get(raw_issues, descriptionSelected)
		//println(value.String())

		description := gjson.Get(value.String(), "description")
		desc_stripped := strip.StripTags(description.String())
		remediation := gjson.Get(value.String(), "remediation")
		rem_stripped := strip.StripTags(remediation.String())

		fmt.Fprintf(color.Output, "\t %v %v\n", cyanBG(" [*] DESCRIPTION:"), desc_stripped)
		fmt.Fprintf(color.Output, "\t %v %v\n", greenBG(" [*] REMEDIATION:"), rem_stripped)

		// var VulnNames []string
		// for _, test := range value.Array() {
		// 	//println(test.String())
		// 	VulnNames = append(VulnNames, test.String())
		// 	println(test.String())

		// }

		//fmt.Println(VulnNames)
		//raw_issues := value.String()[1 : len(value.String())-1]
	}

}
