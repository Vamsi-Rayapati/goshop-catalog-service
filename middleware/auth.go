package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smartbot/catalog/pkg/errors"
)

type JWKS struct {
	Keys []JWK `json:"keys"`
}
type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func authenticationMiddleware(c *gin.Context) {

	log.Printf("authenticationMiddleware")
	k, err := keyfunc.NewDefaultCtx(c, []string{"http://fusionauth-service:9011/.well-known/jwks.json"}) // Context is used to end the refresh goroutine.
	if err != nil {
		c.JSON(http.StatusInternalServerError, &errors.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to authenticate",
		})
	}

	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	token, err := jwt.Parse(tokenString, k.Keyfunc)

	if err != nil || !token.Valid {

		log.Printf("ERRR %+v", err)
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	c.Set("user_id", claims["sub"])
	c.Set("roles", claims["roles"])
	c.Set("username", claims["username"])
	c.Next()

}

func Authenticate() gin.HandlerFunc {
	return authenticationMiddleware
}
