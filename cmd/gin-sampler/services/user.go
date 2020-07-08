package services

import (
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/models"
)

type userDAO interface {
	Get(id uint) (*models.User, error)
	GetAll() (*[]models.User, error)
	Create(input *models.UserCreate) (*models.User, error)
	Update(id uint, input *models.UserUpdate) (*models.User, error)
	Delete(id uint) error
}

// UserService contains a userDAO
type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for
// processing data retrieved by DAOs
func (s *UserService) Get(id uint) (*models.User, error) {
	return s.dao.Get(id)
}

// GetAll just retrieves user using User DAO, here can be additional logic for
// processing data retrieved by DAOs
func (s *UserService) GetAll() (*[]models.User, error) {
	return s.dao.GetAll()
}

// Create takes input and instantiates it in a database
func (s *UserService) Create(input *models.UserCreate) (*models.User, error) {
	return s.dao.Create(input)
}

// Update takes an id and data input and updates the user with the specified id
// with the provided data
func (s *UserService) Update(id uint, input *models.UserUpdate) (*models.User, error) {
	return s.dao.Update(id, input)
}

// Delete takes a user id and deletes the user associated with the id if the
// user exists
func (s *UserService) Delete(id uint) error {
	return s.dao.Delete(id)
}
