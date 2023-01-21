package api

import (
	_ "add/api/docs"
	"add/api/handler"
	"add/config"
	"add/storage"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg config.Config, storage storage.StorageI) {

	handlerV1 := handler.NewHandler(cfg, storage)

	r.POST("/branch", handlerV1.CreateBranch)

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetByIDUser)
	r.GET("/user", handlerV1.GetListUser)
	r.DELETE("/user/:id", handlerV1.DeleteUser)
	r.PUT("/user/:id", handlerV1.UpdateUser)

	r.POST("/login", handlerV1.Login)

	r.POST("/product", handlerV1.CreateProduct)
	r.GET("/product/:id", handlerV1.GetByIDProduct)
	r.GET("/product", handlerV1.GetListProduct)
	r.PUT("/product/:id", handlerV1.UpdateProduct)
	r.DELETE("/product/:id", handlerV1.DeleteProduct)
	r.POST("/client", handlerV1.CreateClient)

	r.POST("/bucket", handlerV1.CreateBucket)
	r.POST("/casier", handlerV1.CreateCasier)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetByIdCategory)
	r.GET("/category", handlerV1.GetListCategory)
	r.DELETE("/category/:id", handlerV1.DeleteCategory)
	r.PUT("/category/:id", handlerV1.UpdateCategory)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
