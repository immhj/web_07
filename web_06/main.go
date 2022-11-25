package main

import (
	dao "web_06/dao/mysql"
	dao2 "web_06/dao/redis"
	"web_06/router"
)

func main() {
	err := dao.Initmysql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	err = dao2.Initredis()
	if err != nil {
		panic(err)
	}
	defer dao2.Rdb.Close()
	router.Initrouter()
}
