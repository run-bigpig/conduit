package utils

import (
	"bytes"
	"github.com/run-bigpig/conduit/internal/config"
	"github.com/run-bigpig/conduit/internal/constants"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

func SendRequest(url string, method string, header map[string]any, body []byte) (*http.Response, error) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value.(string))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetDomain() string {
	domain := config.ReadString("domain")
	if domain == "" {
		return constants.Domain
	}
	return domain
}

func GetPort() int {
	port := cast.ToInt(config.Read("port"))
	if port != 443 {
		return port
	}
	return 443
}
