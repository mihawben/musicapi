package model

import (
	
	"music/util"
	"strings"

	"github.com/jinzhu/gorm"
)

type MusicLink struct {
	gorm.Model
	Linkname string `gorm:"type:varchar(250);not null" json:"linkname" label:"链接名称"`
	Link     string `gorm:"type:varchar(500);not null" json:"link" label:"链接"`
}

type Musiclinkout struct {
	MusicName string `json:"musicname"`
	MusicLink string `json:"musiclink"`
}

//返回歌单及链接

func Musiclinklist(page string, pagesize string, name string) ([]Musiclinkout, int) {

	var MusiclikList []MusicLink
	var Kuwomodellist []util.Kuwomodel
	var Musiclinkoutlist []Musiclinkout
	var pageCount int
	err = db.Select("Linkname", "Link").Find(&MusiclikList).Error
	for _, s := range MusiclikList {
		if s.Linkname == "酷狗" {
			var linke = s.Link
			var build strings.Builder
			build.WriteString(linke)
			build.WriteString("?format=json&keyword=")
			build.WriteString(name)
			build.WriteString("&page=")
			build.WriteString(page)
			build.WriteString("&pagesize=")
			build.WriteString(pagesize)
			s3 := build.String()
			Kuwomodellist = util.Kuwomusic(s3)
			for ids, v := range Kuwomodellist {
				
				Musiclinkoutlist[ids].MusicName = v.Data
				Musiclinkoutlist[ids].MusicLink = v.Data

			}
			return Musiclinkoutlist, pageCount
		}

	}
	return Musiclinkoutlist, pageCount
}
