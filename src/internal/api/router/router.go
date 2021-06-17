package router

import (
	"fmt"
	"io"
	"os"

	"github.com/aceslick911/bcg-checkout/internal/api/controllers"
	"github.com/aceslick911/bcg-checkout/internal/api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	app.GET("/", controllers.GatewayRedirect) // Redirect to docs
	// Routes
	// ================== Login Routes
	app.POST("/api/login", controllers.Login)
	app.POST("/api/login/add", controllers.CreateUser)
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ================== User Routes
	app.GET("/api/users", controllers.GetUsers)
	app.GET("/api/users/:id", controllers.GetUserById)
	app.POST("/api/users", controllers.CreateUser)
	app.PUT("/api/users/:id", controllers.UpdateUser)
	app.DELETE("/api/users/:id", controllers.DeleteUser)
	// ================== Tasks Routes
	app.GET("/api/tasks/:id", controllers.GetTaskById)
	app.GET("/api/tasks", controllers.GetTasks)
	app.POST("/api/tasks", controllers.CreateTask)
	app.PUT("/api/tasks/:id", controllers.UpdateTask)
	app.DELETE("/api/tasks/:id", controllers.DeleteTask)
	// ================== Product Routes
	app.GET("/api/products/:id", controllers.GetProductById)
	app.GET("/api/products", controllers.GetProducts)
	app.POST("/api/products", controllers.CreateProduct)
	app.PUT("/api/products/:id", controllers.UpdateProduct)
	app.DELETE("/api/products/:id", controllers.DeleteProduct)
	// ================== Discount Routes
	app.GET("/api/discounts/:id", controllers.GetDiscountById)
	app.GET("/api/discounts", controllers.GetDiscounts)
	app.POST("/api/discounts", controllers.CreateDiscount)
	app.PUT("/api/discounts/:id", controllers.UpdateDiscount)
	app.DELETE("/api/discounts/:id", controllers.DeleteDiscount)
	// ================== Discount Routes
	app.GET("/api/basket_items/:id", controllers.GetBasket_ItemById)
	app.GET("/api/basket_items", controllers.GetBasket_Items)
	app.POST("/api/basket_items", controllers.CreateBasket_Item)
	app.PUT("/api/basket_items/:id", controllers.UpdateBasket_Item)
	app.DELETE("/api/basket_items/:id", controllers.DeleteBasket_Item)
	// ================== Discount Routes
	app.GET("/api/baskets/:id", controllers.GetBasketById)
	app.GET("/api/baskets", controllers.GetBaskets)
	app.POST("/api/baskets", controllers.CreateBasket)
	app.PUT("/api/baskets/:id", controllers.UpdateBasket)
	app.DELETE("/api/baskets/:id", controllers.DeleteBasket)

	return app
}
