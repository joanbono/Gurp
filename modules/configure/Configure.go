package configure

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
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

func ScanConfig(target, port, urls string) (ScanLocation string) {
	var endpoint string = "http://" + target + ":" + port + "/v0.1/scan"

	// At the moment, this only allows 1 url to be scanned
	var url_string string = `{"urls":["` + urls + `"]}`

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
