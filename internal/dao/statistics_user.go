package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var StatisticsUser = NewStatisticsUserDao()

type StatisticsUserDao struct {
	*MongoDB[entity.StatisticsUser]
}

func NewStatisticsUserDao(database ...string) *StatisticsUserDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &StatisticsUserDao{
		MongoDB: NewMongoDB[entity.StatisticsUser](database[0], STATISTICS_USER),
	}
}
