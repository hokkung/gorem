package gorem_test

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/hokkung/gorem/v2"
	"github.com/hokkung/gorem/v2/tests"
	"github.com/stretchr/testify/assert"
)

func (ts *UIDModelRepositorySuite) TestDeleteByID() {
	model := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name: "test",
	}
	err := ts.repo.Create(ts.ctx, &model)
	assert.NoError(ts.T(), err)

	err = ts.repo.DeleteByID(ts.ctx, model.ID)
	assert.NoError(ts.T(), err)
}

func (ts *UIDModelRepositorySuite) TestDelete() {
	model := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name: "test",
	}
	err := ts.repo.Create(ts.ctx, &model)
	assert.NoError(ts.T(), err)

	err = ts.repo.Delete(ts.ctx, &model)
	assert.NoError(ts.T(), err)

	ents, err := ts.repo.FindAll(ts.ctx)
	assert.NoError(ts.T(), err)
	assert.Empty(ts.T(), ents)
}

func (ts *UIDModelRepositorySuite) TestDeleteAll() {
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

	err := ts.repo.Creates(ts.ctx, []*tests.TestUIDModel{&model, &model2})
	assert.NoError(ts.T(), err)

	err = ts.repo.DeleteAll(ts.ctx)
	assert.NoError(ts.T(), err)

	ents, err := ts.repo.FindAll(ts.ctx)
	assert.NoError(ts.T(), err)
	assert.Empty(ts.T(), ents)
}
