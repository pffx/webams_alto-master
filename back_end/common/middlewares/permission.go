package middlewares

import (
	"alto_server/common/pkg/e"
	"alto_server/common/utils"
	"fmt"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
// used for permission management for different account level (admin,common,customer)
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetRoleName gets the user role from the context.
func (a *BasicAuthorizer) GetRoleName(c *gin.Context) interface{} {
	inRole := c.GetString("role")
	fmt.Printf("GetRoleName   ^_^ role  : %+v \n", inRole)

	//inUsername := c.GetString("username")
	//return "customer"
	return inRole
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	// if _, ok := c.Get("islogin"); ok {
	// 	fmt.Printf("CheckPermission   ^_^  login !\n")
	// 	//c.Next()
	// 	return true
	// }
	role := a.GetRoleName(c)
	method := c.Request.Method
	path := c.Request.URL.Path
	return a.enforcer.Enforce(role, path, method)
}

// RequirePermission returns the error information
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	//c.AbortWithStatus(403)
	fmt.Printf("RequirePermission      \n")
	utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_PERMISSION, e.GetMessage(e.ERROR_PERMISSION))
}
