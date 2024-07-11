package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var Midjourney = NewMidjourneyDao()

type MidjourneyDao struct {
	*MongoDB[entity.Midjourney]
}

func NewMidjourneyDao(database ...string) *MidjourneyDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &MidjourneyDao{
		MongoDB: NewMongoDB[entity.Midjourney](database[0], do.MIDJOURNEY_COLLECTION),
	}
}
