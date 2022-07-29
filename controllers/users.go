package controllers 

 import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dblab/questlist/models"
	"time"
	"os"
	"dblab/questlist/initializers"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
 )



type SignupInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

 func Signup(c *gin.Context) {
	// Validate input
	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var user models.User
	if err := initializers.DB.Where("Email = ?", input.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Email already exists!"})
		return
	}

	// Create user
	user = models.User{Email: input.Email, Password: string(hashedPassword), Username: input.Username, Role: models.UserRoleCommon}
	initializers.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context){
	var user LoginInput
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFromDB models.User
	if err := initializers.DB.Where("email = ?", user.Email).First(&userFromDB).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password!"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password!"})
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userFromDB.ID,
		"email": userFromDB.Email,
		"username": userFromDB.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"role": userFromDB.Role,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "accessToken": tokenString })

}

func GetUser(c *gin.Context) {
	user, _ := c.Get("user")
	myUser := user.(models.User)
	c.JSON(http.StatusOK, gin.H{"data": myUser})

}

func VerifyUser(c *gin.Context) {
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	user.Verified = true
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UpdateUserInput struct {
	Email string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Birthday string `json:"birthday"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Email = input.Email
	user.Username = input.Username
	user.Birthday = input.Birthday
	user.Phone = input.Phone
	user.Address = input.Address

	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func ChangePassword(c *gin.Context) {
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hashedPassword)

	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}


