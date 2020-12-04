package config

import (
	"disbursement-service/model"
	"fmt"

	"github.com/spf13/viper"
)

var configJSONFileName = "./config.json"

func init() {
	viper.SetConfigFile(configJSONFileName)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
}

// NewConfig ...
func NewConfig() (defConfig *model.Config, err error) {
	defConfig = &model.Config{}
	appEnv := viper.GetString(`APP_ENV`)
	appPort := viper.GetString(`APP_PORT`)
	debug := viper.GetBool(`DEBUG`)

	flipHost := viper.GetString(`FLIP_HOST`)
	flipAuthorization := viper.GetString(`FLIP_AUTHORIZATION`)

	dbHost := viper.GetString(`DB_HOST`)
	dbPort := viper.GetInt(`DB_PORT`)
	dbUser := viper.GetString(`DB_USER`)
	dbPass := viper.GetString(`DB_PASS`)
	dbName := viper.GetString(`DB_NAME`)
	dbMaxOpenConn := viper.GetInt(`DB_MAX_OPEN_CONN`)
	dbMaxIdleConn := viper.GetInt(`DB_MAX_IDLE_CONN`)

	if appEnv == "" || appPort == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section APP on %s", configJSONFileName)
		return
	}

	defConfig.AppEnv = appEnv
	defConfig.AppPort = appPort
	defConfig.Debug = debug

	if flipHost == "" || flipAuthorization == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section FLIP on %s", configJSONFileName)
		return
	}

	flipConfig := &model.Flip{
		Host:          flipHost,
		Authorization: flipAuthorization,
	}

	defConfig.Flip = flipConfig

	if dbHost == "" || dbPort == 0 || dbUser == "" || dbPass == "" || dbName == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section DB on %s", configJSONFileName)
		return
	}

	dbConfig := &model.Postgres{
		Host:         dbHost,
		Port:         dbPort,
		User:         dbUser,
		Pass:         dbPass,
		Name:         dbName,
		MaxIdleCount: dbMaxIdleConn,
		MaxOpenConn:  dbMaxOpenConn,
	}

	defConfig.DB = dbConfig

	return defConfig, nil

}
