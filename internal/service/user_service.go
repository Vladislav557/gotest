package service

import (
	"gotest/internal/model"
	"gotest/internal/resources/postgres"
	"log"
)

type UserService struct {
	db *postgres.Database
}

func (us *UserService) Remove(id int) {
	_, err := us.db.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func (us *UserService) Update(u *model.User, userID int) {
	_, err := us.db.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, userID)
	if err != nil {
		log.Fatal(err)
	}
}

func (us *UserService) Create(u *model.User) {
	err := us.db.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email)
	if err != nil {
		log.Fatal(err)
	}
}

func (us *UserService) GetUserByID(id int) model.User {
	u := model.User{}
	err := us.db.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func (us *UserService) GetUsers() []model.User {
	rows, err := us.db.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	users := []model.User{}
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