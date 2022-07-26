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
	var Kuwomodellist util.Kugoinfo
	
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
			//初始化数组  添加数组长度
			 Musiclinkoutlist:= make([]Musiclinkout,len(Kuwomodellist.Info))
			for ids, v := range Kuwomodellist.Info{
				
				if v.Filename!="" && v.Sqhash!=""{
					Musiclinkoutlist[ids].MusicName = v.Filename
					Musiclinkoutlist[ids].MusicLink = v.Sqhash
				}
				    

			}
			return Musiclinkoutlist, pageCount
		}

	}
	return nil, pageCount
}
