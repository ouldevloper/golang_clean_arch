package infrastructure

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() *gorm.DB {
	dsn := "go_clean:user@tcp(127.0.0.1:3306)/go_clean?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could Not Connect to db")
	}

	sh := new(SqlHandler)
	sh.db = db
	return db
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(&obj)
}

func (handler *SqlHandler) Find(obj []interface{}) {
	handler.db.Create(&obj)
}

func (handler *SqlHandler) Delete(obj interface{}, id string) {
	handler.db.Delete(&obj, id)
}
