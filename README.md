![](img/Gurp_logo.png)

## Usage

```bash
$ go run Gurp.go -h
Gurp - Interact with Burp API  Flags:
    -h --help  Displays help with available flag, subcommand, and positional value parameters.
    -t --target  Burp Address. Default 127.0.0.1
    -p --port  Burp API Port. Default 1337
    -U --username  Username for an authenticated scan
    -P --password  Password for an authenticated scan
    -s --scan  URLs to scan
    -S --scan-id  Scanned URL identifier
    -M --metrics  Provides metrics for a given task
    -D --description  Provides description for a given issue
    -d --description-names  Returns vulnerability names from PortSwigger
    -I --issues  Provides issues for a given task
    -e --export  Export issues json.
    -k --key  Api Key
    -v --version  Gurp version
```

+ Create a scan

```bash
go run Gurp.go -s "localhost.com/WebGoat/attack"
 [+] SUCCESS: Found Burp API endpoint on 127.0.0.1:1337.
 [i] INFO Setting up scanner...
 [+] SUCCESS: Scanning localhost.com/WebGoat/attack over 8.
```

+ Get Scan Metrics

```bash
go run Gurp.go -S 8 -M
 [+] SUCCESS: Found Burp API endpoint on 127.0.0.1:1337.
 [!] ALERT Retrieving Metrics from task 8
          [i] INFO: Scan status succeeded
          [i] INFO: 181 Requests made
          [i] INFO: 0 Requests queued
          [i] INFO: 6 Audit items completed
          [i] INFO: 0 Audit items waiting
          [i] INFO: 20058 Audit requests made
          [i] INFO: 2 Audit network errors
          [i] INFO: 5 Issue events
```

+ Get Issues from scan

```bash
go run Gurp.go -S 8 -I
 [+] SUCCESS: Found Burp API endpoint on 127.0.0.1:1337.
 [!] ALERT: Retrieving Issues from task 8
         [i] INFO: Frameable response (potential Clickjacking)
         [*] HIGH: Cleartext submission of password
         [*] LOW: Password field with autocomplete enabled
         [*] MEDIUM: Host header poisoning
         [i] INFO: Path-relative style sheet import
```

+ Export Issues' json

```bash
go run Gurp.go -S 8 -e /tmp
 [+] SUCCESS: Found Burp API endpoint on 127.0.0.1:1337.
 [!] ALERT: Retrieving Issues from task 8
         [i] INFO: Frameable response (potential Clickjacking)
         [*] HIGH: Cleartext submission of password
         [*] LOW: Password field with autocomplete enabled
         [*] MEDIUM: Host header poisoning
         [i] INFO: Path-relative style sheet import
 [!] ALERT: Exporting raw json to /tmp
```

+ Launch an authenticated scan with user/password

```bash
go run Gurp.go -s test.com -U admin -P 1234
 [+] SUCCESS: Found Burp API endpoint on 127.0.0.1:1337.
 [i] INFO Setting up scanner using credentials admin:1234
 [+] SUCCESS: Scanning test.com over 13.
```

+ Connect to Burp using API Key

```bash
go run Gurp.go -k "APIKEY" -d | grep -i SQL
         [2] SQL injection
         [3] SQL injection (second order)
         [35] Client-side SQL injection (DOM-based)
         [36] Client-side SQL injection (reflected DOM-based)
         [37] Client-side SQL injection (stored DOM-based)
         [68] SQL statement in request parameter
```

***

## Requirements

+ BurpSuite Pro `v2.0.0beta` or greater