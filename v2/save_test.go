package gorem_test

import (
	"github.com/google/uuid"
	"github.com/hokkung/gorem/v2"
	"github.com/hokkung/gorem/v2/tests"
	"github.com/stretchr/testify/assert"
)

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
