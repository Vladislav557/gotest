package service

import (
	"gotest/internal/model"
	"gotest/internal/resources/postgres"
	"log"
)

type UserService struct{}

func (us *UserService) Remove(id int) {
	_, err := postgres.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func (us *UserService) Update(u *model.User, userID int) {
	_, err := postgres.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, userID)
	if err != nil {
		log.Fatal(err)
	}
}

func (us *UserService) Create(u *model.User) {
	row := postgres.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email)
	if row.Err() != nil {
		log.Fatal(row.Err())
	}
}

func (us *UserService) GetUserByID(id int) model.User {
	u := model.User{}
	row := postgres.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
	if row.Err() != nil {
		log.Fatal(row.Err())
	}
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return model.User{}
	}
	return u
}

func (us *UserService) GetUsers() []model.User {
	rows, err := postgres.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return users
}
