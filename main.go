package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtKey = []byte(os.Getenv("API_KEY"))
var tokens []string

type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type Expense struct {
	Price   float64 `json:"price"`
	Article string  `json:"article"`
}

var Expenses = []Expense{
	{1.99, "tjard"},
	{2.99, "jard"},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()

	router.POST("/register", gin.BasicAuth(gin.Accounts{
		"tjard": "123456",
	}), func(c *gin.Context) {
		token, _ := generateJWT()
		tokens = append(tokens, token)

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	router.GET("/resource", func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		reqToken := strings.Split(bearerToken, " ")[1]
		claims := &Claims{}
		tkn, _ := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "resource data"})
	})

	router.Run("localhost:8080")
}

func generateJWT() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Name: "tjard",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
