package dal

import (
	"github.com/oldmoonni/gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
