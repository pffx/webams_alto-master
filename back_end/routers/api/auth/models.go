package auth

import (
	DbModule "alto_server/common/db"
	"alto_server/common/pkg/e"
	. "alto_server/common/utils"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// func init() {
// 	db, err := sql.Open("sqlite3", "./db/alto.db")
// 	PanickErr(err)
// }

type UserInfo struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Department string `json:"department"`
}

// Login is login struct
type Login struct {
	Username string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Users 是用户列表
type users []UserInfo

func checkDBNotNull() error {
	if db == nil {
		return e.ErrDBInvalid
	}
	return nil
}

// Validator .
func (login *Login) validator() (*UserInfo, string, bool) {
	user := &UserInfo{}
	db := DbModule.DbOriginal
	//err := checkDBNotNull()
	rows, err := db.Query("SELECT * FROM user")
	var msg string
	PanickErr(err)
	for rows.Next() {
		var userName string
		var password string
		var role string
		var department string
		err = rows.Scan(&userName, &password, &role, &department)
		PanickErr(err)
		if userName == login.Username {
			if password == login.Password {
				user.Username = userName
				user.Password = "****"
				user.Role = role
				user.Department = department
				msg = "success"
				return user, msg, true
			} else {
				msg = "wrong password"
				return nil, msg, false
			}
		}
	}
	msg = "account not exist"
	return nil, msg, false
}
