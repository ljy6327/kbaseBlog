package request

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_const "kbase/common/const"
	"kbase/model/dto"
	"net/http"
)

/**
  Token验证
*/
func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Token")
		if tokenString == "" {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{"message": "unToken"})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(_const.SecretKey), nil
		})
		if token != nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("author", dto.Author{
					Id: uint(claims["authorId"].(float64)),
				})
				c.Next()
			} else {
				c.Abort()
				fmt.Println(err)
			}
		} else {
			fmt.Println("noToken")
		}
	}
}
