package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type Kuwomodel struct {
	//Filename string `json:"filename"`
	//Sqhash   string `json:"sqhash"`
	//Key      string `json:"key"`
	Status  float64                `json:"status"`
	Errcode float64                `json:"errcode"`
	Data    map[string]interface{} `json:"data"`
}
type Kugoinfo struct {
	Info []Kugohash `json:"info"`
}
type Kugohash struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
	Key      string `json:"key"`
}

//返回音乐名称和音乐链接及歌词
type Musiclinkout struct {
	Author_name string `json:"author_name"`
	Song_name   string `json:"song_name"`
	Play_url    string `json:"play_url"`
	Lyrics      string `json:"lyrics"`
}

//调用接口 返回音乐链接
func Kuwomusic(urlname string) []Musiclinkout {

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", urlname, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	var Kuwomodellist Kugoinfo
	var Kuwomodelone Kuwomodel

	reqest.Header.Set("Accept", "*/*")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "requests+bs4.BeautifulSoup")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "no-cache")
	reqest.Header.Set("Connection", "keep-alive")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

	response, _ := client.Do(reqest)
	bodystr := ""
	if response.StatusCode == 200 {

		body, _ := ioutil.ReadAll(response.Body)
		bodystr = string(body)

		log.Printf("调用成功 数据为%s 读取完成/n", bodystr)
		fmt.Println(bodystr)
		returnMap := ParseResponse(body)
		err := mapstructure.Decode(returnMap, &Kuwomodelone)

		if err == nil {
			err := mapstructure.Decode(Kuwomodelone.Data, &Kuwomodellist)
			if err == nil {
				fmt.Println(err)
				//return Kuwomodellist
			}
		}
		 Musiclinkoutlist :=make([]Musiclinkout,len(Kuwomodellist.Info))
		if len(Kuwomodellist.Info) != 0 {
			for n, v := range Kuwomodellist.Info {
				//	databyd:=[]byte(v.Sqhash+"kgcloud")
				var Musiclinkout Musiclinkout
				v.Key = "484a7efeea23ffd3e7192dd7fc6bedb0"
				urls := "https://wwwapi.kugou.com/yy/index.php?r=play/getdata&hash=" + v.Hash + "&mid=" + v.Key + "&platid=4&album_id=973367"

				reqestlink, _ := http.NewRequest("GET", urls, nil)
				reqestlink.Header.Set("Accept", "*/*")
				reqestlink.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
				reqestlink.Header.Set("Accept-Encoding", "requests+bs4.BeautifulSoup")
				reqestlink.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
				reqestlink.Header.Set("Cache-Control", "no-cache")
				reqestlink.Header.Set("Connection", "keep-alive")
				reqestlink.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62")

				responselink, _ := client.Do(reqestlink)
				bodylink, _ := ioutil.ReadAll(responselink.Body)
				//bodylinkstr := string(bodylink)

				bodymap := ParseResponse(bodylink)
				err := mapstructure.Decode(bodymap, &Kuwomodelone)

				if err != nil {
					fmt.Println("链接错误")
				}

				errs := mapstructure.Decode(Kuwomodelone.Data, &Musiclinkout)
				if errs != nil {
					fmt.Println("转化错误")
				}
				Musiclinkoutlist[n].Author_name = Musiclinkout.Author_name
				Musiclinkoutlist[n].Lyrics = Musiclinkout.Lyrics
				Musiclinkoutlist[n].Play_url = Musiclinkout.Play_url
				Musiclinkoutlist[n].Song_name = Musiclinkout.Song_name

				fmt.Println("写完")

			}
			
		}
		return Musiclinkoutlist
	}

	return nil

}

//md5处理
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

//返回值 map处理
func ParseResponse(Body []byte) map[string]interface{} {
	var result map[string]interface{}

	_ = json.Unmarshal(Body, &result)

	return result
}
