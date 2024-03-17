package controller

import (
	"fmt"
	"net/http"
	"sesi11/database"
	"sesi11/helpers"
	"sesi11/model"

	"github.com/gin-gonic/gin"
)

var APPJSON = "application/json"

func UserRegis(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	user := model.User{}

	if contentType == APPJSON {
		c.ShouldBindJSON(&user)

	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.Fullname,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	user := model.User{}
	pass := ""
	if contentType == APPJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}
	//Password nya disimpan ke variabel dulu
	pass = user.Password
	// Lalu user.password di overwrite dengan pass dari database
	err := db.Debug().Where("email=?", user.Email).Take(&user).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	//Dibandingkan password yang dari inputan dan dari database
	comparePass := helpers.ComparePass([]byte(user.Password), []byte(pass))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return

	}

	token := helpers.GenerateToken(user.ID, user.Email)
	fmt.Printf("user.ID: %v\n", user.ID)
	fmt.Printf("user.Email: %v\n", user.Email)
	fmt.Printf("helpers.GenerateToken(user.ID, user.Email): %v\n", helpers.GenerateToken(user.ID, user.Email))
	ctx.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "Behasil Login",
	})

}

func UserLogout(ctx *gin.Context) {

}
