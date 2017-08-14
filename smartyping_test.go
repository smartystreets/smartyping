package smartyping

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t.Parallel()

	domains := []string{
		"international-street.api.smartystreets.com",
		"us-street.api.smartystreets.com",
		"us-zipcode.api.smartystreets.com",
		"us-extract.api.smartystreets.com",
		"us-autocomplete.api.smartystreets.com",
		"download.api.smartystreets.com",
	}

	for _, domain := range domains {

		ips, err := net.LookupIP(domain)
		if err != nil {
			t.Fatalf("Could not resolve IP addresses for %s: %s", domain, err)
		}

		for _, ip := range ips {
			t.Run(domain+"--"+ip.String(), func(t *testing.T) {
				t.Parallel()

				if response, err := ping(ip, domain); err != nil {
					t.Error(err)
				} else {
					cleanup(response)
				}
			})
		}
	}
}
func ping(ip net.IP, domain string) (*http.Response, error) {
	return buildClient(ip.String(), domain).Get("https://" + domain)
}

func buildClient(ip, domain string) *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			DialTLS: func(network, address string) (net.Conn, error) {
				return tls.Dial(network, ip+":443", &tls.Config{ServerName: domain})
			},
		},
	}
}

func cleanup(response *http.Response) {
	ioutil.ReadAll(response.Body)
	response.Body.Close()
}
