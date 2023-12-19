package repository

import (
	"fmt"

	"github.com/Vanv1k/web-course/internal/app/ds"
)

func (r *Repository) GetUserName(userID uint) (string, error) {
	user := &ds.User{}
	fmt.Println("adfvasfvadfvadsfvadfvadfvs")
	fmt.Println(userID)
	if userID == 0 {
		return "Не установлен", nil
	}
	err := r.db.First(user, "id = ?", userID).Error
	if err != nil {
		return "", err
	}

	return user.Name, nil
}
