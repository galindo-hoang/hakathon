package routes

import (
	"image-service/handlers"
	// "image-servicdde/helpers"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	setRoutes(r)
	// port := helpers.GetValueFromEnv("PORT")
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}

func setRoutes(route *gin.Engine) {
	detect := route.Group("/v1/detect")
	{
		detect.POST("/images", handlers.DetectImage)
	}

	shop := route.Group("/v1/shop")
	{
		shop.POST("", handlers.CreateShop)
		shop.GET("/all", handlers.GetAllShops)
		shop.GET("/:id", handlers.GetShopById)
		shop.GET("/:name", handlers.GetShopsByName)
		shop.PUT("/:id", handlers.UpdateShopInfo)
	}

	item := route.Group("/v1/item")
	{
		item.POST("")
		item.GET("/:name")
		item.PUT("/:id")
		item.DELETE("/:id")
	}

}
