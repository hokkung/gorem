package tests_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hokkung/gorem/v2"
	"github.com/hokkung/gorem/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(&tests.TestUIDModel{})
	if err != nil {
		t.Fatalf("failed to auto migrate test uid model: %v", err)
	}

	return db
}

func TestUIDModelRepositorySuite(t *testing.T) {
	suite.Run(t, new(UIDModelRepositorySuite))
}

type UIDModelRepositorySuite struct {
	suite.Suite

	ctx  context.Context
	db   *gorm.DB
	repo *tests.TestUIDModelRepository
}

func (ts *UIDModelRepositorySuite) SetupSuite() {
	ts.ctx = context.Background()
	ts.db = setupTestDB(ts.T())
}

func (ts *UIDModelRepositorySuite) SetupTest() {
	ts.repo = tests.NewTestUIDModelRepository(ts.db)
	ts.repo.DeleteAll(ts.ctx)
}

func (ts *UIDModelRepositorySuite) TestCreate() {
	nullTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	model := tests.TestUIDModel{
		Name:     "test",
		Age:      10,
		Money:    100.01,
		IsActive: true,
		NullableName: sql.NullString{
			String: "test",
			Valid:  true,
		},
		NullableAge: sql.NullInt64{
			Int64: 10,
			Valid: true,
		},
		NullableMoney: sql.NullFloat64{
			Float64: 100,
			Valid:   true,
		},
		NullableIsActive: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		NullableTime: nullTime,
	}

	err := ts.repo.Create(ts.ctx, &model)
	assert.NoError(ts.T(), err)

	ents, err := ts.repo.FindAll(ts.ctx)
	assert.NoError(ts.T(), err)

	ent := ents[0]
	assert.Len(ts.T(), ents, 1)
	assert.Equal(ts.T(), "test", ent.Name)
	assert.Equal(ts.T(), 10, ent.Age)
	assert.Equal(ts.T(), 100.01, ent.Money)
	assert.Equal(ts.T(), true, ent.IsActive)
	assert.Equal(ts.T(), "test", ent.NullableName.String)
	assert.Equal(ts.T(), int64(10), ent.NullableAge.Int64)
	assert.Equal(ts.T(), float64(100), ent.NullableMoney.Float64)
	assert.Equal(ts.T(), true, ent.NullableIsActive.Bool)
	assert.Equal(ts.T(), nullTime.Time.UTC(), ent.NullableTime.Time.UTC())

	assert.NotZero(ts.T(), ent.CreatedAt)
	assert.NotZero(ts.T(), ent.UpdatedAt)
	assert.Zero(ts.T(), ent.DeletedAt)
}

func (ts *UIDModelRepositorySuite) TestFindAll() {
	ents, err := ts.repo.FindAll(ts.ctx)

	assert.Empty(ts.T(), ents)
	assert.NoError(ts.T(), err)

	model1 := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name:     "test",
		Age:      10,
		Money:    100.01,
		IsActive: true,
	}
	model2 := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name: "test2",
	}

	err = ts.repo.Creates(ts.ctx, []*tests.TestUIDModel{&model1, &model2})
	assert.NoError(ts.T(), err)

	ents, err = ts.repo.FindAll(ts.ctx)
	assert.NoError(ts.T(), err)
	assert.Len(ts.T(), ents, 2)
	assert.Equal(ts.T(), model1.Name, ents[0].Name)
	assert.Equal(ts.T(), model2.Name, ents[1].Name)
}

func (ts *UIDModelRepositorySuite) TestFindByKey() {
	ent, found, err := ts.repo.FindByKey(ts.ctx, "123")

	assert.False(ts.T(), found)
	assert.NoError(ts.T(), err)
	assert.Nil(ts.T(), ent)

	model := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name:     "test",
		Age:      10,
		Money:    100.01,
		IsActive: true,
	}
	err = ts.repo.Create(ts.ctx, &model)
	assert.NoError(ts.T(), err)

	ent, found, err = ts.repo.FindByKey(ts.ctx, model.ID)
	assert.True(ts.T(), found)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), model.ID, ent.ID)
	assert.Equal(ts.T(), model.Name, ent.Name)
	assert.Equal(ts.T(), model.Age, ent.Age)
	assert.Equal(ts.T(), model.Money, ent.Money)
	assert.Equal(ts.T(), model.IsActive, ent.IsActive)
}

func (ts *UIDModelRepositorySuite) TestUpdate() {
	model := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name:     "test",
		Age:      10,
		Money:    100.01,
		IsActive: true,
	}

	err := ts.repo.Create(ts.ctx, &model)
	assert.NoError(ts.T(), err)

	model.Name = "test2"
	model.Age = 20
	model.Money = 200.02
	model.IsActive = false

	err = ts.repo.Save(ts.ctx, &model)

	assert.NoError(ts.T(), err)

	ent, found, err := ts.repo.FindByKey(ts.ctx, model.ID)
	assert.True(ts.T(), found)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), model.ID, ent.ID)
	assert.Equal(ts.T(), model.Name, ent.Name)
	assert.Equal(ts.T(), model.Age, ent.Age)
	assert.Equal(ts.T(), model.Money, ent.Money)
	assert.Equal(ts.T(), model.IsActive, ent.IsActive)
}
