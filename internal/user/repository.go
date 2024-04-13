package user

import (
	"database/sql"
	"log"
)

type UserRepositoryInterface interface {
	GetAll() ([]*User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetAll() ([]*User, error) {
	rows, err := repo.db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

func scanIntoUser(rows *sql.Rows) (*User, error) {

	user := User{}

	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Bio,
		&user.Avatar,
		&user.Phone,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt)

	return &user, err

}
