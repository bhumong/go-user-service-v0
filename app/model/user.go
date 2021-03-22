package model

import (
	"database/sql"
	"time"

	"github.com/bhumong/go-user-service-v0/app/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    string         `json: "user_id"`
	Email     string         `json: "email"`
	Address   sql.NullString `json: "address"`
	Password  string         `json: "password"`
	CreatedAt mysql.NullTime `json: "created_at"`
	UpdatedAt mysql.NullTime `json: "updated_at"`
	DeletedAt mysql.NullTime `json: "deleted_at"`
}

type Users struct {
	Users []User `json: "products"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UserClaims struct {
	UserId string `json:"userid"`
	jwt.StandardClaims
}

func (user *UserRequest) EncryptPassword() {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(newPassword)
}

func (user User) VerifyPassword(requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestPassword))
	return err
}

func (user User) CreateToken() (string, error) {
	claims := UserClaims{
		UserId: user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "my site",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config("JWT_SECRET")))
}
