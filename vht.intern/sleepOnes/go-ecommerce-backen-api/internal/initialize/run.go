package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	//test to see config loaded
	fmt.Printf("Mysql config loaded %s, %s \n ", global.Config.MySQL.Username, global.Config.MySQL.Password)
	InitLogger()
	InitMySQL()
	InitRedis()

	r := InitRouter()
	r.Run(":8022")
}
