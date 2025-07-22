package repo

import (
	"fmt"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetInfoUser(id int, name string) string {
	return fmt.Sprintf("name: %s, id: %d", name, id)
}
