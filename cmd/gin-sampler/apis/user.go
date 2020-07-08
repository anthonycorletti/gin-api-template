package apis

import (
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/daos"
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/models"
	"github.com/anthcor/gin-sampler/cmd/gin-sampler/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUser retrieves a user
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found."})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUsers retrieves a user
func GetUsers(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	if users, err := s.GetAll(); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "No users."})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// CreateUser creates a user
func CreateUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	var input models.UserCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if user, err := s.Create(&input); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Oops we failed to create the user."})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if _, err := s.Get(uint(id)); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found."})
	} else {
		var input models.UserUpdate
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if _, err := s.Update(uint(id), &input); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Oops we failed to update the user."})
		} else {
			GetUser(c)
		}
	}
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found."})
	} else {
		if err := s.Delete(uint(id)); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Oops we failed to delete the user."})
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}
