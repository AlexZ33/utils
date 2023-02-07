package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var httpCli = NewSimpleHttpClient()

func TestHttpClient(t *testing.T) {

	data := bytes.NewBufferString("aaa")
	req, err := http.NewRequest("POST", "http://127.0.0.1:12345", data)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := httpCli.DoWithTimeout(10*time.Second, req)
	if err != nil {
		fmt.Println("Timeout:", err)
		return
	}
	fmt.Println("httpcode:", resp.StatusCode)
	respdata, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respdata))
}
