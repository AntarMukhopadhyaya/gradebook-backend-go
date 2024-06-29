package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func createAssignment(c *gin.Context) {
	var body struct {
		Topic             string    `json:"topic" binding:"required"`
		DueDate           time.Time `json:"due_date" binding:"required"`
		AssignedClass     string    `json:"assigned_class" binding:"required"`
		NumberOfQuestions int       `json:"number_of_questions" binding:"required"`
	}
	if c.Bind(&body) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// generate the question from the open ai api
	
}
