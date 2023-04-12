package routers

import (
	"challenge-9/controllers"
	"challenge-9/middlewares"
	"challenge-9/services"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func StartServer(db *gorm.DB) *gin.Engine {
	app := gin.Default()

	var (
		userService = services.UserService{
			DB: db,
		}

		userController = controllers.UserController{
			UserService: &userService,
		}

		productService = services.ProductService{
			DB: db,
		}

		productController = controllers.ProductController{
			ProductService: &productService,
		}

		auth        = middlewares.Authentication(&userService)
		productAuth = middlewares.ProductAuth(&productService)
	)

	userRouter := app.Group("/users")
	{
		// Register
		userRouter.POST("/register", userController.Register)
		// Login
		userRouter.POST("/login", userController.Login)

		// Admin-only, get All Users
		userRouter.GET("/", auth, userController.GetAllUser)

		// Update user
		userRouter.PUT("/account", auth, userController.EditUser)

		//Delete
		userRouter.DELETE("/account", auth, userController.DeleteUserByEmail)
	}

	productRouter := app.Group("/products")
	productRouter.Use(auth)
	{

		// Create
		productRouter.POST("/", productController.CreateProduct)

		// Read
		productRouter.GET("/", productController.GetAllProductByUserId)

		// Update
		productRouter.PUT("/:productId", productAuth, productController.UpdateProduct)

		// Delete
		productRouter.DELETE("/:productId", productAuth, productController.DeleteProduct)
	}

	return app
}
