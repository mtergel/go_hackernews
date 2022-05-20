package users

import (
	"database/sql"
	"log"

	"github.com/tergelm/go_hackernews/internal/pkg/db/migrations/postgresql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type WrongUsernameOrPasswordError struct{}

func (user *User) Create() {
	query, err := db_handler.DBClient.Prepare(`
          INSERT INTO users (username, password)
          VALUES ($1, $2);
        `)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, _ := HashPassword(user.Password)

	_, err = query.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func (user *User) Authenticate() bool {
	query, err := db_handler.DBClient.Prepare(`
    SELECT password FROM users WHERE username = $1;
  `)
	if err != nil {
		log.Fatal(err)
	}

	row := query.QueryRow(user.Username)
	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

func GetUserIdByUsername(username string) (int, error) {
	query, err := db_handler.DBClient.Prepare(`
    SELECT id FROM users WHERE username = $1;
  `)
	if err != nil {
		log.Fatal(err)
	}

	row := query.QueryRow(username)
	var Id int
	err = row.Scan(&Id)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}

		return 0, err
	}

	return Id, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "Wrong username or password"
}
