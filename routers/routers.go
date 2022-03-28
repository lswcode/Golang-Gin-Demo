package routers

import (
	"1-gin_project/controllers"
	"1-gin_project/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()        // 创建一个默认的gin服务器
	router.LoadHTMLGlob("views/*") // 设置ctx.HTML返回html文件时的路径前缀，这里只写了一个/，所以views下面只能有一级子路径
	// router.LoadHTMLGlob("views/**/*") // 此时views下的文件夹必须是二级的，只有一级或者有三级都会报错
	// -----------------------------------------
	router.GET("/", controllers.TestController) // 在全局中间件被注册之前的路由请求不会触发全局中间件
	router.Use(middleware.M1)                   // 全局中间件，参数函数会自动被调用，除非要传参数，否则不需要加小括号

	router.GET("/test/:id", controllers.ParamsController)  // 此时前端发送的请求url中必须参参数: /test/123，不带参数会报错
	router.GET("/test2/*id", controllers.ParamsController) // 使用*+参数时，表示参数是可选的，可以写也可以不写
	router.POST("/form", controllers.FormController)       // 获取用户表单输入的数据
	router.POST("/axios", controllers.QueryController)     // 获取axios发送的查询参数
	router.POST("/file", controllers.FileController)       // 处理文件上传
	router.GET("/re", controllers.RedirectController)      // 重定向

	// 需要加小括号调用return的函数
	router.GET("/mw", middleware.M2(), middleware.M3(), controllers.UserController) // 单个路由独享的局部中间件

	router.StaticFile("/app", "static/index.html") // 作为web服务器，返回静态资源
	// --------------------------------------------------
	return router // 返回配置完成的路由
}
