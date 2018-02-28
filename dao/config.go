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

	prepare()
}

func prepare() {
	var tables []interface{}

	// create tables if not exists
	tables = append(tables, &Transaction{})
	tables = append(tables, &Transfer{})

	for _, t := range tables {
		if ok := db.HasTable(t); !ok {
			if err := db.CreateTable(t).Error; err != nil {
				log.Fatalf("create mysql table error:%s", err.Error())
			}
		}
	}

	// auto migrate to keep schema update to date
	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON'T change existing column's type or delete unused columns to protect your data
	db.AutoMigrate(tables...)
}
