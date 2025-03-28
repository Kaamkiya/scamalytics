package db

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"-"`
	Joined       time.Time `json:"joined"`
	SID          string    `json:"-"`
	Data         []Data    `json:"data"`
}

func (u User) CheckPassword(attempt string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(attempt)) == nil
}

func GetUserByID(id string) (User, error) {
	var u User
	var joined string

	err := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		RunWith(db).
		QueryRow().
		Scan(&u.ID, &u.Name, &u.PasswordHash, &joined, &u.SID)

	if err != nil {
		return u, err
	}

	u.Joined, err = time.Parse(time.UnixDate, joined)

	return u, err
}

func GetUserBySID(sid string) (User, error) {
	var u User
	var joined string

	err := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"sid": sid}).
		RunWith(db).
		QueryRow().
		Scan(&u.ID, &u.Name, &u.PasswordHash, &joined, &u.SID)

	if err != nil {
		return u, err
	}

	u.Joined, err = time.Parse(time.UnixDate, joined)

	return u, err
}

func GetUserByName(name string) (User, error) {
	var u User
	var joined string

	err := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"name": name}).
		RunWith(db).
		QueryRow().
		Scan(&u.ID, &u.Name, &u.PasswordHash, &joined, &u.SID)

	if err != nil {
		return u, err
	}

	u.Joined, err = time.Parse(time.UnixDate, joined)

	return u, err
}

func AddUser(id, name, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = sq.
		Insert("users").
		Columns("id", "name", "passwordhash", "joined", "sid").
		Values(id, name, string(passwordHash), time.Now().Format(time.UnixDate), "").
		RunWith(db).
		Query()

	return err
}

func SetUserSID(id, sid string) error {
	_, err := sq.
		Update("users").
		Set("sid", sid).
		Where("id = ?", id).
		RunWith(db).
		Query()

	return err
}
