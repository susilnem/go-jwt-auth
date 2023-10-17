package controller

import (
	"go-jwt/database"
	"go-jwt/internal/format_errors"
	"go-jwt/internal/pagination"
	model "go-jwt/models"
	"go-jwt/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
// @BasePath /api/auth
// @Summary Register
// @Description Register new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body AuthenticationInput true "Register"
// @Success 200 {object} string "ok"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var input AuthenticationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user already exists
	var user model.User
	database.Database.Where("username = ?", input.Username).First(&user)

	if user.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// hash password before saving to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return
	}

	newUser := model.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	// save user to database
	database.Database.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}
// @BasePath /api/auth
// @Summary Login
// @Description Login user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body AuthenticationInput true "Login"
// @Success 200 {object} string "ok"
// @Router /auth/login [post]
func Login(c *gin.Context) {

	var input AuthenticationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user exists
	var user model.User
	database.Database.Where("username = ?", input.Username).First(&user)

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	// compare password and if true then generate token
	err = utils.ValidatePassword(user.Password, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate token and return
	token, err := utils.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
// @BasePath /api/users
// @Summary Get Users
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param perPage query int false "Per Page"
// @Success 200 {object} string "ok"
// @Router /users [get]
// @Security Bearer
func GetUsers(c *gin.Context) {
	var users []model.User

	//set default page and perPage
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, _ := strconv.Atoi(perPageStr)

	result, err := pagination.Paginate(database.Database, page, perPage, nil, &users)
	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
// @BasePath /api/users
// @Summary Get User
// @Description Get user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} string "ok"
// @Router /users/{id} [get]
// @Security Bearer
func EditUser(c *gin.Context) {
	// get user id from url
	id := c.Param("id")

	var user model.User
	result := database.Database.First(&user, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @BasePath /api/users
// @Summary Update User
// @Description Update user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body AuthenticationInput true "Update"
// @Success 200 {object} string "ok"
// @Router /users/{id}/update [put]
// @Security Bearer
func UpdateUser(c *gin.Context) {
	// get user id from url
	id := c.Param("id")

	var user model.User
	result := database.Database.First(&user, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	var input model.User
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @BasePath /api/users
// @Summary Delete User
// @Description Delete user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} string "ok"
// @Router /users/{id}/delete [delete]
// @Security Bearer
func DeleteUser(c *gin.Context) {
	// get user id from url
	id := c.Param("id")

	var user model.User
	result := database.Database.First(&user, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	database.Database.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "The user has been deleted successfully",
	})
}