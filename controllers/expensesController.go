package controllers

import (
	"ExpenseAPI/initializers"
	"ExpenseAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

func AddExpense(c *gin.Context) {
	authCookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(401, gin.H{"error": "Authorization Cookie not found!"})
		return
	}

	token, err := jwt.Parse(authCookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_TOKEN")), nil
	})
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(500, gin.H{"error": "Failed to get valid claims"})
	}

	userIDVal, ok := claims["sub"].(float64)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID := int(userIDVal)

	var body struct {
		Expense string
		Price   float64
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	expense := models.Expense{UserID: userID, Expense: body.Expense, Price: body.Price}

	result := initializers.DB.Create(&expense)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create expense"})
		return
	}
	c.JSON(200, gin.H{"message": "Expense created successfully"})

}
