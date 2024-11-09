package dal

import (
	"github.com/oldmoonni/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/oldmoonni/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
