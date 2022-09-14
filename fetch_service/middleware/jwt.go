package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type JwtMiddleware struct{}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}
type MyClaims struct {
	jwt.StandardClaims
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("Unauthorized")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New("error get secret key")
	}

	secretTokenJwt := os.Getenv("JWT_KEY")
	token, err := jwt.ParseWithClaims(jwtToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretTokenJwt), nil
	})
	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func JwtTokenCheck(c *gin.Context) string {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return ""
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return ""
	}

	claims, OK := token.Claims.(*MyClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return ""
	}
	return claims.Role
}

func (ctrl JwtMiddleware) AdminRole(c *gin.Context) {
	roleJwt := JwtTokenCheck(c)
	if roleJwt == "admin" {
		c.Next()
	} else {
		return
	}
}

func (ctrl JwtMiddleware) UserRole(c *gin.Context) {
	roleJwt := JwtTokenCheck(c)
	if roleJwt == "user" {
		c.Next()
	} else {
		return
	}
}
