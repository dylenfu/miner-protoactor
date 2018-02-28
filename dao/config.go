package dao

import (
	"github.com/dylenfu/miner-protoactor/config"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func NewDb(options *config.MysqlOptions) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return options.TablePrefix + defaultTableName
	}

	url := options.User + ":" + options.Password + "@tcp(" + options.Hostname + ":" + options.Port + ")/" + options.DbName + "?charset=utf8&parseTime=True"
	dao, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatalf("mysql connection error:%s", err.Error())
	}

	dao.LogMode(options.Debug)

	db = dao
}
