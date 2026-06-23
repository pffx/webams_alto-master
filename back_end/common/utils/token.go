package utils

import (
	"alto_server/common/pkg/e"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username string `json:"username"`
	//Password             string `json:"password"`
	Role                 string `json:"role"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 720

var MySecret = []byte("654321")

func GetToken(c *gin.Context) (tokenString string) {
	//method := c.Request.Method
	// if method == "GET" {
	// 	return c.DefaultQuery("token", "")
	// } else {
	// 	return c.DefaultPostForm("token", "")

	// }
	for k, v := range c.Request.Header {
		if k == "Token" {
			//return c.Request.Header.Get("token")
			return strings.Join(v, "")
		}
	}
	return ""
}

func GetUserInfoFromToken(c *gin.Context) (*MyClaims, error) {
	token := GetToken(c)
	claims, err := ParseToken(token)

	return claims, err
}

// GenToken 生成JWT
func GenToken(username, role string) (tokenString string, err error) {
	claim := MyClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间2小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(MySecret)
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil // 这是我的secret
	}
}

func ParseToken(tokenss string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &MyClaims{}, Secret())
	// fmt.Println("token   ^_^  ParseToken: ", err.Error())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, e.ErrTokenAuthFailed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, e.ErrTokenTimeOut
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, e.ErrTokenAuthFailed
			} else {
				return nil, e.ErrTokenInvalid
			}
		}
	}
	// fmt.Printf("ParseToken   ^_^ whole JWT token: \n %+v \n", token)
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		//fmt.Printf("ParseToken   ^_^ whole JWT claims: \n %+v \n", claims)
		return claims, nil
	}
	return nil, e.ErrTokenInvalid
}
