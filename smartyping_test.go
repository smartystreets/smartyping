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
		"international-street",
		"us-street",
		"us-zipcode",
		"us-extract",
		"us-autocomplete",
		"download",
	}

	for _, domain := range domains {
		domain += ".api.smartystreets.com"

		ips, err := net.LookupIP(domain)
		if err != nil {
			t.Errorf("Could not resolve IP addresses for %s: %s", domain, err)
		} else {
			for _, ip := range ips {
				t.Run(domain+"--"+ip.String(), func(t *testing.T) {
					t.Parallel()

					if response, err := ping(ip.String(), domain); err != nil {
						t.Error(err)
					} else {
						cleanup(response)
					}
				})
			}
		}
	}
}
func ping(ip string, domain string) (*http.Response, error) {
	return buildClient(ip, domain).Get("https://" + domain)
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
