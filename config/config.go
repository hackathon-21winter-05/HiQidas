package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Client_ID        string `json:"client_id,omitempty"`        // 本番環境向けのクライアントID (default: "")
	Client_ID_Dev    string `json:"client_id_dev,omitempty"`    // ローカル開発環境向けのクライアントID (default: "")
	MariaDB_Hostname string `json:"mariadb_hostname,omitempty"` // DB のホスト (default: "mariadb")
	MariaDB_Database string `json:"mariadb_database,omitempty"` // DB の DB 名 (default: "SchMes")
	MariaDB_Username string `json:"mariadb_username,omitempty"` // DB のユーザー名 (default: "root")
	MariaDB_Password string `json:"mariadb_password,omitempty"` // DB のパスワード (default: "password")
}

func GetConfig() (*Config, error) {
	viper.SetDefault("Client_ID", "")
	viper.SetDefault("Client_ID_Dev", "")
	viper.SetDefault("MariaDB_Hostname", "mariadb")
	viper.SetDefault("MariaDB_Database", "SchMes")
	viper.SetDefault("MariaDB_Username", "root")
	viper.SetDefault("MariaDB_Password", "password")

	viper.AutomaticEnv()

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("Unable to find config.json, default settings or environmental variables are to be used.")
		} else {
			return nil, fmt.Errorf("Error: failed to load config.json - %s ", err)
		}
	}

	var c *Config

	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to parse configs - %s ", err)
	}

	return c, nil
}
