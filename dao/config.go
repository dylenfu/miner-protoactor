package dao

import "github.com/jinzhu/gorm"

var db *gorm.DB

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return options.TablePrefix + defaultTableName
	}

	url := options.User + ":" + options.Password + "@tcp(" + options.Hostname + ":" + options.Port + ")/" + options.DbName + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatalf("mysql connection error:%s", err.Error())
	}

	db.LogMode(options.Debug)

	impl.db = db

	return impl
}
