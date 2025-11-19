package main

import "log"

func main() {
	//连接数据库
	err := connectDB()
	if err != nil {
		log.Printf("[ERROR]连接数据库失败, 退出程序")
		return
	}
}
