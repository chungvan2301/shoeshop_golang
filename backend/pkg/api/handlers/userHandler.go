package handlers

import (
	"net/http"
	"time"

	services "github.com/chungvan2301/shoeshop/backend/pkg/repository/interface"
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/models"
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/response"
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type UserHandler struct {
	userService services.UserRepository
}

func NewUserHandler(ser services.UserRepository) *UserHandler {
	return &UserHandler{
		userService: ser,
	}
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	userID, exist := c.Get("ID")
	if !exist {
		errResponse := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "Unauthorized")
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	useIDStr, ok := userID.(string)
	if !ok {
		errResponse := response.ClientResponse(http.StatusBadRequest, "ID not a string", nil, "ID not a string")
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	user, errG := h.userService.GetUserDetail(useIDStr)
	if errG != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "ID not found!", nil, "ID not found!")
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Edit user thành công", user, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.UserInput

	if err := c.ShouldBind(&user); err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Invalid requestl!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if err := validate.Struct(&user); err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Validated input!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Hash password fail!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	user.Password = string(hashedPassword)

	if err := h.userService.RegisterUser(user); err != nil {
		errResponse := response.ClientResponse(http.StatusInternalServerError, "Lỗi thêm user vào database", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Thêm user thành công", user, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *UserHandler) EditUser(c *gin.Context) {
	var user models.UserUpdate

	if err := c.ShouldBind(&user); err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Invalid requestl!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if err := validate.Struct(&user); err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Validated input!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	password, errP := h.userService.GetUserPassword(user.ID)
	if errP != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Not found password!", nil, errP.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	errC := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.OldPassword))
	if errC != nil {
		errResponse := response.ClientResponse(http.StatusUnauthorized, "Wrong password!", nil, errC.Error())
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Hash password fail!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	user.NewPassword = string(hashedPassword)

	if errE := h.userService.EditUser(user); errE != nil {
		errResponse := response.ClientResponse(http.StatusInternalServerError, "Lỗi edit user vào database", nil, errE.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Edit user thành công", user, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, exist := c.Get("ID")

	if !exist {
		errResponse := response.ClientResponse(http.StatusBadRequest, "ID not found in context!", nil, "ID not found in context!")
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	useIDStr, ok := userID.(string)
	if !ok {
		errResponse := response.ClientResponse(http.StatusBadRequest, "ID not a string", nil, "ID not a string")
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err := h.userService.DeleteUser(useIDStr)
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "ID not found!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Xóa user thành công", nil, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *UserHandler) Login(c *gin.Context) {
	var userLogin models.UserLogin

	if err := c.ShouldBind(&userLogin); err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Invalid input!", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	user, errG := h.userService.GetUserByEmail(userLogin.Email)
	if errG != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Not found email!", nil, errG.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	errC := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if errC != nil {
		errResponse := response.ClientResponse(http.StatusUnauthorized, "Invalid email or password!", nil, errC.Error())
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}

	accessToken, errA := token.GenerateJWT(user.ID, "access", 15*time.Minute)
	refreshToken, errR := token.GenerateJWT(user.ID, "refresh", 24*time.Hour)
	if errA != nil || errR != nil {
		errResponse := response.ClientResponse(http.StatusUnauthorized, "Could not generate token", nil, "Could not generate token	")
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Edit user thành công", models.TokenResponse{
		AccessToken: models.AccessToken{
			Token:     accessToken,
			ExpiresAt: time.Now().Add(15 * time.Minute),
		},
		RefreshToken: models.RefreshToken{
			Token:     refreshToken,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		},
	}, nil)

	c.JSON(http.StatusOK, successResponse)
}
