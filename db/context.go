package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var UserDbContext *gorm.DB

func InitModelDbContext() {
	UserDbContext = DB.Table("users")
}
