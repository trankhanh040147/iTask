package middleware

import (
	"log"
	"net/http"
	"paradise-booking/common"
	jwtprovider "paradise-booking/provider/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeader(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")

	parts := strings.Split(bearerToken, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", jwtprovider.ErrInvalidToken
	}
	return parts[1], nil
}

// check if token is valid or not
// B1: Get token from header
// B2: Validate token and get payload
// B3: From payload, use email to find user in db
func (m *middlewareManager) RequiredAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.Request)
		if err != nil {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			// return
			log.Printf("err: %v, when extractTokenFromHeader", err)
			panic(common.ErrAuthorized(err))
		}

		payload, err := jwtprovider.ValidateJWT(token, m.cfg)
		if err != nil {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			// return
			log.Printf("err: %v, when ValidateJWT", err)
			panic(common.ErrAuthorized(err))
		}
		account, err := m.accountSto.GetAccountByEmail(c.Request.Context(), payload.Email)
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{"error": err})
			// return
			panic(common.ErrBadRequest(err))
		}

		c.Set("Account", account)

		c.Next()
	}
}
