package middleware

import (
	"errors"
	"fmt"
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
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(secretTokenJwt), nil
	})
	fmt.Println(secretTokenJwt)
	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func (ctrl JwtMiddleware) JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()
}
