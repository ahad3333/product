package handler

import (
	"context"
	"log"
	"net/http"

	"add/models"

	"github.com/gin-gonic/gin"
)

// CreateBucket godoc
// @ID CreateBucket
// @Router /bucket [POST]
// @Summary CreateBucket
// @Description CreateBucket
// @Tags Bucket_7
// @Accept json
// @Produce json
// @Param Bucket body models.CreateBucket true "CreateBucketRequestBody"
// @Success 201 {object} models.Bucket "GetBuckettBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateBucket(c *gin.Context) {

	var Bucket models.CreateBucket

	err := c.ShouldBindJSON(&Bucket)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Order().CreateBucket(context.Background(), &Bucket)
	if err != nil {
		log.Println("error whiling create Bucket:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

// CreateCasier godoc
// @ID CreateCasier
// @Router /casier [POST]
// @Summary CreateCasier
// @Description CreateCasier
// @Tags Bucket_7
// @Accept json
// @Produce json
// @Param Bucket body models.GetBucketByClientID true "GetBucketByClientIDBody"
// @Success 201 {object} models.GetBucketByClientResponse "GetBucketByClientResponseBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateCasier(c *gin.Context) {

	var clientBucket models.GetBucketByClientID

	err := c.ShouldBindJSON(&clientBucket)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.storage.Order().GetBucketByClientID(context.Background(), &clientBucket)
	if err != nil {
		log.Println("error whiling GetBucketByClientID:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.storage.Order().Create(context.Background(),
		&models.CreateOrder{
			TotalPrice: resp.TotalPrice,
			BranchId:   clientBucket.BranchId,
			ClientId:   clientBucket.ClientId,
		},
	)
	if err != nil {
		log.Println("error whiling create order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
		c.JSON(http.StatusCreated, resp)
		
}
