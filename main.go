package main

import "1-gin_project/routers"

func main() {
	router := routers.InitRouter()
	router.Static("/static", "./static") // 注册静态资源路径，导入JS,CSS时，路径以/static开头
	router.Run(":8080")                  // 第一种启动服务器的格式
	// http.ListenAndServe(":8080", router) // 第二种格式

	// s := &http.Server {
	//     Addr:             ":8080",
	//     Handler:          router,
	//     ReadTimeout:      10 * time.Second,
	//     WriteTimeout:     10 * time.Second,
	//     MaxHeaderBytes:   1 << 20,
	// }
	// s.ListenAndServe()   第三种格式，可以自定义http配置
}
