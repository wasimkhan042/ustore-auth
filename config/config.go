package config

import "github.com/spf13/viper"

//keys for database configuration.
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"
	DbUser = "db.user"
	DbPass = "db.pass"

	ServerHost = "server.host"
	ServerPort = "server.port"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")
	_ = viper.BindEnv(DbUser, "DB_USER")
	_ = viper.BindEnv(DbPass, "DB_PASS")

	// env var for server
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	// defaults
	viper.SetDefault(DbName, "ustore")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
	viper.SetDefault(DbPass, "MYpassword100")
}
