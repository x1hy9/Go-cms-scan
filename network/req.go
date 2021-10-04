package network

import (
	"encoding/json"
	"fmt"

	"net/http"

	"scan/json_new"
	"time"

	"github.com/asmcos/requests"
)

func Reqdata(url string) (*json_new.FetchResult, error) {
	req := requests.Requests()

	req.SetTimeout(time.Duration(10))

	resp, err := req.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	var headerString string

	req_data := json_new.FetchResult{
		Url:          url,
		Content:      resp.Content(),
		Headers:      resp.R.Header,
		HeaderString: headerString,
		Certs:        getCerts(resp.R),
	}
	return &req_data, nil

}

// 获取证书内容，参考byro07/fwhatweb
func getCerts(resp *http.Response) []byte {
	var certs []byte
	if resp.TLS != nil {
		cert := resp.TLS.PeerCertificates[0]
		var str string
		if js, err := json.Marshal(cert); err == nil {
			certs = js
		}
		str = string(certs) + cert.Issuer.String() + cert.Subject.String()
		certs = []byte(str)
	}
	return certs
}
