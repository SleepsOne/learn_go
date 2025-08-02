package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	//test to see config loaded
	fmt.Printf("Mysql config loaded %s, %s \n ", global.Config.MySQL.Username, global.Config.MySQL.Password)
	InitLogger()
	global.Logger.Log(zap.DebugLevel, "config log okee", zap.String("ok", "success"), zap.Int("age", 15))
	InitMySQL()
	InitRedis()

	r := InitRouter()
	fmt.Println("Port server: ", global.Config.ServerConfig.Port)
	r.Run(fmt.Sprintf(":%v", global.Config.ServerConfig.Port))
}
