package controllers

import (
	"net/http"

	"effective_mobile/models"
	"effective_mobile/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func InitPeopleController(database *gorm.DB) {
	db = database
}

func AddPerson(c *gin.Context) {
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	age, err := services.GetAge(person.Name)
	if err != nil {
		logrus.Error("Error fetching age", err)
	}

	gender, err := services.GetGender(person.Name)
	if err != nil {
		logrus.Error("Error fetching gender", err)
	}

	nationality, err := services.GetNationality(person.Name)
	if err != nil {
		logrus.Error("Error fetching nationality", err)
	}

	person.Age = age
	person.Gender = gender
	person.Nationality = nationality

	if err := db.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving person"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func GetPeople(c *gin.Context) {
	var people []models.Person
	if err := db.Find(&people).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching people"})
		return
	}
	c.JSON(http.StatusOK, people)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	if err := db.Where("id = ?", id).Delete(&models.Person{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := db.Model(&models.Person{}).Where("id = ?", id).Updates(person).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}
