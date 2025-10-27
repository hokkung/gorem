package gorem_test

import (
	"context"
	"testing"

	"github.com/hokkung/gorem/v2/tests"
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
