package routes

import (
	"gin1gorm/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		user.CreateUser(db)
	}

}
