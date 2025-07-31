package main

// type Config struct {
// 	Server struct {
// 		Port int `mapstructure:"port"`
// 	} `mapstructure:"server"`
// 	Databases []struct {
// 		Host     string `mapstructure:"host"`
// 		DbName   string `mapstructure:"dbname"`
// 		User     string `mapstructure:"user"`
// 		Password string `mapstructure:"password"`
// 	} `mapstructure:"databases"`
// }

func main() {
	// viper := viper.New()
	// viper.AddConfigPath("./config")
	// viper.SetConfigName("local")
	// viper.SetConfigType("yaml")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Server Port:", viper.GetInt("server.port"))
	// fmt.Println("security key:", viper.GetString("security.jwt.key"))

	// var config Config
	// err = viper.Unmarshal(&config)
	// if err != nil {
	// 	panic(fmt.Errorf("unable to decode into struct, %v", err))
	// }
	// fmt.Println("Database Host:", config.Databases[0].DbName)
}
