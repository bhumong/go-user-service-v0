package repository

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bhumong/go-user-service-v0/app/database"
	"github.com/bhumong/go-user-service-v0/app/model"
	"github.com/google/uuid"
)

func FindUserById(id string) (model.User, error) {
	row, err := database.DB.Query("SELECT * FROM users WHERE user_id = ? and deleted_at is null LIMIT 1", id)
	user := model.User{}
	if err != nil {
		return user, err
	}
	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&user.UserId, &user.Email, &user.Address, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err {
		case nil:
			return user, nil
		default:
			return user, err
		}
	}
	return user, nil
}

func FindUserByEmail(email string) (model.User, error) {
	row, err := database.DB.Query("SELECT * FROM users WHERE email = ? and deleted_at is null LIMIT 1", email)
	user := model.User{}
	if err != nil {
		return user, err
	}
	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&user.UserId, &email, &user.Address, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err {
		case nil:
			return user, nil
		default:
			return user, err
		}
	}
	return user, nil
}

func FindAll() (model.Users, error) {
	rows, err := database.DB.Query("SELECT * FROM users where deleted_at is null")
	users := model.Users{}
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		switch err := rows.Scan(&user.UserId, &user.Email, &user.Address, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err {
		case nil:
			users.Users = append(users.Users, user)
		default:
			return users, err
		}
	}
	return users, err
}

func DeleteUserById(id string) error {
	row, err := database.DB.Query("UPDATE users SET deleted_at = now() where user_id = ? and deleted_at is null", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer row.Close()
	return nil
}

func UpdateUserById(id string, pu model.UserRequest) error {
	pu.EncryptPassword()
	fields, values := getFieldAndValueFromRequest(pu, " = ?")
	set := strings.Join(fields, " , ")
	values = append(values, id)
	_, err := database.DB.Query("UPDATE users SET "+set+" where user_id = ?", values...)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(pu model.UserRequest) error {
	pu.EncryptPassword()
	id := uuid.NewString()
	fields, values := getFieldAndValueFromRequest(pu, "")

	fields = append(fields, "user_id")
	fields = append(fields, "created_at")
	fields = append(fields, "updated_at")
	values = append(values, id)
	values = append(values, time.Now())
	values = append(values, time.Now())
	cols := strings.Join(fields, " , ")
	_, err := database.DB.Query("INSERT INTO users ("+cols+") VALUES (?,?,?,?,?,?)", values...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getFieldDbByRequest(fieldRequest string) string {
	switch fieldRequest {
	case "Email":
		return "email"
	case "Password":
		return "password"
	case "Address":
		return "address"
	default:
		return ""
	}
}

func getFieldAndValueFromRequest(pu model.UserRequest, seperator string) ([]string, []interface{}) {
	fields := []string{}
	values := []interface{}{}

	v := reflect.ValueOf(pu)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface().(string)
		if value != "" {
			field := getFieldDbByRequest(v.Type().Field(i).Name)
			if field != "" {
				fields = append(fields, field+seperator)
				values = append(values, value)
			}
		}
	}
	return fields, values
}
