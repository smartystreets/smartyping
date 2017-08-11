package http

import (
	"net/http"
)

func Send(request *http.Request) ([]byte, error) {
	if _, err := http.DefaultClient.Do(request); err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
