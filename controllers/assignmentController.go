package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"gradebook-api/helpers"
	"gradebook-api/models"
	"net/http"
)

type ResponseData struct {
	Questions []struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	} `json:"questions"`
}
type Payload struct {
	Topic             string `json:"topic"`
	NumberOfQuestions int    `json:"num_questions"`
}

func CreateAssignment(c *gin.Context) {

	var body struct {
		Topic string `json:"topic" binding:"required"`
		//DueDate           time.Time `json:"due_date" binding:"required"`
		AssignedClass     string `json:"assigned_class" binding:"required"`
		NumberOfQuestions int    `json:"number_of_questions" binding:"required"`
	}
	if c.Bind(&body) != nil {
		fmt.Println("Failed to Bind body")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, exists := c.Get("user")
	if exists == false {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//Checking if the user role is not teacher
	if user.(models.User).Role != "teacher" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	url := "http://127.0.0.1:5000/generate-qa"
	payload := Payload{
		NumberOfQuestions: body.NumberOfQuestions,
		Topic:             body.Topic,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	var data ResponseData

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//Creating a transaction due to ACID properties
	tx := helpers.DB.Begin()
	//At First creating the assignment with the help of database
	assignment := models.Assignment{Topic: body.Topic, UserID: user.(models.User).ID, AssignedClass: body.AssignedClass, NumberOfQuestions: body.NumberOfQuestions, DueDate: time.Now()}
	result := tx.Create(&assignment)
	if result.Error != nil {
		tx.Rollback()
		fmt.Println(result.Error)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//Now creating the questions
	for _, q := range data.Questions {
		question := models.Question{
			Topic:        body.Topic,
			UserID:       user.(models.User).ID,
			Question:     q.Question,
			Answer:       q.Answer,
			AssignmentID: assignment.ID,
		}
		result := tx.Create(&question)
		if result.Error != nil {
			tx.Rollback()
			fmt.Println(result.Error)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
func ShowAssignment(c *gin.Context) {
	user, exists := c.Get("user")
	if exists == false {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var assignments []models.Assignment
	helpers.DB.Where("user_id = ?", user.(models.User).ID).Preload("Questions").Find(&assignments)
	c.JSON(http.StatusOK, gin.H{
		"data": assignments,
	})

}
func IndexAssignment(c *gin.Context) {
	var assignments []models.Assignment
	helpers.DB.Preload("Questions").Find(&assignments)
	c.JSON(http.StatusOK, gin.H{
		"data": assignments,
	})
}
