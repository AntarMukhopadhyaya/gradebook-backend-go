package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gradebook-api/models"
	"io/ioutil"
	"net/http"
)

type ResponseQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
type Payload struct {
	Topic             string `json:"topic"`
	NumberOfQuestions int    `json:"num_questions"`
}

func CreateAssignment(c *gin.Context) {
	fmt.Println("In Controller")
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

	// generate the question from the open ai api
	// create the payload
	//payload := Payload{
	//	NumberOfQuestions: body.NumberOfQuestions,
	//	Topic:             body.Topic,
	//}
	//jsonPayload, err := json.Marshal(payload)
	//if err != nil {
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	return
	//}
	url := "http://127.0.0.1:5000/generate-qa"
	var jsonData = []byte(`{
		"topic": "Data Warehouse",
		"num_questions":1
	}`)
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	respbody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response body", string(respbody))
}
