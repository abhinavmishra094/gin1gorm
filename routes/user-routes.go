package routes

import (
	"gin1gorm/models"
	"gin1gorm/services"
	"gin1gorm/util"
	"strconv"

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
		var user models.User
		err := c.BindJSON(&login)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		u, err := user.GetUserByUserName(login.UserName, db)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else if util.DoPasswordsMatch(u.Password, login.Password) {
			token, err := jwtService.GenerateToken(u.UserName, false)
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
		number, err := strconv.ParseUint(string(id), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		u, err := user.GetUserByID(number, db)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			userOut := models.UserOut{
				ID:       u.ID,
				UserName: u.UserName,
				Email:    u.Email,
				Age:      u.Age,
			}
			c.JSON(200, userOut)
		}
	}
}

func GetUserByUserName(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		userName := c.Param("username")
		var user models.User
		u, err := user.GetUserByUserName(userName, db)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			userOut := models.UserOut{
				ID:       u.ID,
				UserName: u.UserName,
				Email:    u.Email,
				Age:      u.Age,
			}
			c.JSON(200, userOut)
		}
	}
}

func DeleteUserByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		number, err := strconv.ParseUint(string(id), 10, 64)
		_, e := user.DeleteUserByID(number, db)
		if e != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"message": "User deleted"})
		}
	}
}

func DeleteUserByUseName(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		username := c.Param("username")
		var user models.User

		_, e := user.DeleteUserByUserName(username, db)
		if e != nil {
			c.JSON(400, gin.H{"error": e.Error()})
		} else {
			c.JSON(200, gin.H{"message": "User deleted"})
		}
	}
}
