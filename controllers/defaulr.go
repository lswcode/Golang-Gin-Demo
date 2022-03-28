package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestController(ctx *gin.Context) { // gin把request和response都封装到了gin.Context
	// ctx.String(200, "hello gin")                                 // 返回一个字符串
	ctx.HTML(http.StatusOK, "index.html", nil) // 返回一个页面，第三个参数是传递给HTML文件的数据，可以省略
}

func ParamsController(ctx *gin.Context) {
	id := ctx.Param("id") // 使用Param()获取路径参数
	ctx.String(200, id)   // 显示某个字符串，第二个参数就是要显示的内容

}

type User struct {
	Id   int    `json:"id" form:"id"` // 设置tag，在不同的格式下对应不同的键名
	Name string `form:"username"`
}

func FormController(ctx *gin.Context) {
	name := ctx.PostForm("username")
	id := ctx.PostForm("id")
	ctx.String(200, "用户名:%s,id为:%s", name, id) // %s表示字符串占位符，只能给字符串使用 %T表示Go语法占位符
	// context.DefaultPostForm("username", "小白") 如果没有获取form提交的参数，则使用默认值
	// context.PostFormArray("name") // 如果提交的数据有多个相同的name，则获取数组

	// 可以使用ShouldBind自动提取form表单，JSON或QueryString中的数据，将数据存储在结构体中
	var u User
	err := ctx.ShouldBind(&u) // 将u的指针传入，如果获取数据失败会返回一个error，如果获取数据成功则返回nil，表示错误不存在，此时数据就会存入结构体u中
	if err != nil {
		fmt.Println(err)
	}
	// 1  ShouldBind
	// 2  ShouldBindWith  在第二个参数中可以指定要绑定的类型(binding.Form只绑定form，binding.JSON只绑定JSON数据)
	// 3  ShouldBindQuery  只绑定查询参数而忽略post参数
	// 4  Mustbind和ShouldBind区别:
	// 当绑定发生错误，也就是获取数据失败时，Mustbind会自动处理错误，返回400状态码
	// ShouldBind: 只会返回错误，由开发者自行处理错误和请求
	// 5  Bind，BindJSON，BindXML，BindQuery，BindYAML，这几个属于MustBind的具体调用，就是少了Should的就属于Must bind
	// -------------------------------------------------
	// 返回JSON字符串给前端
	data := map[string]interface{}{
		"code": 200,
		"msg":  "成功",
		"data": u,
	}
	ctx.JSON(http.StatusOK, data) // 将数据转换为JSON格式

	// 也可以直接使用gin.H方法，参数格式为 map[string]interface{}{}，即可以使用任意类型的值
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"msg":  "提交成功",
	// })

}

func QueryController(ctx *gin.Context) {
	username := ctx.Query("username") // 获取url中的查询参数 http://127.0.0.1:8080/user/?name=xxx
	fmt.Println("username:", username)
	ctx.String(200, username)
	if username == "" {
		newName := ctx.DefaultQuery("username", "默认值") // DefaultQuery 可以设置一个默认值，当没有获取到查询参数时，默认使用这个参数
		ctx.String(200, newName)
		// 如果参数类型是数组: user?name=1,2,3,4,5 ，使用contxt.QueryArray获取
		// 如果参数类型是map: user?name[1]=hallen1&name[2]=hallen2 ，使用contxt.QueryMap获取
	}

}

func UserController(ctx *gin.Context) {
	ctx.String(200, "user")
}
