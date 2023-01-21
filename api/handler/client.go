package handler

import (
	"context"
	"log"
	"net/http"

	"add/models"

	"github.com/gin-gonic/gin"
)

// CreateClient godoc
// @ID CreateClient
// @Router /client [POST]
// @Summary CreateClient
// @Description CreateClient
// @Tags Client_6
// @Accept json
// @Produce json
// @Param Client body models.CreateClient true "CreateClientRequestBody"
// @Success 201 {object} models.Client "GetClienttBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateClient(c *gin.Context) {

	var Client models.CreateClient

	err := c.ShouldBindJSON(&Client)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Client().Create(context.Background(), &Client)
	if err != nil {
		log.Println("error whiling create Client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Client().GetByID(context.Background(), &models.ClientPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id Client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}
