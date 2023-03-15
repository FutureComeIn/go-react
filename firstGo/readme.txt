安装gdb：http://c.biancheng.net/view/8296.html （gdb在 bin、share里面）

命令：
	go env -w GOPROXY=https://goproxy.cn --下载不了依赖，设置代理
	gf init demo -u  -- 创建goframe框架项目，-u指定是否更新项目中使用的goframe框架为最新版本
	go get -u github.com/gogf/gf --创建后拉取依赖，也可跑go.mod文件
	go mod tidy ：添加需要用到但go.mod中查不到的模块
	运行：cd demo && gf run main.go（也可直接运行main.go）
	
语法：https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md
	1）&：是对变量取地址，如&a
	   *：对指针取值，如*&a，就是a变量所在地址的值，也就是a的值
	所以 ：*&可以抵消掉，a=*&a