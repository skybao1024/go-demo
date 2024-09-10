package db

import (
	"gorm.io/gorm"
)

// Database 定义了一个更简洁的数据库操作接口
type Database interface {
	Create(value interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Update(columnMap interface{}, value ...interface{}) *gorm.DB
	Delete(value interface{}, deleteOptions ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}
