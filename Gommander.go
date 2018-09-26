package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/integrii/flaggy"

	"github.com/joanbono/Gommander/modules/commander"
	"github.com/joanbono/Gommander/modules/configure"
)

// Defining colors
var yellow = color.New(color.Bold, color.FgYellow).SprintfFunc()
var red = color.New(color.Bold, color.FgRed).SprintfFunc()
var cyan = color.New(color.Bold, color.FgCyan).SprintfFunc()
var green = color.New(color.Bold, color.FgGreen).SprintfFunc()

var VERSION = `1.0.0`

//var BurpAPI, username, password, ApiToken string
var target, port string = "127.0.0.1", "1337"
var key, issue_type, issue_name string
var metrics, description, issues string
var scan, scan_id, username, password string

func init() {
	flaggy.SetName("Gommander")
	flaggy.SetDescription("Interact with Burp API")
	flaggy.DefaultParser.ShowVersionWithVersionFlag = false

	flaggy.String(&target, "t", "target", "Burp Address. Default 127.0.0.1")
	flaggy.String(&port, "p", "port", "Burp API Port. Default 1337")

	flaggy.String(&username, "U", "username", "Username for an authenticated scan")
	flaggy.String(&password, "P", "password", "Password for an authenticated scan")

	flaggy.String(&scan, "s", "scan", "URLs to scan")
	flaggy.String(&scan_id, "S", "scan-id", "Scanned URL identifier")

	flaggy.String(&metrics, "m", "metrics", "Provides metrics for a given task")
	flaggy.String(&description, "d", "description", "Provides description for a given issue")
	flaggy.String(&description, "I", "issues", "Provides issues for a given task")

	flaggy.String(&key, "k", "key", "Api Key")
	flaggy.String(&issue_type, "i", "issue-type", "String to search for")
	flaggy.String(&issue_name, "n", "issue-name", "String to search for")

}

func main() {
	flaggy.Parse()

	//configure.Configurer("test test")
	//fmt.Println(target, port)
	if configure.CheckBurp(target, port) == true {
		fmt.Fprintf(color.Output, "%v Found Burp API endpoint on %v.\n", green(" [+] SUCCESS:"), target+":"+port)
	} else {
		fmt.Fprintf(color.Output, "%v No Burp API endpoint found on %v.\n", red(" [-] ERROR:"), target+":"+port)
		os.Exit(0)
	}

	if scan != "" {
		//fmt.Println(configure.ScanConfig(target, port, scan))
		Location := configure.ScanConfig(target, port, scan)
		if Location != "" {
			fmt.Fprintf(color.Output, "%v Scanning %v over %v.\n", green(" [+] SUCCESS:"), scan, Location)
		} else {
			fmt.Fprintf(color.Output, "%v Can't start scan .\n", red(" [-] ERROR:"))
			os.Exit(0)
		}

	}

	if scan == "" && scan_id != "" {
		commander.GetScan(target, port, scan_id)
	}
	commander.Printer("test test")

}
