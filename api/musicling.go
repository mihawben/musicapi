package api

import (
	"music/model"
	"net/http"
	

	"github.com/gin-gonic/gin"
)

//音乐名称和播放地址列表
func Musiclinklist(c *gin.Context){
page:=c.Query("page")
pagesize:=c.Query("pagesize")
name:=c.Query("name")
data,pageCount:=model.Musiclinklist(page,pagesize,name)
c.JSON(
	http.StatusOK, gin.H{
		//"status":  code,
		"data":    data,
		"total":   pageCount,
		// "message": errmsg.GetErrMsg(code),
	},

)
}
