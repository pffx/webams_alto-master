package auth

import (
	"alto_server/common/feature"
	"alto_server/common/pkg/e"
	"alto_server/common/utils"

	//"database/sql"

	"fmt"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateTokenHandler godoc
// @BasePath /nokia-alto
// @Summary update the token
// @schemes http https
// @Description update the token and getback the new token by header
// @Tags Auth
// @Accept json
// @Produce json
// @Param token header string true "登录信息"
// @name                        token
// @Success 200 {string} successful!
// @Router /auth/token [get]
func UpdateTokenHandler(c *gin.Context) {
	fmt.Printf("update the token")
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success"})
}

// LoginHandler godoc
// @BasePath /nokia-alto
// @Summary login action
// @schemes http https
// @Description login
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param        account   formData      string  true  "Account"
// @Param        password   formData      string  true  "Password"
// @Success      200  {json}   {"message":"login success","status":200,"timestamp":"xxxx","token":"xxxxx","userInfo":{"username":"admin","password":"****","role":"admin","department":"A"}}
// @Failure      20001  {string}  {Token鉴权失败}
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var login Login
	if err := c.ShouldBind(&login); err != nil {
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, e.GetMessage(e.INVALID_PARAMS))
	}
	// fmt.Printf("login info: %#v \n", login)

	user, _, result := login.validator()
	if result {
		token, err := utils.GenToken(user.Username, user.Role)
		if err != nil {
			utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_AUTH_CHECK_TOKEN_FAIL, e.GetMessage(e.ERROR_AUTH_CHECK_TOKEN_FAIL))
		}

		var featurelist = feature.GetFeatureList()
		utils.RES(c, e.SUCCESS, gin.H{
			"token":    token,
			"userInfo": user,
			"featurelist": featurelist,
			"status":   e.SUCCESS,
			"message":  "You are logged in.",
		})
	} else {
		//c.JSON(http.StatusOK, gin.H{"status": e.ERROR_AUTH_CHECK_TOKEN_FAIL, "message": "The username or password are incorrect"})
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_AUTH, e.GetMessage(e.ERROR_AUTH))
	}

}

// LoginHandler godoc
// @BasePath /nokia-alto
// @Summary logout action
// @schemes http https
// @Description logout
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {string} successful!
// @Router /auth/logout [post]
func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS, "message": "You are logged out"})
}
