package dal

import (
	"github.com/oldmoonni/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/oldmoonni/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
