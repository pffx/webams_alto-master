package user

import (
	"alto_server/common/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"account name"`
}

// registerHandler godoc
// @BasePath /nokia-alto
// @Summary new user account register
// @Schemes
// @Description regist a new user account
// @Tags User
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {array}   Account
// @Failure      400  {string}  httputil.HTTPError
// @Failure      404  {string}  httputil.HTTPError
// @Router /user/register [post]
func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS, "message": "Account created!"})
}

// roleHandler godoc
// @BasePath /nokia-alto
// @Summary modify the role for an exist user account
// @Schemes
// @Description modify the role for an exist user account
// @Tags User
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Account ID"
// @Success 200 {string} Your role is modified successfully!
// @Router/user/role [post]
func roleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS, "message": "Your role is modified successfully!"})
}
