package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 第一种中间件格式
func M1(ctx *gin.Context) {
	fmt.Println("全局中间件11111开始") // Next之前的代码是中间件对目标的操作
	startTime := time.Now()
	// request请求被路由接收前
	ctx.Next()
	// 如果没有其它中间件了，表示对请求的预处理已经完成，请求被对应的路由以及控制器接收处理
	timeUse := time.Since(startTime)
	fmt.Println("全局中间件11111结束--耗费时间: ", timeUse) // 此时请求已经执行完成
}

// 第二种中间件格式，return一个gin.HandlerFunc函数，使用这个方法时，调用中间件时需要加小括号，来调用返回的函数
func M2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("自定义中间件22222开始")
		// if 3 < 4 {
		// 假设这里发生token验证失败，此时就要调用Abort()，终止剩下的中间件的执行
		// 因为此时已经验证失败了，表示这个请求没通过，那就没有必要对这个请求进行其它处理了，直接在这结束中间件的调用
		// 	context.Abort() // Abort()用来终止整个中间件链条，不会停止当前的函数执行，终止的是当前请求的其它还未执行的中间件
		// }
		ctx.Next()
		fmt.Println("自定义中间件22222结束")
	}
}

func M3() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("自定义中间件33333开始")
		time.Sleep(5 * time.Second) // 暂停五秒
		ctx.Next()
		fmt.Println("自定义中间件33333结束")
	}
}
