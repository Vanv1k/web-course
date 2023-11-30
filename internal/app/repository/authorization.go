package repository

import (
	"github.com/Vanv1k/web-course/internal/app/ds"
)

func (r *Repository) Login() (*ds.User, error) {
	return nil, nil
}

func (r *Repository) Register(user *ds.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) Logout() (*ds.User, error) {
	return nil, nil
}

func (r *Repository) GetUserByLogin(login string) (*ds.User, error) {
	user := &ds.User{
		Login: login,
	}

	err := r.db.Where(user).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
