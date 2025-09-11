package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var SysConfig = NewSysConfigDao()

type SysConfigDao struct {
	*MongoDB[entity.SysConfig]
}

func NewSysConfigDao(database ...string) *SysConfigDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SysConfigDao{
		MongoDB: NewMongoDB[entity.SysConfig](database[0], SYS_CONFIG),
	}
}
