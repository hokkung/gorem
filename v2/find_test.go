package gorem_test

import (
	"github.com/google/uuid"
	"github.com/hokkung/gorem/v2"
	"github.com/hokkung/gorem/v2/tests"
	"github.com/stretchr/testify/assert"
)

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

func (ts *UIDModelRepositorySuite) TestFindByFilter() {
	ents, err := ts.repo.FindByFilter(ts.ctx, map[string]any{
		"name": []string{"test"},
	})
	assert.NoError(ts.T(), err)
	assert.Empty(ts.T(), ents)

	model := tests.TestUIDModel{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		Name: "test",
	}
	err = ts.repo.Create(ts.ctx, &model)
	assert.NoError(ts.T(), err)

	ents, err = ts.repo.FindByFilter(ts.ctx, map[string]any{
		"name": []string{"test"},
	})
	assert.NoError(ts.T(), err)
	assert.Len(ts.T(), ents, 1)
	assert.Equal(ts.T(), model.ID, ents[0].ID)
	assert.Equal(ts.T(), model.Name, ents[0].Name)
}
