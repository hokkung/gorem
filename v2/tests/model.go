package tests

import (
	"database/sql"

	"github.com/hokkung/gorem/v2"
	"gorm.io/gorm"
)

type TestUIDModel struct {
	gorem.UIDModel
	Name             string
	NullableName     sql.NullString
	Age              int
	NullableAge      sql.NullInt64
	Money            float64
	NullableMoney    sql.NullFloat64
	IsActive         bool
	NullableIsActive sql.NullBool
	NullableTime     sql.NullTime
}

func (m TestUIDModel) TableName() string {
	return "test_uid_models"
}

func (m TestUIDModel) PrimaryKey() string {
	return "id"
}

type TestUIDModelRepository struct {
	*gorem.BaseRepository[TestUIDModel]
}

func NewTestUIDModelRepository(db *gorm.DB) *TestUIDModelRepository {
	return &TestUIDModelRepository{
		BaseRepository: gorem.NewBaseRepository[TestUIDModel](db),
	}
}
