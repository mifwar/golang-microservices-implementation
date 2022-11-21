package database

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	u "users/pkg/utils"
)

var db *gorm.DB

type config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
}

func getConfig(c *config) {
	conf_dir, _ := filepath.Abs("pkg/database/config/dbconfig.yml")
	conf, err := ioutil.ReadFile(conf_dir)

	u.LogError(err, "read file")

	err = yaml.Unmarshal(conf, c)

	u.LogError(err, "unmarshal yaml")
}

func Connect() {
	dbConfig := config{}
	additionalConfig := " sslmode=disable TimeZone=Asia/Shanghai"

	getConfig(&dbConfig)
	gormCommand := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%d%s", dbConfig.Username, dbConfig.Password, dbConfig.DBname, dbConfig.Port, additionalConfig)

	d, err := gorm.Open(postgres.Open(gormCommand), &gorm.Config{})
	u.LogError(err, "gorm")

	db = d
}

func GetDB() *gorm.DB {
	Connect()
	return db
}
