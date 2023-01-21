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

// CreateUser godoc
// @ID CreateUser
// @Router /user [POST]
// @Summary CreateUser
// @Description CreateUser
// @Tags User_2
// @Accept json
// @Produce json
// @Param user body models.UpdateUserSwag true "CreateUserRequestBody"
// @Success 201 {object} models.CreateUser "GetUsertBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateUser(c *gin.Context) {

	var user models.CreateUser

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.User().Insert(context.Background(), &user)
	if err != nil {
		log.Println("error whiling create praduct :", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.User().GetByID(context.Background(), &models.UserPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id product:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDUser godoc
// @ID Get_By_IDUser
// @Router /user/{id} [GET]
// @Summary GetByID User
// @Description GetByID User
// @Tags User_2
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.User "GetByIDUserBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDUser(c *gin.Context) {

	id := c.Param("id")
	res, err := h.storage.User().GetByID(context.Background(), &models.UserPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id product:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListUser godoc
// @ID UserPrimeryKey
// @Router /user [GET]
// @Summary Get List User
// @Description Get List user
// @Tags User_2
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListUserResponse "GetUserListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListUser(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
		search    = c.Query("search")
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

	res, err := h.storage.User().GetList(context.Background(), &models.GetListUserRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})

	if err != nil {
		log.Println("error whiling get list User:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateUser godoc
// @ID UpdateUser
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User_2
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.UpdateUserSwag true "UpdateUserRequestBody"
// @Success 202 {object} models.User "UpdateUserBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var (
		user models.UpdateUser
	)

	id := c.Param("id")

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = h.storage.User().Update(context.Background(), &models.UpdateUser{
		Id:        id,
		Name:      user.Name,
		Login:     user.Login,
		Password:  user.Password,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		log.Printf("error whiling update user : %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}
	c.JSON(http.StatusAccepted, id)
}

// DeleteUser godoc
// @ID DeleteUser
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User_2
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteUserBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.User().Delete(context.Background(), &models.UserPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  User:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "delete User")
}
