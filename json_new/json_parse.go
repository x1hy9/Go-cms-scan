package json_new

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FofaFinger struct {
	RuleID         string `json:"rule_id"`
	Level          string `json:"level"`
	Softhard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Rules          [][]struct {
		Match   string `json:"match"`
		Content string `json:"content"`
	} `json:"rules"`
}

type FetchResult struct {
	Url          string
	Content      []byte
	Headers      http.Header
	HeaderString string
	Certs        []byte
}

//返回切片数据
func Parse(filename string) ([]FofaFinger, error) {

	Json, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var dataArray []FofaFinger
	err = json.Unmarshal(Json, &dataArray)
	if err != nil {
		return nil, err
	}

	return dataArray, nil
}
