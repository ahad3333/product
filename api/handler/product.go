package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"add/models"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @ID CreateProduct
// @Router /product [POST]
// @Summary CreateProduct
// @Description CreateProduct
// @Tags Product_5
// @Accept json
// @Produce json
// @Param product body models.CreateProduct true "CreateProductRequestBody"
// @Success 201 {object} models.CreateProduct "GetProductBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateProduct(c *gin.Context) {

	var praduct models.CreateProduct

	err := c.ShouldBindJSON(&praduct)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Product().Insert(context.Background(), &praduct)
	if err != nil {
		log.Println("error whiling create praduct :", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// resp, err := h.storage.Product().GetByID(context.Background(), &models.ProductPrimeryKey{
	// 	Id: id,
	// })
	// if err != nil {
	// 	log.Println("error whiling get by id book:", err.Error())
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusCreated, id)
}

// GetByIDProduct godoc
// @ID Get_By_IDProduct
// @Router /product/{id} [GET]
// @Summary GetByID Product
// @Description GetByID Product
// @Tags Product_5
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Product "GetByIDProductBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDProduct(c *gin.Context) {

	id := c.Param("id")
	res, err := h.storage.Product().GetByID(context.Background(), &models.ProductPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id product:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListProduct godoc
// @ID ProductPrimeryKey
// @Router /product [GET]
// @Summary Get List Product
// @Description Get List product
// @Tags Product_5
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListProductResponse "GetProdctListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListProduct(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
		search = c.Query("search")

	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res, err := h.storage.Product().GetList(context.Background(),&models.GetListProductRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})

	if err != nil {
		log.Println("error whiling get list book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateProduct godoc
// @ID UpdateProduct
// @Router /product/{id} [PUT]
// @Summary Update Product
// @Description Update Product
// @Tags Product_5
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param product body models.UpdateProductPut true "UpdateProductRequestBody"
// @Success 202 {object} models.Product "UpdateProductBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateProduct(c *gin.Context) {

	var (
		praduct models.UpdateProduct
	)

	id := c.Param("id")

	err := c.ShouldBindJSON(&praduct)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	 err = h.storage.Product().Update(context.Background(),&models.UpdateProduct{
		Id: id,
		Name: praduct.Name,
		Price: praduct.Price,
		Description: praduct.Description,
		Photo: praduct.Photo,
		CategoryId: praduct.CategoryId,
	})
	if err != nil {
		log.Printf("error whiling update 2: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	// resp, err := h.storage.Product().GetByID(context.Background(),&models.ProductPrimeryKey{Id: book.Id})
	// if err != nil {
	// 	log.Printf("error whiling get by id: %v\n", err)
	// 	c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
	// 	return
	// }
	c.JSON(http.StatusAccepted, id)
}

// DeleteProduct godoc
// @ID DeleteProduct
// @Router /product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product_5
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteProductBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Product().Delete(context.Background(),&models.ProductPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  product:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "delete product")
}
