package auth

import (
	"database/sql"
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
	"math/rand"
	"strings"
	"time"
)

type DBUser struct {
	Id            int                `json:"id" db:"id"`
	Username      string             `json:"username" db:"username"`
	Password      string             `json:"password" db:"password"`
	Scope         int                `json:"scope" db:"scope"`
	RememberToken sql.NullString     `json:"remember_token" db:"remember_token"`
	CreatedAt     database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt     database.TimeStamp `json:"updatedAt" db:"updated_at"`
}

type FormAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (db *DBUser) ToUser(scope Scope) User {
	return User{
		Id:            db.Id,
		Username:      db.Username,
		Password:      db.Password,
		Scope:         scope,
		RememberToken: db.RememberToken.String,
		CreatedAt:     db.CreatedAt,
		UpdatedAt:     db.UpdatedAt,
	}
}

type User struct {
	Id            int                `json:"id" db:"id"`
	Username      string             `json:"username" db:"username"`
	Password      string             `json:"password" db:"password"`
	Scope         Scope              `json:"scope"`
	RememberToken string             `json:"rememberToken" db:"remember_token"`
	CreatedAt     database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt     database.TimeStamp `json:"updatedAt" db:"updated_at"`
}

func (u *User) ConvertToAuthJSON() AuthJSON {
	return AuthJSON{
		Id:        u.Id,
		Username:  u.Username,
		Scope:     u.Scope,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) CreateAuthToken() string {
	token := GenerateRandomToken()
	return token
}

type Permission struct {
	Id         int                `json:"id" db:"id"`
	Permission string             `json:"permission" db:"permission"`
	CreatedAt  database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt  database.TimeStamp `json:"updatedAt" db:"updated_at"`
}

type PermissionScope struct {
	Id         int                `json:"id" db:"id"`
	Scope      int                `json:"scope" db:"scope"`
	Permission int                `json:"permission" db:"permission"`
	CreatedAt  database.TimeStamp `json:"createdAt" db:"created_at"`
}

type Scope struct {
	Id          int                `json:"id" db:"id"`
	Scope       string             `json:"scope" db:"scope"`
	Permissions []Permission       `json:"permissions"`
	CreatedAt   database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt   database.TimeStamp `json:"updatedAt" db:"updated_at"`
}

func (scope *Scope) hasPermission(permission string) bool {
	for _, p := range scope.Permissions {
		if strings.ToLower(p.Permission) == strings.ToLower(permission) {
			return true
		}
	}
	return false
}

type DBScope struct {
	Id        int                `json:"id" db:"id"`
	Scope     string             `json:"scope" db:"scope"`
	CreatedAt database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt database.TimeStamp `json:"updatedAt" db:"updated_at"`
}

type AuthJSON struct {
	Id        int                `json:"id"`
	Username  string             `json:"username"`
	Scope     Scope              `json:"scope"`
	CreatedAt database.TimeStamp `json:"created_at"`
	UpdatedAt database.TimeStamp `json:"updated_at"`
}

// GenerateRandomToken generates a unique random token 100% of the time
func GenerateRandomToken() string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 100)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
