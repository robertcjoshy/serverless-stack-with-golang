package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robert/serverless-stack-with-golang/src/models"
)

// type createemail struct {
// 	MAIL string `json:"mail"`
// }

type Response struct {
	Res string `json:"res"`
}

func GetUser(c *gin.Context) {

	email := c.Param("id")

	email_result, err := models.GetEmail(email)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, email_result)

}

func CreateUser(c *gin.Context) {

	var req models.User

	//email_valid := validaters.IsEmailValid(req.Email)

	// if !email_valid {
	// 	errr := errors.New("invalid email")
	// 	c.AbortWithError(http.StatusInternalServerError, errr)
	// 	return
	// }

	//log.Println("email=", email_valid)

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println("json binded with no error")
	//mail := req.MAIL
	log.Println("json := ", req)
	created_email, err := req.SaveEmail()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(created_email.Email) == 0 {
		log.Println(created_email)
	}
	c.JSON(201, created_email)

}

func UpdateUser() {

}

func DeleteUser() {

}
