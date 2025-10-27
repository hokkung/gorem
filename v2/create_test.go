package gorem_test

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/hokkung/gorem/v2"
	"github.com/hokkung/gorem/v2/tests"
	"github.com/stretchr/testify/assert"
)

func (ts *UIDModelRepositorySuite) TestCreate() {
	nullTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	model := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
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

func (ts *UIDModelRepositorySuite) TestCreates() {
	nullTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	model := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
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
	model2 := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name:     "test2",
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
	expected := []tests.TestUIDModel{model, model2}

	err := ts.repo.Creates(ts.ctx, []*tests.TestUIDModel{&model, &model2})
	assert.NoError(ts.T(), err)

	ents, err := ts.repo.FindAll(ts.ctx)
	assert.NoError(ts.T(), err)

	assert.Len(ts.T(), ents, 2)
	for i := 0; i < len(ents); i++ {
		assert.Equal(ts.T(), expected[i].Name, ents[i].Name)
		assert.Equal(ts.T(), expected[i].Age, ents[i].Age)
		assert.Equal(ts.T(), expected[i].Money, ents[i].Money)
		assert.Equal(ts.T(), expected[i].IsActive, ents[i].IsActive)
		assert.Equal(ts.T(), expected[i].NullableName.String, ents[i].NullableName.String)
		assert.Equal(ts.T(), expected[i].NullableAge.Int64, ents[i].NullableAge.Int64)
		assert.Equal(ts.T(), expected[i].NullableMoney.Float64, ents[i].NullableMoney.Float64)
		assert.Equal(ts.T(), expected[i].NullableIsActive.Bool, ents[i].NullableIsActive.Bool)
		assert.Equal(ts.T(), expected[i].NullableTime.Time.UTC(), ents[i].NullableTime.Time.UTC())

		assert.NotZero(ts.T(), ents[i].CreatedAt)
		assert.NotZero(ts.T(), ents[i].UpdatedAt)
		assert.Zero(ts.T(), ents[i].DeletedAt)
	}
}
