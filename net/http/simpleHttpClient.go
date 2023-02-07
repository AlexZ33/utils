package http

import (
	"bytes"
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type SimpleHttpClient struct {
	header  map[string]string
	httpCli *http.Client
}

func NewSimpleHttpClient() *SimpleHttpClient {
	return &SimpleHttpClient{
		httpCli: &http.Client{},
		header:  make(map[string]string),
	}
}

func (client *SimpleHttpClient) GetClient() *http.Client {
	return client.httpCli
}

func (client *SimpleHttpClient) NewTransPort() *http.Transport {
	return &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		ResponseHeaderTimeout: 30 * time.Second,
	}
}

func (client *SimpleHttpClient) SetTimeOut(timeOut time.Duration) {
	client.httpCli.Timeout = timeOut
}

func (client *SimpleHttpClient) SetHeader(key, value string) {
	client.header[key] = value
}

func (client *SimpleHttpClient) GET(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "GET", header, data)

}

func (client *SimpleHttpClient) POST(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "POST", header, data)
}

func (client *SimpleHttpClient) DELETE(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "DELETE", header, data)
}

func (client *SimpleHttpClient) PUT(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "PUT", header, data)
}

func (client *SimpleHttpClient) GETEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "GET", header, data)
}

func (client *SimpleHttpClient) POSTEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "POST", header, data)
}

func (client *SimpleHttpClient) DELETEEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "DELETE", header, data)
}

func (client *SimpleHttpClient) PUTEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "PUT", header, data)
}

func (client *SimpleHttpClient) Request(url, method string, header http.Header, data []byte) ([]byte, error) {
	var req *http.Request
	var errReq error
	if data != nil {
		req, errReq = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		req, errReq = http.NewRequest(method, url, nil)
	}

	if errReq != nil {
		return nil, errReq
	}

	req.Close = true

	if header != nil {
		req.Header = header
	}

	for key, value := range client.header {
		req.Header.Set(key, value)
	}

	rsp, err := client.httpCli.Do(req)
	if err != nil {
		return nil, err
	}

	/*if rsp.StatusCode >= http.StatusBadRequest {
		 return 0, nil, fmt.Errorf("statuscode:%d, status:%s", rsp.StatusCode, rsp.Status)
	 }*/

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)

	return body, err
}

func (client *SimpleHttpClient) RequestEx(url, method string, header http.Header, data []byte) (int, []byte, error) {
	var req *http.Request
	var errReq error
	if data != nil {
		req, errReq = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		req, errReq = http.NewRequest(method, url, nil)
	}

	if errReq != nil {
		return 0, nil, errReq
	}

	req.Close = true

	if header != nil {
		req.Header = header
	}

	for key, value := range client.header {
		req.Header.Set(key, value)
	}

	rsp, err := client.httpCli.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)

	return rsp.StatusCode, body, err
}

func (client *SimpleHttpClient) DoWithTimeout(timeout time.Duration, req *http.Request) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	return client.httpCli.Do(req)
}
