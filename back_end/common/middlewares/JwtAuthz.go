package middlewares

import (
	"alto_server/common/db"
	"alto_server/common/pkg/e"
	"alto_server/common/utils"
	"fmt"
	"time"

	models "alto_server/common/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func withinLimit(s int64, l int64) bool {
	e := time.Now().Unix()
	return s-e < l && s-e > 0
}

// JWT Auth的中间件，校验token的合法性
func JWTAuthz() gin.HandlerFunc {
	userInfo := models.NewUserInfo()

	return func(c *gin.Context) {
		db := db.DbOriginal
		token := utils.GetToken(c)

		if token != "" {
			claims, err := utils.ParseToken(token)
			if err != nil {
				// fmt.Println("JWTAuthz   ^_^  ParseToken: ", err.Error())
				if err == e.ErrTokenTimeOut {
					utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, e.GetMessage(e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT))
				} else {
					utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_AUTH, e.GetMessage(e.ERROR_AUTH))
				}
			} else {
				//判断token中的用户是否数据库中的用户，如果是，保存用户信息
				rows, err := db.Query("SELECT * FROM user")
				utils.PanickErr(err)
				hasAccount := false
				// fmt.Println("JWTAuthz   ^_^  claims.Username: ", claims.Username)
				// logger.SystemLogger.Error("JWTAuthz   ^_^  claims.Username: ", claims.Username)
				for rows.Next() {
					err = rows.Scan(&userInfo.Account, &userInfo.Password, &userInfo.Role, &userInfo.Department)
					utils.PanickErr(err)
					if userInfo.Account == claims.Username && userInfo.Role == claims.Role {
						hasAccount = true
						c.Set("username", userInfo.Account)
						c.Set("role", userInfo.Role)
						c.Set("password", userInfo.Password)
						//c.Set("department", userInfo.Department)
					}
				}
				//fmt.Printf("JWTAuthz   ^_^ hasAccount: \n %+v \n", hasAccount)

				if !hasAccount {
					utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_AUTH, e.GetMessage(e.ERROR_AUTH))
				}
				if withinLimit(claims.RegisteredClaims.ExpiresAt.Unix(), 1800) {
					// if withinLimit(claims.RegisteredClaims.ExpiresAt.Unix(), 7200) { // keep for testing
					fmt.Printf("renew the token  \n")
					newToken, err := utils.GenToken(userInfo.Account, userInfo.Role) //renew the token when deadline nearby
					if err != nil {
						utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_AUTH_CHECK_TOKEN_FAIL, e.GetMessage(e.ERROR_AUTH_CHECK_TOKEN_FAIL))
					} else {
						c.Header("newtoken", newToken)
						c.Request.Header.Set("Authorization", newToken)
						c.Next()
					}
				}
			}
		} else {
			fmt.Printf("Check TOken   ^_^  please pass the token !\n")
			utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_AUTH_TOKEN, e.GetMessage(e.ERROR_AUTH_TOKEN))
		}
	}
}
