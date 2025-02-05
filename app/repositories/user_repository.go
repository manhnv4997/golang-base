package repositories

import (
	"database/sql"
	"demo/app/models"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := userRepository.DB.Query("SELECT id, username FROM users")

	if err != nil {
		log.Println("Lỗi truy vấn", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	log.Println(users, "users")

	return users, nil
}
