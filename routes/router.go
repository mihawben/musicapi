package routes

import (
	"music/api"
	"music/util"

	"github.com/gin-gonic/gin"
)

//func createMyRender() multitemplate.Renderer {
//	p := multitemplate.NewRenderer()
//	p.AddFromFiles("admin", "web/admin/dist/index.html")
//	p.AddFromFiles("front", "web/front/dist/index.html")
//	return p
//}

func InitRouter() {
	gin.SetMode(util.AppMode)

	r := gin.New()
	//r.HTMLRender = createMyRender()
	//	r.Use(middleware.Log())

	r.Use(gin.Recovery())
	//	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})
	/*
		后端页面接口
	*/

	/*
		前端展示页面接口
	*/
	router := r.Group("api")
	{

		//查看所有音乐
		router.GET("musice/getlist", api.Musiclinklist)

	}

	_ = r.Run(util.HttpPort)

}
