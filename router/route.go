package router

import (
	"sesi11/controller"
	"sesi11/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegis)
		userRouter.POST("/login", controller.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controller.CreateProduct)
		productRouter.POST("/:productId", middleware.ProductAuthorization(), controller.CreateProduct)

	}
	return r
}
