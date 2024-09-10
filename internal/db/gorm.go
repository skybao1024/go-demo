package db

import (
	"gorm.io/gorm"
)

// GormDB 是一个包装 gorm.DB 的结构体，实现了 Database 接口
type GormDB struct {
	*gorm.DB
}

// Create 实现 Database 接口的 Create 方法
func (d *GormDB) Create(value interface{}) *gorm.DB {
	return d.DB.Create(value)
}

// First 实现 Database 接口的 First 方法
func (d *GormDB) First(out interface{}, where ...interface{}) *gorm.DB {
	return d.DB.First(out, where...)
}

// Find 实现 Database 接口的 Find 方法
func (d *GormDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return d.DB.Find(out, where...)
}

// Update 实现 Database 接口的 Update 方法
func (d *GormDB) Update(columnMap interface{}, value ...interface{}) *gorm.DB {
	switch v := columnMap.(type) {
	case string:
		// 单个字段更新
		if len(value) == 0 {
			return d.DB
		}
		return d.DB.Update(v, value[0])
	case map[string]interface{}:
		// 多个字段更新
		return d.DB.Updates(v)
	default:
		// 如果类型不匹配，则返回原始的 *gorm.DB 指针
		return d.DB
	}
}

// Delete 实现 Database 接口的 Delete 方法
func (d *GormDB) Delete(value interface{}, deleteOptions ...interface{}) *gorm.DB {
	return d.DB.Delete(value, deleteOptions...)
}

// Where 实现 Database 接口的 Where 方法
func (d *GormDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return d.DB.Where(query, args...)
}

// NewGormDB 创建一个新的 GormDB 实例
func NewGormDB(db *gorm.DB) *GormDB {
	return &GormDB{DB: db}
}

func (d *GormDB) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
