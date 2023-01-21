package handler

import (
	"context"
	"log"
	"net/http"

	"add/models"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @ID CreateBranch
// @Router /branch [POST]
// @Summary CreateBranch
// @Description CreateBranch
// @Tags Branch_1
// @Accept json
// @Produce json
// @Param Branch body models.CreateBranch true "CreateBranchRequestBody"
// @Success 201 {object} models.Branch "GetBranchtBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateBranch(c *gin.Context) {

	var Branch models.CreateBranch

	err := c.ShouldBindJSON(&Branch)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Branch().Create(context.Background(), &Branch)
	if err != nil {
		log.Println("error whiling create branch:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Branch().GetByID(context.Background(), &models.BranchPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id branch:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}
