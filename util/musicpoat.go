package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/mitchellh/mapstructure"
)

type Kuwomodel struct {
	Filename string `json:"filename"`
	Sqhash   string `json:"sqhash"`
	Key      string `json:"key"`
}

//调用接口 返回音乐链接
func Kuwomusic(urlname string) []Kuwomodel {

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", urlname, nil)
	if err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }
	var Kuwomodellist []Kuwomodel
	reqest.Header.Set("Accept", "*/*")
	//reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "no-cache")
	reqest.Header.Set("Connection", "keep-alive")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

	response, _ := client.Do(reqest)
	bodystr := ""
	if response.StatusCode == 200 {

		returnMap, _ := ParseResponse(response)

		body, _ := ioutil.ReadAll(response.Body)
		bodystr = string(body)
		fmt.Println(bodystr)
		err := mapstructure.Decode(returnMap, &Kuwomodellist)
		if err == nil {
			return Kuwomodellist
		}

	}

	return Kuwomodellist

}

//返回值 map处理
func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}
