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

//返回歌单及链接

func Musiclinklist(page string, pagesize string, name string) ([]util.Musiclinkout, int) {

	var MusiclikList []MusicLink
	var Musiclinklist []util.Musiclinkout

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
			Musiclinklist = util.Kuwomusic(s3)
			//初始化数组  添加数组长度
			Musiclinkoutlist := make([]util.Musiclinkout, len(Musiclinklist))
			for ids, v := range Musiclinklist {

				if v.Author_name != "" && v.Song_name != "" {
					Musiclinkoutlist[ids].Author_name = v.Author_name
					Musiclinkoutlist[ids].Song_name = v.Song_name
					Musiclinkoutlist[ids].Lyrics = v.Lyrics
					Musiclinkoutlist[ids].Play_url = v.Play_url
				}

			}
			return Musiclinkoutlist, pageCount
		}

	}
	return nil, pageCount
}
