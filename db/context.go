package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var UserDbContext *gorm.DB
var CostDbContext *gorm.DB
var PaymentDbContext *gorm.DB

func InitModelDbContext() {
	UserDbContext = DB.Table("users")
	CostDbContext = DB.Table("costs")
	PaymentDbContext = DB.Table("payments")
}
