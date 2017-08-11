package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {

	smartyDomains := [...]string{
		"international-street.api.smartystreets.com",
		"us-street.api.smartystreets.com",
		"us-zipcode.api.smartystreets.com",
		"us-extract.api.smartystreets.com",
		"us-autocomplete.api.smartystreets.com",
		"download.api.smartystreets.com",
	}

	numFailedIps := 0

	for _, domain := range smartyDomains {
		dialTLS := func(network, addr string) (net.Conn, error) {
			return tls.Dial(network, addr, &tls.Config{ServerName: domain})
		}

		client := &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				DialTLS: dialTLS,
			},
		}

		ipAddresses, _ := net.LookupIP(domain) // take from 1st argument

		for _, ipAddress := range ipAddresses {
			request, err := http.NewRequest("GET", "https://"+ipAddress.String(), nil)

			if err != nil {
				fmt.Println("Request builder error: ", err)
			}

			message := "Resolving " + domain + " on " + ipAddress.String() + " ... "
			if _, err := client.Do(request); err != nil {
				message += "FAIL"
				numFailedIps++
			} else {
				message += "OK"
			}
			fmt.Println(message)
		}
	}
	fmt.Println(numFailedIps, " failures")
}
