package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	DBName     = "auth.db.name"
	DBHost     = "auth.db.ip"
	DBPort     = "auth.db.port"
	DBUser     = "auth.db.user"
	DBPassword = "auth.db.password"
	DBDriver   = "auth.db.driver"

	ServerHost = "server.host"
	ServerPort = "server.port"

	ErrorMessage401 = "error.401"
	ErrorMessage422 = "error.422"
	ErrorMessage500 = "error.500"
	ErrorMessage404 = "error.404"
	ErrorMessage409 = "error.404"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DBName, "AUTH_DB_NAME")
	_ = viper.BindEnv(DBHost, "AUTH_DB_HOST")
	_ = viper.BindEnv(DBUser, "AUTH_DB_USER")
	_ = viper.BindEnv(DBPassword, "AUTH_DB_PASSWORD")
	_ = viper.BindEnv(DBPort, "AUTH_DB_PORT")
	_ = viper.BindEnv(DBDriver, "AUTH_DB_DRIVER")

	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	// defaults
	viper.SetDefault(DBName, "database_name")
	viper.SetDefault(DBHost, "127.0.0.1")
	viper.SetDefault(DBPort, "3306")
	viper.SetDefault(DBDriver, "mysql")
	viper.SetDefault(DBUser, "root")
	viper.SetDefault(DBPassword, "myproject123")

	viper.SetDefault(ServerHost, "0.0.0.0")
	viper.SetDefault(ServerPort, "8091")

	// error message
	viper.SetDefault(ErrorMessage401, "failed to verify token")
	viper.SetDefault(ErrorMessage422, "Password pattern does not match")
	viper.SetDefault(ErrorMessage500, "Sorry! Something went wrong on making request")
	viper.SetDefault(ErrorMessage404, "Unable to find the resource")
	viper.SetDefault(ErrorMessage409, "requested resource is already Modify")
}
