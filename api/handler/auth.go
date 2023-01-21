package handler

import (
	"errors"
	"log"
	"net/http"

	"add/models"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID Login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Auth_3
// @Accept json
// @Produce json
// @Param login body models.Login true "loginBodyRequest"
// @Success 201 {object} models.LoginResponse "LoginResponseBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) Login(c *gin.Context) {

	var login models.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	exists, err := h.storage.User().CheckUser(c.Request.Context(), &login)
	if err != nil {
		log.Println("error whiling storage check user:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		log.Println("error whiling not found user:", errors.New("login and password invalid").Error())
		c.JSON(http.StatusUnauthorized, errors.New("login and password invalid").Error())
		return
	}

	c.JSON(http.StatusCreated, models.LoginResponse{
		Key: h.cfg.SecretKey,
	})
}
