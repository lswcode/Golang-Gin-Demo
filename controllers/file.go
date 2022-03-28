package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func FileController(ctx *gin.Context) { // 上传单个文件时
	file, err := ctx.FormFile("file") // 第一个参数是表单组件的name属性值，需要和前端沟通确定
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("获取到的用户上传文件名称:", file.Filename)

	timeStampInt := time.Now().Unix()                   // 获取当前时间戳，是秒，int类型，为了防止文件重名一直被覆盖，使用时间戳生成不同的文件名
	fmt.Println(time.Now())                             // 2022-01-26 17:22:43.5631045 +0800 CST m=+25.359952301
	fmt.Println(timeStampInt)                           // 1643188963
	timeStampStr := strconv.FormatInt(timeStampInt, 10) // 将int类型转为string类型，这样就可以直接进行字符串拼接，第二个参数是进制

	filePath := "upload/" + timeStampStr + "--" + file.Filename // 设置保存文件的完整路径，要加上文件名
	ctx.SaveUploadedFile(file, filePath)                        // 保存文件

	ctx.String(http.StatusOK, "文件上传成功") // 返回给前端的内容
}

func MoreFileController(ctx *gin.Context) { // 如果上传文件有多个
	form, _ := ctx.MultipartForm() // 获取整个form
	files := form.File["file"]     // 获取form中的所有文件，返回的是一个map

	for _, file := range files { // 循环执行以下操作
		timeStampInt := time.Now().Unix()
		fmt.Println(time.Now())
		fmt.Println(timeStampInt)
		timeStampStr := strconv.FormatInt(timeStampInt, 10)

		filePath := "upload/" + timeStampStr + file.Filename
		ctx.SaveUploadedFile(file, filePath)
	}

	ctx.String(http.StatusOK, "多文件上传成功")
}

// ---路由重定向--------------------------------------------------------
func RedirectController(ctx *gin.Context) {
	fmt.Println("路由重定向")
	ctx.Redirect(http.StatusFound, "/")                      // StatusFound就是302，第二个参数就是新的路由路径
	ctx.Redirect(http.StatusFound, "https://www.baidu.com/") // 可以跳转到任意网址，优先匹配上面先写的重定向
}
