package routes

import (
	"gin1gorm/models"
	"gin1gorm/services"
	"gin1gorm/util"

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
		user.Password = util.HashPassword(user.Password)
		u, e := user.CreateUser(db)
		userOut := models.UserOut{
			ID:       u.ID,
			UserName: u.UserName,
			Email:    u.Email,
			Age:      u.Age,
		}
		if e != nil {
			c.JSON(400, gin.H{"error": e.Error()})
		}
		c.JSON(200, userOut)
	}

}

func Login(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		jwtService := services.NewJWTService()
		var login models.Login

		err := c.BindJSON(&login)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		user, err := login.GetUserByUserName(login.UserName, db)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		if util.DoPasswordsMatch(user.Password, login.Password) {
			token, err := jwtService.GenerateToken(user.UserName, false)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
			}

			c.JSON(200, gin.H{"message": "Login Successful", "token": token})
		} else {
			c.JSON(400, gin.H{"error": "Invalid credentials"})
		}

	}
}

func GetUserByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		err := db.Debug().Where("id = ?", id).Take(&user).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		userOut := models.UserOut{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			Age:      user.Age,
		}
		c.JSON(200, userOut)
	}
}

func GetUserByUserName(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		userName := c.Param("username")
		var user models.User
		err := db.Debug().Where("user_name = ?", userName).Take(&user).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		userOut := models.UserOut{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			Age:      user.Age,
		}
		c.JSON(200, userOut)
	}
}
