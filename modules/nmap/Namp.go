package nmap

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	nmap "github.com/tomsteele/go-nmap"
)

//ParseNmap returns a slice of targets
func ParseNmap(fileName string) (result []string, err error) {

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	contentType := http.DetectContentType(dat)
	isXML := strings.Split(contentType, ";")

	if isXML[0] == "text/xml" {
		scan, err := nmap.Parse(dat)
		if err != nil {
			log.Fatal(err)
		}
		for i, host := range scan.Hosts {
			for _, port := range host.Ports {
				if port.Service.Name == "http" || port.Service.Name == "https" {
					target := port.Service.Name + "://" + host.Addresses[i].Addr + ":" + strconv.Itoa(port.PortId)
					result = append(result, target)
					// fmt.Printf("Ip: %s Port: %v \n", host.Addresses[i].Addr, port.PortId)
				}

			}

		}

	} else {
		err = errors.New("This file is not an xml file")
	}
	return result, err

}
