package daos

import (
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/config"
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

var userByID = "id = ?"

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get does the actual query to database, if user with specified id is not
// found error is returned
func (dao *UserDAO) Get(id uint) (*models.User, error) {
	var user models.User
	err := config.Config.DB.Where(userByID, id).First(&user).Error
	return &user, err
}

// GetAll does the actual query to database, if user with specified id is not
// found error is returned
func (dao *UserDAO) GetAll() (*[]models.User, error) {
	var users []models.User
	err := config.Config.DB.Find(&users).Error
	return &users, err
}

// Create creates a user
func (dao *UserDAO) Create(input *models.UserCreate) (*models.User, error) {
	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		UserName:  input.UserName,
		Email:     input.Email,
	}
	err := config.Config.DB.Create(&user).Error
	return &user, err
}

// Update updates a user
func (dao *UserDAO) Update(id uint, input *models.UserUpdate) (*models.User, error) {
	var user models.User
	err := config.Config.DB.Where(userByID, id).First(&user).Updates(input).Error
	return &user, err
}

// Delete deletes a user
func (dao *UserDAO) Delete(id uint) error {
	var user models.User
	err := config.Config.DB.Where(userByID, id).First(&user).Delete(&user).Error
	return err
}
