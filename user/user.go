package user

import (
	"errors"
	"fmt"
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//User holds data for a single user
type User struct {
	ID bson.ObjectId `json:"id" storm:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

const (
	dbPath = "users.db"
)

// Errors
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

// All retrieves all users from the database
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// One returns a single user record from the database
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id , u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Delete removes a given record from the database
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id , u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

// Save updates or creates a given record in the database
func (u *User) Save() error {
	start := time.Now()
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return  err
	}
	defer db.Close()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	return db.Save(u)


}

// validate makes sure that the record contains valid data
func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}