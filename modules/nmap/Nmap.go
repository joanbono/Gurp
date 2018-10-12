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

package nmap

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joanbono/color"
	nmap "github.com/tomsteele/go-nmap"
)

var red = color.New(color.Bold, color.FgRed).SprintfFunc()

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

// ParseFile return a slice of targets
func ParseFile(filename string) (result []string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Buf := make([]byte, 0, 2048*2048) //bytes.Buffer
	scanner.Buffer(Buf, 10*2048*2048)

	for scanner.Scan() {
		target := scanner.Text()
		result = append(result, target)
	}

	return result

}
