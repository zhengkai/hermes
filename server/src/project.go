package project

import (
	"project/build"
	"project/config"
	"project/db"
	"project/pb"
	"project/zj"

	"github.com/zhengkai/zu"
)

// Start ...
func Start() {

	build.DumpBuildInfo()

	zj.J(zu.JSON(&pb.Demo{
		ID:   43,
		Name: `rpg`,
	}))

	db.WaitConn(config.MySQL)

	select {}
}

// Prod ...
func Prod() {

	config.Prod = true

	zj.J(`prod start`)

	Start()
}
