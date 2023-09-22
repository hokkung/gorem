package example

import "github.com/hokkung/gorem/repository/gorm"

type User struct {
	ID uint
}

type UserRepository struct {
	gorm.GormBaseRepository[User, int64]
}
