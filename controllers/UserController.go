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
