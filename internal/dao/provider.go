package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var Provider = NewProviderDao()

type ProviderDao struct {
	*MongoDB[entity.Provider]
}

func NewProviderDao(database ...string) *ProviderDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &ProviderDao{
		MongoDB: NewMongoDB[entity.Provider](database[0], do.PROVIDER_COLLECTION),
	}
}
