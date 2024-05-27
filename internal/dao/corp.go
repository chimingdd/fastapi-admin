package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var Corp = NewCorpDao()

type CorpDao struct {
	*MongoDB[entity.Corp]
}

func NewCorpDao(database ...string) *CorpDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &CorpDao{
		MongoDB: NewMongoDB[entity.Corp](database[0], do.CORP_COLLECTION),
	}
}
