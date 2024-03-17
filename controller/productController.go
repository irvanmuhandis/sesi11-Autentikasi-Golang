package controller

import (
	"net/http"
	"sesi11/database"
	"sesi11/helpers"
	"sesi11/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	_, _ = db, contentType

	product := model.Product{}
	userId := uint(userData["id"].(float64))
	if contentType == APPJSON {
		c.ShouldBindJSON(&product)

	} else {
		c.ShouldBind(&product)
	}
	product.UserID = userId

	err := db.Debug().Create(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	product := model.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))
	if contentType == APPJSON {
		c.ShouldBindJSON(&product)

	} else {
		c.ShouldBind(&product)
	}
	product.UserID = userId
	product.ID = uint(productId)

	err := db.Model(&product).Where("id=?", productId).Updates(model.Product{
		Title: product.Title,
		Desc:  product.Desc,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
